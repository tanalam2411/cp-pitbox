package cluster

import "fmt"


type Cmd struct {

	Type string `arg:"" default:"kind" help:"The local cluster type to install, such as Kind"`
	 
}


func(c *Cmd) Help() string {
	return `
CrossplanePitbox local install cluster	
`
}


func (c *Cmd) Run() error {

	fmt.Println("Executing: installation of local kind cluster")
	fmt.Printf("\nType: %s\n", c.Type)
	return nil
}