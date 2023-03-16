package cmd

import (
	"fmt"
	"github.com/dfds/iam-anywhere-ninja/acmpcaService"
	"github.com/dfds/iam-anywhere-ninja/argsValidationHandler"
	"github.com/dfds/iam-anywhere-ninja/flags"
	"github.com/dfds/iam-anywhere-ninja/revocationReasons"
	"github.com/spf13/cobra"
)

var revokeCertificateCmd = &cobra.Command{
	Use:   "revoke-certificate",
	Short: "revoke certificate",
	Long:  `Revokes the certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certArn, _ := cmd.Flags().GetString(flags.CertificateArn)
		pcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
		revocationReason, _ := cmd.Flags().GetString(flags.RevocationReason)

		err := argsValidationHandler.IsValidRevocationReason(revocationReason)
		if err != nil {
			fmt.Println(err)
			return
		}
		acmpcaService.RevokeCertificate(profileName, certArn, pcaArn, revocationReason)
	},
}

func init() {
	rootCmd.AddCommand(revokeCertificateCmd)

	revokeCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", "Name of the profile to be used for access to the PCA")
	revokeCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", "ARN of the certificate to be revoked")
	revokeCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", "ARN of the private CA that issues the certificate")
	revokeCertificateCmd.PersistentFlags().StringP(flags.RevocationReason, "r", revocationReasons.Unspecified, "Reason why the certificate is revoked")

	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
}
