package dockerfile

import (
	"reflect"
	"strings"
	"testing"
	"github.com/moby/buildkit/frontend/dockerfile/command"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func TestFindAll(t *testing.T) {
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
	instructions := `FROM scratch
LABEL version=1.0
FROM busybox
ENV PATH=/bin
`
	node, err := parser.Parse(strings.NewReader(instructions))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for cmd, want := range map[string][]int{command.From: {0, 2}, command.Label: {1}, command.Env: {3}, command.Maintainer: nil, "UnknownCommand": nil} {
		got := FindAll(node.AST, cmd)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FindAll(node, %q) = %#v; want %#v", cmd, got, want)
		}
	}
}
func TestFindAllNilNode(t *testing.T) {
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
	cmd := command.From
	got := FindAll(nil, cmd)
	if got != nil {
		t.Errorf("FindAll(nil, %q) = %#v; want nil", cmd, got)
	}
}
func TestInsertInstructions(t *testing.T) {
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
	testCases := map[string]struct {
		original	string
		index		int
		newInstructions	string
		want		string
	}{"insert nothing": {original: `FROM busybox
ENV PATH=/bin
`, index: 0, newInstructions: ``, want: `FROM busybox
ENV PATH=/bin
`}, "insert instruction in empty file": {original: ``, index: 0, newInstructions: `FROM busybox`, want: `FROM busybox
`}, "prepend single instruction": {original: `FROM busybox
ENV PATH=/bin
`, index: 0, newInstructions: `FROM scratch`, want: `FROM scratch
FROM busybox
ENV PATH=/bin
`}, "append single instruction": {original: `FROM busybox
ENV PATH=/bin
`, index: 2, newInstructions: `FROM scratch`, want: `FROM busybox
ENV PATH=/bin
FROM scratch
`}, "insert single instruction in the middle": {original: `FROM busybox
ENV PATH=/bin
`, index: 1, newInstructions: `LABEL version=1.0`, want: `FROM busybox
LABEL version=1.0
ENV PATH=/bin
`}}
	for name, tc := range testCases {
		got, err := parser.Parse(strings.NewReader(tc.original))
		if err != nil {
			t.Errorf("InsertInstructions: %s: parse error: %v", name, err)
			continue
		}
		err = InsertInstructions(got.AST, tc.index, tc.newInstructions)
		if err != nil {
			t.Errorf("InsertInstructions: %s: %v", name, err)
			continue
		}
		want, err := parser.Parse(strings.NewReader(tc.want))
		if err != nil {
			t.Errorf("InsertInstructions: %s: parse error: %v", name, err)
			continue
		}
		if !reflect.DeepEqual(got.AST.Dump(), want.AST.Dump()) {
			t.Errorf("InsertInstructions: %s: got %s, want %s", name, got.AST.Dump(), want.AST.Dump())
		}
	}
}
func TestInsertInstructionsNilNode(t *testing.T) {
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
	err := InsertInstructions(nil, 0, "")
	if err == nil {
		t.Errorf("InsertInstructions: got nil; want error")
	}
}
func TestInsertInstructionsPosOutOfRange(t *testing.T) {
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
	original := `FROM busybox
ENV PATH=/bin
`
	node, err := parser.Parse(strings.NewReader(original))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, pos := range []int{-1, 3, 4} {
		err := InsertInstructions(node.AST, pos, "")
		if err == nil {
			t.Errorf("InsertInstructions(node, %d, \"\"): got nil; want error", pos)
		}
	}
}
func TestInsertInstructionsUnparseable(t *testing.T) {
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
	original := `FROM busybox
ENV PATH=/bin
`
	node, err := parser.Parse(strings.NewReader(original))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for name, instructions := range map[string]string{"env without value": `ENV PATH`, "nested json": `CMD [ "echo", [ "nested json" ] ]`} {
		err = InsertInstructions(node.AST, 1, instructions)
		if err == nil {
			t.Errorf("InsertInstructions: %s: got nil; want error", name)
		}
	}
}
func TestLastBaseImage(t *testing.T) {
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
	testCases := map[string]struct {
		in	string
		want	string
	}{"empty Dockerfile": {in: ``, want: ""}, "FROM missing argument": {in: `FROM`, want: ""}, "single FROM": {in: `FROM centos:7`, want: "centos:7"}, "multiple FROM": {in: `FROM scratch
COPY . /boot
FROM centos:7`, want: "centos:7"}}
	for name, tc := range testCases {
		node, err := parser.Parse(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%s: parse error: %v", name, err)
			continue
		}
		got := LastBaseImage(node.AST)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("LastBaseImage: %s: got %#v; want %#v", name, got, tc.want)
		}
	}
}
func TestLastBaseImageNilNode(t *testing.T) {
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
	want := ""
	if got := LastBaseImage(nil); got != want {
		t.Errorf("LastBaseImage(nil) = %#v; want %#v", got, want)
	}
}
func TestBaseImages(t *testing.T) {
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
	testCases := map[string]struct {
		in	string
		want	[]string
	}{"empty Dockerfile": {in: ``, want: nil}, "FROM missing argument": {in: `FROM`, want: nil}, "single FROM": {in: `FROM centos:7`, want: []string{"centos:7"}}, "multiple FROM": {in: `FROM scratch
COPY . /boot
FROM centos:7`, want: []string{"scratch", "centos:7"}}}
	for name, tc := range testCases {
		node, err := parser.Parse(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%s: parse error: %v", name, err)
			continue
		}
		got := baseImages(node.AST)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("baseImages: %s: got %#v; want %#v", name, got, tc.want)
		}
	}
}
func TestBaseImagesNilNode(t *testing.T) {
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
	if got := baseImages(nil); got != nil {
		t.Errorf("baseImages(nil) = %#v; want nil", got)
	}
}
func TestLastExposedPorts(t *testing.T) {
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
	testCases := map[string]struct {
		in	string
		want	[]string
	}{"empty Dockerfile": {in: ``, want: nil}, "EXPOSE missing argument": {in: `EXPOSE`, want: nil}, "EXPOSE no FROM": {in: `EXPOSE 8080`, want: nil}, "single EXPOSE after FROM": {in: `FROM centos:7
		EXPOSE 8080`, want: []string{"8080"}}, "multiple EXPOSE and FROM": {in: `# EXPOSE before FROM should be ignore
EXPOSE 777
FROM busybox
EXPOSE 8080
COPY . /boot
FROM rhel
# no EXPOSE instruction
FROM centos:7
EXPOSE 8000
EXPOSE 9090 9091
`, want: []string{"8000", "9090", "9091"}}}
	for name, tc := range testCases {
		node, err := parser.Parse(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%s: parse error: %v", name, err)
			continue
		}
		got := LastExposedPorts(node.AST)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("LastExposedPorts: %s: got %#v; want %#v", name, got, tc.want)
		}
	}
}
func TestLastExposedPortsNilNode(t *testing.T) {
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
	if got := LastExposedPorts(nil); got != nil {
		t.Errorf("LastExposedPorts(nil) = %#v; want nil", got)
	}
}
func TestExposedPorts(t *testing.T) {
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
	testCases := map[string]struct {
		in	string
		want	[][]string
	}{"empty Dockerfile": {in: ``, want: nil}, "EXPOSE missing argument": {in: `EXPOSE`, want: nil}, "EXPOSE no FROM": {in: `EXPOSE 8080`, want: nil}, "single EXPOSE after FROM": {in: `FROM centos:7
		EXPOSE 8080`, want: [][]string{{"8080"}}}, "multiple EXPOSE and FROM": {in: `# EXPOSE before FROM should be ignore
EXPOSE 777
FROM busybox
EXPOSE 8080
COPY . /boot
FROM rhel
# no EXPOSE instruction
FROM centos:7
EXPOSE 8000
EXPOSE 9090 9091
`, want: [][]string{{"8080"}, nil, {"8000", "9090", "9091"}}}}
	for name, tc := range testCases {
		node, err := parser.Parse(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%s: parse error: %v", name, err)
			continue
		}
		got, _ := exposedPorts(node.AST)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("exposedPorts: %s: got %#v; want %#v", name, got, tc.want)
		}
	}
}
func TestLastExposedPortsVarAndRange(t *testing.T) {
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
	testCases := map[string]struct {
		in	string
		want	[]string
	}{"EXPOSE with ENV": {in: `FROM centos:7
ENV PORT 8080
EXPOSE $PORT`, want: []string{"8080"}}, "EXPOSE with RANGE": {in: `FROM centos:7
EXPOSE 8080-9000`, want: []string{"8080"}}, "EXPOSE with default ARG": {in: `FROM centos:7
ARG PORT=8080
EXPOSE $PORT`, want: []string{"8080"}}, "EXPOSE recursive ARGS, ENV and range": {in: `FROM centos:7
ARG PORT=8080-8085
ENV PORT2 "$PORT"
ENV PORT3="$PORT2"
ENV PORT4=$PORT3
ENV PORT5=$PORT4
EXPOSE $PORT5`, want: []string{"8080"}}, "EXPOSE ignore post eval": {in: `FROM centos:7
ENV PORT=8080
EXPOSE $PORT
ENV PORT=10`, want: []string{"8080"}}, "EXPOSE parameter substitution syntax": {in: `FROM centos:7
ENV PORT2 8081
ENV PORT=8080 PORT1="${PORT2}"
EXPOSE ${PORT} "${PORT2:-8085}" ${PORT_MISSING:-8082}`, want: []string{"8080", "8081", "8082"}}, "EXPOSE parameter substitution syntax with +": {in: `FROM centos:7
ENV PORT2 8081
ENV PORT=8080 PORT1="${PORT2}"
EXPOSE ${PORT} "${PORT2:+8085}" ${PORT_MISSING:+8082}`, want: []string{"8080", "8085", ""}}, "EXPOSE shadow scope ENV": {in: `FROM centos:7
ENV PORT 8080
FROM centos:7
ENV PORT 8081
EXPOSE $PORT`, want: []string{"8081"}}, "EXPOSE out of scope ENV": {in: `FROM centos:7
ENV PORT 8080
FROM centos:7
EXPOSE $PORT`, want: []string{""}}, "EXPOSE parse by line": {in: `FROM centos:7
ENV PORT 8080
ENV PORT=8090 PORT1=$PORT PORT=9080
EXPOSE $PORT1`, want: []string{"8080"}}, "EXPOSE partial ENV range": {in: `FROM centos:7
ENV PORT 8080
EXPOSE ${PORT}-8090`, want: []string{"8080"}}, "EXPOSE redefined ENV": {in: `FROM centos:7
ENV PORT 8080
ENV PORT 8081
EXPOSE $PORT`, want: []string{"8081"}}, "Multiple EXPOSE": {in: `FROM centos:7
ENV PORT 8080
EXPOSE $PORT
ENV PORT2 8081
EXPOSE $PORT2`, want: []string{"8080", "8081"}}, "Multiple EXPOSE and ENV redefined": {in: `FROM centos:7
ENV PORT 8080
EXPOSE $PORT
ENV PORT2 8081
ENV PORT 8082
EXPOSE $PORT2 $PORT
ENV PORT=8083 PORT2=8084
EXPOSE $PORT $PORT2`, want: []string{"8080", "8081", "8082", "8083", "8084"}}}
	for name, tc := range testCases {
		node, err := parser.Parse(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%s: parse error: %v", name, err)
			continue
		}
		got := LastExposedPorts(node.AST)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("exposedPorts: %s: got %#v; want %#v", name, got, tc.want)
		}
	}
}
func TestExposedPortsNilNode(t *testing.T) {
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
	if got, _ := exposedPorts(nil); got != nil {
		t.Errorf("exposedPorts(nil) = %#v; want nil", got)
	}
}
func TestNextValues(t *testing.T) {
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
	testCases := map[string][]string{`FROM busybox:latest`: {"busybox:latest"}, `MAINTAINER nobody@example.com`: {"nobody@example.com"}, `LABEL version=1.0`: {"version", "1.0"}, `EXPOSE 8080`: {"8080"}, `VOLUME /var/run/www`: {"/var/run/www"}, `ENV PATH=/bin`: {"PATH", "/bin"}, `ADD file /home/`: {"file", "/home/"}, `COPY dir/ /tmp/`: {"dir/", "/tmp/"}, `RUN echo "Hello world!"`: {`echo "Hello world!"`}, `ENTRYPOINT /bin/sh`: {"/bin/sh"}, `CMD ["-c", "env"]`: {"-c", "env"}, `USER 1001`: {"1001"}, `WORKDIR /home`: {"/home"}}
	for original, want := range testCases {
		node, err := parser.Parse(strings.NewReader(original))
		if err != nil {
			t.Fatalf("parse error: %s: %v", original, err)
		}
		if len(node.AST.Children) != 1 {
			t.Fatalf("unexpected number of children in test case: %s", original)
		}
		newNode := node.AST.Children[0]
		if got := nextValues(newNode); !reflect.DeepEqual(got, want) {
			t.Errorf("nextValues(%+v) = %#v; want %#v", newNode, got, want)
		}
	}
}
func TestNextValuesOnbuild(t *testing.T) {
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
	testCases := map[string][]string{`ONBUILD ADD . /app/src`: {".", "/app/src"}, `ONBUILD RUN echo "Hello universe!"`: {`echo "Hello universe!"`}}
	for original, want := range testCases {
		node, err := parser.Parse(strings.NewReader(original))
		if err != nil {
			t.Fatalf("parse error: %s: %v", original, err)
		}
		if len(node.AST.Children) != 1 {
			t.Fatalf("unexpected number of children in test case: %s", original)
		}
		nextNode := node.AST.Children[0].Next
		if nextNode == nil || len(nextNode.Children) != 1 {
			t.Fatalf("unexpected number of children in ONBUILD instruction of test case: %s", original)
		}
		nextNode = nextNode.Children[0]
		if got := nextValues(nextNode); !reflect.DeepEqual(got, want) {
			t.Errorf("nextValues(%+v) = %#v; want %#v", nextNode, got, want)
		}
	}
}
func TestNextValuesNilNode(t *testing.T) {
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
	if got := nextValues(nil); got != nil {
		t.Errorf("nextValues(nil) = %#v; want nil", got)
	}
}
