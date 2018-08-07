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
	flag.Parse()

	// assign a role
	var role rbac.Role
	switch r := strings.ToLower(*roleName); r {
	case "guest":
		role = NewGuestRole()
	case "member":
		role = NewMemberRole(Articles()[0].AuthorID)
	case "admin":
		role = NewAdminRole()
	default:
		log.Fatalf("Role '%s' not recognized. Only 'guest', 'member', or 'admin' may be used.", r)
	}

	// print role permissions
	w := tabwriter.NewWriter(os.Stdout, 20, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Role: %s\n", role.RoleID)
	fmt.Fprintln(w, "Action\tArticleID\tAuthorID\tAllowed")
	fmt.Fprintln(w, "-------------------------------------------------------------------")

	canCreate, err := role.Can("CreateArticle", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "CreateArticle\t-\t-\t%t\n", canCreate)

	for _, article := range Articles() {
		canRead, err := role.Can("ReadArticle", article.ArticleID)
		if err != nil {
			log.Fatal(err)
		}

		canEdit, err := role.Can("EditArticle", article.ArticleID)
		if err != nil {
			log.Fatal(err)
		}

		canDelete, err := role.Can("DeleteArticle", article.ArticleID)
		if err != nil {
			log.Fatal(err)
		}

		canRate, err := role.Can("RateArticle", article.ArticleID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "ReadArticle\t%s\t%s\t%t\n", article.ArticleID, article.AuthorID, canRead)
		fmt.Fprintf(w, "EditArticle\t%s\t%s\t%t\n", article.ArticleID, article.AuthorID, canEdit)
		fmt.Fprintf(w, "DeleteArticle\t%s\t%s\t%t\n", article.ArticleID, article.AuthorID, canDelete)
		fmt.Fprintf(w, "RateArticle\t%s\t%s\t%t\n", article.ArticleID, article.AuthorID, canRate)
		w.Flush()
	}
}
