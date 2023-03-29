package cmd

import (
	"fmt"

	"github.com/dfds/roles-anywhere-helper/acmpcaService"
	"github.com/dfds/roles-anywhere-helper/argsValidationHandler"
	"github.com/dfds/roles-anywhere-helper/flags"
	"github.com/dfds/roles-anywhere-helper/revocationReasons"
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
		region, _ := cmd.Flags().GetString(flags.PcaRegion)

		err := argsValidationHandler.IsValidRevocationReason(revocationReason)
		if err != nil {
			fmt.Println(err)
			return
		}
		acmpcaService.RevokeCertificate(profileName, certArn, pcaArn, revocationReason, region)
	},
}

func init() {
	rootCmd.AddCommand(revokeCertificateCmd)

	revokeCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "default", flags.ProfNameAcmPcaDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.CertificateArn, "c", "", flags.CertArnDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", flags.AcmPcaArnDesc)
	revokeCertificateCmd.PersistentFlags().StringP(flags.RevocationReason, "r", revocationReasons.Unspecified, flags.RevocReasonDesc)
	revokeCertificateCmd.PersistentFlags().String(flags.PcaRegion, "eu-east-1", flags.RegionNameAcmPcaDesc)

	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.CertificateArn)
	cobra.MarkFlagRequired(revokeCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
}
