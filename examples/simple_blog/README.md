# Simple Blog Example
This example shows how one can use `rbac` to manage permissions for a simple blog application.
The application requires the following roles and permissions:
  * The **Guest** role can view and rate articles.
  * The **Admin** role can create, read, edit, delete, and rate articles. 
  
| Role  | Create Article | Read Article | Edit Article   | Delete Article | Rate Article |
|-------|----------------|--------------|----------------|----------------|--------------|
| Guest | -              | Allow        | -              | -              | Allow        |
| Admin | Allow          | Allow        | Allow          | Allow          | Allow        |

 
## Creating the Roles
The [roles.go](/examples/simple_blog/roles.go) file shows how one can implement this permission set.

#### Admin Role
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
The [rbac.NewGlobPermission](https://godoc.org/github.com/zpatrick/rbac#NewGlobPermission) function takes arguments: `actionPattern` and `targetPattern`. 
Then, it creates a permission that will return true if the requested action glob matches `actionPattern`, and if the requested target glob matches `targetPattern`. 
Since `*`is a wildcard in glob matching, we've created a permission that will return true for _any_ action on _any_ target; or, to put it shortly, the **Admin** role is allowed to do anything. 

In the following example, `admin.Can` will always return true no matter what we pass in as the `action` or the `target`:
```go
admin := NewAdminRole()
admin.Can("do some action", "to some target") // will always return true!
```
