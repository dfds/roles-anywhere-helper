package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/acmpcaService"
	"github.com/dfds/iam-anywhere-ninja/flags"

	"github.com/spf13/cobra"
)

var getCertificateCmd = &cobra.Command{
	Use:   "get-certificate",
	Short: "Get certificate",
	Long:  `Get certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		acmpcaService.ImportCertificate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(getCertificateCmd)
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationalUnit, "u", "", "The organization unit for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.OrganizationName, "n", "", "The organization name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.CommonName, "c", "", "The common name for the X509 certificate")
	getCertificateCmd.PersistentFlags().StringP(flags.AcmpcaArn, "a", "", "Arn for the ACM PCA")
	getCertificateCmd.PersistentFlags().StringP(flags.ProfileName, "p", "", "Profile of the ACM PCA")

	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.CommonName)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.AcmpcaArn)
	cobra.MarkFlagRequired(getCertificateCmd.PersistentFlags(), flags.ProfileName)
}
