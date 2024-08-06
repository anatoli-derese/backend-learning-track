    # Library Service Documentation

## Overview

The `LibraryService` interface and `Library` struct provide a solution for managing books and members in a library system.

## LibraryService Interface

The `LibraryService` interface defines the following methods:

- **AddBook(book models.Book)**
  - Adds a book to the library.

- **RemoveBook(bookID int)**
  - Removes a book from the library by its ID.

- **BorrowBook(bookID int, memberID int) error**
  - Allows a member to borrow a book.
  - Returns an error if the book or member is not found, or if the book is already borrowed.

- **ReturnBook(bookID int, memberID int) error**
  - Allows a member to return a borrowed book.
  - Returns an error if the book is not found, already available, or if the member is not found.

- **ListAvailableBooks() []*models.Book**
  - Lists all available books in the library.

- **AddMember(member *models.Member)**
  - Adds a member to the library.

- **ListBorrowedBooks() []*models.Book**
  - Lists all books currently borrowed by members.

- **Format() string**
  - Returns a formatted string representation of the library's books and members.

## Library Struct

### Fields

- **Books**: A map from integer IDs to pointers of `models.Book`.
- **Members**: A map from integer IDs to pointers of `models.Member`.

### Methods

#### NewLibrary

```go
func NewLibrary() *Library
