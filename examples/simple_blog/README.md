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
The [roles.go](/examples/simple_blog/roles.go) file shows how one can implement this permission set.

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
Then, it creates a permission that will return true if the requested action glob matches `actionPattern`, and if the requested target glob matches `targetPattern`. 
Since `*`is a wildcard in glob matching, we've created a permission that will return true for _any_ action on _any_ target. 
To put it more simply: this permission allows the **Admin** role to do anything.

```go
admin := NewAdminRole()

// rbac.NewGlobPermission("*", "*") will cause this to return true since
// "ReadArticle" glob matches the first "*", and "article_id" glob matches the second "*"
admin.Can("ReadArticle", "article_id")

// rbac.NewGlobPermission("*", "*") will cause this to return true since
// "DeleteArticle" glob matches the first "*", and "article_id" glob matches the second "*"
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
// "ReadArticle" glob matches "ReadArticle", and "article_id" glob matches the second "*"
guest.Can("ReadArticle", "article_id") 

// this will return false because there are no permissions for this role
// that glob match "DeleteArticle"
guest.Can("DeleteArticle", "article_id") 
```
