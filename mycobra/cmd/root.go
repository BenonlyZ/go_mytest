package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobrademo",
	Short: "sparkdev's cobra demo",
	Long:  "the demo show how to use cobra package",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cobra demo program, with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var Verbose bool

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobrademo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobrademo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobrademo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
