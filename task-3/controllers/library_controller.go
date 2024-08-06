package controllers

import (
	"backend-learning-track/task-3/models"
	"backend-learning-track/task-3/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var memberID int
var bookID int

func displayMenu() {
	fmt.Println("Welcome to the Library Management System")
	fmt.Println("0. Create a Member")
	fmt.Println("1. Add a book")
	fmt.Println("2. Remove a book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List available books")
	fmt.Println("6. List borrowed books")
	fmt.Println("7. List Formated Library")
	fmt.Println("8. Exit")
}

func addBook(library *services.Library, reader bufio.Reader) {
	mu := sync.Mutex{}
	mu.Lock()
	bookID++
	defer mu.Unlock()

	fmt.Print("Enter the title of the book: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading title")
		addBook(library, reader)
	}
	title = strings.TrimSpace(title)
	fmt.Print("Enter the author of the book: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading author")
		addBook(library, reader)

	}
	author = strings.TrimSpace(author)
	newBook := models.Book{
		ID:     bookID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	newBookPointer := *&newBook
	library.AddBook(&newBookPointer)
	fmt.Println("Added Book")
	TakeInputAndDelegate(library)
}

func removeBook(library *services.Library, reader bufio.Reader) {
	fmt.Print("Input the Id of the Book: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading")
		removeBook(library, reader)
	}
	id = strings.TrimSpace(id)
	numberId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		removeBook(library, reader)
	}

	deleteError := library.RemoveBook(numberId)
	if deleteError != nil {
		fmt.Println(deleteError)
	}
	fmt.Println("Succesfully deleted")
	TakeInputAndDelegate(library)
}

func createMember(library *services.Library, reader bufio.Reader) {
	fmt.Print("Enter the name of Member: ")
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading")
		createMember(library, reader)
	}
	name = strings.TrimSpace(name)
	member := models.Member{
		ID:            memberID,
		Name:          name,
		BorrowedBooks: []*models.Book{},
	}
	library.AddMember(&member)
	fmt.Println("Member created with id ", memberID)
	memberID++
	TakeInputAndDelegate(library)
}

func borrowBook(library *services.Library, reader bufio.Reader) {
	fmt.Print("Enter the ID of the Book: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading")
		borrowBook(library, reader)
	}
	input = strings.TrimSpace(input)
	borrowBookID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting ID to string")
		borrowBook(library, reader)
	}
	fmt.Print("Enter the ID of the Member: ")

	mem, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading")
		borrowBook(library, reader)
	}
	mem = strings.TrimSpace(mem)
	borrowMemberId, err := strconv.Atoi(mem)
	if err != nil {
		fmt.Println("Error converting ID to string")
		borrowBook(library, reader)
	}

	err = library.BorrowBook(borrowBookID, borrowMemberId)
	if err != nil {
		fmt.Println(err)
		borrowBook(library, reader)
	}

	fmt.Println("Book Borrowed Succesfully")
	TakeInputAndDelegate(library)
}

func returnBook(library *services.Library, reader bufio.Reader) {
	fmt.Print("Enter the ID of the book: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading the input")
		returnBook(library, reader)
	}
	input = strings.TrimSpace(input)
	theBookId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("error converting the input")
		returnBook(library, reader)
	}
	fmt.Print("Enter the ID of the member: ")
	meminput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading the input")
		returnBook(library, reader)
	}
	meminput = strings.TrimSpace(meminput)
	theMemberId, err := strconv.Atoi(meminput)
	if err != nil {
		fmt.Print("error converting the input")
		returnBook(library, reader)
	}
	err = library.ReturnBook(theMemberId, theBookId)
	if err != nil {
		fmt.Println(err)
		returnBook(library, reader)
	}
	fmt.Println("Succesfully returned book")

	TakeInputAndDelegate(library)

}

func listAvailable(lib *services.Library) {
	books := lib.ListAvailableBooks()
	fmt.Print(lib.Format())
	var myBooks []models.Book
	for _, bookRef := range books {
		myBooks = append(myBooks, *bookRef)
	}
	fmt.Println(myBooks)
	TakeInputAndDelegate(lib)
}
func listBorrowedBooks(lib *services.Library) {
	books := lib.ListBorrowedBooks()
	var myBooks []models.Book
	for _, bookRef := range books {
		myBooks = append(myBooks, *bookRef)
	}
	fmt.Println(myBooks)
	TakeInputAndDelegate(lib)

}
func TakeInputAndDelegate(lib *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	displayMenu()
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	input = strings.TrimSpace(input)
	switch input {
	case "0":
		createMember(lib, *reader)
	case "1":
		addBook(lib, *reader)

	case "2":
		removeBook(lib, *reader)
	case "3":
		borrowBook(lib, *reader)
	case "4":
		returnBook(lib, *reader)
	case "5":
		listAvailable(lib)
	case "6":
		listBorrowedBooks(lib)
	case "7":
		fmt.Println(lib.Format())
		TakeInputAndDelegate(lib)
	case "8":
		return // Exit

	default:
		fmt.Print("Unknown Input")
		TakeInputAndDelegate(lib)
	}

}
