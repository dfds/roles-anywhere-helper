package cmd

import (
	"iam-anywhere-ninja/CredentialHandler"
	"iam-anywhere-ninja/Flags"

	"github.com/spf13/cobra"
)

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

	configureCmd.PersistentFlags().StringP(Flags.ProfileName, "n", "default", "Name of the profile to that the credentials will be created under")
	configureCmd.PersistentFlags().StringP(Flags.CertificateDirectory, "x", "", "Destination of the Certificate on the machine")
	configureCmd.PersistentFlags().StringP(Flags.PrivateKeyDirectory, "k", "", "Destination of the unincrypted private key on the machine")
	configureCmd.PersistentFlags().StringP(Flags.TrustAnchor, "t", "", "The Arn of the AWS IAM roles anywhere trust anchor")
	configureCmd.PersistentFlags().StringP(Flags.ProfileArn, "p", "", "The Arn of the AWS IAM roles Anywhere profile")
	configureCmd.PersistentFlags().StringP(Flags.RoleArn, "i", "", "The Arn of the role to be assumed with AWS IAM roles Anywhere")
	configureCmd.PersistentFlags().StringP(Flags.Region, "r", "us-east-1", "The region for the credential profile")

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.CertificateDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.PrivateKeyDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.TrustAnchor)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.ProfileArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.RoleArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.Region)
}
