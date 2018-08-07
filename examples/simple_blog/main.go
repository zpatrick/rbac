package main

import (
	"fmt"

	"github.com/zpatrick/rbac"
)

type Article struct {
	ArticleID string
	Text      string
}

func printArticle(role rbac.Role, article Article) error {
	can, err := role.Can("ReadArticle", article.ArticleID)
	if err != nil {
		return err
	}

	if !can {
		return fmt.Errorf("role '%s' is not allowed to read article '%s'", role.RoleID, article.ArticleID)
	}

	fmt.Printf("[Role:%s] [Article:%s] %s\n", role.RoleID, article.ArticleID, article.Text)
	return nil
}

func main() {
	articles := []Article{
		{
			ArticleID: "welcome",
			Text:    "Welcome to...",
		},
		{
			ArticleID: "a123",
			Text:      "Five-star WR Blake Miller signs letter of intent...",
		},
		{
			ArticleID: "a456",
			Text:      "Late in the fourth quarter, senior quarterback Riley...",
		},
	}

	// the guest role is only allowed to read the 'welcome' article
	guest := rbac.Role{
		RoleID: "guest",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("ReadArticle", "welcome"),
		},
	}

	// the member role can read any article
	member := rbac.Role{
		RoleID: "member",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("ReadArticle", "*"),
		},
	}

	for _, role := range []rbac.Role{guest, member} {
		for _, article := range articles {
			if err := printArticle(role, article); err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
			}
		}
	}
}
