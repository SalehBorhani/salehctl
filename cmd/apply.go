package cmd

import (
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration by file name",
	Long:  `Shows the data of file which has been applied.`,
	Run: func(cmd *cobra.Command, args []string) {
		loadYaml(args[0])
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().BoolP("filename", "f", false, "salehctl apply -f FILENAME")

}
