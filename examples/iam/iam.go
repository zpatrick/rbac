package iam

import (
	"fmt"

	"github.com/zpatrick/rbac"
)

func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "*"),
		},
	}
}

func NewReadOnlyRole() rbac.Role {
	return rbac.Role{
		RoleID: "Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("read:*", "*"),
		},
	}
}

func NewEC2AdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "EC2Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "arn:aws:ec2:*"),
		},
	}
}

func NewS3BucketReadOnlyRole(bucket string) rbac.Role {
	return rbac.Role{
		RoleID: "S3BucketReadOnly",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("read:*", fmt.Sprintf("arn:aws:s3:::%s*", bucket)),
		},
	}
}
