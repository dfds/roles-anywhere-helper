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

	configureCmd.PersistentFlags().StringP(Flags.ProfileName, "n", "", "")
	configureCmd.PersistentFlags().StringP(Flags.CertificateDirectory, "x", "", "")
	configureCmd.PersistentFlags().StringP(Flags.PrivateKeyDirectory, "k", "", "")
	configureCmd.PersistentFlags().StringP(Flags.TrustAnchor, "t", "", "")
	configureCmd.PersistentFlags().StringP(Flags.ProfileArn, "p", "", "")
	configureCmd.PersistentFlags().StringP(Flags.RoleArn, "i", "", "")
	configureCmd.PersistentFlags().StringP(Flags.Region, "r", "", "")

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.CertificateDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.PrivateKeyDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.TrustAnchor)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.ProfileArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.RoleArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.Region)
}
