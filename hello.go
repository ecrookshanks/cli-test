package main

import (
	"cli-test/cmd"
)

func main() {
	runCmd := cmd.GetCmd()
	cmd.Execute(runCmd)
}
