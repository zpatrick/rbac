package main

import (
	"fmt"

	"github.com/zpatrick/rbac"
)

func main() {
	iamStyleRoles := []rbac.Role{
		{
			RoleID: "Admin",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("*", "*"),
			},
		},
		{
			RoleID: "ReadOnly",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("read:*", "*"),
			},
		},
		{
			RoleID: "EC2Admin",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("*", "arn:aws:ec2:*"),
			},
		},
		{
			RoleID: "S3BucketReadOnly",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("read:*", "arn:aws:s3:::my_bucket*"),
			},
		},
	}

	fmt.Println(iamStyleRoles)
}
