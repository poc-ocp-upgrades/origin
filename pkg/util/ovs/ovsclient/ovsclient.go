package ovsclient

import (
	"fmt"
	goformat "fmt"
	"io"
	"io/ioutil"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type Client struct {
	*rpc.Client
	conn net.Conn
}

func New(conn net.Conn) *Client {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Client{Client: jsonrpc.NewClient(conn), conn: conn}
}
func DialTimeout(network, addr string, timeout time.Duration) (*Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return New(conn), nil
}
func (c *Client) Ping() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var result interface{}
	if err := c.Call("echo", []string{"hello"}, &result); err != nil {
		return err
	}
	return nil
}
func (c *Client) WaitForDisconnect() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n, err := io.Copy(ioutil.Discard, c.conn)
	if err != nil && err != io.EOF {
		return err
	}
	if n > 0 {
		return fmt.Errorf("unexpected bytes read waiting for disconnect: %d", n)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
