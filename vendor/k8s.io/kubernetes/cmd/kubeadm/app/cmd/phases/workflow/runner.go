package workflow

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

const phaseSeparator = "/"

type RunnerOptions struct {
	FilterPhases []string
	SkipPhases   []string
}
type RunData = interface{}
type Runner struct {
	Options            RunnerOptions
	Phases             []Phase
	runDataInitializer func(*cobra.Command) (RunData, error)
	runData            RunData
	runCmd             *cobra.Command
	cmdAdditionalFlags *pflag.FlagSet
	phaseRunners       []*phaseRunner
}
type phaseRunner struct {
	Phase
	parent        *phaseRunner
	level         int
	selfPath      []string
	generatedName string
	use           string
}

func NewRunner() *Runner {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Runner{Phases: []Phase{}}
}
func (e *Runner) AppendPhase(t Phase) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.Phases = append(e.Phases, t)
}
func (e *Runner) computePhaseRunFlags() (map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	phaseRunFlags := map[string]bool{}
	phaseHierarchy := map[string][]string{}
	e.visitAll(func(p *phaseRunner) error {
		phaseRunFlags[p.generatedName] = true
		phaseHierarchy[p.generatedName] = []string{}
		parent := p.parent
		for parent != nil {
			phaseHierarchy[parent.generatedName] = append(phaseHierarchy[parent.generatedName], p.generatedName)
			parent = parent.parent
		}
		return nil
	})
	if len(e.Options.FilterPhases) > 0 {
		for i := range phaseRunFlags {
			phaseRunFlags[i] = false
		}
		for _, f := range e.Options.FilterPhases {
			if _, ok := phaseRunFlags[f]; !ok {
				return phaseRunFlags, errors.Errorf("invalid phase name: %s", f)
			}
			phaseRunFlags[f] = true
			for _, c := range phaseHierarchy[f] {
				phaseRunFlags[c] = true
			}
		}
	}
	for _, f := range e.Options.SkipPhases {
		if _, ok := phaseRunFlags[f]; !ok {
			return phaseRunFlags, errors.Errorf("invalid phase name: %s", f)
		}
		phaseRunFlags[f] = false
		for _, c := range phaseHierarchy[f] {
			phaseRunFlags[c] = false
		}
	}
	return phaseRunFlags, nil
}
func (e *Runner) SetDataInitializer(builder func(cmd *cobra.Command) (RunData, error)) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.runDataInitializer = builder
}
func (e *Runner) InitData() (RunData, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if e.runData == nil && e.runDataInitializer != nil {
		var err error
		if e.runData, err = e.runDataInitializer(e.runCmd); err != nil {
			return nil, err
		}
	}
	return e.runData, nil
}
func (e *Runner) Run() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.prepareForExecution()
	phaseRunFlags, err := e.computePhaseRunFlags()
	if err != nil {
		return err
	}
	var data RunData
	if data, err = e.InitData(); err != nil {
		return err
	}
	err = e.visitAll(func(p *phaseRunner) error {
		if run, ok := phaseRunFlags[p.generatedName]; !run || !ok {
			return nil
		}
		if p.RunAllSiblings && (p.RunIf != nil || p.Run != nil) {
			return errors.Wrapf(err, "phase marked as RunAllSiblings can not have Run functions %s", p.generatedName)
		}
		if p.RunIf != nil {
			ok, err := p.RunIf(data)
			if err != nil {
				return errors.Wrapf(err, "error execution run condition for phase %s", p.generatedName)
			}
			if !ok {
				return nil
			}
		}
		if p.Run != nil {
			if err := p.Run(data); err != nil {
				return errors.Wrapf(err, "error execution phase %s", p.generatedName)
			}
		}
		return nil
	})
	return err
}
func (e *Runner) Help(cmdUse string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.prepareForExecution()
	maxLength := 0
	e.visitAll(func(p *phaseRunner) error {
		if !p.Hidden && !p.RunAllSiblings {
			length := len(p.use)
			if maxLength < length {
				maxLength = length
			}
		}
		return nil
	})
	line := fmt.Sprintf("The %q command executes the following phases:\n", cmdUse)
	line += "```\n"
	offset := 2
	e.visitAll(func(p *phaseRunner) error {
		if !p.Hidden && !p.RunAllSiblings {
			padding := maxLength - len(p.use) + offset
			line += strings.Repeat(" ", offset*p.level)
			line += p.use
			line += strings.Repeat(" ", padding)
			line += p.Short
			line += "\n"
		}
		return nil
	})
	line += "```"
	return line
}
func (e *Runner) SetAdditionalFlags(fn func(*pflag.FlagSet)) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.cmdAdditionalFlags = pflag.NewFlagSet("phaseAdditionalFlags", pflag.ContinueOnError)
	fn(e.cmdAdditionalFlags)
}
func (e *Runner) BindToCommand(cmd *cobra.Command) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(e.Phases) == 0 {
		return
	}
	e.prepareForExecution()
	phaseCommand := &cobra.Command{Use: "phase", Short: fmt.Sprintf("use this command to invoke single phase of the %s workflow", cmd.Name()), Args: cobra.NoArgs}
	cmd.AddCommand(phaseCommand)
	subcommands := map[string]*cobra.Command{}
	e.visitAll(func(p *phaseRunner) error {
		if p.Hidden {
			return nil
		}
		phaseSelector := p.generatedName
		if p.RunAllSiblings {
			phaseSelector = p.parent.generatedName
		}
		phaseCmd := &cobra.Command{Use: strings.ToLower(p.Name), Short: p.Short, Long: p.Long, Example: p.Example, Aliases: p.Aliases, Run: func(cmd *cobra.Command, args []string) {
			if len(p.Phases) > 0 {
				cmd.Help()
				return
			}
			e.runCmd = cmd
			e.Options.FilterPhases = []string{phaseSelector}
			if err := e.Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}, Args: cobra.NoArgs}
		inheritsFlags(cmd.Flags(), phaseCmd.Flags(), p.InheritFlags)
		if e.cmdAdditionalFlags != nil {
			inheritsFlags(e.cmdAdditionalFlags, phaseCmd.Flags(), p.InheritFlags)
		}
		if p.LocalFlags != nil {
			p.LocalFlags.VisitAll(func(f *pflag.Flag) {
				phaseCmd.Flags().AddFlag(f)
			})
		}
		if p.level == 0 {
			phaseCommand.AddCommand(phaseCmd)
		} else {
			subcommands[p.parent.generatedName].AddCommand(phaseCmd)
		}
		subcommands[p.generatedName] = phaseCmd
		return nil
	})
	if cmd.Long != "" {
		cmd.Long = fmt.Sprintf("%s\n\n%s\n", cmd.Long, e.Help(cmd.Use))
	} else {
		cmd.Long = fmt.Sprintf("%s\n\n%s\n", cmd.Short, e.Help(cmd.Use))
	}
	cmd.Flags().StringSliceVar(&e.Options.SkipPhases, "skip-phases", nil, "List of phases to be skipped")
	e.runCmd = cmd
}
func inheritsFlags(sourceFlags, targetFlags *pflag.FlagSet, cmdFlags []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cmdFlags == nil {
		return
	}
	sourceFlags.VisitAll(func(f *pflag.Flag) {
		for _, c := range cmdFlags {
			if f.Name == c {
				targetFlags.AddFlag(f)
			}
		}
	})
}
func (e *Runner) visitAll(fn func(*phaseRunner) error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, currentRunner := range e.phaseRunners {
		if err := fn(currentRunner); err != nil {
			return err
		}
	}
	return nil
}
func (e *Runner) prepareForExecution() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.phaseRunners = []*phaseRunner{}
	var parentRunner *phaseRunner
	for _, phase := range e.Phases {
		if phase.RunAllSiblings {
			continue
		}
		addPhaseRunner(e, parentRunner, phase)
	}
}
func addPhaseRunner(e *Runner, parentRunner *phaseRunner, phase Phase) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	generatedName := strings.ToLower(phase.Name)
	use := generatedName
	selfPath := []string{generatedName}
	if parentRunner != nil {
		generatedName = strings.Join([]string{parentRunner.generatedName, generatedName}, phaseSeparator)
		use = fmt.Sprintf("%s%s", phaseSeparator, use)
		selfPath = append(parentRunner.selfPath, selfPath...)
	}
	currentRunner := &phaseRunner{Phase: phase, parent: parentRunner, level: len(selfPath) - 1, selfPath: selfPath, generatedName: generatedName, use: use}
	e.phaseRunners = append(e.phaseRunners, currentRunner)
	for _, childPhase := range phase.Phases {
		addPhaseRunner(e, currentRunner, childPhase)
	}
}
