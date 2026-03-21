package main

import (
	"fmt"
	"sort"
)

// 1. Basic struct - group related data together
type Person struct {
	name  string
	age   int
	email string
}

// 2. Struct with multiple field types
type Product struct {
	id       int
	name     string
	price    float64
	quantity int
	inStock  bool
}

// 3. Anonymous struct (inline, no type name)
func demonstrateAnonymousStruct() {
	book := struct {
		title  string
		author string
		pages  int
	}{
		title:  "Go Programming",
		author: "John Doe",
		pages:  350,
	}
	fmt.Println("Anonymous struct:", book)
}

// 4. Struct with embedded struct (composition)
type Address struct {
	street string
	city   string
	zip    string
}

type Employee struct {
	name    string
	id      int
	Address // embedded (promoted fields)
}

// 5. Struct with methods
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 6. Struct with pointer receiver (can modify)
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

// 7. Struct literal - multiple ways to initialize
func demonstrateStructLiterals() {
	// Full initialization
	p1 := Person{"Alice", 30, "alice@email.com"}
	fmt.Println("Full init:", p1)

	// Named fields
	p2 := Person{
		name:  "Bob",
		age:   25,
		email: "bob@email.com",
	}
	fmt.Println("Named fields:", p2)

	// Empty struct
	p3 := Person{}
	fmt.Println("Empty struct:", p3)
}

// 8. Slice of structs
func demonstrateSliceOfStructs() {
	people := []Person{
		{"Alice", 30, "alice@email.com"},
		{"Bob", 25, "bob@email.com"},
		{"Charlie", 28, "charlie@email.com"},
	}

	fmt.Println("Slice of structs:")
	for _, p := range people {
		fmt.Printf("  %s (%d) - %s\n", p.name, p.age, p.email)
	}
}

// 9. Map of structs
func demonstrateMapOfStructs() {
	employees := map[string]Employee{
		"E001": {"Alice", 1, Address{"123 Main", "NYC", "10001"}},
		"E002": {"Bob", 2, Address{"456 Oak", "LA", "90001"}},
	}

	fmt.Println("Map of structs:")
	for id, emp := range employees {
		fmt.Printf("  %s: %s - %s, %s\n", id, emp.name, emp.city, emp.zip)
	}
}

// 10. Compare structs
func demonstrateStructComparison() {
	p1 := Person{"Alice", 30, "alice@email.com"}
	p2 := Person{"Alice", 30, "alice@email.com"}
	p3 := Person{"Bob", 25, "bob@email.com"}

	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false
}

func main() {
	// 1. Create and access struct
	person := Person{"Alice", 30, "alice@email.com"}
	fmt.Println("Struct:", person)
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)

	// 2. Modify struct fields
	person.age = 31
	fmt.Println("After modification:", person.age)

	// 3. Create struct with all fields
	product := Product{
		id:       1,
		name:     "Laptop",
		price:    999.99,
		quantity: 5,
		inStock:  true,
	}
	fmt.Println("Product:", product)

	// 4. Pointer to struct
	ptr := &person
	fmt.Println("Via pointer:", ptr.name) // auto-dereferences
	fmt.Println("Explicit dereference:", (*ptr).age)

	// 5. Embedded struct
	emp := Employee{
		name: "Bob",
		id:   101,
		Address: Address{
			street: "789 Pine",
			city:   "Chicago",
			zip:    "60601",
		},
	}
	fmt.Println("Employee:", emp.name, "in", emp.city) // city promoted field

	// 6. Struct methods
	rect := Rectangle{10, 5}
	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())

	rect.Scale(2) // pointer receiver
	fmt.Printf("After scale: Area: %.2f\n", rect.Area())

	// 7. Anonymous struct
	demonstrateAnonymousStruct()

	// 8. Struct literals
	demonstrateStructLiterals()

	// 9. Slice of structs
	demonstrateSliceOfStructs()

	// 10. Map of structs
	demonstrateMapOfStructs()

	// 11. Struct comparison
	demonstrateStructComparison()

	// PROJECT: Company Employee Management System
	fmt.Println("\n--- PROJECT: Company Employee Management System ---")
	employeeManagementSystem()
}

// PROJECT: Employee management system with multiple operations
func employeeManagementSystem() {
	// Type Company: stores employee records
	// - employees: map where KEY is Employee ID, VALUE is pointer to Employee
	// - nextID: auto-incremented ID for new employees
	type Company struct {
		name      string
		employees map[int]*Employee // Key: Employee ID, Value: Employee data
		nextID    int               // Next ID to assign (auto-increments)
	}

	company := Company{
		name:      "TechCorp",
		employees: make(map[int]*Employee),
		nextID:    1001, // First employee will get ID 1001
	}

	// Add initial employees with IDs 1001 and 1002
	company.employees[1001] = &Employee{
		name:    "Alice Johnson",
		id:      1001, // Unique Employee ID
		Address: Address{"123 Tech St", "San Francisco", "94107"},
	}
	company.employees[1002] = &Employee{
		name:    "Bob Smith",
		id:      1002, // Unique Employee ID
		Address: Address{"456 Dev Ave", "Seattle", "98101"},
	}

	fmt.Printf("=== %s Management System ===\n", company.name)
	fmt.Println("\n📌 Employee IDs: 1001, 1002")
	fmt.Println("Commands: add, view, list, delete, update, help, quit\n")

	for {
		fmt.Print("Command: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "add":
			fmt.Print("Name: ")
			var name string
			fmt.Scanln(&name)

			fmt.Print("Street: ")
			var street string
			fmt.Scanln(&street)

			fmt.Print("City: ")
			var city string
			fmt.Scanln(&city)

			fmt.Print("Zip: ")
			var zip string
			fmt.Scanln(&zip)

			company.employees[company.nextID] = &Employee{
				name:    name,
				id:      company.nextID,
				Address: Address{street, city, zip},
			}
			fmt.Printf("✓ Employee added with ID: %d\n", company.nextID)
			company.nextID++

		case "view":
			fmt.Print("Employee ID: ")
			var id int
			fmt.Scanln(&id)
			// ID is the key to look up employee in the map
			// Example: 1001, 1002, 1003, etc.

			if emp, exists := company.employees[id]; exists {
				fmt.Printf("\n--- Employee Details ---\n")
				fmt.Printf("ID: %d\n", emp.id)
				fmt.Printf("Name: %s\n", emp.name)
				fmt.Printf("Address: %s, %s %s\n", emp.street, emp.city, emp.zip)
			} else {
				fmt.Println("❌ Employee not found")
			}

		case "list":
			if len(company.employees) == 0 {
				fmt.Println("No employees in system")
			} else {
				fmt.Println("\n--- All Employees ---")
				// Sort by ID
				ids := make([]int, 0, len(company.employees))
				for id := range company.employees {
					ids = append(ids, id)
				}
				sort.Ints(ids)

				for _, id := range ids {
					emp := company.employees[id]
					fmt.Printf("📍 ID: %d | Name: %s | City: %s\n", emp.id, emp.name, emp.city)
				}
				fmt.Println("\n💡 Use these IDs for view, delete, or update commands")
			}

		case "delete":
			fmt.Print("Employee ID to delete: ")
			var id int
			fmt.Scanln(&id)

			if _, exists := company.employees[id]; exists {
				delete(company.employees, id)
				fmt.Println("✓ Employee deleted")
			} else {
				fmt.Println("❌ Employee not found")
			}

		case "update":
			fmt.Print("Employee ID: ")
			var id int
			fmt.Scanln(&id)

			emp, exists := company.employees[id]
			if !exists {
				fmt.Println("❌ Employee not found")
				break
			}

			fmt.Print("New name: ")
			var newName string
			fmt.Scanln(&newName)

			fmt.Print("New city: ")
			var newCity string
			fmt.Scanln(&newCity)

			emp.name = newName
			emp.city = newCity
			fmt.Println("✓ Employee updated")

		case "help":
			fmt.Println("\n=== HELP: Understanding Employee IDs ===")
			fmt.Println("📌 WHAT IS AN EMPLOYEE ID?")
			fmt.Println("   - Unique number assigned to each employee automatically")
			fmt.Println("   - First employee: ID 1001")
			fmt.Println("   - Second employee: ID 1002")
			fmt.Println("   - Third employee: ID 1003, etc.")
			fmt.Println("\n📋 HOW TO USE IDs:")
			fmt.Println("   1. 'list' command → Shows all employees with their IDs")
			fmt.Println("   2. 'view' command → Enter the ID to see full details")
			fmt.Println("   3. 'delete' command → Enter the ID to remove employee")
			fmt.Println("   4. 'update' command → Enter the ID to modify employee info")
			fmt.Println("\n📍 EXAMPLE:")
			fmt.Println("   Command: view")
			fmt.Println("   Employee ID: 1001 ← (enter this number)")
			fmt.Println()

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Try: add, view, list, delete, update, help, quit")
		}
	}
}
