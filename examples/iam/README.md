# IAM Example
This examples shows how to use `rbac` to implement a permissions model similar to [AWS IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json) policies. 
The permission model makes use of the following patterns:
* Each **action** is a string with the following format: `"action_type:object_type"`, e.g. `"list:users"` or `"delete:comment"`. 
* Each **target** is a unique identifier for the specified `object_type` in the **action** string (where applicable). 
IAM uses their concept of [ARNs](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html), which works well in this model as ARNs typically contain metadata about the object they represent. 
A typical ARN has the following pattern: 
```
arn:aws:<service_name>:<region>:<account_id>:<object_type>:<object_id>
```

## Roles
This section contains some example `rbac.Role` objects that can be created by using these patterns. 

#### Administrator
This `rbac.Role` can do any action on any target. 

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

#### ReadOnly
This `rbac.Role` can do any `read` action on any target.

```go
func NewReadOnlyRole() rbac.Role {
        return rbac.Role{
                RoleID: "Admin",
                Permissions: []rbac.Permission{
                        rbac.NewGlobPermission("read:*", "*"),
                },
        }
}
```

#### EC2Admin
This `rbac.Role` can do any action as long as the target belongs to the `ec2` service. 

```go
func NewEC2AdminRole() rbac.Role {
        return rbac.Role{
                RoleID: "EC2Admin",
                Permissions: []rbac.Permission{
                        rbac.NewGlobPermission("*", "arn:aws:ec2:*"),
                },
        }
}
```

#### S3BucketReadOnly
This `rbac.Role` can do any `read` action as long as the target belongs to the specified S3 bucket. 

```go
func NewS3BucketReadOnlyRole(bucket string) rbac.Role {
        return rbac.Role{
                RoleID: "S3BucketReadOnly",
                Permissions: []rbac.Permission{
                        rbac.NewGlobPermission("read:*", fmt.Sprintf("arn:aws:s3:::%s*", bucket)),
                },
        }
}
```
