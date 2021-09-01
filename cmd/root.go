package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "injectenv",
	Short: "Wraps and executes commands with additional environmental variables.",
	Long: `Injectenv adds profile defined variables to your environment for a single command.

Example:
  injectenv exec profile1 -- printenv | grep key1
`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "0.1.0",
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
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".injectenv" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".injectenv")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Warning: No config file found")
		fmt.Println()
	} else {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		fmt.Println()
	}
}