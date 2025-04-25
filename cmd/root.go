package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "nexuscli",
	Short: "Nexus CLI",
	Long:  "Nexus CLI tool to interact with nexus repository",
	Run: func(cmd *cobra.Command, args []string) {
		println("hello")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nexuscli.yaml)")

	// rootCmd.AddCommand(authCmd)
	// rootCmd.AddCommand(configCmd)
	// rootCmd.AddCommand(repoCmd)
	// rootCmd.AddCommand(repoDeleteCmd)
	// rootCmd.AddCommand(repoListCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".nexuscli" (without extension).
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nexuscli")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	notFound := &viper.ConfigFileNotFoundError{}

	switch {
	case err != nil && !errors.As(err, notFound):
		cobra.CheckErr(err)
	case err != nil && errors.As(err, notFound):
		break

	default:
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
