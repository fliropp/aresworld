package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "root",
}

var firstCmd = &cobra.Command{
	Use:   "are1",
	Short: "Ares World cmd 1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
	},
}

var secondCmd = &cobra.Command{
	Use:   "are2",
	Short: "Ares World cmd 2",
	Run: func(cmd *cobra.Command, args []string) {
		for _, b := range args {
			fmt.Println(b)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := firstCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(firstCmd)
	rootCmd.AddCommand(secondCmd)
}
