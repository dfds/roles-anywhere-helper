package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/acmpcaService"
	"github.com/dfds/iam-anywhere-ninja/flags"

	"github.com/spf13/cobra"
)

var getCertificateCmd = &cobra.Command{
	Use:   "generate-certificate",
	Short: "Generate certificate",
	Long:  `Generate certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		acmpcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
		commonName, _ := cmd.Flags().GetString(flags.CommonName)
		organizationName, _ := cmd.Flags().GetStringArray(flags.OrganizationName)
		organizationalUnit, _ := cmd.Flags().GetStringArray(flags.OrganizationalUnit)
		certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)

		acmpcaService.ImportCertificate(profileName, acmpcaArn, commonName, organizationName, organizationalUnit, certificateDirectory)
	},
}

func init() {
	rootCmd.AddCommand(getCertificateCmd)
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", "The organization unit for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "n", "", "The organization name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.CommonName, "c", "", "The common name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", "Arn for the ACM PCA")
	getCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", "Profile of the ACM PCA")
	getCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", "Name of the profile to that the credentials will be created under")

	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
}
