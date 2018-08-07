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
The [admin.go](/admin.go) file 
