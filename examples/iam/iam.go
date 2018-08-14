package iam

import (
	"fmt"

	"github.com/zpatrick/rbac"
)

// NewAdminRole returns a rbac.Role that can do any action on any target.
func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "*"),
		},
	}
}

// NewReadOnlyRole returns a rbac.Role that can do any "read" action on any target.
func NewReadOnlyRole() rbac.Role {
	return rbac.Role{
		RoleID: "ReadOnly",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("read:*", "*"),
		},
	}
}

// NewEC2AdminRole returns a rbac.Role that can do any action
// as long as the target belongs to the "ec2" service.
func NewEC2AdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "EC2Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "arn:aws:ec2:*"),
		},
	}
}

// NewS3BucketReadOnlyRole returns a rbac.Role that can do any "read" action
// as long as the target belongs to the specified S3 bucket.
func NewS3BucketReadOnlyRole(bucket string) rbac.Role {
	return rbac.Role{
		RoleID: "S3BucketReadOnly",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("read:*", fmt.Sprintf("arn:aws:s3:::%s*", bucket)),
		},
	}
}
