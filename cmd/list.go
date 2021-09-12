package cmd

import (
	"fmt"

	"github.com/dspec12/injectenv/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists profiles and optionally their variables",
	Long: `Lists profiles and optionally their variables

Example:
  injectenv list
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Set output colors.
		boldColor := color.New(color.FgBlue).Add(color.Bold).Add(color.Underline)
		keyColor := color.New(color.FgYellow)

		// Set verbose flag.
		verboseFlag, err := cmd.Flags().GetBool("verbose")
		cobra.CheckErr(err)

		switch verboseFlag {
		case true:
			// Verbose output.
			for profile := range config.EnvMap {
				boldColor.Println(profile)
				for k, v := range config.EnvMap[profile] {
					keyColor.Printf("%s: ", k)
					fmt.Println(v)
				}
				fmt.Println()
			}
		default:
			// Normal output.
			boldColor.Println("Profiles")
			for k := range config.EnvMap {
				fmt.Println(k)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "Shows both profiles and variables.")
}
