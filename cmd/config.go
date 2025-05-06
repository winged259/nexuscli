package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/winged259/nexuscli/model"
)

var (
	nexusClient model.Registry
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config commands",
	Long:  "Config commands",
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	nexusClient = newClient()
	// },

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(nexusClient)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringVarP(&nexusUser, "user", "u", nexusUser, "your Nexus Repository Manager user name.")
	configCmd.PersistentFlags().StringVarP(&nexusPassword, "password", "p", nexusPassword, "your Nexus Repository Manager password.")
	configCmd.PersistentFlags().StringVarP(&nexusHost, "server", "s", nexusHost, "the address of the Nexus Repository Manager server to use.")
	configCmd.PersistentFlags().StringVarP(&nexusRepository, "repository", "r", nexusRepository, "the registry for Nexus Repository Manager server to use.")
}

func newClient(host, user, password, repo string) (string, error) {
	// nexusUser = viper.GetString("nexusServer")
	// nexusPassword = viper.GetString("nexusPassword")
	// nexusHost = viper.GetString("nexusHost")
	model := model.Registry{
		Hostname:   host,
		Username:   user,
		Password:   password,
		Repository: repo,
	}
	token, err := model.GetToken()
	return token, err
}
