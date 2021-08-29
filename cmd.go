package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/fatih/color"
)

// Const Do Not Mutate
var bold = color.New(color.FgBlue).Add(color.Bold).Add(color.Underline)

const helpPage = `
InjectEnv
Wraps and executes commands with additional environmental variables.

Usage: injectenv <command> [<flags> ...]
Example: injectenv exec -profile nomad-dev -- nomad job status

Commands:
  help
      Show help.

  list [<flags>]
    List profiles.

	Flags:
	  --profile			If specified the program will list all variables under the target profile.

  exec [<flags>] [<cmd>] [<args>...]
    Executes a command with additional profile vars in the environment

	Flags:
	  --profile			Name of the profile to use (Required).

`

func handleList(listCmd *flag.FlagSet, listProfile *string, ec envConfig) {
	listCmd.Parse(os.Args[2:])

	if *listProfile == "" {
		bold.Println("Available Profiles")
		for k := range ec {
			fmt.Println(k)
		}
		os.Exit(0)
	}

	if *listProfile == "help" {
		fmt.Println(helpPage)
		os.Exit(0)
	}

	if env, ok := ec[*listProfile]; ok {
		bold.Println(*listProfile)
		for k, v := range env {
			fmt.Printf("%s: %s\n", k, v)
		}
	} else {
		bold.Println("Available Profiles")
		for k := range ec {
			fmt.Println(k)
		}
		fmt.Println()
		log.Fatalf("Profile: \"%s\" not found. Please refer to the above profiles.", *listProfile)
	}
}

func handleExec(execCmd *flag.FlagSet, execProfile *string, ec envConfig) {
	execCmd.Parse(os.Args[2:])

	if *execProfile == "" {
		fmt.Println(helpPage)
		log.Fatalln("A profile must be specified.")
	}

	if len(os.Args) <= 5 {
		fmt.Println(helpPage)
		log.Fatalln("Malformated command. Not enough arguments.")
	}

	if os.Args[4] != "--" {
		fmt.Println(helpPage)
		log.Fatalln("Malformated command. Place \"--\" between the command flags and exec arguments.")
	}

	cmd := os.Args[5]
	bin, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatal(err)
	}

	ec.setENV(*execProfile)
	if err := syscall.Exec(bin, os.Args[5:], os.Environ()); err != nil {
		log.Fatal(err)
	}
}

func (ec envConfig) setENV(p string) {
	if env, ok := ec[p]; ok {
		for k, v := range env {
			err := os.Setenv(k, v)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		bold.Println("Available Profiles")
		for k := range ec {
			fmt.Println(k)
		}
		fmt.Println()
		log.Fatalf("Profile: \"%s\" not found. Please refer to the above profiles.", p)
	}
}
