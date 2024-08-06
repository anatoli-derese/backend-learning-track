package services

import (
	"backend-learning-track/task-3/models"
	"errors"
	"strconv"
)

type LibraryService interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks(id int) []models.Book
	AddMember(Member *models.Member)
	ListBorrowedBooks(memberID int) []models.Book
	Format() string
}

type Library struct {
	Books   map[int]*models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*models.Book),
		Members: make(map[int]*models.Member),
	}
}

func (lib *Library) AddMember(Member *models.Member) {
	id := Member.ID
	lib.Members[id] = Member
}
func (lib *Library) AddBook(book *models.Book) {

	id := book.ID

	lib.Books[id] = book
}

func (lib *Library) RemoveBook(bookID int) error {
	// Implement this method
	for id := range lib.Books {
		if id == bookID {
			delete(lib.Books, bookID)
			return nil
		}
	}
	return errors.New("book not found")

}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	// Implement this method
	book, ok := lib.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, ok := lib.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	if lib.Books[bookID].Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}
	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil

}

func (lib *Library) Format() string {
	fs := "Books " + "\n"
	for _, book := range lib.Books {
		fs += strconv.Itoa(book.ID) + " " + book.Title + " " + book.Author + " " + book.Status + "\n"
	}
	fs += "Members " + "\n"
	for _, member := range lib.Members {
		fs += member.Name + " " + "\n"
	}
	return fs
}

func (lib *Library) ReturnBook(memberID int, bookId int) error {
	member, ok := lib.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	book, ok := lib.Books[bookId]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book was already available")
	}
	book.Status = "Available"
	var newBooks []*models.Book

	for _, book := range member.BorrowedBooks {
		if book.ID != bookId {
			newBooks = append(newBooks, book)
		}
	}
	return nil

}

func (lib *Library) ListAvailableBooks() []*models.Book {
	// Implement this method
	var availableBooks []*models.Book
	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (lib *Library) ListBorrowedBooks() []*models.Book {
	var borrowedBooks []*models.Book
	for _, book := range lib.Books {
		if book.Status == "Borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	return borrowedBooks

}
