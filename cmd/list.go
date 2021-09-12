package cmd

import (
	"fmt"

	"github.com/dspec12/injectenv/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List profiles.",
	Long: `Lists the profiles defined in config file.

Example:
  injectenv list
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		boldColor := color.New(color.FgBlue).Add(color.Bold).Add(color.Underline)
		keyColor := color.New(color.FgYellow)
		verboseFlag, err := cmd.Flags().GetBool("verbose")
		cobra.CheckErr(err)

		switch verboseFlag {
		case true:
			for profile := range config.EnvMap {
				boldColor.Println(profile)
				for k, v := range config.EnvMap[profile] {
					keyColor.Printf("%s: ", k)
					fmt.Println(v)
				}
				fmt.Println()
			}
		default:
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
