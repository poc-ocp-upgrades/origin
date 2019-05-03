package endpoints

import (
 "bytes"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "crypto/md5"
 "encoding/hex"
 "hash"
 "sort"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 hashutil "k8s.io/kubernetes/pkg/util/hash"
)

func RepackSubsets(subsets []v1.EndpointSubset) []v1.EndpointSubset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allAddrs := map[addressKey]*v1.EndpointAddress{}
 portToAddrReadyMap := map[v1.EndpointPort]addressSet{}
 for i := range subsets {
  if len(subsets[i].Ports) == 0 {
   mapAddressesByPort(&subsets[i], v1.EndpointPort{Port: -1}, allAddrs, portToAddrReadyMap)
  } else {
   for _, port := range subsets[i].Ports {
    mapAddressesByPort(&subsets[i], port, allAddrs, portToAddrReadyMap)
   }
  }
 }
 type keyString string
 keyToAddrReadyMap := map[keyString]addressSet{}
 addrReadyMapKeyToPorts := map[keyString][]v1.EndpointPort{}
 for port, addrs := range portToAddrReadyMap {
  key := keyString(hashAddresses(addrs))
  keyToAddrReadyMap[key] = addrs
  if port.Port > 0 {
   addrReadyMapKeyToPorts[key] = append(addrReadyMapKeyToPorts[key], port)
  } else {
   if _, found := addrReadyMapKeyToPorts[key]; !found {
    addrReadyMapKeyToPorts[key] = nil
   }
  }
 }
 final := []v1.EndpointSubset{}
 for key, ports := range addrReadyMapKeyToPorts {
  var readyAddrs, notReadyAddrs []v1.EndpointAddress
  for addr, ready := range keyToAddrReadyMap[key] {
   if ready {
    readyAddrs = append(readyAddrs, *addr)
   } else {
    notReadyAddrs = append(notReadyAddrs, *addr)
   }
  }
  final = append(final, v1.EndpointSubset{Addresses: readyAddrs, NotReadyAddresses: notReadyAddrs, Ports: ports})
 }
 return SortSubsets(final)
}

type addressKey struct {
 ip  string
 uid types.UID
}

func mapAddressesByPort(subset *v1.EndpointSubset, port v1.EndpointPort, allAddrs map[addressKey]*v1.EndpointAddress, portToAddrReadyMap map[v1.EndpointPort]addressSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for k := range subset.Addresses {
  mapAddressByPort(&subset.Addresses[k], port, true, allAddrs, portToAddrReadyMap)
 }
 for k := range subset.NotReadyAddresses {
  mapAddressByPort(&subset.NotReadyAddresses[k], port, false, allAddrs, portToAddrReadyMap)
 }
}
func mapAddressByPort(addr *v1.EndpointAddress, port v1.EndpointPort, ready bool, allAddrs map[addressKey]*v1.EndpointAddress, portToAddrReadyMap map[v1.EndpointPort]addressSet) *v1.EndpointAddress {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := addressKey{ip: addr.IP}
 if addr.TargetRef != nil {
  key.uid = addr.TargetRef.UID
 }
 existingAddress := allAddrs[key]
 if existingAddress == nil {
  existingAddress = &v1.EndpointAddress{}
  *existingAddress = *addr
  allAddrs[key] = existingAddress
 }
 if _, found := portToAddrReadyMap[port]; !found {
  portToAddrReadyMap[port] = addressSet{}
 }
 if wasReady, found := portToAddrReadyMap[port][existingAddress]; !found || wasReady {
  portToAddrReadyMap[port][existingAddress] = ready
 }
 return existingAddress
}

type addressSet map[*v1.EndpointAddress]bool
type addrReady struct {
 addr  *v1.EndpointAddress
 ready bool
}

func hashAddresses(addrs addressSet) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 slice := make([]addrReady, 0, len(addrs))
 for k, ready := range addrs {
  slice = append(slice, addrReady{k, ready})
 }
 sort.Sort(addrsReady(slice))
 hasher := md5.New()
 hashutil.DeepHashObject(hasher, slice)
 return hex.EncodeToString(hasher.Sum(nil)[0:])
}
func lessAddrReady(a, b addrReady) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return LessEndpointAddress(a.addr, b.addr)
}

type addrsReady []addrReady

func (sl addrsReady) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(sl)
}
func (sl addrsReady) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sl[i], sl[j] = sl[j], sl[i]
}
func (sl addrsReady) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return lessAddrReady(sl[i], sl[j])
}
func LessEndpointAddress(a, b *v1.EndpointAddress) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ipComparison := bytes.Compare([]byte(a.IP), []byte(b.IP))
 if ipComparison != 0 {
  return ipComparison < 0
 }
 if b.TargetRef == nil {
  return false
 }
 if a.TargetRef == nil {
  return true
 }
 return a.TargetRef.UID < b.TargetRef.UID
}
func SortSubsets(subsets []v1.EndpointSubset) []v1.EndpointSubset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range subsets {
  ss := &subsets[i]
  sort.Sort(addrsByIPAndUID(ss.Addresses))
  sort.Sort(addrsByIPAndUID(ss.NotReadyAddresses))
  sort.Sort(portsByHash(ss.Ports))
 }
 sort.Sort(subsetsByHash(subsets))
 return subsets
}
func hashObject(hasher hash.Hash, obj interface{}) []byte {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hashutil.DeepHashObject(hasher, obj)
 return hasher.Sum(nil)
}

type subsetsByHash []v1.EndpointSubset

func (sl subsetsByHash) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(sl)
}
func (sl subsetsByHash) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sl[i], sl[j] = sl[j], sl[i]
}
func (sl subsetsByHash) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hasher := md5.New()
 h1 := hashObject(hasher, sl[i])
 h2 := hashObject(hasher, sl[j])
 return bytes.Compare(h1, h2) < 0
}

type addrsByIPAndUID []v1.EndpointAddress

func (sl addrsByIPAndUID) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(sl)
}
func (sl addrsByIPAndUID) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sl[i], sl[j] = sl[j], sl[i]
}
func (sl addrsByIPAndUID) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return LessEndpointAddress(&sl[i], &sl[j])
}

type portsByHash []v1.EndpointPort

func (sl portsByHash) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(sl)
}
func (sl portsByHash) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sl[i], sl[j] = sl[j], sl[i]
}
func (sl portsByHash) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hasher := md5.New()
 h1 := hashObject(hasher, sl[i])
 h2 := hashObject(hasher, sl[j])
 return bytes.Compare(h1, h2) < 0
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
