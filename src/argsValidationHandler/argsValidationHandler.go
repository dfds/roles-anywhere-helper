package argsValidationHandler

import (
	"fmt"
	"github.com/dfds/iam-anywhere-ninja/revocationReasons"
)

func IsValidRevocationReason(revocationReason string) error {
	validReasons := []string{
		revocationReasons.Unspecified,
		revocationReasons.KeyCompromise,
		revocationReasons.CertificateAuthorityCompromise,
		revocationReasons.AffiliationChanged,
		revocationReasons.Superseded,
		revocationReasons.CessationOfOperation,
		revocationReasons.PrivilegeWithdrawn,
		revocationReasons.AACompromise
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
