package cmd

import (
	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/flags"

	"github.com/spf13/cobra"
)

var getCertificateCmd = &cobra.Command{
	Use:   "generate-certificate",
	Short: "Generates a X509 certificate and issues with AWS PCM",
	Long:  `Generates a Sha256 X509 certificate and issues with AWS PCM, It will then store the certificate files on the drive`,
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
		region, _ := cmd.Flags().GetString(flags.PcaRegion)
		expiryDays, _ := cmd.Flags().GetInt64(flags.CertificateExpiryDays)

		accessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcmPca)
		secretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcmPca)
		sessionToken, _ := cmd.Flags().GetString(flags.SessionTokenAcmPca)

		creds := awsService.NewAwsCredentialsObject(accessKey, secretAccessKey, sessionToken, profileName)

		_, err := acmpcaService.GenerateCertificate(creds, acmpcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory, region, expiryDays)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(getCertificateCmd)

	getCertificateCmd.PersistentFlags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)
	getCertificateCmd.PersistentFlags().String(flags.ProfileNameAcm, "default", flags.ProfNameAcmDesc)
	getCertificateCmd.PersistentFlags().String(flags.AccessKeyAcmPca, "", flags.AccessKeyAcmPcaDesc)
	getCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcmPca, "", flags.SecretAccessKeyAcmPcaDesc)
	getCertificateCmd.PersistentFlags().String(flags.SessionTokenAcmPca, "", flags.SessionTokenAcmPcaDesc)

	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", flags.OrgUnitDesc)
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "o", "", flags.OrgNameDesc)
	getCertificateCmd.PersistentFlags().StringP(flags.CommonName, "n", "", flags.CommonNameDesc)
	getCertificateCmd.PersistentFlags().StringP(flags.Country, "k", "", flags.CountryDesc)
	getCertificateCmd.PersistentFlags().StringP(flags.Locality, "l", "", flags.LocalityDesc)
	getCertificateCmd.PersistentFlags().StringP(flags.Province, "s", "", flags.Province)
	getCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", flags.AcmPcaArnDesc)
	getCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "d", "", flags.CertDirDesc)
	getCertificateCmd.PersistentFlags().Int64P(flags.CertificateExpiryDays, "e", 365, flags.CertificateExpiryDaysDesc)

	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
}
