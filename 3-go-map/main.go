package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. Declare an empty map (nil map)
	var person map[string]string
	fmt.Println("Empty map:", person, "len:", len(person))

	// 2. Initialize map with make
	scores := make(map[string]int)
	fmt.Println("Map created with make:", scores)

	// 3. Map literal - declare and initialize together
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println("Colors map:", colors)

	// 4. Add/Update key-value pairs
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Println("After adding:", scores)

	// 5. Access values by key
	fmt.Println("Alice's score:", scores["Alice"])
	fmt.Println("David's score:", scores["David"]) // returns 0 if key doesn't exist

	// 6. Check if key exists (ok idiom)
	value, exists := scores["David"]
	fmt.Println("David exists:", exists, "Value:", value)
	if exists {
		fmt.Println("Found David with score:", value)
	} else {
		fmt.Println("David not found in map")
	}

	// 7. Update existing value
	scores["Alice"] = 98
	fmt.Println("Updated Alice's score:", scores)

	// 8. Delete a key
	delete(scores, "Bob")
	fmt.Println("After deleting Bob:", scores)

	// 9. Iterate over map (range)
	fmt.Println("\nIterating over scores:")
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// 10. Map length
	fmt.Println("Total entries:", len(scores))

	// 11. Map with different value types (integers)
	inventory := map[string]int{
		"apples":  15,
		"oranges": 8,
		"bananas": 12,
	}
	fmt.Println("Inventory:", inventory)

	// 12. Map with slice as value type
	hobbies := map[string][]string{
		"Alice":  {"reading", "gaming"},
		"Bob":    {"sports", "cooking"},
		"Claire": {"music", "painting"},
	}
	fmt.Println("Hobbies (map with slice):", hobbies)
	fmt.Println("Alice's hobbies:", hobbies["Alice"])

	// 13. Nested maps
	students := map[string]map[string]int{
		"Alice": {"Math": 90, "Science": 85},
		"Bob":   {"Math": 78, "Science": 92},
	}
	fmt.Println("Nested map - Alice's Math score:", students["Alice"]["Math"])

	// 14. Clear a map (delete all entries)
	temp := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", temp)
	for k := range temp {
		delete(temp, k)
	}
	fmt.Println("After clear:", temp)

	// PROJECT: Phone Book Contact Manager
	fmt.Println("\n--- PROJECT: Phone Book Contact Manager ---")
	phoneBookManager()
}

// PROJECT: Interactive phone book for storing and managing contacts
func phoneBookManager() {
	// Map: name -> Contact details
	type Contact struct {
		phone string
		email string
		city  string
	}
	contacts := make(map[string]Contact)

	// Pre-populated contacts
	contacts["Alice"] = Contact{"555-1001", "alice@email.com", "NYC"}
	contacts["Bob"] = Contact{"555-2002", "bob@email.com", "LA"}
	contacts["Charlie"] = Contact{"555-3003", "charlie@email.com", "Chicago"}

	fmt.Println("=== Phone Book Manager ===")
	fmt.Println("(Type 'list' to see all, 'add' to add contact, 'search' to find, 'delete' to remove, 'quit' to exit)")

	for {
		fmt.Print("\nCommand: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "list":
			if len(contacts) == 0 {
				fmt.Println("Phone book is empty")
			} else {
				fmt.Println("\n--- All Contacts ---")
				// Sort names for consistent output
				names := make([]string, 0, len(contacts))
				for name := range contacts {
					names = append(names, name)
				}
				sort.Strings(names)

				for _, name := range names {
					c := contacts[name]
					fmt.Printf("%s:\n", name)
					fmt.Printf("  Phone: %s\n", c.phone)
					fmt.Printf("  Email: %s\n", c.email)
					fmt.Printf("  City: %s\n", c.city)
				}
			}

		case "add":
			fmt.Print("Enter name: ")
			var name string
			fmt.Scanln(&name)

			fmt.Print("Enter phone: ")
			var phone string
			fmt.Scanln(&phone)

			fmt.Print("Enter email: ")
			var email string
			fmt.Scanln(&email)

			fmt.Print("Enter city: ")
			var city string
			fmt.Scanln(&city)

			contacts[name] = Contact{phone, email, city}
			fmt.Println("✓ Contact added!")

		case "search":
			fmt.Print("Enter name to search: ")
			var name string
			fmt.Scanln(&name)

			contact, exists := contacts[name]
			if exists {
				fmt.Printf("\n%s found:\n", name)
				fmt.Printf("  Phone: %s\n", contact.phone)
				fmt.Printf("  Email: %s\n", contact.email)
				fmt.Printf("  City: %s\n", contact.city)
			} else {
				fmt.Printf("❌ %s not found in phone book\n", name)
			}

		case "delete":
			fmt.Print("Enter name to delete: ")
			var name string
			fmt.Scanln(&name)

			_, exists := contacts[name]
			if exists {
				delete(contacts, name)
				fmt.Println("✓ Contact deleted!")
			} else {
				fmt.Printf("❌ %s not found\n", name)
			}

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Try: list, add, search, delete, quit")
		}
	}
}
