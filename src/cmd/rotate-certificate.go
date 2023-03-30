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
		acmRegion, _ := cmd.Flags().GetString(flags.AcmRegion)
		acmPcaRegion, _ := cmd.Flags().GetString(flags.RegionNameAcmPcaDesc)
    expiryDays, _ := cmd.Flags().GetInt64(flags.CertificateExpiryDays)

		acmpcaService.GenerateCertificate(profileName, acmpcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory, acmPcaRegion, expiryDays)
		acmService.ImportCertificate(profileName, certificateDirectory, acmRegion)
		acmpcaService.RevokeCertificate(profileName, certArn, acmpcaArn, revocationReasons.Superseded, acmPcaRegion)
	},
}

func init() {
	rootCmd.AddCommand(rotateCertificateCmd)

	rotateCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", flags.ProfNameAcmPcaDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", flags.CertArnDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", flags.AcmPcaArnDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", flags.OrgUnitDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "o", "", flags.OrgNameDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.CommonName, "n", "", flags.CommonNameDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", flags.ProfNameRolesAnywhereDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.Country, "k", "", flags.CountryDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.Locality, "l", "", flags.LocalityDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.Province, "s", "", flags.ProvinceDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.AcmRegion, "eu-east-1", flags.RegionNameAcmDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)	
  rotateCertificateCmd.PersistentFlags().Int64P(flags.CertificateExpiryDays, "e", 365, flags.CertificateExpiryDaysDesc)

	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
}
