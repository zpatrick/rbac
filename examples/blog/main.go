package main

import (
	"fmt"
	"log"

	"github.com/zpatrick/rbac"
)

// EXAMPLE: A simple book website.
// Guests can view books
// Authors can create, edit, and delete their own books
// Administrators can edit and delete any book

const (
	CreateBook = "CreateBook"
	ReadBook   = "ReadBook"
	UpdateBook = "UpdateBook"
	DeleteBook = "DeleteBook"
)

type User struct {
	UserID string
	Role   rbac.Role
}

type Book struct {
	Title   string
	Author  string
	Content string
}

func NewGuestRole() rbac.Role {
	return rbac.Role{
		RoleID: "Guest",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission(ReadBook, "*"),
		},
	}
}

func NewBookAuthorPermission(action, userID string) rbac.Permission {
	return func(a, target string) (bool, error) {
		if action != a {
			return false, nil
		}

		for _, book := range books {
			if book.Title == target {
				return book.Author == userID, nil
			}
		}

		return false, nil
	}
}

func NewAuthorRole(books []Book, userID string) rbac.Role {
	return rbac.Role{
		RoleID: "Author",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission(CreateBook, "*"),
			rbac.NewGlobPermission(ReadBook, "*"),
			NewBookAuthorPermission(UpdateBook, userID),
			NewBookAuthorPermission(DeleteBook, userID),
		},
	}
}

func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "Administrator",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission(CreateBook, "*"),
			rbac.NewGlobPermission(ReadBook, "*"),
			rbac.NewGlobPermission(UpdateBook, "*"),
			rbac.NewGlobPermission(DeleteBook, "*"),
		},
	}
}

var books = []Book{
	{
		Title:   "The Adventures of Tom Sawyer",
		Author:  "Mark Twain",
		Content: "Once upon a bye, there was a mischievous boy named Tom Sawyer...",
	},
	{
		Title:   "1984",
		Author:  "George Orwell",
		Content: "It was a bright day in April, and the clocks were striking thirteen....",
	},
	{
		Title:   "Old Man and the Sea",
		Author:  "Ernest Hemingway",
		Content: "He was an old man who fished alone in a skiff in the Gulf Stream...",
	},
}

var users = []User{
	{
		UserID: "Administrator",
		Role:   NewAdminRole(),
	},
	{
		UserID: "Mark Twain",
		Role:   NewAuthorRole(books, "Mark Twain"),
	},
	{
		UserID: "Ernest Hemingway",
		Role:   NewAuthorRole(books, "Ernest Hemingway"),
	},
	{
		UserID: "J. K. Rowling",
		Role:   NewAuthorRole(books, "J. K. Rowling"),
	},
	{
		UserID: "Guest",
		Role:   NewGuestRole(),
	},
}

func main() {
	for _, book := range books {
		fmt.Println("=======================================================")
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Println()
		fmt.Printf("User\t\t\tRead\tUpdate\tDelete\n")
		fmt.Println("-----------------------------------------------")
		for _, user := range users {
			read, err := user.Role.Can(ReadBook, book.Title)
			if err != nil {
				log.Fatal(err)
			}

			update, err := user.Role.Can(UpdateBook, book.Title)
			if err != nil {
				log.Fatal(err)
			}

			delete, err := user.Role.Can(DeleteBook, book.Title)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%-20s\t%t\t%t\t%t\n", user.UserID, read, update, delete)
		}
		fmt.Println()
	}
}
