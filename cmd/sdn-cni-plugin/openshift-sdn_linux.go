package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/containernetworking/plugins/pkg/ipam"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/openshift/origin/pkg/network/node/cniserver"
	"github.com/vishvananda/netlink"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type cniPlugin struct {
	socketPath string
	hostNS     ns.NetNS
}

func NewCNIPlugin(socketPath string, hostNS ns.NetNS) *cniPlugin {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cniPlugin{socketPath: socketPath, hostNS: hostNS}
}
func newCNIRequest(args *skel.CmdArgs) *cniserver.CNIRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	envMap := make(map[string]string)
	for _, item := range os.Environ() {
		idx := strings.Index(item, "=")
		if idx > 0 {
			envMap[strings.TrimSpace(item[:idx])] = item[idx+1:]
		}
	}
	return &cniserver.CNIRequest{Env: envMap, Config: args.StdinData}
}
func (p *cniPlugin) doCNI(url string, req *cniserver.CNIRequest) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CNI request %v: %v", req, err)
	}
	client := &http.Client{Transport: &http.Transport{Dial: func(proto, addr string) (net.Conn, error) {
		return net.Dial("unix", p.socketPath)
	}}}
	var resp *http.Response
	err = p.hostNS.Do(func(ns.NetNS) error {
		resp, err = client.Post(url, "application/json", bytes.NewReader(data))
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send CNI request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read CNI result: %v", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("CNI request failed with status %v: '%s'", resp.StatusCode, string(body))
	}
	return body, nil
}
func (p *cniPlugin) doCNIServerAdd(req *cniserver.CNIRequest, hostVeth string) (*current.Result, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req.HostVeth = hostVeth
	body, err := p.doCNI("http://dummy/", req)
	if err != nil {
		return nil, err
	}
	result, err := current.NewResult(body)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response '%s': %v", string(body), err)
	}
	return result.(*current.Result), nil
}
func (p *cniPlugin) testCmdAdd(args *skel.CmdArgs) (types.Result, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result, err := p.doCNIServerAdd(newCNIRequest(args), "dummy0")
	if err != nil {
		return nil, err
	}
	return convertToRequestedVersion(args.StdinData, result)
}
func (p *cniPlugin) CmdAdd(args *skel.CmdArgs) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := newCNIRequest(args)
	config, err := cniserver.ReadConfig(cniserver.CNIServerConfigFilePath)
	if err != nil {
		return err
	}
	var hostVeth, contVeth net.Interface
	err = ns.WithNetNSPath(args.Netns, func(hostNS ns.NetNS) error {
		hostVeth, contVeth, err = ip.SetupVeth(args.IfName, int(config.MTU), hostNS)
		if err != nil {
			return fmt.Errorf("failed to create container veth: %v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	result, err := p.doCNIServerAdd(req, hostVeth.Name)
	if err != nil {
		return err
	}
	if err != nil || len(result.IPs) != 1 || result.IPs[0].Version != "4" {
		return fmt.Errorf("Unexpected IPAM result: %v", err)
	}
	defaultGW := result.IPs[0].Gateway
	result.IPs[0].Gateway = nil
	result.Interfaces = []*current.Interface{{Name: args.IfName, Mac: contVeth.HardwareAddr.String(), Sandbox: args.Netns}}
	result.IPs[0].Interface = current.Int(0)
	err = ns.WithNetNSPath(args.Netns, func(hostNS ns.NetNS) error {
		if err := ip.SetHWAddrByIP(args.IfName, result.IPs[0].Address.IP, nil); err != nil {
			return fmt.Errorf("failed to set pod interface MAC address: %v", err)
		}
		if err := ipam.ConfigureIface(args.IfName, result); err != nil {
			return fmt.Errorf("failed to configure container IPAM: %v", err)
		}
		link, err := netlink.LinkByName("lo")
		if err == nil {
			err = netlink.LinkSetUp(link)
		}
		if err != nil {
			return fmt.Errorf("failed to configure container loopback: %v", err)
		}
		link, err = netlink.LinkByName("macvlan0")
		if err == nil {
			err = netlink.LinkSetUp(link)
			if err != nil {
				return fmt.Errorf("failed to enable macvlan device: %v", err)
			}
			var addrs []netlink.Addr
			err = hostNS.Do(func(ns.NetNS) error {
				parent, err := netlink.LinkByIndex(link.Attrs().ParentIndex)
				if err != nil {
					return err
				}
				addrs, err = netlink.AddrList(parent, netlink.FAMILY_V4)
				return err
			})
			if err != nil {
				return fmt.Errorf("failed to configure macvlan device: %v", err)
			}
			var dsts []*net.IPNet
			for _, addr := range addrs {
				dsts = append(dsts, &net.IPNet{IP: addr.IP, Mask: net.CIDRMask(32, 32)})
			}
			_, serviceIPNet, err := net.ParseCIDR(config.ServiceNetworkCIDR)
			if err != nil {
				return fmt.Errorf("failed to parse ServiceNetworkCIDR: %v", err)
			}
			dsts = append(dsts, serviceIPNet)
			for _, dst := range dsts {
				route := &netlink.Route{Dst: dst, Gw: defaultGW}
				if err := netlink.RouteAdd(route); err != nil && !os.IsExist(err) {
					return fmt.Errorf("failed to add route to dst: %v via SDN: %v", dst, err)
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	convertedResult, err := convertToRequestedVersion(req.Config, result)
	if err != nil {
		return err
	}
	return convertedResult.Print()
}
func convertToRequestedVersion(stdinData []byte, result *current.Result) (types.Result, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	versionDecoder := &version.ConfigDecoder{}
	confVersion, err := versionDecoder.Decode(stdinData)
	if err != nil {
		return nil, err
	}
	newResult, err := result.GetAsVersion(confVersion)
	if err != nil {
		return nil, err
	}
	return newResult, nil
}
func (p *cniPlugin) CmdDel(args *skel.CmdArgs) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := p.doCNI("http://dummy/", newCNIRequest(args))
	return err
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rand.Seed(time.Now().UTC().UnixNano())
	hostNS, err := ns.GetCurrentNS()
	if err != nil {
		panic(fmt.Sprintf("could not get current kernel netns: %v", err))
	}
	defer hostNS.Close()
	p := NewCNIPlugin(cniserver.CNIServerSocketPath, hostNS)
	skel.PluginMain(p.CmdAdd, p.CmdDel, version.All)
}
