/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "image subcommand",
	Long:  `Subcommand to interract with Nexus repository image`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image called")
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
