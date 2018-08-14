# Complex Blog Example
This example is a continuation of the [simple blog](https://github.com/zpatrick/rbac/tree/master/examples/blog_simple) example.
Please view that example before continuing. 


In this example, we add a new role named **Member**.
This role is allowed to create, read, and rate any article.
It is also allowed to edit and delete articles that were authored by the user. 

| Role   | Create Article | Read Article | Edit Article   | Delete Article | Rate Article |
|--------|----------------|--------------|----------------|----------------|--------------|
| Guest  | -              | Allow        | -              | -              | Allow        |
| Member | Allow          | Allow        | IfAuthor       | IfAuthor       | Allow        |
| Admin  | Allow          | Allow        | Allow          | Allow          | Allow        |

## Creating the Roles
The [roles.go](/examples/blog_complex/roles.go) file shows how one can implement this permission set.
Please see the [simple blog](https://github.com/zpatrick/rbac/tree/master/examples/blog_simple) example for information on how to create the **Guest** and **Admin** roles. 

### Member Role
Since the **Member** role is allowed to create, read, and rate any article, we can define those permissions in the following way:

```go
func NewMemberRole() rbac.Role {
	return rbac.Role{
		RoleID: "Member",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("CreateArticle", "*"),
			rbac.NewGlobPermission("ReadArticle", "*"),
			rbac.NewGlobPermission("RateArticle", "*"),
		},
	}
}
```

This role is also allowed to edit and delete articles as long as the article was authored by the the user who is assuming the **Member** role. 
In order to implement this sort logic, we need to create a custom [Matcher](https://godoc.org/github.com/zpatrick/rbac#Matcher). 
A matcher is a function that returns a bool representing whether or not the target matches some pre-defined pattern.
In this context, we need a function that returns true if and only if the specified target (an article's ID) was authored by some specified user:
```go
// ifArticleAuthor returns a rbac.Matcher that will only return true if
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
```

Now, we can create a new permission that utilizes this Matcher:
```go
rbac.NewPermission(rbac.GlobMatch("EditArticle"), ifArticleAuthor(userID))
```
This creates a new permission that will only return true in the following circumstance:
* The specified action glob matches `"EditArticle"`
* The specified target matches an article's `ArticleID`, and that article's author matches the specified `userID`. 

We can put this all together to generate the final `NewMemberRole` function:
```go
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
```

```go
member := NewMemberRole("u1")

// rbac.NewGlobPermission("ReadArticle", "*") will cause this to always return true.
member.Can("ReadArticle", "a1")

// rbac.NewPermission(rbac.GlobMatch("DeleteArticle"), ifArticleAuthor(userID)) will cause 
// this to return true if and only if article "a1" exists, and that article's author is "u1".
member.Can("DeleteArticle", "a1")
```

## Try It Out
You can run this program yourself to view the permission with the following commands:
```console
$ go run *.go -role=member
Role: Member(u1)
Action              ArticleID           AuthorID            Allowed
-------------------------------------------------------------------
CreateArticle       -                   -                   true
ReadArticle         a1                  u1                  true
EditArticle         a1                  u1                  true
DeleteArticle       a1                  u1                  true
RateArticle         a1                  u1                  true
ReadArticle         a2                  u2                  true
EditArticle         a2                  u2                  false
DeleteArticle       a2                  u2                  false
RateArticle         a2                  u2                  true
ReadArticle         a3                  u3                  true
EditArticle         a3                  u3                  false
DeleteArticle       a3                  u3                  false
RateArticle         a3                  u3                  true
```








