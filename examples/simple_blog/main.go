package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/zpatrick/rbac"
)

/*
$ go run main.go [-role=guest]
Current Role: "Guest"
Action		Target	Allowed?
CreateArticle	""	No
ReadArticle	a1	Yes
DeleteArticle	a1	No
EditArticle	a1	No
RateArticle	a1	Yes

$ go run main.go -role=admin
Current Role: "Admin"
Action          Target  Allowed?
CreateArticle   ""      No
ReadArticle     a1      Yes
DeleteArticle   a1      No
EditArticle     a1      No
RateArticle     a1      Yes
*/

func main() {
	roleName := flag.String("role", "guest", "the role to use")
	flag.Parse()

	// assign a role
	var role rbac.Role
	switch r := strings.ToLower(*roleName); r {
	case "guest":
		role = NewGuestRole()
	case "admin":
		role = NewAdminRole()
	default:
		log.Fatalf("Role '%s' not recognized. Only 'guest' or 'admin' may be used.", r)
	}

	// calculate role permissions
	canCreate, err := role.Can("CreateArticle", "")
	if err != nil {
		log.Fatal(err)
	}

	canRead, err := role.Can("ReadArticle", "a1")
	if err != nil {
		log.Fatal(err)
	}

	canEdit, err := role.Can("CreateArticle", "a1")
	if err != nil {
		log.Fatal(err)
	}

	canDelete, err := role.Can("DeleteArticle", "a1")
	if err != nil {
		log.Fatal(err)
	}

	canRate, err := role.Can("RateArticle", "a1")
	if err != nil {
		log.Fatal(err)
	}

	// print role permissions
	w := tabwriter.NewWriter(os.Stdout, 20, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Role: %s\n", role.RoleID)
	fmt.Fprintln(w, "Action\tTarget\tAllowed")
	fmt.Fprintln(w, "-----------------------------------------------")
	fmt.Fprintf(w, "CreateArticle\t-\t%t\n", canCreate)
	fmt.Fprintf(w, "ReadArticle\ta1\t%t\n", canRead)
	fmt.Fprintf(w, "EditArticle\ta1\t%t\n", canEdit)
	fmt.Fprintf(w, "DeleteArticle\ta1\t%t\n", canDelete)
	fmt.Fprintf(w, "RateArticle\ta1\t%t\n", canRate)
	w.Flush()
}
