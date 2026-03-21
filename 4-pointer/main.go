package main

import "fmt"

// Account struct for bank account system (defined at package level for use in all functions)
type Account struct {
	holder  string
	balance float64
	pin     string
}

func main() {
	// 1. Declare a pointer variable
	var age int = 25
	var ptr *int // pointer to int (nil initially)
	ptr = &age   // & gets the address (reference) of variable
	fmt.Println("Age:", age)
	fmt.Println("Pointer ptr:", ptr)
	fmt.Println("Address of age:", &age)

	// 2. Dereference a pointer (access value it points to)
	fmt.Println("Dereferenced value (*ptr):", *ptr)

	// 3. Modify value through pointer
	*ptr = 30 // changes the original value
	fmt.Println("After modifying through pointer, age:", age)

	// 4. Pointer to different types
	name := "Alice"
	ptrName := &name
	fmt.Println("String pointer:", ptrName)
	fmt.Println("Dereferenced string:", *ptrName)

	// 5. Nil pointer (pointer that doesn't point to anything)
	var emptyPtr *int
	fmt.Println("Nil pointer:", emptyPtr)
	fmt.Println("Is nil?:", emptyPtr == nil)

	// 6. Pointer arithmetic (not allowed in Go like C, but can use indexing)
	nums := []int{10, 20, 30, 40}
	ptrNums := &nums[1] // pointer to second element
	fmt.Println("Pointing to nums[1]:", *ptrNums)

	// 7. Create pointer with new() function
	ptrInt := new(int)
	*ptrInt = 100
	fmt.Println("Pointer created with new():", *ptrInt)

	// 8. Function receives pointer parameter (pass by reference)
	value := 5
	fmt.Println("Before function (pass by reference):", value)
	modifyByReference(&value) // pass pointer
	fmt.Println("After function:", value)

	// 9. Function receives regular parameter (pass by value)
	value = 5
	fmt.Println("Before function (pass by value):", value)
	modifyByValue(value) // pass value
	fmt.Println("After function:", value)

	// 10. Pointer to struct
	type Person struct {
		name string
		age  int
	}
	person := Person{"Bob", 28}
	ptrPerson := &person
	fmt.Println("Struct via pointer:", ptrPerson)
	fmt.Println("Access struct field via pointer:", ptrPerson.name)

	// 11. Modify struct fields through pointer
	ptrPerson.age = 29 // automatically dereferences
	fmt.Println("Modified age through pointer:", person.age)

	// 12. Pointer to array
	arr := [3]string{"a", "b", "c"}
	ptrArr := &arr
	fmt.Println("Array pointer:", ptrArr)
	fmt.Println("First element via pointer:", ptrArr[0])

	// 13. Slice contains pointers (slice elements are pointers)
	var ptrSlice []*int
	x, y, z := 10, 20, 30
	ptrSlice = append(ptrSlice, &x, &y, &z)
	fmt.Println("Slice of pointers:")
	for i, p := range ptrSlice {
		fmt.Printf("  Index %d: %d\n", i, *p)
	}

	// 14. Multiple levels of pointers (pointer to pointer)
	num := 42
	ptr1 := &num  // pointer to int
	ptr2 := &ptr1 // pointer to pointer to int
	fmt.Println("Direct value:", num)
	fmt.Println("Pointer to pointer dereference:", **ptr2)

	// PROJECT: Bank Account System (with pointer-based operations)
	fmt.Println("\n--- PROJECT: Bank Account System ---")
	bankAccountSystem()
}

// Function that modifies value through pointer
func modifyByReference(ptr *int) {
	*ptr = *ptr * 2 // multiply by 2 and modify original
}

// Function that doesn't modify original (pass by value)
func modifyByValue(val int) {
	val = val * 2 // modifies local copy only
}

// PROJECT: Bank account system demonstrating pointer usage
func bankAccountSystem() {
	// Create accounts
	accounts := make(map[string]*Account)

	accounts["alice@bank"] = &Account{"Alice Brown", 5000, "1234"}
	accounts["bob@bank"] = &Account{"Bob Smith", 3500, "5678"}
	accounts["charlie@bank"] = &Account{"Charlie Davis", 7200, "9012"}

	fmt.Println("=== Bank Account Manager ===")
	fmt.Println("Commands: create, balance, deposit, withdraw, transfer, list, quit\n")

	for {
		fmt.Print("Command: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "create":
			fmt.Print("Account ID (e.g., username@bank): ")
			var accID string
			fmt.Scanln(&accID)

			if _, exists := accounts[accID]; exists {
				fmt.Println("❌ Account already exists")
				break
			}

			fmt.Print("Full name: ")
			var holder string
			fmt.Scanln(&holder)

			fmt.Print("Starting balance: ")
			var balance float64
			fmt.Scanln(&balance)

			fmt.Print("PIN: ")
			var pin string
			fmt.Scanln(&pin)

			// Create new account using pointer
			accounts[accID] = &Account{holder, balance, pin}
			fmt.Printf("✓ Account created! Account ID: %s\n", accID)

		case "balance":
			fmt.Print("Account ID: ")
			var accID string
			fmt.Scanln(&accID)

			acc, exists := accounts[accID]
			if !exists {
				fmt.Println("❌ Account not found")
				break
			}

			fmt.Print("Enter PIN: ")
			var pin string
			fmt.Scanln(&pin)

			if pin != acc.pin {
				fmt.Println("❌ Invalid PIN")
				break
			}

			fmt.Printf("Hello %s, Your balance: $%.2f\n", acc.holder, acc.balance)

		case "deposit":
			fmt.Print("Account ID: ")
			var accID string
			fmt.Scanln(&accID)

			acc, exists := accounts[accID]
			if !exists {
				fmt.Println("❌ Account not found")
				break
			}

			fmt.Print("Enter PIN: ")
			var pin string
			fmt.Scanln(&pin)

			if pin != acc.pin {
				fmt.Println("❌ Invalid PIN")
				break
			}

			fmt.Print("Amount to deposit: ")
			var amount float64
			fmt.Scanln(&amount)

			if amount > 0 {
				deposit(acc, amount) // pass pointer to modify
				fmt.Printf("✓ Deposited $%.2f. New balance: $%.2f\n", amount, acc.balance)
			} else {
				fmt.Println("❌ Invalid amount")
			}

		case "withdraw":
			fmt.Print("Account ID: ")
			var accID string
			fmt.Scanln(&accID)

			acc, exists := accounts[accID]
			if !exists {
				fmt.Println("❌ Account not found")
				break
			}

			fmt.Print("Enter PIN: ")
			var pin string
			fmt.Scanln(&pin)

			if pin != acc.pin {
				fmt.Println("❌ Invalid PIN")
				break
			}

			fmt.Print("Amount to withdraw: ")
			var amount float64
			fmt.Scanln(&amount)

			if withdraw(acc, amount) { // pass pointer to modify
				fmt.Printf("✓ Withdrawn $%.2f. New balance: $%.2f\n", amount, acc.balance)
			} else {
				fmt.Println("❌ Insufficient funds")
			}

		case "transfer":
			fmt.Print("From account ID: ")
			var fromID string
			fmt.Scanln(&fromID)

			fmt.Print("To account ID: ")
			var toID string
			fmt.Scanln(&toID)

			fromAcc, exists1 := accounts[fromID]
			toAcc, exists2 := accounts[toID]

			if !exists1 || !exists2 {
				fmt.Println("❌ One or both accounts not found")
				break
			}

			fmt.Print("Enter PIN for source account: ")
			var pin string
			fmt.Scanln(&pin)

			if pin != fromAcc.pin {
				fmt.Println("❌ Invalid PIN")
				break
			}

			fmt.Print("Amount to transfer: ")
			var amount float64
			fmt.Scanln(&amount)

			if withdraw(fromAcc, amount) {
				deposit(toAcc, amount)
				fmt.Printf("✓ Transferred $%.2f from %s to %s\n", amount, fromAcc.holder, toAcc.holder)
			} else {
				fmt.Println("❌ Insufficient funds")
			}

		case "list":
			fmt.Println("\n--- All Accounts ---")
			for id, acc := range accounts {
				fmt.Printf("%s: %s - Balance: $%.2f\n", id, acc.holder, acc.balance)
			}
			fmt.Println()

		case "quit":
			fmt.Println("Thank you for using Bank System. Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Try: create, balance, deposit, withdraw, transfer, list, quit")
		}
	}
}

// Deposit: modify account balance through pointer
func deposit(acc *Account, amount float64) {
	acc.balance += amount
}

// Withdraw: modify account balance through pointer
func withdraw(acc *Account, amount float64) bool {
	if acc.balance >= amount {
		acc.balance -= amount
		return true
	}
	return false
}
