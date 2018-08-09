# Simple Blog Example
This example shows how one can use `rbac` to manage permissions for a simple blog application.
The application requires the following roles and permissions:
  * The **Guest** role can view and rate any article.
  * The **Admin** role can create, read, edit, delete, and rate any article. 
  
| Role  | Create Article | Read Article | Edit Article   | Delete Article | Rate Article |
|-------|----------------|--------------|----------------|----------------|--------------|
| Guest | -              | Allow        | -              | -              | Allow        |
| Admin | Allow          | Allow        | Allow          | Allow          | Allow        |

 
## Creating the Roles
The [roles.go](/examples/blog_simple/roles.go) file shows how one can implement this permission set.

### Admin Role
Since the **Admin** role is allowed to do any action (`CreateArticle`, `ReadArticle`, `EditArticle`, `DeleteArticle`, and `RateArticle`), on any target (e.g. on any article), we can define that role's permissions in the following way:

```go
func NewAdminRole() rbac.Role {
        return rbac.Role{
                RoleID: "Admin",
                Permissions: []rbac.Permission{
                        rbac.NewGlobPermission("*", "*"),
                },
        }
}
```
The [rbac.NewGlobPermission](https://godoc.org/github.com/zpatrick/rbac#NewGlobPermission) function takes two arguments: `actionPattern` and `targetPattern`. 
Then, it creates a permission that will return true if the requested action matches `actionPattern`, and if the requested target matches `targetPattern`. 
Since `*`is a wildcard in glob matching, we've created a permission that will return true for _any_ action on _any_ target. 
To put it more simply: this permission allows the **Admin** role to do anything.

```go
admin := NewAdminRole()

// rbac.NewGlobPermission("*", "*") will cause this to return true since
// the "ReadArticle" action glob matches the "*" actionPattern in the permission
// and the "article_id" target glob matches the "*" targetPattern in the permission. 
admin.Can("ReadArticle", "article_id")

// rbac.NewGlobPermission("*", "*") will cause this to return true since
// the "DeleteArticle" action glob matches the "*" actionPattern in the permission
// and the "article_id" target glob matches the "*" targetPattern in the permission. 
admin.Can("DeleteArticle", "article_id")
```

### Guest Role
Since the **Guest** role is only allowed to do the `ReadArticle` and `RateArticle` actions on any target (e.g. on any article), we can define that role's permissions in the following way:

```go
ffunc NewGuestRole() rbac.Role {
        return rbac.Role{
                RoleID: "Guest",
                Permissions: []rbac.Permission{
                        rbac.NewGlobPermission("ReadArticle", "*"),
                        rbac.NewGlobPermission("RateArticle", "*"),
                },
        }
}
```
The first permission we define, `rbac.NewGlobPermission("ReadArticle", "*")`, allows the role to perform the `"ReadArticle"` action on `*` (any) target. 
To put it more simply: this permission allows the **Guest** role to read any article. 

The second permission we define, `rbac.NewGlobPermission("RateArticle", "*")`, allows the role to perform the `"RateArticle"` action on `*` (any) target. 
To put it more simply: this permission allows the **Guest** role to rate any article.  

```go
guest := NewGuestRole()

// rbac.NewGlobPermission("ReadArticle", "*") will cause this to return true since
// the "ReadArticle" action glob matches the "ReadArticle" actionPattern in the permission
// and the "article_id" target glob matches the "*" targetPattern in the permission. 
guest.Can("ReadArticle", "article_id") 

// this will return false beacause the guest role has no permissions 
// that match the "DeleteArticle" action
guest.Can("DeleteArticle", "article_id") 
```

## Try It Out
You can run this program yourself to view the permission with the following commands:
```console
$ go run *.go
Role: Guest
Action              ArticleID           Allowed
-----------------------------------------------
CreateArticle       -                   false
ReadArticle         a1                  true
EditArticle         a1                  false
DeleteArticle       a1                  false
RateArticle         a1                  true
```

```console
$ go run *.go -role=admin
Role: Admin
Action              ArticleID           Allowed
-----------------------------------------------
CreateArticle       -                   true
ReadArticle         a1                  true
EditArticle         a1                  true
DeleteArticle       a1                  true
RateArticle         a1                  true
```
