package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/credentialHandler"
	"github.com/dfds/iam-anywhere-ninja/flags"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure-credential",
	Short: "Configure your AWS credential file for iam roles",
	Long:  `Configure your AWS credential file for iam roles using a basic model`,
	Run: func(cmd *cobra.Command, args []string) {

		credentialHandler.Configure(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.PersistentFlags().StringP(flags.ProfileName, "n", "default", "Name of the profile to that the credentials will be created under")
	configureCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", "Directory of the SSH certificates and private key on the machine")
	configureCmd.PersistentFlags().StringP(flags.TrustAnchor, "t", "", "The Arn of the AWS IAM roles anywhere trust anchor")
	configureCmd.PersistentFlags().StringP(flags.ProfileArn, "p", "", "The Arn of the AWS IAM roles Anywhere profile")
	configureCmd.PersistentFlags().StringP(flags.RoleArn, "i", "", "The Arn of the role to be assumed with AWS IAM roles Anywhere")
	configureCmd.PersistentFlags().StringP(flags.Region, "r", "eu-east-1", "The region for the credential profile")

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.CertificateDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.PrivateKeyDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.TrustAnchor)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.ProfileArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.RoleArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.Region)
}
