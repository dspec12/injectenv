package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <profile>",
	Short: "Executes a command with specified profile variables in the environment.",
	Long: `Example:
  injectenv exec profile1 -- printenv | grep key1
`,
	Args: func(cmd *cobra.Command, args []string) error {
		v := viper.AllSettings()

		// Verifies that correct the ammount of args are passed.
		if len(args) < 2 {
			return fmt.Errorf("command \"exec\" requires at least two args")
		}

		// Ensures the profile argument exists in config file.
		if _, ok := v[args[0]]; !ok {
			return fmt.Errorf("invalid argument \"%s\" not found in config", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		targetProfile := viper.GetStringMapString(args[0])

		bin, err := exec.LookPath(args[1])
		cobra.CheckErr(err)

		for k, v := range targetProfile {
			err := os.Setenv(k, v)
			cobra.CheckErr(err)
		}

		err = syscall.Exec(bin, args[1:], os.Environ())
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
