package main

import (
	"fmt"
	"time"
)

// ========== COMPOSITION LEARNING EXAMPLES - TYPE DEFINITIONS ==========

// Example 1: Basic Embedding
type Address struct {
	street string
	city   string
	zip    string
}

type Person struct {
	name    string
	Address // Anonymous field (embedded)
}

// Example 2: Method Promotion
type Engine struct {
	power int
}

func (e *Engine) Start() {
	fmt.Printf("  🏎️  Engine started (power: %d)\n", e.power)
}

type Car struct {
	make   string
	Engine // Embedded
}

// Example 3: Multiple Embedding
type Tracked struct {
	createdAt time.Time
	updatedAt time.Time
}

type Audited struct {
	createdBy string
	updatedBy string
}

type Document struct {
	title   string
	Tracked // Embedding 1
	Audited // Embedding 2
}

// Example 4: Field Shadowing
type Base struct {
	name string
}

type Derived struct {
	name string // Shadows Base.name
	Base
}

// Example 5: Embedding Interfaces
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader // Embedded interface
	Writer // Embedded interface
}

// Example 6: Code Reuse
type Logger struct{}

func (l *Logger) Log(msg string) {
	fmt.Printf("  📝 LOG: %s\n", msg)
}

type Service struct {
	name   string
	Logger // Embedded
}

type Database struct {
	connStr string
	Logger  // Embedded
}

// Example 7: Polymorphism
type Vehicle interface {
	Drive()
}

type Motorcycle struct {
	Engine
}

func (c *Car) Drive() {
	fmt.Println("  🚗 Car driving")
}

func (m *Motorcycle) Drive() {
	fmt.Println("  🏍️  Motorcycle driving")
}

// Example 9: Encapsulation
type Money struct {
	amount int
}

func (m Money) Add(other Money) Money {
	return Money{amount: m.amount + other.amount}
}

type BankAccount struct {
	accountID string
	balance   Money
}

// ========== EXAMPLE FUNCTIONS ==========

// 1. BASIC EMBEDDING: Embed struct into another struct
// Composition: Include one struct inside another
// Go's way of sharing behavior (not inheritance)
func example1_BasicEmbedding() {
	fmt.Println("1. Basic Embedding:")

	p := Person{
		name: "Alice",
		Address: Address{
			street: "123 Main St",
			city:   "NYC",
			zip:    "10001",
		},
	}

	fmt.Printf("  Name: %s\n", p.name)
	fmt.Printf("  City: %s (accessed from embedded Address)\n", p.city) // Direct access!
}

// 2. METHOD PROMOTION: Embedded type's methods become available
func example2_MethodPromotion() {
	fmt.Println("2. Method Promotion:")

	car := &Car{
		make: "Tesla",
		Engine: Engine{
			power: 400,
		},
	}

	fmt.Printf("  Car: %s\n", car.make)
	car.Start() // Method promoted from Engine!
}

// 3. MULTIPLE EMBEDDING: Embed multiple structs
func example3_MultipleEmbedding() {
	fmt.Println("3. Multiple Embedding:")

	doc := Document{
		title: "Report",
		Tracked: Tracked{
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		Audited: Audited{
			createdBy: "Alice",
			updatedBy: "Bob",
		},
	}

	fmt.Printf("  Title: %s\n", doc.title)
	fmt.Printf("  Created by: %s at %s\n", doc.createdBy, doc.createdAt.Format("2006-01-02"))
}

// 4. FIELD SHADOWING: Override embedded field names
func example4_FieldShadowing() {
	fmt.Println("4. Field Shadowing:")

	d := Derived{
		name: "Derived name",
		Base: Base{
			name: "Base name",
		},
	}

	fmt.Printf("  d.name: %s (Derived's own field)\n", d.name)
	fmt.Printf("  d.Base.name: %s (accesses Base's field)\n", d.Base.name)
}

// 5. EMBEDDING INTERFACES: Interface embedding for composition
func example5_EmbeddingInterfaces() {
	fmt.Println("5. Embedding Interfaces:")

	fmt.Println("  Embedded interfaces combine multiple capabilities")
	fmt.Println("  ReadWriter includes both Read() and Write() methods")
}

// 6. STRUCT EMBEDDING FOR CODE REUSE: Share common functionality
func example6_CodeReuse() {
	fmt.Println("6. Struct Embedding for Code Reuse:")

	svc := Service{name: "UserService"}
	svc.Log("Service started") // Embedded method available

	db := Database{connStr: "postgres://..."}
	db.Log("Database connected") // Same method available
}

// 7. POLYMORPHISM WITH COMPOSITION: Treat different types uniformly
func example7_Polymorphism() {
	fmt.Println("7. Polymorphism with Composition:")

	var vehicles []Vehicle
	vehicles = append(vehicles, &Car{
		make:   "Tesla",
		Engine: Engine{power: 400},
	})
	vehicles = append(vehicles, &Motorcycle{
		Engine: Engine{power: 150},
	})

	for _, v := range vehicles {
		v.Drive() // Polymorphic behavior
	}
}

// 8. COMPOSITION vs INHERITANCE: Why Go uses composition
func example8_CompositionPattern() {
	fmt.Println("8. Composition vs Inheritance:")
	fmt.Println("  Go philosophy: Composition over Inheritance")
	fmt.Println("  Why:")
	fmt.Println("    • Simpler code structure")
	fmt.Println("    • Easier to test")
	fmt.Println("    • Avoids inheritance diamond problem")
	fmt.Println("    • Explicit over implicit")
	fmt.Println("    • Mix and match behaviors easily")
}

// 9. ENCAPSULATION WITH COMPOSITION: Bundle data and methods
func example9_Encapsulation() {
	fmt.Println("9. Encapsulation with Composition:")

	acc := BankAccount{
		accountID: "ACC-001",
		balance:   Money{amount: 1000},
	}

	newBalance := acc.balance.Add(Money{amount: 500})
	fmt.Printf("  Account: %s\n", acc.accountID)
	fmt.Printf("  New balance: %d\n", newBalance.amount)
}

// 10. BEST PRACTICES: Summary of composition patterns
func example10_BestPractices() {
	fmt.Println("10. Best Practices:")
	fmt.Println("  ✓ Use composition to share behavior")
	fmt.Println("  ✓ Embed interface types for flexibility")
	fmt.Println("  ✓ Keep embedded type names short or anonymous")
	fmt.Println("  ✓ Avoid multiple embedding of same type")
	fmt.Println("  ✓ Document what each embedded type provides")
	fmt.Println("  ✓ Use field names for clarity when needed")
	fmt.Println("  ✓ Prefer explicit over implicit method promotion")
	fmt.Println("  ✓ Consider interfaces for defining behavior")
}

func main() {
	fmt.Println("========== LEARNING COMPOSITION ==========\n")

	// 1. Basic embedding
	fmt.Println("--- 1. Basic Embedding ---")
	example1_BasicEmbedding()

	// 2. Method promotion
	fmt.Println("\n--- 2. Method Promotion ---")
	example2_MethodPromotion()

	// 3. Multiple embedding
	fmt.Println("\n--- 3. Multiple Embedding ---")
	example3_MultipleEmbedding()

	// 4. Field shadowing
	fmt.Println("\n--- 4. Field Shadowing ---")
	example4_FieldShadowing()

	// 5. Embedding interfaces
	fmt.Println("\n--- 5. Embedding Interfaces ---")
	example5_EmbeddingInterfaces()

	// 6. Code reuse
	fmt.Println("\n--- 6. Code Reuse ---")
	example6_CodeReuse()

	// 7. Polymorphism
	fmt.Println("\n--- 7. Polymorphism ---")
	example7_Polymorphism()

	// 8. Composition pattern
	fmt.Println("\n--- 8. Composition Pattern ---")
	example8_CompositionPattern()

	// 9. Encapsulation
	fmt.Println("\n--- 9. Encapsulation ---")
	example9_Encapsulation()

	// 10. Best practices
	fmt.Println("\n--- 10. Best Practices ---")
	example10_BestPractices()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Online Store Product System ==========\n")
	productStoreProject()
}

// ========== PROJECT: ONLINE STORE PRODUCT SYSTEM ==========
// Real-world scenario: E-commerce system with different product types
// Shows: Composition of Product with type-specific attributes
// Demonstrates: Embedding common properties, polymorphism, flexible design

type Product struct {
	id       int
	name     string
	price    float64
	quantity int
}

func (p Product) Info() string {
	return fmt.Sprintf("[%d] %s - $%.2f (Stock: %d)", p.id, p.name, p.price, p.quantity)
}

type Electronics struct {
	Product
	warranty int // months
	brand    string
}

type Book struct {
	Product
	author string
	isbn   string
	pages  int
}

type Clothing struct {
	Product
	size     string
	color    string
	material string
}

func (e Electronics) Details() string {
	return fmt.Sprintf("%s | Brand: %s | Warranty: %d months", e.Info(), e.brand, e.warranty)
}

func (b Book) Details() string {
	return fmt.Sprintf("%s | Author: %s | ISBN: %s | Pages: %d", b.Info(), b.author, b.isbn, b.pages)
}

func (c Clothing) Details() string {
	return fmt.Sprintf("%s | Size: %s | Color: %s | Material: %s", c.Info(), c.size, c.color, c.material)
}

func productStoreProject() {
	fmt.Println("🏪 Online Store Product System - Composition Example\n")

	// Create products using composition
	products := map[string]interface{}{
		"E001": Electronics{
			Product: Product{
				id:       1,
				name:     "Laptop",
				price:    999.99,
				quantity: 5,
			},
			warranty: 24,
			brand:    "Dell",
		},
		"B001": Book{
			Product: Product{
				id:       2,
				name:     "Go Programming",
				price:    39.99,
				quantity: 20,
			},
			author: "Jon Bodner",
			isbn:   "978-1492067542",
			pages:  624,
		},
		"C001": Clothing{
			Product: Product{
				id:       3,
				name:     "T-Shirt",
				price:    24.99,
				quantity: 50,
			},
			size:     "M",
			color:    "Blue",
			material: "Cotton",
		},
	}

	for {
		fmt.Print("\nCommand (list/search/add-to-cart/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "list":
			fmt.Println("\n📦 Available Products:")
			for id, prod := range products {
				switch p := prod.(type) {
				case Electronics:
					fmt.Printf("  [%s] %s\n", id, p.Details())
				case Book:
					fmt.Printf("  [%s] %s\n", id, p.Details())
				case Clothing:
					fmt.Printf("  [%s] %s\n", id, p.Details())
				}
			}

		case "search":
			fmt.Print("Search by product ID (E001/B001/C001): ")
			var searchID string
			fmt.Scanln(&searchID)

			if prod, exists := products[searchID]; exists {
				fmt.Println("\n✅ Found:")
				switch p := prod.(type) {
				case Electronics:
					fmt.Printf("  %s\n", p.Details())
					fmt.Printf("  Technical Specs:\n")
					fmt.Printf("    - Brand: %s\n", p.brand)
					fmt.Printf("    - Warranty: %d months\n", p.warranty)
				case Book:
					fmt.Printf("  %s\n", p.Details())
					fmt.Printf("  Book Details:\n")
					fmt.Printf("    - Author: %s\n", p.author)
					fmt.Printf("    - Total Pages: %d\n", p.pages)
				case Clothing:
					fmt.Printf("  %s\n", p.Details())
					fmt.Printf("  Fit Details:\n")
					fmt.Printf("    - Size: %s\n", p.size)
					fmt.Printf("    - Material: %s\n", p.material)
				}
			} else {
				fmt.Println("❌ Product not found")
			}

		case "add-to-cart":
			fmt.Print("Product ID: ")
			var prodID string
			fmt.Scanln(&prodID)

			fmt.Print("Quantity: ")
			var qty int
			fmt.Scanln(&qty)

			if prod, exists := products[prodID]; exists {
				var baseProd Product
				var typename string

				switch p := prod.(type) {
				case Electronics:
					baseProd = p.Product
					typename = "Electronics"
				case Book:
					baseProd = p.Product
					typename = "Book"
				case Clothing:
					baseProd = p.Product
					typename = "Clothing"
				}

				if qty > baseProd.quantity {
					fmt.Printf("❌ Only %d in stock\n", baseProd.quantity)
				} else {
					total := float64(qty) * baseProd.price
					fmt.Printf("✅ Added %d x %s (%s) to cart\n", qty, baseProd.name, typename)
					fmt.Printf("   Total: $%.2f\n", total)

					// Update stock
					switch p := prod.(type) {
					case Electronics:
						p.quantity -= qty
						products[prodID] = p
					case Book:
						p.quantity -= qty
						products[prodID] = p
					case Clothing:
						p.quantity -= qty
						products[prodID] = p
					}
				}
			} else {
				fmt.Println("❌ Product not found")
			}

		case "quit":
			fmt.Println("Thank you for shopping!")
			return

		default:
			fmt.Println("Unknown command. Use: list, search, add-to-cart, quit")
		}
	}
}
