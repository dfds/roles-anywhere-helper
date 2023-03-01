/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"go-aws-iam-roles-anywhere-credential-helper/CredentialHandler"
	"go-aws-iam-roles-anywhere-credential-helper/Flags"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your AWS credential file for iam roles",
	Long:  `Configure your AWS credential file for iam roles using a basic model`,
	Run: func(cmd *cobra.Command, args []string) {
		CredentialHandler.Configure(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.PersistentFlags().StringP(Flags.ProfileName, "u", "", "")
	configureCmd.PersistentFlags().StringP(Flags.CertificateDirectory, "u", "", "")
	configureCmd.PersistentFlags().StringP(Flags.PrivateKeyDirectory, "u", "", "")
	configureCmd.PersistentFlags().StringP(Flags.TrustAnchor, "u", "", "")
	configureCmd.PersistentFlags().StringP(Flags.ProfileArn, "u", "", "")
	configureCmd.PersistentFlags().StringP(Flags.RoleArn, "u", "", "")
	configureCmd.PersistentFlags().StringP(Flags.Region, "u", "", "")

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.CertificateDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.PrivateKeyDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.TrustAnchor)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.ProfileArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.RoleArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.Region)
}
