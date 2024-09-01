package local

import "github.com/tanalam2411/cp-pitbox/cmd/local/install"


type Cmd struct {
	Install install.Cmd `cmd:"" help:"Install kind, Crossplane etc"`
}

func (c *Cmd) Help() string {
	return "Local: helps to perform local setup"
}