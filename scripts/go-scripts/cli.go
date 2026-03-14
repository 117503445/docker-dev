package main

var cli struct {
	Test cmdTest `cmd:"" help:"Test container environment after build"`
}

type cmdTest struct {
	ContainerName string `help:"Name of the container to test" default:"dev-desktop"`
}

func (c *cmdTest) Run() error {
	testContainer(c.ContainerName)
	return nil
}