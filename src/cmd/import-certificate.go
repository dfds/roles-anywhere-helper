package cmd

import (
	"github.com/dfds/roles-anywhere-helper/acmService"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/flags"

	"github.com/spf13/cobra"
)

var importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "Uploads a certificate to ACM",
	Long:  `Adds your certificate to ACM`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)
		region, _ := cmd.Flags().GetString(flags.AcmRegion)
		accessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcm)
		secretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcm)
		sessionToken, _ := cmd.Flags().GetString(flags.SessionTokenAcm)

		creds := awsService.NewAwsCredentialsObject(accessKey, secretAccessKey, sessionToken, profileName)

		_, err := acmService.ImportCertificate(creds, certificateDirectory, region)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(importCertificateCmd)

	importCertificateCmd.PersistentFlags().String(flags.ProfileName, "default", flags.ProfNameAcmDesc)
	importCertificateCmd.PersistentFlags().String(flags.AcmRegion, "eu-east-1", flags.RegionNameAcmDesc)
	importCertificateCmd.PersistentFlags().String(flags.AccessKeyAcm, "", flags.AccessKeyAcmDesc)
	importCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcm, "", flags.SecretAccessKeyAcmDesc)
	importCertificateCmd.PersistentFlags().String(flags.SessionTokenAcm, "", flags.SessionTokenAcmDesc)

	importCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", flags.ProfNameRolesAnywhereDesc)
	importCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", flags.CertArnDesc)

	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), flags.CertificateArn)
}
