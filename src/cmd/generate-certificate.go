package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/certificateService"
	"github.com/dfds/iam-anywhere-ninja/flags"

	"github.com/spf13/cobra"
)

var generateCertificateCmd = &cobra.Command{
	Use:   "create-certificate",
	Short: "create RSA certificate",
	Long:  `Creates RSA certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		certificateService.Generate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(generateCertificateCmd)

	generateCertificateCmd.PersistentFlags().StringP(flags.CertificateDirectory, "x", "", "Directory of the Certificate on the machine")
	generateCertificateCmd.PersistentFlags().StringP(flags.PrivateKeyDirectory, "k", "", "Directory of the unincrypted private key on the machine")

	cobra.MarkFlagRequired(generateCertificateCmd.PersistentFlags(), flags.CertificateDirectory)
	cobra.MarkFlagRequired(generateCertificateCmd.PersistentFlags(), flags.PrivateKeyDirectory)
}
