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

var setupAllCmd = &cobra.Command{
	Use:   "setup-all",
	Short: "Generate certificate, Import Certificate, Configure Credentials",
	Long: `Setup the whole process of configuring AWS IAM Roles Anywhere.
	List of operations:
		- Generates certificate and issues that in AWS ACM PCA.
		- Imports certificate into AWS ACM.
		- Configures AWS profile in credentials file to use AWS Signing Helper with created certificate.
	`,
	Run: func(cmd *cobra.Command, args []string) {

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

	},
}

func init() {
	rootCmd.AddCommand(setupAllCmd)

	setupAllCmd.Flags().String(flags.ProfileNameAcmPca, "default", flags.ProfNameAcmPcaDesc)
	setupAllCmd.Flags().String(flags.ProfileNameAcm, "default", flags.ProfNameAcmDesc)
	setupAllCmd.Flags().String(flags.ProfileNameRolesAnywhere, "roles-anywhere", flags.ProfNameRolesAnywhereDesc)
	setupAllCmd.Flags().String(flags.AcmpcaArn, "", flags.AcmPcaArnDesc)

	setupAllCmd.Flags().String(flags.OrganizationalUnit, "", flags.OrgUnitDesc)
	setupAllCmd.Flags().String(flags.OrganizationName, "", flags.OrgNameDesc)
	setupAllCmd.Flags().String(flags.CommonName, "", flags.CommonNameDesc)
	setupAllCmd.Flags().String(flags.Country, "", flags.CountryDesc)
	setupAllCmd.Flags().String(flags.Locality, "", flags.LocalityDesc)
	setupAllCmd.Flags().String(flags.Province, "", flags.ProvinceDesc)

	setupAllCmd.Flags().String(flags.CertificateDirectory, "", flags.CertDirDesc)

	setupAllCmd.Flags().String(flags.TrustAnchor, "", flags.TrustAnchorArnDesc)
	setupAllCmd.Flags().String(flags.ProfileArn, "", flags.ProfileArnDesc)
	setupAllCmd.Flags().String(flags.RoleArn, "", flags.RoleArnDesc)
	setupAllCmd.Flags().String(flags.Region, "eu-east-1", flags.RegionDesc)

	setupAllCmd.MarkFlagRequired(flags.CommonName)
	setupAllCmd.MarkFlagRequired(flags.AcmpcaArn)
	setupAllCmd.MarkFlagRequired(flags.CertificateDirectory)

	setupAllCmd.MarkFlagRequired(flags.CertificatePath)
	setupAllCmd.MarkFlagRequired(flags.PrivateKeyPath)
	setupAllCmd.MarkFlagRequired(flags.TrustAnchor)
	setupAllCmd.MarkFlagRequired(flags.ProfileArn)
	setupAllCmd.MarkFlagRequired(flags.RoleArn)
	setupAllCmd.MarkFlagRequired(flags.Region)
}
