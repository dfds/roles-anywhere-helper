package profileHandler

import (
	"fmt"
)

func SetProfileName(profileName string) string {
	fmt.Println("Checking profile name")
	if len(profileName) == 0 {
		profileName = "default"
		fmt.Println("No profile name provided setting to .... ")
	}
	fmt.Println(profileName)
	return profileName
}
