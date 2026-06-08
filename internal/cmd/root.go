package cmd

import (
	"flag"
	"fmt"
	"os"
)

type command struct {
	name string
	args func(*flag.FlagSet)
	run  func() error
}

func Run() {
	commands := []*command{
		fixturesCmd(),
		lsCmd(),
	}

	if len(os.Args) < 2 {
		printUsage(commands)
		os.Exit(1)
	}

	name := os.Args[1]
	subArgs := os.Args[2:]

	for _, c := range commands {
		if c.name != name {
			continue
		}

		fs := flag.NewFlagSet(c.name, flag.ExitOnError)
		c.args(fs)
		err := fs.Parse(subArgs)
		if err != nil {
			fmt.Printf("Bad flags passed to %s: %v\n", c.name, err)
			os.Exit(1)
		}
		err = c.run()
		if err != nil {
			fmt.Printf("Command %s failed: %v\n", c.name, err)
			os.Exit(1)
		}
		return
	}

	fmt.Printf("Unknown command: %s\n", name)
	printUsage(commands)
	os.Exit(1)
}

func printUsage(cmds []*command) {
	fmt.Println("Usage: app <command> [flags]")
	fmt.Println("Available commands:")
	for _, cmd := range cmds {
		fmt.Printf("  - %s\n", cmd.name)
	}
}
