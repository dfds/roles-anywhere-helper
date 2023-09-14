package cmd

import (
	"github.com/dfds/roles-anywhere-helper/acmService"
	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/flags"
	"github.com/dfds/roles-anywhere-helper/revocationReasons"
	"github.com/spf13/cobra"
)

var rotateCertificateCmd = &cobra.Command{
	Use:   "rotate-certificate",
	Short: "Rotate certificate",
	Long:  `Rotates the certificate by first creating the new certificate then revoking the old certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		profileNameAcm, _ := cmd.Flags().GetString(flags.ProfileNameAcm)
		profileNamePca, _ := cmd.Flags().GetString(flags.ProfileNameAcmPca)
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

		accessKeyPca, _ := cmd.Flags().GetString(flags.AccessKeyAcmPca)
		secretAccessKeyPca, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcmPca)
		sessionTokenPca, _ := cmd.Flags().GetString(flags.SessionTokenAcmPca)

		accessKeyAcm, _ := cmd.Flags().GetString(flags.AccessKeyAcm)
		secretAccessKeyAcm, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcm)
		sessionTokenAcm, _ := cmd.Flags().GetString(flags.SessionTokenAcm)

		acmCreds := awsService.NewAwsCredentialsObject(accessKeyAcm, secretAccessKeyAcm, sessionTokenAcm, profileNameAcm)
		acmPcaCreds := awsService.NewAwsCredentialsObject(accessKeyPca, secretAccessKeyPca, sessionTokenPca, profileNamePca)

		_, err := acmpcaService.GenerateCertificate(acmPcaCreds, acmpcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory, acmPcaRegion, expiryDays)
		cobra.CheckErr(err)
		_, err = acmService.ImportCertificate(acmCreds, certificateDirectory, acmRegion)
		cobra.CheckErr(err)
		_, err = acmpcaService.RevokeCertificate(acmPcaCreds, certArn, acmpcaArn, revocationReasons.Superseded, acmPcaRegion)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(rotateCertificateCmd)
	
	rotateCertificateCmd.PersistentFlags().String(flags.ProfileNameAcm, "default", flags.ProfNameAcmDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.ProfileNameAcmPca, "default", flags.ProfNameAcmPcaDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.AcmRegion, "eu-east-1", flags.RegionNameAcmDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.AccessKeyAcm, "", flags.AccessKeyAcmDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcm, "", flags.SecretAccessKeyAcmDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.SessionTokenAcm, "", flags.SessionTokenAcmDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.AccessKeyAcmPca, "", flags.AccessKeyAcmPcaDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcmPca, "", flags.SecretAccessKeyAcmPcaDesc)
	rotateCertificateCmd.PersistentFlags().String(flags.SessionTokenAcmPca, "", flags.SessionTokenAcmPcaDesc)

	rotateCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", flags.CertArnDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", flags.AcmPcaArnDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", flags.OrgUnitDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "o", "", flags.OrgNameDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.CommonName, "n", "", flags.CommonNameDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", flags.ProfNameRolesAnywhereDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.Country, "k", "", flags.CountryDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.Locality, "l", "", flags.LocalityDesc)
	rotateCertificateCmd.PersistentFlags().StringP(flags.Province, "s", "", flags.ProvinceDesc)
	rotateCertificateCmd.PersistentFlags().Int64P(flags.CertificateExpiryDays, "e", 365, flags.CertificateExpiryDaysDesc)

	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(rotateCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
}
