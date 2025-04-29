package cmd

import (
	"fmt"
	u "os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile         string
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
		println(cfgFile, nexusUser, nexusPassword, nexusHost, nexusRepository)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := u.Current()
	cobra.CheckErr(err)

	configPath := fmt.Sprintf("%s/.nexuscli.json", home.HomeDir)

	nexusUser = viper.GetString("nexusUser")
	nexusPassword = viper.GetString("nexusPassword")
	nexusHost = viper.GetString("nexusHost")
	nexusRepository = viper.GetString("nexusRepository")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", configPath, "config file (default is $HOME/.nexuscli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&nexusUser, "user", "u", nexusUser, "your Nexus Repository Manager user name.")
	rootCmd.PersistentFlags().StringVarP(&nexusPassword, "password", "p", nexusPassword, "your Nexus Repository Manager password.")
	rootCmd.PersistentFlags().StringVarP(&nexusHost, "server", "s", nexusHost, "the address of the Nexus Repository Manager server to use.")
	rootCmd.PersistentFlags().StringVarP(&nexusRepository, "repository", "r", nexusRepository, "the registry for Nexus Repository Manager server to use.")
}

func initConfig() {
	viper.SetConfigType("json")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".nexuscli" (without extension).
		home, err := u.Current()
		cobra.CheckErr(err)

		viper.AddConfigPath(home.HomeDir)
		viper.SetConfigFile(".nexuscli.json")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		//Used for Debug
		//fmt.Printf(viper.GetString("user"))
	}

}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
