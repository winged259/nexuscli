package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
 
var (
	cfgFile string
	nexusUser       string
	nexusPassword   string
	nexusHost       string
	nexusRepository string
)
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".nexuscli.json", "config file (default is $HOME/.nexuscli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&nexusUser, "user", "u", "cicd_system", "your Nexus Repository Manager user name.")
	rootCmd.PersistentFlags().StringVarP(&nexusPassword, "password", "p", "password", "your Nexus Repository Manager password.")
	rootCmd.PersistentFlags().StringVarP(&nexusHost, "server", "s", "https://nexus.abbank.vn", "the address of the Nexus Repository Manager server to use.")
	rootCmd.PersistentFlags().StringVarP(&nexusRepository, "repository", "r", "abb-registry", "the registry for Nexus Repository Manager server to use.")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		fmt.Println(cfgFile)
	} else {
		// Search config in home directory with name ".nexuscli" (without extension).
		fmt.Println("No config file found")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("json")
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
