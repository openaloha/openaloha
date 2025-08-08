package runhandler

type RunHandler interface {
	Run(cmds []string) error
}