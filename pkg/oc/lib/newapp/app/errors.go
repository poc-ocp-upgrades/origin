package app

import (
	"bytes"
	"fmt"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

type ErrNoMatch struct {
	Value		string
	Type		string
	Qualifier	string
	Errs		[]error
}

func (e ErrNoMatch) Error() string {
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
	if len(e.Qualifier) != 0 {
		return fmt.Sprintf("unable to locate any %s with name %q: %s", e.Type, e.Value, e.Qualifier)
	}
	return fmt.Sprintf("unable to locate any %s with name %q", e.Type, e.Value)
}
func (e ErrNoMatch) Suggestion(commandName string) string {
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
	return fmt.Sprintf("%[3]s - does a Docker image with that name exist?", e.Value, commandName, e.Error())
}

type ErrPartialMatch struct {
	Value	string
	Match	*ComponentMatch
	Errs	[]error
}

func (e ErrPartialMatch) Error() string {
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
	return fmt.Sprintf("only a partial match was found for %q: %q", e.Value, e.Match.Name)
}
func (e ErrPartialMatch) Suggestion(commandName string) string {
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
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "* %s\n", e.Match.Description)
	fmt.Fprintf(buf, "  Use %[1]s to specify this image or template\n\n", e.Match.Argument)
	return fmt.Sprintf(`%[3]s
The argument %[1]q only partially matched the following Docker image or OpenShift image stream:

%[2]s
`, e.Value, buf.String(), cmdutil.MultipleErrors("error: ", e.Errs))
}

type ErrNoTagsFound struct {
	Value	string
	Match	*ComponentMatch
	Errs	[]error
}

func (e ErrNoTagsFound) Error() string {
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
	imageStream := fmt.Sprintf("%s/%s", e.Match.ImageStream.Namespace, e.Match.ImageStream.Name)
	return fmt.Sprintf("no tags found on matching image stream: %q", imageStream)
}
func (e ErrNoTagsFound) Suggestion(commandName string) string {
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
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "* %s\n", e.Match.Description)
	fmt.Fprintf(buf, "  Use --allow-missing-imagestreamtags to use this image stream\n\n")
	return fmt.Sprintf(`%[3]s
The argument %[1]q matched the following OpenShift image stream which has no tags:

%[2]s
`, e.Value, buf.String(), cmdutil.MultipleErrors("error: ", e.Errs))
}

type ErrMultipleMatches struct {
	Value	string
	Matches	[]*ComponentMatch
	Errs	[]error
}

func (e ErrMultipleMatches) Error() string {
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
	return fmt.Sprintf("multiple images or templates matched %q", e.Value)
}

var ErrNameRequired = fmt.Errorf("you must specify a name for your app")

type CircularOutputReferenceError struct{ Reference string }

func (e CircularOutputReferenceError) Error() string {
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
	return fmt.Sprintf("output image of %q should be different than input", e.Reference)
}

type CircularReferenceError struct{ Reference string }

func (e CircularReferenceError) Error() string {
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
	return fmt.Sprintf("image stream tag reference %q is a circular loop of image stream tags", e.Reference)
}
