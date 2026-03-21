package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== SIMPLE MUTEX LEARNING EXAMPLES ==========

// 1. BASIC MUTEX: Protect shared variable from race conditions
// Mutex = Mutual Exclusion lock - only one goroutine can access at a time
func example1_BasicMutex() {
	fmt.Println("1. Basic Mutex:")

	var counter int
	var mu sync.Mutex

	// Without mutex: race condition!
	// With mutex: safe concurrent access
	mu.Lock()
	counter = 0
	mu.Unlock()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++ // Safe: only one goroutine at a time
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("  Counter: %d (always 5 - SAFE)\n", counter)
}

// 2. LOCK and UNLOCK: Explicit control
func example2_LockUnlock() {
	fmt.Println("2. Lock and Unlock:")

	var mu sync.Mutex
	var value string

	// Lock: Acquire exclusive access
	mu.Lock()
	value = "protected data"
	fmt.Printf("  Inside lock: %s\n", value)
	mu.Unlock() // Unlock: Release lock, other goroutines can acquire

	fmt.Println("  Lock released")
}

// 3. DEFER UNLOCK: Ensure unlock happens even on error
// Best practice: Use defer to guarantee unlock
func example3_DeferUnlock() {
	fmt.Println("3. Defer Unlock (best practice):")

	var mu sync.Mutex
	var counter int

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock() // Guarantees unlock even if panic

			counter++
			fmt.Printf("  Incremented to: %d\n", counter)
			// Unlock happens automatically here via defer
		}()
	}

	wg.Wait()
	fmt.Printf("  Final: %d\n", counter)
}

// 4. MULTIPLE MUTEXES: Protect different resources
// Each mutex protects specific data
func example4_MultipleMutexes() {
	fmt.Println("4. Multiple Mutexes:")

	var muName sync.Mutex
	var muAge sync.Mutex

	var name string
	var age int

	var wg sync.WaitGroup

	// Goroutine 1: Modify name
	wg.Add(1)
	go func() {
		defer wg.Done()
		muName.Lock()
		defer muName.Unlock()
		name = "Alice"
		fmt.Println("  Set name: Alice")
	}()

	// Goroutine 2: Modify age
	wg.Add(1)
	go func() {
		defer wg.Done()
		muAge.Lock()
		defer muAge.Unlock()
		age = 30
		fmt.Println("  Set age: 30")
	}()

	wg.Wait()
	fmt.Printf("  Name: %s, Age: %d\n", name, age)
}

// 5. RACE CONDITION WITHOUT MUTEX: Show the problem
func example5_RaceCondition() {
	fmt.Println("5. Race Condition (UNSAFE without mutex):")

	var counter int
	// NO MUTEX - UNSAFE!

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Race condition!
		}()
	}

	wg.Wait()
	fmt.Printf("  Counter: %d (expected 100, might be less!)\n", counter)
	fmt.Println("  ⚠️  This might not always be 100 - demonstrates race condition")
}

// 6. RWMUTEX: Read-Write mutex for read-heavy workloads
// Multiple readers OR one writer (not both)
type Account struct {
	balance float64
	mu      sync.RWMutex
}

func (acc *Account) Read() float64 {
	acc.mu.RLock() // Read lock - multiple readers allowed
	defer acc.mu.RUnlock()
	return acc.balance
}

func (acc *Account) Write(amount float64) {
	acc.mu.Lock() // Write lock - exclusive access
	defer acc.mu.Unlock()
	acc.balance = amount
}

func example6_RWMutex() {
	fmt.Println("6. RWMutex (Read-Write lock):")

	acc := &Account{balance: 1000}

	var wg sync.WaitGroup

	// Multiple readers (allowed concurrently)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			balance := acc.Read()
			fmt.Printf("  Reader %d: balance = %.2f\n", id, balance)
		}(i)
	}

	// One writer (exclusive access)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		acc.Write(1500)
		fmt.Printf("  Writer: updated balance to 1500.00\n")
	}()

	wg.Wait()
	fmt.Printf("  Final balance: %.2f\n", acc.Read())
}

// 7. CHANNEL vs MUTEX: Different synchronization approaches
func example7_ChannelVsMutex() {
	fmt.Println("7. Channel vs Mutex:")
	fmt.Println("  Mutex:")
	fmt.Println("    - Protects shared data")
	fmt.Println("    - Low-level synchronization")
	fmt.Println("    - Better for shared state")
	fmt.Println("  Channel:")
	fmt.Println("    - Passes data between goroutines")
	fmt.Println("    - High-level communication")
	fmt.Println("    - Better for work distribution")
}

// 8. DEADLOCK: Common mutex mistake
func example8_DeadlockRisk() {
	fmt.Println("8. Deadlock Risk:")
	fmt.Println("  Example: Two mutexes locked in different order")
	fmt.Println("  Goroutine 1: Lock A, then Lock B")
	fmt.Println("  Goroutine 2: Lock B, then Lock A")
	fmt.Println("  Result: Both wait forever - DEADLOCK!")
	fmt.Println("  Solution: Always lock in same order")
}

// 9. SYNC.ONCE: Execute code exactly once
// Useful for initialization
var initOnce sync.Once
var initValue string

func initializeExpensive() {
	fmt.Println("  (Expensive initialization happening...)")
	time.Sleep(100 * time.Millisecond)
	initValue = "Initialized!"
}

func example9_SyncOnce() {
	fmt.Println("9. Sync.Once (execute once):")

	var wg sync.WaitGroup

	// Multiple goroutines try to initialize
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			initOnce.Do(initializeExpensive) // Runs only once
			fmt.Printf("  Goroutine %d sees: %s\n", id, initValue)
		}(i)
	}

	wg.Wait()
	fmt.Println("  (Initialization ran only once despite 3 goroutines)")
}

// 10. MUTEX BEST PRACTICES: Summary
func example10_BestPractices() {
	fmt.Println("10. Mutex Best Practices:")
	fmt.Println("  ✓ Always use defer to unlock")
	fmt.Println("  ✓ Keep critical section small")
	fmt.Println("  ✓ Never hold lock during I/O")
	fmt.Println("  ✓ Lock in consistent order (prevent deadlock)")
	fmt.Println("  ✓ Use RWMutex for read-heavy workloads")
	fmt.Println("  ✓ Prefer channels for communication")
	fmt.Println("  ✓ Don't deadlock: avoid nested locks")
	fmt.Println("  ✓ Use sync.Once for one-time initialization")
}

func main() {
	fmt.Println("========== LEARNING MUTEX ==========\n")

	// 1. Basic mutex
	fmt.Println("--- 1. Basic Mutex ---")
	example1_BasicMutex()

	// 2. Lock unlock
	fmt.Println("\n--- 2. Lock and Unlock ---")
	example2_LockUnlock()

	// 3. Defer unlock
	fmt.Println("\n--- 3. Defer Unlock ---")
	example3_DeferUnlock()

	// 4. Multiple mutexes
	fmt.Println("\n--- 4. Multiple Mutexes ---")
	example4_MultipleMutexes()

	// 5. Race condition
	fmt.Println("\n--- 5. Race Condition ---")
	example5_RaceCondition()

	// 6. RWMutex
	fmt.Println("\n--- 6. RWMutex ---")
	example6_RWMutex()

	// 7. Channel vs Mutex
	fmt.Println("\n--- 7. Channel vs Mutex ---")
	example7_ChannelVsMutex()

	// 8. Deadlock
	fmt.Println("\n--- 8. Deadlock Risk ---")
	example8_DeadlockRisk()

	// 9. Sync.Once
	fmt.Println("\n--- 9. Sync.Once ---")
	example9_SyncOnce()

	// 10. Best practices
	fmt.Println("\n--- 10. Best Practices ---")
	example10_BestPractices()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Bank Account System ==========\n")
	bankAccountProject()
}

// ========== PROJECT: BANK ACCOUNT SYSTEM ==========
// Real-world scenario: Multiple goroutines accessing shared bank account
// Shows: Mutex protecting account balance, transfer operations, thread-safe operations

type BankAccount struct {
	accountID string
	balance   float64
	mu        sync.Mutex
}

func (acc *BankAccount) Deposit(amount float64) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if amount <= 0 {
		return false
	}

	acc.balance += amount
	return true
}

func (acc *BankAccount) Withdraw(amount float64) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if amount <= 0 || amount > acc.balance {
		return false
	}

	acc.balance -= amount
	return true
}

func (acc *BankAccount) GetBalance() float64 {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.balance
}

func (from *BankAccount) Transfer(to *BankAccount, amount float64) bool {
	// Lock both accounts (always same order to prevent deadlock)
	// Order: lock account with smaller ID first
	first, second := from, to
	if from.accountID > to.accountID {
		first, second = to, from
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	second.mu.Lock()
	defer second.mu.Unlock()

	if amount <= 0 || amount > from.balance {
		return false
	}

	from.balance -= amount
	to.balance += amount
	return true
}

func bankAccountProject() {
	fmt.Println("🏦 Bank Account System - Thread-safe operations\n")

	// Create accounts
	alice := &BankAccount{accountID: "ACC-001", balance: 5000}
	bob := &BankAccount{accountID: "ACC-002", balance: 3000}
	charlie := &BankAccount{accountID: "ACC-003", balance: 2000}

	accounts := map[string]*BankAccount{
		"ACC-001": alice,
		"ACC-002": bob,
		"ACC-003": charlie,
	}

	for {
		fmt.Print("\nCommand (deposit/withdraw/transfer/list/stress-test/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "deposit":
			fmt.Print("Account ID (ACC-001/ACC-002/ACC-003): ")
			var accID string
			fmt.Scanln(&accID)

			if acc, exists := accounts[accID]; exists {
				fmt.Print("Amount: ")
				var amount float64
				fmt.Scanln(&amount)

				if acc.Deposit(amount) {
					fmt.Printf("✅ Deposited $%.2f to %s\n", amount, accID)
					fmt.Printf("   New balance: $%.2f\n", acc.GetBalance())
				} else {
					fmt.Println("❌ Invalid amount")
				}
			} else {
				fmt.Println("❌ Account not found")
			}

		case "withdraw":
			fmt.Print("Account ID: ")
			var accID string
			fmt.Scanln(&accID)

			if acc, exists := accounts[accID]; exists {
				fmt.Print("Amount: ")
				var amount float64
				fmt.Scanln(&amount)

				if acc.Withdraw(amount) {
					fmt.Printf("✅ Withdrew $%.2f from %s\n", amount, accID)
					fmt.Printf("   New balance: $%.2f\n", acc.GetBalance())
				} else {
					fmt.Println("❌ Insufficient funds or invalid amount")
				}
			} else {
				fmt.Println("❌ Account not found")
			}

		case "transfer":
			fmt.Print("From Account ID: ")
			var fromID string
			fmt.Scanln(&fromID)

			fmt.Print("To Account ID: ")
			var toID string
			fmt.Scanln(&toID)

			fromAcc, fromExists := accounts[fromID]
			toAcc, toExists := accounts[toID]

			if !fromExists || !toExists {
				fmt.Println("❌ One or both accounts not found")
				break
			}

			fmt.Print("Amount: ")
			var amount float64
			fmt.Scanln(&amount)

			if fromAcc.Transfer(toAcc, amount) {
				fmt.Printf("✅ Transferred $%.2f from %s to %s\n", amount, fromID, toID)
				fmt.Printf("   %s balance: $%.2f\n", fromID, fromAcc.GetBalance())
				fmt.Printf("   %s balance: $%.2f\n", toID, toAcc.GetBalance())
			} else {
				fmt.Println("❌ Transfer failed - insufficient funds or invalid amount")
			}

		case "list":
			fmt.Println("\n💰 Account Balances:")
			for accID, acc := range accounts {
				fmt.Printf("  %s: $%.2f\n", accID, acc.GetBalance())
			}

		case "stress-test":
			fmt.Println("\n⚡ Running stress test...")
			fmt.Println("Creating 10 concurrent deposit operations")

			var wg sync.WaitGroup
			startTime := time.Now()

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					alice.Deposit(float64(id * 100))
				}(i)
			}

			wg.Wait()
			elapsed := time.Since(startTime)

			fmt.Printf("✅ Stress test completed in %v\n", elapsed)
			fmt.Printf("   Final balance: $%.2f\n", alice.GetBalance())
			fmt.Println("   All operations completed safely with mutex protection")

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
