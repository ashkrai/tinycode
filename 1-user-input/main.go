package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Basic string input (single word only)
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	// 2. Read entire line including spaces (using bufio)
	fmt.Print("Enter your full address: ")
	address := readLine()
	fmt.Println("Your address:", address)

	// 3. Convert string to integer
	fmt.Print("Enter your age: ")
	ageStr := readLine()
	age, _ := strconv.Atoi(ageStr)
	fmt.Println("Next year you'll be:", age+1)

	// 4. Convert string to float
	fmt.Print("Enter price (e.g. 19.99): ")
	priceStr := readLine()
	price, _ := strconv.ParseFloat(priceStr, 64)
	fmt.Printf("Total with tax: %.2f\n", price*1.1)

	// 5. Split input by spaces
	fmt.Print("Enter two colors separated by space: ")
	input := readLine()
	colors := strings.Fields(input) // splits by any whitespace
	fmt.Println("First color:", colors[0], "Second color:", colors[1])

	// 6. Case-insensitive input check
	fmt.Print("Are you a student? (yes/no): ")
	answer := strings.ToLower(readLine())
	if answer == "yes" || answer == "y" {
		fmt.Println("You are a student!")
	} else {
		fmt.Println("You are not a student!")
	}

	// 7. Multiple inputs in a loop
	fmt.Println("\nEnter 3 favorite fruits:")
	var fruits []string
	for i := 1; i <= 3; i++ {
		fmt.Printf("Fruit %d: ", i)
		fruit := readLine()
		fruits = append(fruits, fruit)
	}
	fmt.Println("Your fruits:", fruits)

	// 8. Input validation with loop (retry on error)
	var quantity int
	for {
		fmt.Print("Enter quantity (must be positive): ")
		qtyStr := readLine()
		qty, err := strconv.Atoi(qtyStr)
		if err != nil || qty <= 0 {
			fmt.Println("Invalid! Try again.")
			continue
		}
		quantity = qty
		break
	}
	fmt.Println("Quantity set to:", quantity)

	// 9. Parse comma-separated values
	fmt.Print("Enter 3 numbers separated by commas: ")
	input = readLine()
	numberStrs := strings.Split(input, ",")
	var numbers []int
	for _, numStr := range numberStrs {
		num, _ := strconv.Atoi(strings.TrimSpace(numStr))
		numbers = append(numbers, num)
	}
	fmt.Println("Numbers:", numbers)

	// PROJECT: Simple Shopping List Manager
	fmt.Println("\n--- PROJECT: Shopping List Manager ---")
	shoppingListManager()
}

// PROJECT: Interactive shopping list with add, view, and calculate total
func shoppingListManager() {
	type Item struct {
		name     string
		price    float64
		quantity int
	}

	var items []Item
	total := 0.0

	fmt.Println("=== Shopping List Manager ===")

	for {
		fmt.Println("\nOptions: (1) Add Item  (2) View List  (3) Checkout  (4) Exit")
		fmt.Print("Choose option: ")
		choice := readLine()

		switch choice {
		case "1": // Add item
			fmt.Print("Item name: ")
			itemName := readLine()

			fmt.Print("Unit price: ")
			priceStr := readLine()
			price, _ := strconv.ParseFloat(priceStr, 64)

			fmt.Print("Quantity: ")
			qtyStr := readLine()
			qty, _ := strconv.Atoi(qtyStr)

			items = append(items, Item{itemName, price, qty})
			fmt.Println("✓ Item added!")

		case "2": // View list
			if len(items) == 0 {
				fmt.Println("Shopping list is empty!")
			} else {
				fmt.Println("\nYour Shopping List:")
				total = 0.0
				for i, item := range items {
					itemTotal := item.price * float64(item.quantity)
					total += itemTotal
					fmt.Printf("%d. %s - $%.2f x %d = $%.2f\n",
						i+1, item.name, item.price, item.quantity, itemTotal)
				}
				fmt.Printf("Subtotal: $%.2f\n", total)
			}

		case "3": // Checkout
			if len(items) == 0 {
				fmt.Println("Nothing to checkout!")
			} else {
				total = 0.0
				for _, item := range items {
					total += item.price * float64(item.quantity)
				}

				fmt.Print("Apply discount percentage (0-100): ")
				discStr := readLine()
				discount, _ := strconv.ParseFloat(discStr, 64)
				finalTotal := total * (1 - discount/100)

				fmt.Printf("\nSubtotal: $%.2f\n", total)
				fmt.Printf("Discount: %.0f%%\n", discount)
				fmt.Printf("Final Total: $%.2f\n", finalTotal)
				fmt.Println("Thank you for shopping!")
				break
			}

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option! Try again.")
		}
	}
}

// Helper function: Read entire line from input (includes spaces)
// More reliable than fmt.Scanln for full line input
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
