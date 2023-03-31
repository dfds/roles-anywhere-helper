package cmd

import (
	"path/filepath"

	"github.com/dfds/roles-anywhere-helper/acmService"
	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/credentialService"
	"github.com/dfds/roles-anywhere-helper/fileNames"
	"github.com/dfds/roles-anywhere-helper/flags"

	"github.com/spf13/cobra"
)

var setupAllCmd = getSetupAllCmd()

func getSetupAllCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "setup-all",
		Short: "Generate certificate, import it to ACm and configure credentials",
		Long: `Setup the whole process of configuring AWS IAM Roles Anywhere.
		List of operations:
			- Generates certificate and issues that in AWS ACM PCA.
			- Imports certificate into AWS ACM.
			- Configures AWS profile in credentials file to use AWS Signing Helper with created certificate.
		`,
		Run: setupAllCmdRun,
	}
}

func setupAllCmdRun(cmd *cobra.Command, args []string) {

	// Certificate generation flags values
	profileNameAcmPca, _ := cmd.Flags().GetString(flags.ProfileNameAcmPca)
	acmPcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
	commonName, _ := cmd.Flags().GetString(flags.CommonName)
	organizationName, _ := cmd.Flags().GetString(flags.OrganizationName)
	organizationalUnit, _ := cmd.Flags().GetString(flags.OrganizationalUnit)
	country, _ := cmd.Flags().GetString(flags.Country)
	locality, _ := cmd.Flags().GetString(flags.Locality)
	province, _ := cmd.Flags().GetString(flags.Province)
	expiryDays, _ := cmd.Flags().GetInt64(flags.CertificateExpiryDays)

	certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)

	// Certificate import flags values
	profileNameAcm, _ := cmd.Flags().GetString(flags.ProfileNameAcm)

	// IAM Roles Anywhere flags values
	trustAnchorArn, _ := cmd.Flags().GetString(flags.TrustAnchor)
	profileArn, _ := cmd.Flags().GetString(flags.ProfileArn)
	roleArn, _ := cmd.Flags().GetString(flags.RoleArn)
	profileNameRoles, _ := cmd.Flags().GetString(flags.ProfileNameRolesAnywhere)
	rolesAnywhereRegion, _ := cmd.Flags().GetString(flags.RegionNameRolesAnywhereDesc)
	pcaRegion, _ := cmd.Flags().GetString(flags.RegionNameAcmPcaDesc)
	acmRegion, _ := cmd.Flags().GetString(flags.RegionNameAcmDesc)

	acmAccessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcm)
	acmSecretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcm)
	acmSessionToken, _ := cmd.Flags().GetString(flags.SessionTokenAcm)
	acmPcaAccessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcmPca)
	acmPcaSecretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcmPca)
	acmPcaSessionToken, _ := cmd.Flags().GetString(flags.SessionTokenAcmPca)

	acmCreds := awsService.NewAwsCredentialsObject(acmAccessKey, acmSecretAccessKey, acmSessionToken, profileNameAcm)
	acmPcaCreds := awsService.NewAwsCredentialsObject(acmPcaAccessKey, acmPcaSecretAccessKey, acmPcaSessionToken, profileNameAcmPca)

	_, err := acmpcaService.GenerateCertificate(acmPcaCreds, acmPcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory, pcaRegion, expiryDays)
	cobra.CheckErr(err)

	_, err = acmService.ImportCertificate(acmCreds, certificateDirectory, acmRegion)
	cobra.CheckErr(err)

	var certificatePath = filepath.Join(certificateDirectory, fileNames.Certificate)
	var privateKeyPath = filepath.Join(certificateDirectory, fileNames.PrivateKey)
	credentialService.Configure(profileNameRoles, certificatePath, privateKeyPath, trustAnchorArn, profileArn, roleArn, rolesAnywhereRegion)

}

func setupAllCmdFlags(cmd *cobra.Command) {
	cmd.Flags().String(flags.ProfileNameAcmPca, "default", flags.ProfNameAcmPcaDesc)
	cmd.Flags().String(flags.ProfileNameAcm, "default", flags.ProfNameAcmDesc)
	cmd.Flags().String(flags.ProfileNameRolesAnywhere, "roles-anywhere", flags.ProfNameRolesAnywhereDesc)
	cmd.Flags().String(flags.AcmpcaArn, "", flags.AcmPcaArnDesc)

	cmd.Flags().String(flags.OrganizationalUnit, "", flags.OrgUnitDesc)
	cmd.Flags().String(flags.OrganizationName, "", flags.OrgNameDesc)
	cmd.Flags().String(flags.CommonName, "", flags.CommonNameDesc)
	cmd.Flags().String(flags.Country, "", flags.CountryDesc)
	cmd.Flags().String(flags.Locality, "", flags.LocalityDesc)
	cmd.Flags().String(flags.Province, "", flags.ProvinceDesc)

	cmd.Flags().String(flags.CertificateDirectory, "", flags.CertDirDesc)

	cmd.Flags().String(flags.TrustAnchor, "", flags.TrustAnchorArnDesc)
	cmd.Flags().String(flags.ProfileArn, "", flags.ProfileArnDesc)
	cmd.Flags().String(flags.RoleArn, "", flags.RoleArnDesc)
	cmd.Flags().String(flags.AcmRegion, "eu-east-1", flags.RegionNameAcmDesc)
	cmd.Flags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)
	cmd.Flags().String(flags.RolesAnywhereRegion, "eu-east-1", flags.RegionNameRolesAnywhereDesc)
	cmd.Flags().Int64(flags.CertificateExpiryDays, 365, flags.CertificateExpiryDaysDesc)

	cmd.Flags().String(flags.AccessKeyAcm, "", flags.AccessKeyAcmDesc)
	cmd.Flags().String(flags.SecretAccessKeyAcm, "", flags.SecretAccessKeyAcmDesc)
	cmd.Flags().String(flags.SessionTokenAcm, "", flags.SessionTokenAcmDesc)
	cmd.Flags().String(flags.AccessKeyAcmPca, "", flags.AccessKeyAcmPcaDesc)
	cmd.Flags().String(flags.SecretAccessKeyAcmPca, "", flags.SecretAccessKeyAcmPcaDesc)
	cmd.Flags().String(flags.SessionTokenAcmPca, "", flags.SessionTokenAcmPcaDesc)

	cmd.MarkFlagRequired(flags.CommonName)
	cmd.MarkFlagRequired(flags.AcmpcaArn)
	cmd.MarkFlagRequired(flags.CertificateDirectory)

	cmd.MarkFlagRequired(flags.CertificatePath)
	cmd.MarkFlagRequired(flags.PrivateKeyPath)
	cmd.MarkFlagRequired(flags.TrustAnchor)
	cmd.MarkFlagRequired(flags.ProfileArn)
	cmd.MarkFlagRequired(flags.RoleArn)
}

func init() {
	rootCmd.AddCommand(setupAllCmd)
	setupAllCmdFlags(setupAllCmd)
}
