package cmd

import (
	"fmt"
	"github.com/dfds/iam-anywhere-ninja/acmpcaService"
	"github.com/dfds/iam-anywhere-ninja/flags"
	"github.com/spf13/cobra"
)

func IsValidReason(revocationReason string) error {
	validReasons := []string{
		"UNSPECIFIED",
		"KEY_COMPROMISE",
		"CERTIFICATE_AUTHORITY_COMPROMISE",
		"AFFILIATION_CHANGED",
		"SUPERSEDED",
		"CESSATION_OF_OPERATION",
		"PRIVILEGE_WITHDRAWN",
		"A_A_COMPROMISE",
	}
	result := false
	for i := 0; i < len(validReasons); i++ {
		if validReasons[i] == revocationReason {
			result = true
			break
		}
	}

	if result == false {
		err := fmt.Errorf("%s is not a valid reason", revocationReason)
		return err
	}

	return nil
}

var revokeCertificateCmd = &cobra.Command{
	Use:   "revoke-certificate",
	Short: "revoke certificate",
	Long:  `Revokes the certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString(flags.ProfileName)
		certArn, _ := cmd.Flags().GetString(flags.CertificateArn)
		pcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
		revocationReason, _ := cmd.Flags().GetString(flags.RevocationReason)

		err := IsValidReason(revocationReason)
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
	revokeCertificateCmd.PersistentFlags().StringP(flags.RevocationReason, "r", "", "Reason why the certificate is revoked")

	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.RevocationReason)
}
