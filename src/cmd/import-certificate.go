package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/acmService"
	"github.com/dfds/iam-anywhere-ninja/flags"

	"github.com/spf13/cobra"
)

var importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "upload certificate to ACM",
	Long:  `Adds your certificate to ACM`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)
		acmService.ImportCertificate(profileName, certificateDirectory)
	},
}

func init() {
	rootCmd.AddCommand(importCertificateCmd)

	importCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", "Name of the profile to be used for access to the ACM")
	importCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", "Name of the profile to that the credentials will be created under")

	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), flags.CertificateArn)
}
