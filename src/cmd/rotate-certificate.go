package cmd

import (
	"github.com/dfds/roles-anywhere-helper/acmService"
	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/flags"
	"github.com/dfds/roles-anywhere-helper/revocationReasons"
	"github.com/spf13/cobra"
)

var rotateCertificateCmd = &cobra.Command{
	Use:   "rotate-certificate",
	Short: "rotate certificate",
	Long:  `Rotates the certificate by first creating the new certificate then revokeing the old certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certArn, _ := cmd.Flags().GetString(flags.CertificateArn)
		acmpcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
		commonName, _ := cmd.Flags().GetString(flags.CommonName)
		organizationName, _ := cmd.Flags().GetString(flags.OrganizationName)
		organizationalUnit, _ := cmd.Flags().GetString(flags.OrganizationalUnit)
		country, _ := cmd.Flags().GetString(flags.Country)
		locality, _ := cmd.Flags().GetString(flags.Locality)
		province, _ := cmd.Flags().GetString(flags.Province)
		certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)

		acmpcaService.GenerateCertificate(profileName, acmpcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory)
		acmService.ImportCertificate(profileName, certificateDirectory)
		acmpcaService.RevokeCertificate(profileName, certArn, acmpcaArn, revocationReasons.Superseded)
	},
}

func init() {
	rootCmd.AddCommand(rotateCertificateCmd)

	rotateCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", "Name of the profile to be used for access to the PCA")
	rotateCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", "ARN of the certificate to be revoked")
	rotateCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", "ARN of the private CA that issues the certificate")
	rotateCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", "The organization unit for the X509 certificate")
	rotateCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "o", "", "The organization name for the X509 certificate")
	rotateCertificateCmd.PersistentFlags().StringP(flags.CommonName, "n", "", "The common name for the X509 certificate")
	rotateCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", "Name of the profile to that the credentials will be created under")
	rotateCertificateCmd.PersistentFlags().StringP(flags.Country, "k", "", "The country name for the X509 certificate")
	rotateCertificateCmd.PersistentFlags().StringP(flags.Locality, "l", "", "The locality name for the X509 certificate")
	rotateCertificateCmd.PersistentFlags().StringP(flags.Province, "s", "", "The state or province name for the X509 certificate")

	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
}
