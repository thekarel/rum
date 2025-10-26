package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thekarel/rum/internal"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rum <path to folder or package.json>",
	Short: "TUI to list, filter and run package.json scripts",
	Long: `TUI to list, filter and run package.json scripts.

To list the scripts in the current folder:
  rum

You can also pass relative or absolute paths either to a folder or a file:
  rum ./modules/thing/
  rum /code/project/package.json`,
	Run: func(cmd *cobra.Command, args []string) {
		path := "."

		if len(args) == 1 {
			path = args[0]
		}
		internal.PickAndRun(path)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rum.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
