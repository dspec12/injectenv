package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		bold := color.New(color.FgBlue).Add(color.Bold).Add(color.Underline)
		verboseFlag, err := cmd.Flags().GetBool("verbose")
		cobra.CheckErr(err)

		switch verboseFlag {
		case true:
			bold.Println("Available Profiles")
			for k := range viper.AllSettings() {
				fmt.Println(k)
			}
			fmt.Println()
			for profile := range viper.AllSettings() {
				bold.Println(profile)
				for k, v := range viper.GetStringMapString(profile) {
					fmt.Printf("%s: %s\n", k, v)
				}
				fmt.Println()
			}
		default:
			bold.Println("Available Profiles")
			for k := range viper.AllSettings() {
				fmt.Println(k)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "Shows both profiles and variables.")
}
