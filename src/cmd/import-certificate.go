package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/Flags"
	"github.com/dfds/iam-anywhere-ninja/acmService"

	"github.com/spf13/cobra"
)

var importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "upload certificate to ACM",
	Long:  `Adds your certificate to ACM`,
	Run: func(cmd *cobra.Command, args []string) {
		acmService.ImportCertificate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(importCertificateCmd)

	importCertificateCmd.PersistentFlags().StringP(Flags.ProfileName, "p", "default", "Name of the profile to be used for access to the ACM")
	importCertificateCmd.PersistentFlags().StringP(Flags.CertificateDirectory, "x", "", "Directory of the Certificate on the machine")
	importCertificateCmd.PersistentFlags().StringP(Flags.PrivateKeyDirectory, "k", "", "Directory of the unincrypted private key on the machine")

	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), Flags.CertificateDirectory)
	cobra.MarkFlagRequired(importCertificateCmd.PersistentFlags(), Flags.PrivateKeyDirectory)
}
