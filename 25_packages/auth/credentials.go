package auth

import "fmt"

// is we  want to scope this function for outside file usage we need to give first letter as capitial
func LoginWithCredential(username string, pwd string) {
	fmt.Println("User login : ", username, " Password : ", pwd)
}
