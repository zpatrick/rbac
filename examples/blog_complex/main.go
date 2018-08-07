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

func main() {
	roleName := flag.String("role", "guest", "the role to use")
	articleID := flag.String("article", "a1", "the ID of an article")
	flag.Parse()

	// assign a role
	var role rbac.Role
	switch r := strings.ToLower(*roleName); r {
	case "guest":
		role = NewGuestRole()
	case "member":
		role = NewMemberRole("u1")
	case "admin":
		role = NewAdminRole()
	default:
		log.Fatalf("Role '%s' not recognized. Only 'guest', 'member', or 'admin' may be used.", r)
	}

	// calculate role permissions
	canCreate, err := role.Can("CreateArticle", "")
	if err != nil {
		log.Fatal(err)
	}

	canRead, err := role.Can("ReadArticle", *articleID)
	if err != nil {
		log.Fatal(err)
	}

	canEdit, err := role.Can("EditArticle", *articleID)
	if err != nil {
		log.Fatal(err)
	}

	canDelete, err := role.Can("DeleteArticle", *articleID)
	if err != nil {
		log.Fatal(err)
	}

	canRate, err := role.Can("RateArticle", *articleID)
	if err != nil {
		log.Fatal(err)
	}

	// print role permissions
	w := tabwriter.NewWriter(os.Stdout, 20, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Role: %s\n", role.RoleID)
	fmt.Fprintln(w, "Action\tTarget\tAllowed")
	fmt.Fprintln(w, "-----------------------------------------------")
	fmt.Fprintf(w, "CreateArticle\t-\t%t\n", canCreate)
	fmt.Fprintf(w, "ReadArticle\t%s\t%t\n", *articleID, canRead)
	fmt.Fprintf(w, "EditArticle\t%s\t%t\n", *articleID, canEdit)
	fmt.Fprintf(w, "DeleteArticle\t%s\t%t\n", *articleID, canDelete)
	fmt.Fprintf(w, "RateArticle\t%s\t%t\n", *articleID, canRate)
	w.Flush()
}
