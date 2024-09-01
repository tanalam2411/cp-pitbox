package install

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/tanalam2411/cp-pitbox/cmd/local/install/cluster"
)

type Cmd struct {
	
	Cluster cluster.Cmd `cmd:"" help:"Install Cluster(supported types: kind, )"`
}

func (c *Cmd) Help() string {
	return `
	Local install
	`
}

func (c *Cmd) Run(k *kong.Context, log logging.Logger) error {

	fmt.Println("Running Local Install command .......")
	log.Debug("Running something...")

	return nil
}