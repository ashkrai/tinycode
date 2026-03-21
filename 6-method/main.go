package main

import (
	"fmt"
	"strings"
)

// 1. Basic type definition for methods
type Circle struct {
	radius float64
}

// 2. Method with value receiver (cannot modify receiver)
// Value receiver: receives a COPY of the struct
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// 3. Method with value receiver returning string
func (c Circle) Describe() string {
	return fmt.Sprintf("Circle with radius %.2f", c.radius)
}

// 4. Method with pointer receiver (can modify receiver)
// Pointer receiver: receives the ADDRESS of the struct
func (c *Circle) Scale(factor float64) {
	c.radius *= factor // modifies original
}

// 5. Multiple methods on same type
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// 6. Method chaining (returning receiver)
func (r *Rectangle) SetWidth(w float64) *Rectangle {
	r.width = w
	return r // return self for chaining
}

func (r *Rectangle) SetHeight(h float64) *Rectangle {
	r.height = h
	return r
}

// 7. Methods on built-in types (via type alias)
type MyString string

// String method - special "ToString" method
func (s MyString) UpperCase() string {
	return strings.ToUpper(string(s))
}

func (s MyString) Length() int {
	return len(s)
}

// 8. Method on slice type
type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, val := range s {
		total += val
	}
	return total
}

func (s IntSlice) Average() float64 {
	if len(s) == 0 {
		return 0
	}
	return float64(s.Sum()) / float64(len(s))
}

// 9. Methods with receivers and return values
type Counter struct {
	value int
}

func (c *Counter) Increment() int {
	c.value++
	return c.value
}

func (c *Counter) GetValue() int {
	return c.value
}

// 10. Value receiver vs Pointer receiver explanation
type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	// Value receiver - works with copy
	return fmt.Sprintf("%s (%d years old)", p.name, p.age)
}

func (p *Person) HaveBirthday() {
	// Pointer receiver - modifies original
	p.age++
}

func main() {
	// 1. Basic method call with value receiver
	circle := Circle{radius: 5}
	fmt.Println("Circle Area:", circle.Area())
	fmt.Println("Description:", circle.Describe())

	// 2. Value receiver doesn't modify original
	circle.radius = 3
	fmt.Println("After reassign:", circle.Area())

	// 3. Pointer receiver modifies original
	circlePtr := &Circle{radius: 4}
	fmt.Println("Before scale:", circlePtr.Area())
	circlePtr.Scale(2) // modifies original
	fmt.Println("After scale (radius doubled):", circlePtr.Area())

	// 4. Multiple methods on same type
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Rectangle Area:", rect.Area())
	fmt.Println("Rectangle Perimeter:", rect.Perimeter())

	// 5. Method chaining (returns receiver)
	rect2 := &Rectangle{}
	rect2.SetWidth(15).SetHeight(8) // chain methods
	fmt.Printf("Chained: Area=%.0f, Perimeter=%.0f\n", float64(rect2.Area()), rect2.Perimeter())

	// 6. Methods on custom string type
	myStr := MyString("hello world")
	fmt.Println("Original:", myStr)
	fmt.Println("UpperCase:", myStr.UpperCase())
	fmt.Println("Length:", myStr.Length())

	// 7. Methods on slice type
	numbers := IntSlice{10, 20, 30, 40, 50}
	fmt.Println("Slice Sum:", numbers.Sum())
	fmt.Printf("Slice Average: %.2f\n", numbers.Average())

	// 8. Counter with pointer receiver
	counter := &Counter{}
	fmt.Println("Increment 1:", counter.Increment())
	fmt.Println("Increment 2:", counter.Increment())
	fmt.Println("Current value:", counter.GetValue())

	// 9. Value receiver for display, pointer receiver for modify
	person := Person{"Alice", 25}
	fmt.Println("Person:", person.String())
	person.HaveBirthday()
	fmt.Println("After birthday:", person.String())

	// PROJECT: Library Book Management System
	fmt.Println("\n--- PROJECT: Library Book Management System ---")
	libraryManagementSystem()
}

// PROJECT: Library management with book methods and operations
type Book struct {
	id       int
	title    string
	author   string
	price    float64
	quantity int
	rating   float64 // 1-5 stars
}

// Methods for Book

// 1. Method with value receiver - get book info
func (b Book) Info() string {
	return fmt.Sprintf("📖 %s by %s (Rating: %.1f⭐)", b.title, b.author, b.rating)
}

// 2. Method with pointer receiver - update price
func (b *Book) UpdatePrice(newPrice float64) {
	b.price = newPrice
	fmt.Printf("   Price updated to $%.2f\n", b.price)
}

// 3. Method with pointer receiver - add stock
func (b *Book) AddStock(count int) int {
	b.quantity += count
	return b.quantity
}

// 4. Method with pointer receiver - remove stock
func (b *Book) BorrowBook() bool {
	if b.quantity > 0 {
		b.quantity--
		return true
	}
	return false
}

// 5. Method with pointer receiver - return book and rate it
func (b *Book) ReturnAndRate(rating float64) {
	b.quantity++
	if rating > 0 && rating <= 5 {
		b.rating = rating
	}
}

// 6. Method - check availability
func (b Book) IsAvailable() bool {
	return b.quantity > 0
}

// 7. Method - get book price
func (b Book) GetPrice() float64 {
	return b.price
}

// Library struct
type Library struct {
	name   string
	books  map[int]*Book
	nextID int
}

func libraryManagementSystem() {
	// Create library
	lib := Library{
		name:   "City Library",
		books:  make(map[int]*Book),
		nextID: 1,
	}

	// Add initial books
	lib.books[1] = &Book{
		id:       1,
		title:    "Go Programming",
		author:   "John Doe",
		price:    35.99,
		quantity: 5,
		rating:   4.5,
	}
	lib.books[2] = &Book{
		id:       2,
		title:    "The Go Language",
		author:   "Rob Pike",
		price:    42.50,
		quantity: 3,
		rating:   4.8,
	}
	lib.books[3] = &Book{
		id:       3,
		title:    "Concurrency in Go",
		author:   "Katherine Cox",
		price:    45.00,
		quantity: 0,
		rating:   4.7,
	}
	lib.nextID = 4

	fmt.Printf("=== %s ===\n", lib.name)
	fmt.Println("Commands: list, info, borrow, return, update, add, quit\n")

	for {
		fmt.Print("Command: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "list":
			if len(lib.books) == 0 {
				fmt.Println("No books in library")
			} else {
				fmt.Println("\n--- Books Available ---")
				for id, book := range lib.books {
					status := "✅ Available"
					if !book.IsAvailable() {
						status = "❌ Out of Stock"
					}
					fmt.Printf("ID: %d | %s | Qty: %d | Price: $%.2f | %s\n",
						id, book.Info(), book.quantity, book.price, status)
				}
				fmt.Println()
			}

		case "info":
			fmt.Print("Book ID: ")
			var id int
			fmt.Scanln(&id)

			if book, exists := lib.books[id]; exists {
				fmt.Printf("\n📚 %s\n", book.Info())
				fmt.Printf("   Author: %s\n", book.author)
				fmt.Printf("   Price: $%.2f\n", book.price)
				fmt.Printf("   Available Copies: %d\n", book.quantity)
				if book.IsAvailable() {
					fmt.Println("   Status: ✅ Available to borrow")
				} else {
					fmt.Println("   Status: ❌ Out of Stock")
				}
				fmt.Println()
			} else {
				fmt.Println("❌ Book not found")
			}

		case "borrow":
			fmt.Print("Book ID to borrow: ")
			var id int
			fmt.Scanln(&id)

			if book, exists := lib.books[id]; exists {
				if book.BorrowBook() {
					// BorrowBook() method reduced quantity
					fmt.Printf("✓ %s borrowed! Remaining: %d copies\n", book.title, book.quantity)
				} else {
					fmt.Printf("❌ %s is out of stock\n", book.title)
				}
			} else {
				fmt.Println("❌ Book not found")
			}

		case "return":
			fmt.Print("Book ID to return: ")
			var id int
			fmt.Scanln(&id)

			if book, exists := lib.books[id]; exists {
				fmt.Print("Rate the book (1-5 stars): ")
				var rating float64
				fmt.Scanln(&rating)

				book.ReturnAndRate(rating)
				fmt.Printf("✓ Thanks for returning %s! (Rating: %.1f⭐)\n", book.title, book.rating)
			} else {
				fmt.Println("❌ Book not found")
			}

		case "update":
			fmt.Print("Book ID to update: ")
			var id int
			fmt.Scanln(&id)

			if book, exists := lib.books[id]; exists {
				fmt.Print("New price: ")
				var newPrice float64
				fmt.Scanln(&newPrice)

				book.UpdatePrice(newPrice)
			} else {
				fmt.Println("❌ Book not found")
			}

		case "add":
			fmt.Print("Book title: ")
			var title string
			fmt.Scanln(&title)

			fmt.Print("Author: ")
			var author string
			fmt.Scanln(&author)

			fmt.Print("Price: ")
			var price float64
			fmt.Scanln(&price)

			fmt.Print("Quantity: ")
			var qty int
			fmt.Scanln(&qty)

			lib.books[lib.nextID] = &Book{
				id:       lib.nextID,
				title:    title,
				author:   author,
				price:    price,
				quantity: qty,
				rating:   0,
			}
			fmt.Printf("✓ Book added with ID: %d\n", lib.nextID)
			lib.nextID++

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Try: list, info, borrow, return, update, add, quit")
		}
	}
}
