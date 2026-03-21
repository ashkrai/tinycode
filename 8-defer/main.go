package main

import (
	"fmt"
	"time"
)

// ========== SIMPLE DEFER LEARNING EXAMPLES ==========

// 1. BASIC DEFER: Executes LAST, after function returns
func example1_BasicDefer() {
	fmt.Println("A. Main code executes first")
	defer fmt.Println("C. Deferred code executes LAST")
	fmt.Println("B. More main code")
}

// 2. MULTIPLE DEFERS: Execute in LIFO order (Last In, First Out - like a stack)
func example2_MultipleDefers() {
	defer fmt.Println("1. This executes LAST")
	defer fmt.Println("2. This executes second")
	defer fmt.Println("3. This executes first of deferred")
	fmt.Println("Main code")
}

// 3. DEFER WITH FUNCTIONS: Defer can call functions
func printMessage(msg string) {
	fmt.Println("Message:", msg)
}

func example3_DeferWithFunctions() {
	defer printMessage("Cleanup message")
	fmt.Println("Main logic")
}

// 4. DEFER FOR RESOURCE CLEANUP: Most common use case
// Simulating a database connection
type Database struct {
	name      string
	connected bool
}

func (db *Database) Connect() {
	db.connected = true
	fmt.Printf("🔗 Connected to database: %s\n", db.name)
}

func (db *Database) Disconnect() {
	if db.connected {
		db.connected = false
		fmt.Printf("🔌 Disconnected from database: %s\n", db.name)
	}
}

func example4_DeferCleanup() {
	db := &Database{name: "UserDB"}
	db.Connect()
	// DEFER ensures disconnect happens even if error occurs
	defer db.Disconnect()

	fmt.Println("✓ Performing database operations...")
	// Cleanup automatically happens here when function ends
}

// 5. DEFER USEFUL FOR FILE HANDLING
func example5_FileHandling() {
	// Simulate file operations
	fmt.Println("📂 Opening file...")
	defer fmt.Println("📂 Closing file...")
	fmt.Println("📝 Writing to file...")
}

// 6. DEFER WITH PANIC RECOVERY
func example6_PanicRecovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("✓ Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("This never executes")
}

// 7. DEFER TIMING: Arguments evaluated immediately but function call delayed
func example7_DeferTiming() {
	x := 10
	defer func(value int) {
		fmt.Printf("Deferred sees x as: %d (captured)\n", value)
	}(x)

	x = 20
	fmt.Printf("Current x: %d\n", x)
	// Deferred function sees x=10 because argument was evaluated immediately
}

// 8. DEFER IN LOOPS: Creates multiple defers (all execute in reverse)
func example8_DeferInLoop() {
	fmt.Println("Starting loop with defers:")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  Deferred: i=%d\n", i)
		fmt.Printf("Loop iteration: %d\n", i)
	}
	fmt.Println("Done - defers execute now in reverse order")
}

// 9. SIMULATE LOCK/UNLOCK WITH DEFER
type Mutex struct {
	locked bool
}

func (m *Mutex) Lock() {
	m.locked = true
	fmt.Println("🔒 Lock acquired")
}

func (m *Mutex) Unlock() {
	m.locked = false
	fmt.Println("🔓 Lock released")
}

func example9_LockUnlock() {
	lock := &Mutex{}
	lock.Lock()
	defer lock.Unlock() // Ensures unlock even if error happens

	fmt.Println("✓ Critical section - only one goroutine at a time")
}

// 10. PRACTICAL: Timing function execution with defer
func example10_Timing() {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("⏱️  Function took: %v\n", elapsed)
	}()

	// Simulate work
	time.Sleep(100 * time.Millisecond)
	fmt.Println("✓ Work completed")
}

func main() {
	fmt.Println("========== LEARNING DEFER BASICS ==========\n")

	// 1. Basic defer
	fmt.Println("--- 1. Basic Defer (executes last) ---")
	example1_BasicDefer()

	// 2. Multiple defers
	fmt.Println("\n--- 2. Multiple Defers (LIFO order) ---")
	example2_MultipleDefers()

	// 3. Defer with functions
	fmt.Println("\n--- 3. Defer with Functions ---")
	example3_DeferWithFunctions()

	// 4. Defer for cleanup
	fmt.Println("\n--- 4. Defer for Resource Cleanup ---")
	example4_DeferCleanup()

	// 5. File handling
	fmt.Println("\n--- 5. File Handling Simulation ---")
	example5_FileHandling()

	// 6. Panic recovery
	fmt.Println("\n--- 6. Panic Recovery with Defer ---")
	example6_PanicRecovery()

	// 7. Defer timing
	fmt.Println("\n--- 7. Defer Timing (argument captured) ---")
	example7_DeferTiming()

	// 8. Defer in loops
	fmt.Println("\n--- 8. Defer in Loops ---")
	example8_DeferInLoop()

	// 9. Lock/Unlock
	fmt.Println("\n--- 9. Lock/Unlock Pattern ---")
	example9_LockUnlock()

	// 10. Timing execution
	fmt.Println("\n--- 10. Timing Function Execution ---")
	example10_Timing()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Bank Account Transactions ==========\n")
	bankTransactionProject()
}

// ========== PROJECT: BANK ACCOUNT TRANSACTIONS ==========
// Real-world scenario: Ensure cleanup happens in transactions

type BankAccount struct {
	accountID string
	balance   float64
	locked    bool
}

func (ba *BankAccount) Lock() {
	ba.locked = true
	fmt.Printf("🔒 Account %s locked for transaction\n", ba.accountID)
}

func (ba *BankAccount) Unlock() {
	ba.locked = false
	fmt.Printf("🔓 Account %s unlocked\n", ba.accountID)
}

func (ba *BankAccount) LogTransaction(action string, amount float64) {
	fmt.Printf("📝 [%s] %s: $%.2f (Balance: $%.2f)\n",
		ba.accountID, action, amount, ba.balance)
}

func (ba *BankAccount) LogClose() {
	fmt.Printf("✓ Transaction complete for account %s\n", ba.accountID)
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Lock()
	defer ba.Unlock()   // Ensure unlock happens
	defer ba.LogClose() // Log completion

	if amount <= 0 {
		fmt.Println("❌ Invalid amount")
		return
	}

	ba.LogTransaction("DEPOSIT", amount)
	ba.balance += amount
	fmt.Printf("✅ New balance: $%.2f\n", ba.balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	ba.Lock()
	defer ba.Unlock()   // Ensure unlock happens
	defer ba.LogClose() // Log completion

	if amount <= 0 {
		fmt.Println("❌ Invalid amount")
		return
	}

	if amount > ba.balance {
		fmt.Println("❌ Insufficient funds")
		return
	}

	ba.LogTransaction("WITHDRAW", amount)
	ba.balance -= amount
	fmt.Printf("✅ New balance: $%.2f\n", ba.balance)
}

func (ba *BankAccount) Transfer(toAccount *BankAccount, amount float64) {
	fmt.Printf("\n💸 Transferring $%.2f from %s to %s\n",
		amount, ba.accountID, toAccount.accountID)

	// Lock both accounts (prevent deadlock safety)
	ba.Lock()
	defer ba.Unlock()
	toAccount.Lock()
	defer toAccount.Unlock()

	if amount <= 0 || amount > ba.balance {
		fmt.Println("❌ Transfer failed")
		return
	}

	ba.balance -= amount
	toAccount.balance += amount
	fmt.Printf("✅ Transfer successful!\n")
	fmt.Printf("  %s balance: $%.2f\n", ba.accountID, ba.balance)
	fmt.Printf("  %s balance: $%.2f\n", toAccount.accountID, toAccount.balance)
}

func bankTransactionProject() {
	// Create accounts
	alice := &BankAccount{accountID: "ACC-001", balance: 1000}
	bob := &BankAccount{accountID: "ACC-002", balance: 500}

	fmt.Println("=== BANK TRANSACTION SYSTEM ===")
	fmt.Println("Defer ensures cleanup even if errors occur!\n")

	for {
		fmt.Print("\nCommand (deposit/withdraw/transfer/list/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "deposit":
			fmt.Print("Amount: ")
			var amount float64
			fmt.Scanln(&amount)
			alice.Deposit(amount)

		case "withdraw":
			fmt.Print("Amount: ")
			var amount float64
			fmt.Scanln(&amount)
			alice.Withdraw(amount)

		case "transfer":
			fmt.Print("Amount to transfer to Bob: ")
			var amount float64
			fmt.Scanln(&amount)
			alice.Transfer(bob, amount)

		case "list":
			fmt.Printf("\n💰 Accounts:\n")
			fmt.Printf("  %s: $%.2f\n", alice.accountID, alice.balance)
			fmt.Printf("  %s: $%.2f\n", bob.accountID, bob.balance)

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
