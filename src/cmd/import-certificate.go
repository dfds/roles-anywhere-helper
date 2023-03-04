package cmd

import (
	"github.com/dfds/iam-anywhere-ninja/Flags"
	"github.com/dfds/iam-anywhere-ninja/acmService"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "upload certificate to ACM",
	Long:  `Adds your certificates to ACM`,
	Run: func(cmd *cobra.Command, args []string) {
		acmService.ImportCertificate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.PersistentFlags().StringP(Flags.ProfileName, "n", "default", "Name of the profile to be used for access to the ACM")
	configureCmd.PersistentFlags().StringP(Flags.CertificateDirectory, "x", "", "Destination of the Certificate on the machine")
	configureCmd.PersistentFlags().StringP(Flags.PrivateKeyDirectory, "k", "", "Destination of the unincrypted private key on the machine")

	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.CertificateDirectory)
	cobra.MarkFlagRequired(configureCmd.PersistentFlags(), Flags.PrivateKeyDirectory)
}
