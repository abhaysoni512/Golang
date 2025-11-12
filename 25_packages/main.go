// packages are	 used for code resuse , as we don't have name mangling here so we use package, code recompilation time prevent
// command to import package which we are using and not available in system

package main

import (
	"fmt"

	"github.com/abhaysoni512/mod_ex/auth"
	"github.com/abhaysoni512/mod_ex/myuser"
	"github.com/fatih/color"
)

func main()  {
	auth.LoginWithCredential("abhay","123456")
	fmt.Println( auth.GetSession())

	myuser := myuser.User{
		Email: "adffgg@email.com",
		Name: "abhay",

	}
	// fmt.Println(myuser)
	color.Red(myuser.Name)

	
}