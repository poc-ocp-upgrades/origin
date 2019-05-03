package aws

import (
 "fmt"
 "sort"
 "sync"
)

type ExistingDevices map[mountDevice]EBSVolumeID
type DeviceAllocator interface {
 GetNext(existingDevices ExistingDevices) (mountDevice, error)
 Deprioritize(mountDevice)
 Lock()
 Unlock()
}
type deviceAllocator struct {
 possibleDevices map[mountDevice]int
 counter         int
 deviceLock      sync.Mutex
}

var _ DeviceAllocator = &deviceAllocator{}

type devicePair struct {
 deviceName  mountDevice
 deviceIndex int
}
type devicePairList []devicePair

func (p devicePairList) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(p)
}
func (p devicePairList) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return p[i].deviceIndex < p[j].deviceIndex
}
func (p devicePairList) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 p[i], p[j] = p[j], p[i]
}
func NewDeviceAllocator() DeviceAllocator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 possibleDevices := make(map[mountDevice]int)
 for _, firstChar := range []rune{'b', 'c'} {
  for i := 'a'; i <= 'z'; i++ {
   dev := mountDevice([]rune{firstChar, i})
   possibleDevices[dev] = 0
  }
 }
 return &deviceAllocator{possibleDevices: possibleDevices, counter: 0}
}
func (d *deviceAllocator) GetNext(existingDevices ExistingDevices) (mountDevice, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, devicePair := range d.sortByCount() {
  if _, found := existingDevices[devicePair.deviceName]; !found {
   return devicePair.deviceName, nil
  }
 }
 return "", fmt.Errorf("no devices are available")
}
func (d *deviceAllocator) sortByCount() devicePairList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dpl := make(devicePairList, 0)
 for deviceName, deviceIndex := range d.possibleDevices {
  dpl = append(dpl, devicePair{deviceName, deviceIndex})
 }
 sort.Sort(dpl)
 return dpl
}
func (d *deviceAllocator) Lock() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 d.deviceLock.Lock()
}
func (d *deviceAllocator) Unlock() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 d.deviceLock.Unlock()
}
func (d *deviceAllocator) Deprioritize(chosen mountDevice) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 d.deviceLock.Lock()
 defer d.deviceLock.Unlock()
 if _, ok := d.possibleDevices[chosen]; ok {
  d.counter++
  d.possibleDevices[chosen] = d.counter
 }
}
