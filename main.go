package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type envConfig map[string]map[string]string

func main() {
	// Set logger
	log.SetFlags(0)
	log.SetPrefix("[InjectEnv] ")

	// Hydrate env map from yaml config
	env := make(envConfig)
	env.loadConfigFile()

	// CLI  subcommands and flags
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	listProfile := listCmd.String("profile", "", "Lists all env vars for profile.")
	execProfile := execCmd.String("profile", "", "Executes a command with added ENV vars from profile. (Required) ")

	// Verify that a subcomand has been provided
	if len(os.Args) < 2 {
		fmt.Println(helpPage)
		log.Fatalln("No arguments specified.")
	}

	// Subcommand handlers
	switch os.Args[1] {
	case "list":
		handleList(listCmd, listProfile, env)
	case "exec":
		handleExec(execCmd, execProfile, env)
	case "help":
		println(helpPage)
		os.Exit(0)
	default:
		fmt.Println(helpPage)
		log.Fatalf("%s is not a valid argument.", os.Args[1])
	}
}
