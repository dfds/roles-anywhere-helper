package ProfileHandler

import (
	"testing"
)

func TestSetProfileName(t *testing.T) {
	testCases := []struct {
		name           string
		profileName    string
		expectedResult string
	}{
		{
			name:           "Empty profile name provided",
			profileName:    "",
			expectedResult: "default",
		},
		{
			name:           "Non-empty profile name provided",
			profileName:    "my-profile",
			expectedResult: "my-profile",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SetProfileName(tc.profileName)

			if result != tc.expectedResult {
				t.Errorf("Expected '%s' but got '%s'", tc.expectedResult, result)
			}
		})
	}
}
