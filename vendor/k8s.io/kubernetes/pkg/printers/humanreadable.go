package printers

import (
	"bytes"
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strings"
	"text/tabwriter"
	gotime "time"
)

type TablePrinter interface {
	PrintTable(obj runtime.Object, options PrintOptions) (*metav1beta1.Table, error)
}
type PrintHandler interface {
	Handler(columns, columnsWithWide []string, printFunc interface{}) error
	TableHandler(columns []metav1beta1.TableColumnDefinition, printFunc interface{}) error
	DefaultTableHandler(columns []metav1beta1.TableColumnDefinition, printFunc interface{}) error
}

var withNamespacePrefixColumns = []string{"NAMESPACE"}

type handlerEntry struct {
	columnDefinitions []metav1beta1.TableColumnDefinition
	printRows         bool
	printFunc         reflect.Value
	args              []reflect.Value
}
type HumanReadablePrinter struct {
	handlerMap     map[reflect.Type]*handlerEntry
	defaultHandler *handlerEntry
	options        PrintOptions
	lastType       interface{}
	skipTabWriter  bool
	encoder        runtime.Encoder
	decoder        runtime.Decoder
}

var _ PrintHandler = &HumanReadablePrinter{}

func NewHumanReadablePrinter(decoder runtime.Decoder, options PrintOptions) *HumanReadablePrinter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	printer := &HumanReadablePrinter{handlerMap: make(map[reflect.Type]*handlerEntry), options: options, decoder: decoder}
	return printer
}
func NewTablePrinter() *HumanReadablePrinter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &HumanReadablePrinter{handlerMap: make(map[reflect.Type]*handlerEntry)}
}
func (a *HumanReadablePrinter) AddTabWriter(t bool) *HumanReadablePrinter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.skipTabWriter = !t
	return a
}
func (a *HumanReadablePrinter) With(fns ...func(PrintHandler)) *HumanReadablePrinter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, fn := range fns {
		fn(a)
	}
	return a
}
func (h *HumanReadablePrinter) EnsurePrintHeaders() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h.options.NoHeaders = false
	h.lastType = nil
}
func (h *HumanReadablePrinter) Handler(columns, columnsWithWide []string, printFunc interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var columnDefinitions []metav1beta1.TableColumnDefinition
	for i, column := range columns {
		format := ""
		if i == 0 && strings.EqualFold(column, "name") {
			format = "name"
		}
		columnDefinitions = append(columnDefinitions, metav1beta1.TableColumnDefinition{Name: column, Description: column, Type: "string", Format: format})
	}
	for _, column := range columnsWithWide {
		columnDefinitions = append(columnDefinitions, metav1beta1.TableColumnDefinition{Name: column, Description: column, Type: "string", Priority: 1})
	}
	printFuncValue := reflect.ValueOf(printFunc)
	if err := ValidatePrintHandlerFunc(printFuncValue); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to register print function: %v", err))
		return err
	}
	entry := &handlerEntry{columnDefinitions: columnDefinitions, printFunc: printFuncValue}
	objType := printFuncValue.Type().In(0)
	if _, ok := h.handlerMap[objType]; ok {
		err := fmt.Errorf("registered duplicate printer for %v", objType)
		utilruntime.HandleError(err)
		return err
	}
	h.handlerMap[objType] = entry
	return nil
}
func (h *HumanReadablePrinter) TableHandler(columnDefinitions []metav1beta1.TableColumnDefinition, printFunc interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	printFuncValue := reflect.ValueOf(printFunc)
	if err := ValidateRowPrintHandlerFunc(printFuncValue); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to register print function: %v", err))
		return err
	}
	entry := &handlerEntry{columnDefinitions: columnDefinitions, printRows: true, printFunc: printFuncValue}
	objType := printFuncValue.Type().In(0)
	if _, ok := h.handlerMap[objType]; ok {
		err := fmt.Errorf("registered duplicate printer for %v", objType)
		utilruntime.HandleError(err)
		return err
	}
	h.handlerMap[objType] = entry
	return nil
}
func (h *HumanReadablePrinter) DefaultTableHandler(columnDefinitions []metav1beta1.TableColumnDefinition, printFunc interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	printFuncValue := reflect.ValueOf(printFunc)
	if err := ValidateRowPrintHandlerFunc(printFuncValue); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to register print function: %v", err))
		return err
	}
	entry := &handlerEntry{columnDefinitions: columnDefinitions, printRows: true, printFunc: printFuncValue}
	h.defaultHandler = entry
	return nil
}
func ValidateRowPrintHandlerFunc(printFunc reflect.Value) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if printFunc.Kind() != reflect.Func {
		return fmt.Errorf("invalid print handler. %#v is not a function", printFunc)
	}
	funcType := printFunc.Type()
	if funcType.NumIn() != 2 || funcType.NumOut() != 2 {
		return fmt.Errorf("invalid print handler." + "Must accept 2 parameters and return 2 value.")
	}
	if funcType.In(1) != reflect.TypeOf((*PrintOptions)(nil)).Elem() || funcType.Out(0) != reflect.TypeOf((*[]metav1beta1.TableRow)(nil)).Elem() || funcType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf("invalid print handler. The expected signature is: "+"func handler(obj %v, options PrintOptions) ([]metav1beta1.TableRow, error)", funcType.In(0))
	}
	return nil
}
func ValidatePrintHandlerFunc(printFunc reflect.Value) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if printFunc.Kind() != reflect.Func {
		return fmt.Errorf("invalid print handler. %#v is not a function", printFunc)
	}
	funcType := printFunc.Type()
	if funcType.NumIn() != 3 || funcType.NumOut() != 1 {
		return fmt.Errorf("invalid print handler." + "Must accept 3 parameters and return 1 value.")
	}
	if funcType.In(1) != reflect.TypeOf((*io.Writer)(nil)).Elem() || funcType.In(2) != reflect.TypeOf((*PrintOptions)(nil)).Elem() || funcType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf("invalid print handler. The expected signature is: "+"func handler(obj %v, w io.Writer, options PrintOptions) error", funcType.In(0))
	}
	return nil
}
func (h *HumanReadablePrinter) HandledResources() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	keys := make([]string, 0)
	for k := range h.handlerMap {
		api := strings.Split(k.String(), ".")
		resource := api[len(api)-1]
		if strings.HasSuffix(resource, "List") {
			continue
		}
		resource = strings.ToLower(resource)
		keys = append(keys, resource)
	}
	return keys
}
func (h *HumanReadablePrinter) unknown(data []byte, w io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := fmt.Fprintf(w, "Unknown object: %s", string(data))
	return err
}
func printHeader(columnNames []string, w io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t")); err != nil {
		return err
	}
	return nil
}
func (h *HumanReadablePrinter) PrintObj(obj runtime.Object, output io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	w, found := output.(*tabwriter.Writer)
	if !found && !h.skipTabWriter {
		w = GetNewTabWriter(output)
		output = w
		defer w.Flush()
	}
	if table, ok := obj.(*metav1beta1.Table); ok {
		if err := DecorateTable(table, h.options); err != nil {
			return err
		}
		return PrintTable(table, output, h.options)
	}
	if h.decoder != nil {
		obj, _ = decodeUnknownObject(obj, h.decoder)
	}
	t := reflect.TypeOf(obj)
	if handler := h.handlerMap[t]; handler != nil {
		includeHeaders := h.lastType != t && !h.options.NoHeaders
		if h.lastType != nil && h.lastType != t && !h.options.NoHeaders {
			fmt.Fprintln(output)
		}
		if err := printRowsForHandlerEntry(output, handler, obj, h.options, includeHeaders); err != nil {
			return err
		}
		h.lastType = t
		return nil
	}
	if h.defaultHandler != nil {
		includeHeaders := h.lastType != h.defaultHandler && !h.options.NoHeaders
		if h.lastType != nil && h.lastType != h.defaultHandler && !h.options.NoHeaders {
			fmt.Fprintln(output)
		}
		if err := printRowsForHandlerEntry(output, h.defaultHandler, obj, h.options, includeHeaders); err != nil {
			return err
		}
		h.lastType = h.defaultHandler
		return nil
	}
	return fmt.Errorf("error: unknown type %#v", obj)
}
func hasCondition(conditions []metav1beta1.TableRowCondition, t metav1beta1.RowConditionType) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, condition := range conditions {
		if condition.Type == t {
			return condition.Status == metav1beta1.ConditionTrue
		}
	}
	return false
}
func PrintTable(table *metav1beta1.Table, output io.Writer, options PrintOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !options.NoHeaders {
		if len(table.Rows) == 0 {
			return nil
		}
		first := true
		for _, column := range table.ColumnDefinitions {
			if !options.Wide && column.Priority != 0 {
				continue
			}
			if first {
				first = false
			} else {
				fmt.Fprint(output, "\t")
			}
			fmt.Fprint(output, strings.ToUpper(column.Name))
		}
		fmt.Fprintln(output)
	}
	for _, row := range table.Rows {
		first := true
		for i, cell := range row.Cells {
			if i >= len(table.ColumnDefinitions) {
				break
			}
			column := table.ColumnDefinitions[i]
			if !options.Wide && column.Priority != 0 {
				continue
			}
			if first {
				first = false
			} else {
				fmt.Fprint(output, "\t")
			}
			if cell != nil {
				fmt.Fprint(output, cell)
			}
		}
		fmt.Fprintln(output)
	}
	return nil
}
func DecorateTable(table *metav1beta1.Table, options PrintOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	width := len(table.ColumnDefinitions) + len(options.ColumnLabels)
	if options.WithNamespace {
		width++
	}
	if options.ShowLabels {
		width++
	}
	columns := table.ColumnDefinitions
	nameColumn := -1
	if options.WithKind && !options.Kind.Empty() {
		for i := range columns {
			if columns[i].Format == "name" && columns[i].Type == "string" {
				nameColumn = i
				break
			}
		}
	}
	if width != len(table.ColumnDefinitions) {
		columns = make([]metav1beta1.TableColumnDefinition, 0, width)
		if options.WithNamespace {
			columns = append(columns, metav1beta1.TableColumnDefinition{Name: "Namespace", Type: "string"})
		}
		columns = append(columns, table.ColumnDefinitions...)
		for _, label := range formatLabelHeaders(options.ColumnLabels) {
			columns = append(columns, metav1beta1.TableColumnDefinition{Name: label, Type: "string"})
		}
		if options.ShowLabels {
			columns = append(columns, metav1beta1.TableColumnDefinition{Name: "Labels", Type: "string"})
		}
	}
	rows := table.Rows
	includeLabels := len(options.ColumnLabels) > 0 || options.ShowLabels
	if includeLabels || options.WithNamespace || nameColumn != -1 {
		for i := range rows {
			row := rows[i]
			if nameColumn != -1 {
				row.Cells[nameColumn] = fmt.Sprintf("%s/%s", strings.ToLower(options.Kind.String()), row.Cells[nameColumn])
			}
			var m metav1.Object
			if obj := row.Object.Object; obj != nil {
				if acc, err := meta.Accessor(obj); err == nil {
					m = acc
				}
			}
			if m == nil {
				if options.WithNamespace {
					r := make([]interface{}, 1, width)
					row.Cells = append(r, row.Cells...)
				}
				for j := 0; j < width-len(row.Cells); j++ {
					row.Cells = append(row.Cells, nil)
				}
				rows[i] = row
				continue
			}
			if options.WithNamespace {
				r := make([]interface{}, 1, width)
				r[0] = m.GetNamespace()
				row.Cells = append(r, row.Cells...)
			}
			if includeLabels {
				row.Cells = appendLabelCells(row.Cells, m.GetLabels(), options)
			}
			rows[i] = row
		}
	}
	table.ColumnDefinitions = columns
	table.Rows = rows
	return nil
}
func (h *HumanReadablePrinter) PrintTable(obj runtime.Object, options PrintOptions) (*metav1beta1.Table, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t := reflect.TypeOf(obj)
	handler, ok := h.handlerMap[t]
	if !ok {
		return nil, fmt.Errorf("no table handler registered for this type %v", t)
	}
	if !handler.printRows {
		return h.legacyPrinterToTable(obj, handler)
	}
	args := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(options)}
	results := handler.printFunc.Call(args)
	if !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}
	columns := handler.columnDefinitions
	if !options.Wide {
		columns = make([]metav1beta1.TableColumnDefinition, 0, len(handler.columnDefinitions))
		for i := range handler.columnDefinitions {
			if handler.columnDefinitions[i].Priority != 0 {
				continue
			}
			columns = append(columns, handler.columnDefinitions[i])
		}
	}
	table := &metav1beta1.Table{ListMeta: metav1.ListMeta{ResourceVersion: ""}, ColumnDefinitions: columns, Rows: results[0].Interface().([]metav1beta1.TableRow)}
	if m, err := meta.ListAccessor(obj); err == nil {
		table.ResourceVersion = m.GetResourceVersion()
		table.SelfLink = m.GetSelfLink()
		table.Continue = m.GetContinue()
	} else {
		if m, err := meta.CommonAccessor(obj); err == nil {
			table.ResourceVersion = m.GetResourceVersion()
			table.SelfLink = m.GetSelfLink()
		}
	}
	if err := DecorateTable(table, options); err != nil {
		return nil, err
	}
	return table, nil
}
func printRowsForHandlerEntry(output io.Writer, handler *handlerEntry, obj runtime.Object, options PrintOptions, includeHeaders bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var results []reflect.Value
	if handler.printRows {
		args := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(options)}
		results = handler.printFunc.Call(args)
		if !results[1].IsNil() {
			return results[1].Interface().(error)
		}
	}
	if includeHeaders {
		var headers []string
		for _, column := range handler.columnDefinitions {
			if column.Priority != 0 && !options.Wide {
				continue
			}
			headers = append(headers, strings.ToUpper(column.Name))
		}
		headers = append(headers, formatLabelHeaders(options.ColumnLabels)...)
		headers = append(headers, formatShowLabelsHeader(options.ShowLabels)...)
		if options.WithNamespace {
			headers = append(withNamespacePrefixColumns, headers...)
		}
		printHeader(headers, output)
	}
	if !handler.printRows {
		args := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(output), reflect.ValueOf(options)}
		resultValue := handler.printFunc.Call(args)[0]
		if resultValue.IsNil() {
			return nil
		}
		return resultValue.Interface().(error)
	}
	if results[1].IsNil() {
		rows := results[0].Interface().([]metav1beta1.TableRow)
		printRows(output, rows, options)
		return nil
	}
	return results[1].Interface().(error)
}
func printRows(output io.Writer, rows []metav1beta1.TableRow, options PrintOptions) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, row := range rows {
		if options.WithNamespace {
			if obj := row.Object.Object; obj != nil {
				if m, err := meta.Accessor(obj); err == nil {
					fmt.Fprint(output, m.GetNamespace())
				}
			}
			fmt.Fprint(output, "\t")
		}
		for i, cell := range row.Cells {
			if i != 0 {
				fmt.Fprint(output, "\t")
			} else {
				if options.WithKind && !options.Kind.Empty() {
					fmt.Fprintf(output, "%s/%s", strings.ToLower(options.Kind.String()), cell)
					continue
				}
			}
			fmt.Fprint(output, cell)
		}
		hasLabels := len(options.ColumnLabels) > 0
		if obj := row.Object.Object; obj != nil && (hasLabels || options.ShowLabels) {
			if m, err := meta.Accessor(obj); err == nil {
				for _, value := range labelValues(m.GetLabels(), options) {
					output.Write([]byte("\t"))
					output.Write([]byte(value))
				}
			}
		}
		output.Write([]byte("\n"))
	}
}
func (h *HumanReadablePrinter) legacyPrinterToTable(obj runtime.Object, handler *handlerEntry) (*metav1beta1.Table, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	printFunc := handler.printFunc
	table := &metav1beta1.Table{ColumnDefinitions: handler.columnDefinitions}
	options := PrintOptions{NoHeaders: true, Wide: true}
	buf := &bytes.Buffer{}
	args := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(buf), reflect.ValueOf(options)}
	if meta.IsListType(obj) {
		listInterface, ok := obj.(metav1.ListInterface)
		if ok {
			table.ListMeta.SelfLink = listInterface.GetSelfLink()
			table.ListMeta.ResourceVersion = listInterface.GetResourceVersion()
			table.ListMeta.Continue = listInterface.GetContinue()
		}
		args[0] = reflect.ValueOf(obj)
		resultValue := printFunc.Call(args)[0]
		if !resultValue.IsNil() {
			return nil, resultValue.Interface().(error)
		}
		data := buf.Bytes()
		i := 0
		items, err := meta.ExtractList(obj)
		if err != nil {
			return nil, err
		}
		for len(data) > 0 {
			cells, remainder := tabbedLineToCells(data, len(table.ColumnDefinitions))
			table.Rows = append(table.Rows, metav1beta1.TableRow{Cells: cells, Object: runtime.RawExtension{Object: items[i]}})
			data = remainder
			i++
		}
	} else {
		args[0] = reflect.ValueOf(obj)
		resultValue := printFunc.Call(args)[0]
		if !resultValue.IsNil() {
			return nil, resultValue.Interface().(error)
		}
		data := buf.Bytes()
		cells, _ := tabbedLineToCells(data, len(table.ColumnDefinitions))
		table.Rows = append(table.Rows, metav1beta1.TableRow{Cells: cells, Object: runtime.RawExtension{Object: obj}})
	}
	return table, nil
}
func printUnstructured(unstructured runtime.Unstructured, w io.Writer, additionalFields []string, options PrintOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	metadata, err := meta.Accessor(unstructured)
	if err != nil {
		return err
	}
	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", metadata.GetNamespace()); err != nil {
			return err
		}
	}
	content := unstructured.UnstructuredContent()
	kind := "<missing>"
	if objKind, ok := content["kind"]; ok {
		if str, ok := objKind.(string); ok {
			kind = str
		}
	}
	if objAPIVersion, ok := content["apiVersion"]; ok {
		if str, ok := objAPIVersion.(string); ok {
			version, err := schema.ParseGroupVersion(str)
			if err != nil {
				return err
			}
			kind = kind + "." + version.Version + "." + version.Group
		}
	}
	name := FormatResourceName(options.Kind, metadata.GetName(), options.WithKind)
	if _, err := fmt.Fprintf(w, "%s\t%s", name, kind); err != nil {
		return err
	}
	for _, field := range additionalFields {
		if value, ok := content[field]; ok {
			var formattedValue string
			switch typedValue := value.(type) {
			case []interface{}:
				formattedValue = fmt.Sprintf("%d item(s)", len(typedValue))
			default:
				formattedValue = fmt.Sprintf("%v", value)
			}
			if _, err := fmt.Fprintf(w, "\t%s", formattedValue); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprint(w, AppendLabels(metadata.GetLabels(), options.ColumnLabels)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, AppendAllLabels(options.ShowLabels, metadata.GetLabels())); err != nil {
		return err
	}
	return nil
}
func formatLabelHeaders(columnLabels []string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	formHead := make([]string, len(columnLabels))
	for i, l := range columnLabels {
		p := strings.Split(l, "/")
		formHead[i] = strings.ToUpper((p[len(p)-1]))
	}
	return formHead
}
func formatShowLabelsHeader(showLabels bool) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if showLabels {
		return []string{"LABELS"}
	}
	return nil
}
func labelValues(itemLabels map[string]string, opts PrintOptions) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var values []string
	for _, key := range opts.ColumnLabels {
		values = append(values, itemLabels[key])
	}
	if opts.ShowLabels {
		values = append(values, labels.FormatLabels(itemLabels))
	}
	return values
}
func appendLabelCells(values []interface{}, itemLabels map[string]string, opts PrintOptions) []interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, key := range opts.ColumnLabels {
		values = append(values, itemLabels[key])
	}
	if opts.ShowLabels {
		values = append(values, labels.FormatLabels(itemLabels))
	}
	return values
}
func FormatResourceName(kind schema.GroupKind, name string, withKind bool) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !withKind || kind.Empty() {
		return name
	}
	return strings.ToLower(kind.String()) + "/" + name
}
func AppendLabels(itemLabels map[string]string, columnLabels []string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var buffer bytes.Buffer
	for _, cl := range columnLabels {
		buffer.WriteString(fmt.Sprint("\t"))
		if il, ok := itemLabels[cl]; ok {
			buffer.WriteString(fmt.Sprint(il))
		} else {
			buffer.WriteString("<none>")
		}
	}
	return buffer.String()
}
func AppendAllLabels(showLabels bool, itemLabels map[string]string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var buffer bytes.Buffer
	if showLabels {
		buffer.WriteString(fmt.Sprint("\t"))
		buffer.WriteString(labels.FormatLabels(itemLabels))
	}
	buffer.WriteString("\n")
	return buffer.String()
}
func decodeUnknownObject(obj runtime.Object, decoder runtime.Decoder) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	switch t := obj.(type) {
	case runtime.Unstructured:
		if objBytes, err := runtime.Encode(unstructured.UnstructuredJSONScheme, obj); err == nil {
			if decodedObj, err := runtime.Decode(decoder, objBytes); err == nil {
				obj = decodedObj
			}
		}
	case *runtime.Unknown:
		if decodedObj, err := runtime.Decode(decoder, t.Raw); err == nil {
			obj = decodedObj
		}
	}
	return obj, err
}
func tabbedLineToCells(data []byte, expected int) ([]interface{}, []byte) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var remainder []byte
	max := bytes.Index(data, []byte("\n"))
	if max != -1 {
		remainder = data[max+1:]
		data = data[:max]
	}
	cells := make([]interface{}, expected)
	for i := 0; i < expected; i++ {
		next := bytes.Index(data, []byte("\t"))
		if next == -1 {
			cells[i] = string(data)
			for j := i + 1; j < expected; j++ {
				cells[j] = ""
			}
			break
		}
		cells[i] = string(data[:next])
		data = data[next+1:]
	}
	return cells, remainder
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
