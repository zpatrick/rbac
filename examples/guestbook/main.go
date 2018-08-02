package main

import (
	"fmt"

	"github.com/zpatrick/rbac/examples/guestbook/permissions"
)

/*

comments: list, read, create, edit, delete
entry

read:entry

Guest, User, Admin
*/

func main() {
	admin := permissions.NewAdminRole()
	johnDoe := permissions.NewUserRole("John Doe", nil, nil)
	guest := permissions.NewGuestRole()

	fmt.Println(admin, johnDoe, guest)
}
