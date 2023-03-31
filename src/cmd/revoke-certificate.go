package cmd

import (
	"fmt"

	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/argsValidationHandler"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/flags"
	"github.com/dfds/roles-anywhere-helper/revocationReasons"
	"github.com/spf13/cobra"
)

var revokeCertificateCmd = &cobra.Command{
	Use:   "revoke-certificate",
	Short: "Revoke a certificate in AWS PCA",
	Long:  `Allow to revoke and provided a reason as to why a certificate is revoked in AWS PCA`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certArn, _ := cmd.Flags().GetString(flags.CertificateArn)
		pcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
		revocationReason, _ := cmd.Flags().GetString(flags.RevocationReason)
		region, _ := cmd.Flags().GetString(flags.PcaRegion)
		accessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcmPca)
		secretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcmPca)
		sessionToken, _ := cmd.Flags().GetString(flags.SessionTokenAcmPca)

		err := argsValidationHandler.IsValidRevocationReason(revocationReason)
		if err != nil {
			fmt.Println(err)
			return
		}

		creds := awsService.NewAwsCredentialsObject(accessKey, secretAccessKey, sessionToken, profileName)
		_, err = acmpcaService.RevokeCertificate(creds, certArn, pcaArn, revocationReason, region)
		cobra.CheckErr(err)

	},
}

func init() {
	rootCmd.AddCommand(revokeCertificateCmd)

	revokeCertificateCmd.PersistentFlags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.ProfileName, "default", flags.ProfNameAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.AccessKeyAcmPca, "", flags.AccessKeyAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcmPca, "", flags.SecretAccessKeyAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.SessionTokenAcmPca, "", flags.SessionTokenAcmPcaDesc)

	revokeCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", flags.CertArnDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", flags.AcmPcaArnDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.RevocationReason, "r", revocationReasons.Unspecified, flags.RevocReasonDesc)

	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
}
