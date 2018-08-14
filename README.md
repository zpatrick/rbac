# RBAC

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/rbac/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/rbac)](https://goreportcard.com/report/github.com/zpatrick/rbac)
[![Go Doc](https://godoc.org/github.com/zpatrick/rbac?status.svg)](https://godoc.org/github.com/zpatrick/rbac)

## Overview
RBAC is a package that makes it easy to implement Role Based Access Control (RBAC) models in Go applications. 

## Download
To download this package, run:
```
go get github.com/zpatrick/rbac
```

## Getting Started
This section will go over some of the basic concepts and an example of how to use `rbac` in an application.
For more advanced usage, please see the [examples](/examples) directory. 


* **Action**: An action is a string that represents some desired operation. 
Actions are typically expressed as a verb or a verb-object combination, but it is ultimately up to the user how actions are expressed. 
Some examples are: `"Upvote"`, `"ReadArticle"`, or `"EditComment"`. 
* **Target**: A target is a string that represents what the action is trying to operate on. 
Targets are typically expressed as an object's unique identifier, but it is ultimately up to the user how targets are expressed. 
An example is passing an `articleID` as the target for a `"ReadArticle"` action. 
Not all actions require a target. 
* **Matcher**: A [matcher](https://godoc.org/github.com/zpatrick/rbac#Matcher) is a function that returns a bool representing whether or not the target matches some pre-defined pattern.
This repo comes with some builtin matchers: 
[GlobMatch](https://godoc.org/github.com/zpatrick/rbac#GlobMatch), 
[RegexMatch](https://godoc.org/github.com/zpatrick/rbac#RegexMatch), 
and [StringMatch](https://godoc.org/github.com/zpatrick/rbac#StringMatch). 
Please see the [complex blog](/examples/blog_complex) example to see how one can implement custom matchers for their applications. 
* **Permission**: A [permission](https://godoc.org/github.com/zpatrick/rbac#Permission) is a function that takes an action and a target, and returns true if and only if the action is allowed on the target. 
A permission should always allow (as opposed to deny) action(s) to be made on target(s), since nothing is allowed by default. 
* **Role**: A [role](https://godoc.org/github.com/zpatrick/rbac#Role) is essentially a grouping of permissions. 
The [`role.Can`](https://godoc.org/github.com/zpatrick/rbac#Role.Can) function should be used to determine whether or not a role can do an action on a target. 
A role is only allowed to do something if it has at least one permission that allows it. 

## Usage
```go
package main

import (
        "fmt"

        "github.com/zpatrick/rbac"
)

func main() {
        roles := []rbac.Role{
                {
                        RoleID: "Adult",
                        Permissions: []rbac.Permission{
                                rbac.NewGlobPermission("watch", "*"),
                        },
                },
                {
                        RoleID: "Teenager",
                        Permissions: []rbac.Permission{
                                rbac.NewGlobPermission("watch", "pg-13"),
                                rbac.NewGlobPermission("watch", "g"),
                        },
                },
                {
                        RoleID: "Child",
                        Permissions: []rbac.Permission{
                                rbac.NewGlobPermission("watch", "g"),
                        },
                },
        }

        for _, role := range roles {
                fmt.Println("Role:", role.RoleID)
                for _, rating := range []string{"g", "pg-13", "r"} {
                        canWatch, _ := role.Can("watch", rating)
                        fmt.Printf("Can watch %s? %t\n", rating, canWatch)
                }
        }
}

```

Output:
```console
Role: Adult
Can watch g? true
Can watch pg-13? true
Can watch r? true
Role: Teenager
Can watch g? true
Can watch pg-13? true
Can watch r? false
Role: Child
Can watch g? true
Can watch pg-13? false
Can watch r? false
```

## License
This work is published under the MIT license.

Please see the `LICENSE` file for details.
