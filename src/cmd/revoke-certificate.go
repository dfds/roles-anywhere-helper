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
		profileNameAcm, _ := cmd.Flags().GetString(flags.ProfileNameAcm)
		profileNameAcmPca, _ := cmd.Flags().GetString(flags.ProfileNameAcmPca)
		certArn, _ := cmd.Flags().GetString(flags.CertificateArn)
		pcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
		revocationReason, _ := cmd.Flags().GetString(flags.RevocationReason)
		acmRegion, _ := cmd.Flags().GetString(flags.AcmRegion)
		pcaRegion, _ := cmd.Flags().GetString(flags.PcaRegion)
		acmAccessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcm)
		acmPcaAccessKey, _ := cmd.Flags().GetString(flags.AccessKeyAcmPca)
		acmSecretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcm)
		acmPcaSecretAccessKey, _ := cmd.Flags().GetString(flags.SecretAccessKeyAcmPca)
		sessionToken, _ := cmd.Flags().GetString(flags.SessionTokenAcmPca)

		err := argsValidationHandler.IsValidRevocationReason(revocationReason)
		if err != nil {
			fmt.Println(err)
			return
		}

		acmCreds := awsService.NewAwsCredentialsObject(acmAccessKey, acmSecretAccessKey, sessionToken, profileNameAcm)
		acmPcaCreds := awsService.NewAwsCredentialsObject(acmPcaAccessKey, acmPcaSecretAccessKey, sessionToken, profileNameAcmPca)
		_, err = acmpcaService.RevokeCertificate(acmCreds, acmPcaCreds, certArn, pcaArn, revocationReason, acmRegion, pcaRegion)
		cobra.CheckErr(err)

	},
}

func init() {
	rootCmd.AddCommand(revokeCertificateCmd)

	revokeCertificateCmd.PersistentFlags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.AcmRegion, "eu-east-1", flags.RegionNameAcmDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.ProfileNameAcmPca, "default", flags.ProfNameAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.ProfileNameAcm, "default", flags.ProfNameAcmDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.AccessKeyAcmPca, "", flags.AccessKeyAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.AccessKeyAcm, "", flags.AccessKeyAcmDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcmPca, "", flags.SecretAccessKeyAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.SecretAccessKeyAcm, "", flags.SecretAccessKeyAcmDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.SessionTokenAcmPca, "", flags.SessionTokenAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.SessionTokenAcm, "", flags.SessionTokenAcmDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", flags.CertArnDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", flags.AcmPcaArnDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.RevocationReason, "r", revocationReasons.Unspecified, flags.RevocReasonDesc)

	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
}
