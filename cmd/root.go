package cmd

import (
	"fmt"
	"os"

	"github.com/dspec12/injectenv/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "injectenv",
	Short: "Wraps and executes commands with additional environmental variables",
	Long: `Injectenv adds profile defined variables to your environment for a single command

Example:
  injectenv exec profile1 -- printenv | grep key1
`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "0.1.2",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.injectenv.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the optional flag.
		err := config.EnvMap.LoadConfigFile(cfgFile)
		cobra.CheckErr(err)
	} else {
		// Find home directory.
		homedir, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Load in data from ".injectenv.yaml".
		cfgFile = homedir + "/.injectenv.yaml"
		if config.EnvMap.LoadConfigFile(cfgFile); err != nil {
			cobra.CheckErr(err)
		}
	}

	fmt.Fprintln(os.Stderr, "Using config file:", cfgFile)
	fmt.Fprintln(os.Stderr, "")
}
