package cmd

import (
	"path/filepath"

	"github.com/dfds/iam-anywhere-ninja/acmService"
	"github.com/dfds/iam-anywhere-ninja/acmpcaService"
	"github.com/dfds/iam-anywhere-ninja/credentialService"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
	"github.com/dfds/iam-anywhere-ninja/flags"

	"github.com/spf13/cobra"
)

var setupAllCmd = getSetupAllCmd()

func getSetupAllCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "setup-all",
		Short: "Generate certificate, Import Certificate, Configure Credentials",
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

	certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)

	// Certificate import flags values
	profileNameAcm, _ := cmd.Flags().GetString(flags.ProfileNameAcm)

	// IAM Roles Anywhere flags values
	trustAnchorArn, _ := cmd.Flags().GetString(flags.TrustAnchor)
	profileArn, _ := cmd.Flags().GetString(flags.ProfileArn)
	roleArn, _ := cmd.Flags().GetString(flags.RoleArn)
	profileNameRoles, _ := cmd.Flags().GetString(flags.ProfileNameRolesAnywhere)
	region, _ := cmd.Flags().GetString(flags.Region)

	acmpcaService.GenerateCertificate(profileNameAcmPca, acmPcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory)
	acmService.ImportCertificate(profileNameAcm, certificateDirectory)

	var certificatePath = filepath.Join(certificateDirectory, fileNames.Certificate)
	var privateKeyPath = filepath.Join(certificateDirectory, fileNames.PrivateKey)
	credentialService.Configure(profileNameRoles, certificatePath, privateKeyPath, trustAnchorArn, profileArn, roleArn, region)

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
	cmd.Flags().String(flags.Region, "eu-east-1", flags.RegionDesc)

	cmd.MarkFlagRequired(flags.CommonName)
	cmd.MarkFlagRequired(flags.AcmpcaArn)
	cmd.MarkFlagRequired(flags.CertificateDirectory)

	cmd.MarkFlagRequired(flags.CertificatePath)
	cmd.MarkFlagRequired(flags.PrivateKeyPath)
	cmd.MarkFlagRequired(flags.TrustAnchor)
	cmd.MarkFlagRequired(flags.ProfileArn)
	cmd.MarkFlagRequired(flags.RoleArn)
	cmd.MarkFlagRequired(flags.Region)
}

func init() {
	rootCmd.AddCommand(setupAllCmd)
	setupAllCmdFlags(setupAllCmd)
}