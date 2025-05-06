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
		fmt.Println(viper.AllKeys())
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
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
