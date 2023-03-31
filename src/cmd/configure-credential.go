package cmd

import (
	"github.com/dfds/roles-anywhere-helper/credentialService"
	"github.com/dfds/roles-anywhere-helper/flags"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure-credential",
	Short: "Configures your local AWS credential file for iam roles anywhere",
	Long:  `Configure your local AWS credential file for iam roles using a aws_signing_helper process`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certificatePath, _ := cmd.Flags().GetString(flags.CertificatePath)
		privateKeyPath, _ := cmd.Flags().GetString(flags.PrivateKeyPath)
		trustAnchorArn, _ := cmd.Flags().GetString(flags.TrustAnchor)
		profileArn, _ := cmd.Flags().GetString(flags.ProfileArn)
		roleArn, _ := cmd.Flags().GetString(flags.RoleArn)
		region, _ := cmd.Flags().GetString(flags.RolesAnywhereRegion)

		credentialService.Configure(profileName, certificatePath, privateKeyPath, trustAnchorArn, profileArn, roleArn, region)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.PersistentFlags().StringP(flags.ProfileName, "n", "default", flags.ProfNameRolesAnywhereDesc)
	configureCmd.PersistentFlags().StringP(flags.CertificatePath, "c", "", flags.CertDirDesc)
	configureCmd.PersistentFlags().StringP(flags.PrivateKeyPath, "k", "", flags.PrivKeyPathDesc)
	configureCmd.PersistentFlags().StringP(flags.TrustAnchor, "t", "", flags.TrustAnchorArnDesc)
	configureCmd.PersistentFlags().StringP(flags.ProfileArn, "p", "", flags.ProfileArnDesc)
	configureCmd.PersistentFlags().StringP(flags.RoleArn, "i", "", flags.RoleArnDesc)
	configureCmd.PersistentFlags().StringP(flags.RolesAnywhereRegion, "r", "eu-east-1", flags.RegionNameRolesAnywhereDesc)

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.CertificatePath)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.PrivateKeyPath)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.TrustAnchor)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.ProfileArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.RoleArn)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), flags.RolesAnywhereRegion)
}
