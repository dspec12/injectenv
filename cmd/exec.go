package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/dspec12/injectenv/config"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <profile>",
	Short: "Executes a command with specified profile variables added to the current environment",
	Long: `Executes a command with specified profile variables added to the current environment

Example:
  injectenv exec profile1 -- printenv | grep key1
`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Verifies that the correct the ammount of args are passed.
		if len(args) < 2 {
			return fmt.Errorf("command \"exec\" requires at least two args")
		}

		// Ensures the profile argument exists in config file.
		if _, ok := config.EnvMap[args[0]]; !ok {
			return fmt.Errorf("invalid argument \"%s\" not found in config", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		targetProfile := config.EnvMap[args[0]]

		// Set command executable.
		bin, err := exec.LookPath(args[1])
		cobra.CheckErr(err)

		// Creates a slice and adds the target profile variables to the users already defined environment.
		var envSlice []string

		for k, v := range targetProfile {
			ev := fmt.Sprintf("%s=%s", k, v)
			envSlice = append(envSlice, ev)
		}

		envSlice = append(envSlice, os.Environ()...)

		// Executes the command with all variables from the slice above.
		err = syscall.Exec(bin, args[1:], envSlice)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
