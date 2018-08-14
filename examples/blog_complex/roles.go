package main

import (
	"fmt"

	"github.com/zpatrick/rbac"
)

// NewAdminRole returns a role with admin-level permissions
func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "*"),
		},
	}
}

// NewGuestRole returns a role with guest-level permissions
func NewGuestRole() rbac.Role {
	return rbac.Role{
		RoleID: "Guest",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("ReadArticle", "*"),
			rbac.NewGlobPermission("RateArticle", "*"),
		},
	}
}

// NewMemberRole returns a role with member-level permissions
func NewMemberRole(userID string) rbac.Role {
	return rbac.Role{
		RoleID: fmt.Sprintf("Member(%s)", userID),
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("CreateArticle", "*"),
			rbac.NewGlobPermission("ReadArticle", "*"),
			rbac.NewGlobPermission("RateArticle", "*"),
			rbac.NewPermission(rbac.GlobMatch("EditArticle"), ifArticleAuthor(userID)),
			rbac.NewPermission(rbac.GlobMatch("DeleteArticle"), ifArticleAuthor(userID)),
		},
	}
}

// ifArticleAuthor returns a matcher that will only return true if
// the article's author matches userID.
func ifArticleAuthor(userID string) rbac.Matcher {
	return func(target string) (bool, error) {
		for _, article := range Articles() {
			if article.ArticleID == target {
				return article.AuthorID == userID, nil
			}
		}

		return false, nil
	}
}
