package cmd

import (
	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/flags"

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
		organizationName, _ := cmd.Flags().GetString(flags.OrganizationName)
		organizationalUnit, _ := cmd.Flags().GetString(flags.OrganizationalUnit)
		country, _ := cmd.Flags().GetString(flags.Country)
		locality, _ := cmd.Flags().GetString(flags.Locality)
		province, _ := cmd.Flags().GetString(flags.Province)
		certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)
		expiryDays, _ := cmd.Flags().GetInt64(flags.CertificateExpiryDays)

		_, err := acmpcaService.GenerateCertificate(profileName, acmpcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory, expiryDays)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(getCertificateCmd)
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", "The organization unit for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "o", "", "The organization name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.CommonName, "n", "", "The common name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.Country, "k", "", "The country name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.Locality, "l", "", "The locality name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.Province, "s", "", "The state or province name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", "Arn for the ACM PCA")
	getCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", "Profile of the ACM PCA")
	getCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", "Name of the profile to that the credentials will be created under")
	getCertificateCmd.PersistentFlags().Int64P(flags.CertificateExpiryDays, "e", 365, flags.CertificateExpiryDaysDesc)

	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
}
