package cmd

import (
	"github.com/dfds/roles-anywhere-helper/acmService"
	"github.com/dfds/roles-anywhere-helper/flags"

	"github.com/spf13/cobra"
)

var importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "Uploads the certificate to ACM",
	Long:  `Adds your certificate to ACM`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)
    region, _ := cmd.Flags().GetString(flags.AcmRegion)
    
		_, err := acmService.ImportCertificate(profileName, certificateDirectory, region)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(importCertificateCmd)

	importCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", flags.ProfNameAcmDesc)
	importCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", flags.ProfNameRolesAnywhereDesc)
	importCertificateCmd.PersistentFlags().String(flags.AcmRegion, "eu-east-1", flags.RegionNameAcmDesc)

	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), flags.CertificateArn)
}
