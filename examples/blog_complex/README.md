# Complex Blog Example
Thsi is a continueation...

We add a new role, *Member*

| Role   | Create Article | Read Article | Edit Article   | Delete Article | Rate Article |
|--------|----------------|--------------|----------------|----------------|--------------|
| Guest  | -              | Allow        | -              | -              | Allow        |
| Member | Allow          | Allow        | IfAuthor       | IfAuthor       | Allow        |
| Admin  | Allow          | Allow        | Allow          | Allow          | Allow        |
