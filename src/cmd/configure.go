/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go-aws-iam-roles-anywhere-credential-helper/CredentialHandler"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your AWS credential file for iam roles",
	Long:  `Configure your AWS credential file for iam roles using a basic model`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.PersistentFlags().StringP("profile-Name", "u", "", "")
	configureCmd.PersistentFlags().StringP("certificate-directory", "u", "", "")
	configureCmd.PersistentFlags().StringP("private-key-directory", "u", "", "")
	configureCmd.PersistentFlags().StringP("trust-anchor-arn", "u", "", "")
	configureCmd.PersistentFlags().StringP("profile-arn", "u", "", "")
	configureCmd.PersistentFlags().StringP("role-arn", "u", "", "")
	configureCmd.PersistentFlags().StringP("region", "u", "", "")

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), "certificate-directory")
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), "private-key-directory")
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), "trust-anchor-arn")
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), "profile-arn")
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), "role-arn")

	CredentialHandler.Configure(rootCmd.PersistentFlags().Lookup("certificate-directory"), rootCmd.PersistentFlags().Lookup("private-key-directory"), rootCmd.PersistentFlags().Lookup("trust-anchor-arn"), rootCmd.PersistentFlags().Lookup("profile-arn"), rootCmd.PersistentFlags().Lookup("role-arn"), rootCmd.PersistentFlags().Lookup("profile-Name"), rootCmd.PersistentFlags().Lookup("region"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
