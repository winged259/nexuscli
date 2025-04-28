package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/winged259/nexuscli/model"
)

var (
	nexusClient model.Registry
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
	Long:  "Authentication commands",
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	nexusClient = newClient()
	// },

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(nexusClient)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
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
