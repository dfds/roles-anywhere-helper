package profileHandler

import (
	"fmt"
)

func SetProfileName(profileName string) string {
	if len(profileName) == 0 {
		profileName = "default"
		fmt.Println("No profile name provided setting to default ")
	}
	return profileName
}
