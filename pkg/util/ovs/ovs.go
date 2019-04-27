package ovs

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"k8s.io/klog"
	utilversion "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/utils/exec"
)

type Interface interface {
	AddBridge(properties ...string) error
	DeleteBridge(ifExists bool) error
	AddPort(port string, ofportRequest int, properties ...string) (int, error)
	DeletePort(port string) error
	GetOFPort(port string) (int, error)
	SetFrags(mode string) error
	Create(table string, values ...string) (string, error)
	Destroy(table, record string) error
	Get(table, record, column string) (string, error)
	Set(table, record string, values ...string) error
	Clear(table, record string, columns ...string) error
	Find(table string, column []string, condition string) ([]map[string]string, error)
	FindOne(table, column, condition string) ([]string, error)
	DumpFlows(flow string, args ...interface{}) ([]string, error)
	NewTransaction() Transaction
}
type Transaction interface {
	AddFlow(flow string, args ...interface{})
	DeleteFlows(flow string, args ...interface{})
	Commit() error
}

const (
	OVS_OFCTL	= "ovs-ofctl"
	OVS_VSCTL	= "ovs-vsctl"
)

type ovsExec struct {
	execer	exec.Interface
	bridge	string
}

func New(execer exec.Interface, bridge string, minVersion string) (Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := execer.LookPath(OVS_OFCTL); err != nil {
		return nil, fmt.Errorf("OVS is not installed")
	}
	if _, err := execer.LookPath(OVS_VSCTL); err != nil {
		return nil, fmt.Errorf("OVS is not installed")
	}
	ovsif := &ovsExec{execer: execer, bridge: bridge}
	if minVersion != "" {
		minVer := utilversion.MustParseGeneric(minVersion)
		out, err := ovsif.exec(OVS_VSCTL, "--version")
		if err != nil {
			return nil, fmt.Errorf("could not check OVS version is %s or higher", minVersion)
		}
		lines := strings.Split(out, "\n")
		spc := strings.LastIndex(lines[0], " ")
		instVer, err := utilversion.ParseGeneric(lines[0][spc+1:])
		if err != nil {
			return nil, fmt.Errorf("could not find OVS version in %q", lines[0])
		}
		if !instVer.AtLeast(minVer) {
			return nil, fmt.Errorf("found OVS %v, need %s or later", instVer, minVersion)
		}
	}
	return ovsif, nil
}
func (ovsif *ovsExec) execWithStdin(cmd string, stdinArgs []string, args ...string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logLevel := klog.Level(4)
	switch cmd {
	case OVS_OFCTL:
		if args[0] == "dump-flows" {
			logLevel = klog.Level(5)
		}
		args = append([]string{"-O", "OpenFlow13"}, args...)
	case OVS_VSCTL:
		args = append([]string{"--timeout=30"}, args...)
	}
	kcmd := ovsif.execer.Command(cmd, args...)
	if stdinArgs != nil {
		stdinString := strings.Join(stdinArgs, "\n")
		stdin := bytes.NewBufferString(stdinString)
		kcmd.SetStdin(stdin)
		klog.V(logLevel).Infof("Executing: %s %s <<\n%s", cmd, strings.Join(args, " "), stdinString)
	} else {
		klog.V(logLevel).Infof("Executing: %s %s", cmd, strings.Join(args, " "))
	}
	output, err := kcmd.CombinedOutput()
	if err != nil {
		klog.V(2).Infof("Error executing %s: %s", cmd, string(output))
		return "", err
	}
	outStr := string(output)
	if outStr != "" {
		nl := strings.Index(outStr, "\n")
		if nl == len(outStr)-1 {
			outStr = outStr[:nl]
		}
	}
	return outStr, nil
}
func (ovsif *ovsExec) exec(cmd string, args ...string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ovsif.execWithStdin(cmd, nil, args...)
}
func validateColumns(columns ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, col := range columns {
		end := strings.IndexAny(col, ":=")
		if end != -1 {
			col = col[:end]
		}
		if strings.Contains(col, "-") {
			return fmt.Errorf("bad ovsdb column name %q: should be %q", col, strings.Replace(col, "-", "_", -1))
		}
	}
	return nil
}
func (ovsif *ovsExec) AddBridge(properties ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := []string{"add-br", ovsif.bridge}
	if len(properties) > 0 {
		if err := validateColumns(properties...); err != nil {
			return err
		}
		args = append(args, "--", "set", "Bridge", ovsif.bridge)
		args = append(args, properties...)
	}
	_, err := ovsif.exec(OVS_VSCTL, args...)
	return err
}
func (ovsif *ovsExec) DeleteBridge(ifExists bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := []string{"del-br", ovsif.bridge}
	if ifExists {
		args = append([]string{"--if-exists"}, args...)
	}
	_, err := ovsif.exec(OVS_VSCTL, args...)
	return err
}
func (ovsif *ovsExec) GetOFPort(port string) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ofportStr, err := ovsif.exec(OVS_VSCTL, "get", "Interface", port, "ofport")
	if err != nil {
		return -1, fmt.Errorf("failed to get OVS port for %s: %v", port, err)
	}
	ofport, err := strconv.Atoi(ofportStr)
	if err != nil {
		return -1, fmt.Errorf("could not parse allocated ofport %q: %v", ofportStr, err)
	}
	if ofport == -1 {
		errStr, err := ovsif.exec(OVS_VSCTL, "get", "Interface", port, "error")
		if err != nil || errStr == "" {
			errStr = "unknown error"
		}
		return -1, fmt.Errorf("error on port %s: %s", port, errStr)
	}
	return ofport, nil
}
func (ovsif *ovsExec) AddPort(port string, ofportRequest int, properties ...string) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := []string{"--may-exist", "add-port", ovsif.bridge, port}
	if ofportRequest > 0 || len(properties) > 0 {
		args = append(args, "--", "set", "Interface", port)
		if ofportRequest > 0 {
			args = append(args, fmt.Sprintf("ofport_request=%d", ofportRequest))
		}
		if len(properties) > 0 {
			if err := validateColumns(properties...); err != nil {
				return -1, err
			}
			args = append(args, properties...)
		}
	}
	_, err := ovsif.exec(OVS_VSCTL, args...)
	if err != nil {
		return -1, err
	}
	ofport, err := ovsif.GetOFPort(port)
	if err != nil {
		return -1, err
	}
	if ofportRequest > 0 && ofportRequest != ofport {
		return -1, fmt.Errorf("allocated ofport (%d) did not match request (%d)", ofport, ofportRequest)
	}
	return ofport, nil
}
func (ovsif *ovsExec) DeletePort(port string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := ovsif.exec(OVS_VSCTL, "--if-exists", "del-port", ovsif.bridge, port)
	return err
}
func (ovsif *ovsExec) SetFrags(mode string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := ovsif.exec(OVS_OFCTL, "set-frags", ovsif.bridge, mode)
	return err
}
func (ovsif *ovsExec) Create(table string, values ...string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validateColumns(values...); err != nil {
		return "", err
	}
	args := append([]string{"create", table}, values...)
	return ovsif.exec(OVS_VSCTL, args...)
}
func (ovsif *ovsExec) Destroy(table, record string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := ovsif.exec(OVS_VSCTL, "--if-exists", "destroy", table, record)
	return err
}
func (ovsif *ovsExec) Get(table, record, column string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validateColumns(column); err != nil {
		return "", err
	}
	return ovsif.exec(OVS_VSCTL, "get", table, record, column)
}
func (ovsif *ovsExec) Set(table, record string, values ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validateColumns(values...); err != nil {
		return err
	}
	args := append([]string{"set", table, record}, values...)
	_, err := ovsif.exec(OVS_VSCTL, args...)
	return err
}
func (ovsif *ovsExec) Find(table string, columns []string, condition string) ([]map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validateColumns(columns...); err != nil {
		return nil, err
	}
	if err := validateColumns(condition); err != nil {
		return nil, err
	}
	output, err := ovsif.exec(OVS_VSCTL, "--columns="+strings.Join(columns, ","), "find", table, condition)
	if err != nil {
		return nil, err
	}
	output = strings.TrimSuffix(output, "\n")
	if output == "" {
		return nil, err
	}
	rows := strings.Split(output, "\n\n")
	result := make([]map[string]string, len(rows))
	for i, row := range rows {
		cols := make(map[string]string)
		for _, col := range strings.Split(row, "\n") {
			data := strings.SplitN(col, ":", 2)
			if len(data) != 2 {
				return nil, fmt.Errorf("bad 'ovs-vsctl find' line %q", col)
			}
			name := strings.TrimSpace(data[0])
			val := strings.TrimSpace(data[1])
			if unquoted, err := strconv.Unquote(val); err == nil {
				val = unquoted
			}
			cols[name] = val
		}
		result[i] = cols
	}
	return result, nil
}
func (ovsif *ovsExec) FindOne(table, column, condition string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fullResult, err := ovsif.Find(table, []string{column}, condition)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(fullResult))
	for _, row := range fullResult {
		result = append(result, row[column])
	}
	return result, nil
}
func (ovsif *ovsExec) Clear(table, record string, columns ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validateColumns(columns...); err != nil {
		return err
	}
	args := append([]string{"--if-exists", "clear", table, record}, columns...)
	_, err := ovsif.exec(OVS_VSCTL, args...)
	return err
}
func (ovsif *ovsExec) DumpFlows(flow string, args ...interface{}) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		flow = fmt.Sprintf(flow, args...)
	}
	out, err := ovsif.exec(OVS_OFCTL, "dump-flows", ovsif.bridge, flow)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(out, "\n")
	flows := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.Contains(line, "cookie=") {
			flows = append(flows, line)
		}
	}
	return flows, nil
}
func (ovsif *ovsExec) NewTransaction() Transaction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ovsExecTx{ovsif: ovsif, flows: []string{}}
}
func (ovsif *ovsExec) bundle(flows []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(flows) == 0 {
		return nil
	}
	_, err := ovsif.execWithStdin(OVS_OFCTL, flows, "bundle", ovsif.bridge, "-")
	return err
}

type ovsExecTx struct {
	ovsif	*ovsExec
	flows	[]string
}

func (tx *ovsExecTx) AddFlow(flow string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		flow = fmt.Sprintf(flow, args...)
	}
	tx.flows = append(tx.flows, fmt.Sprintf("flow add %s", flow))
}
func (tx *ovsExecTx) DeleteFlows(flow string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		flow = fmt.Sprintf(flow, args...)
	}
	tx.flows = append(tx.flows, fmt.Sprintf("flow delete %s", flow))
}
func (tx *ovsExecTx) Commit() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := tx.ovsif.bundle(tx.flows)
	tx.flows = []string{}
	return err
}
