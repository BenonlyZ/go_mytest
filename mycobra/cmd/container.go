package cmd

import (
	"fmt"

	"strings"

	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Print containers information",
	Long:  "Print all containers information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("container one is a dog")
		fmt.Println("container two is a cat")
		fmt.Println("container args are : " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
