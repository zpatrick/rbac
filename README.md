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


There are 4 basic concepts in the `rbac` package:
* **Action**: An action is a string that represents some desired operation. 
Actions are typically expressed as a verb or a verb-object combination, but it is ultimately up to the user how actions are expressed. 
Some examples are: `"Upvote"`, `"ReadArticle"`, or `"EditComment"`. 
* **Target**: A target is a string that represents what the action is trying to operate on. 
Targets are typically expressed as an object's unique identifier, but it is ultimately up to the user how targets are expressed. 
An example is passing an `articleID` as the target for a `"ReadArticle"` action. 
Not all actions require a target. 
* **Permission**: A [permission](https://godoc.org/github.com/zpatrick/rbac#Permission) is a function that takes an action and a target, and returns true if and only if the action is allowed on the target. 
A permission should always allow (as opposed to deny) action(s) to be made on target(s), since nothing is allowed by default. 
* **Role**: A [role](https://godoc.org/github.com/zpatrick/rbac#Role) is essentially a grouping of permissions. 
The [`role.Can`](https://godoc.org/github.com/zpatrick/rbac#Role.Can) function should be used to determine whether or not a role can do an action on a target. 
A role is only allowed to do something if it has at least one permission that allows it. 

## License
This work is published under the MIT license.

Please see the `LICENSE` file for details.
