package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
)

// ========== PANIC & RECOVER LEARNING EXAMPLES ==========

// 1. BASIC PANIC: Stop execution with an error message
func example1_BasicPanic() {
	fmt.Println("1. Basic Panic:")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  🛑 Caught panic: %v\n", r)
		}
	}()

	fmt.Println("  Before panic")
	panic("something went wrong!")
}

// 2. RECOVER: Catch and handle panics
func example2_BasicRecover() {
	fmt.Println("2. Basic Recover:")

	safeDiv := func(a, b int) (result int) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Recovered from: %v\n", r)
				result = 0
			}
		}()

		if b == 0 {
			panic("division by zero")
		}
		return a / b
	}

	fmt.Printf("  10 / 2 = %d\n", safeDiv(10, 2))
	fmt.Printf("  10 / 0 = %d (recovered)\n", safeDiv(10, 0))
}

// 3. DEFER WITH RECOVER: Always use defer before calling code
func example3_DeferWithRecover() {
	fmt.Println("3. Defer with Recover:")

	riskOperation := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic caught: %v", r)
			}
		}()

		data := []int{1, 2, 3}
		_ = data[10] // Index out of bounds
		return nil
	}

	err := riskOperation()
	if err != nil {
		fmt.Printf("  Handled: %v\n", err)
	}
}

// 4. PANIC WITH CUSTOM MESSAGE
func example4_CustomPanicMessage() {
	fmt.Println("4. Custom Panic Message:")

	safeArray := func(arr []string, index int) string {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Error: Array access failed - %v\n", r)
			}
		}()

		if index < 0 || index >= len(arr) {
			panic(fmt.Sprintf("index %d out of range [0:%d]", index, len(arr)))
		}
		return arr[index]
	}

	items := []string{"apple", "banana", "orange"}
	_ = safeArray(items, 10)
}

// 5. DETECTING PANIC VALUE TYPE
func example5_PanicValueType() {
	fmt.Println("5. Detecting Panic Value Type:")

	handlePanic := func(fn func()) {
		defer func() {
			if r := recover(); r != nil {
				switch x := r.(type) {
				case string:
					fmt.Printf("  String panic: %s\n", x)
				case error:
					fmt.Printf("  Error panic: %v\n", x)
				case int:
					fmt.Printf("  Integer panic: %d\n", x)
				default:
					fmt.Printf("  Unknown panic type: %T\n", x)
				}
			}
		}()
		fn()
	}

	handlePanic(func() { panic("text error") })
	handlePanic(func() { panic(42) })
	handlePanic(func() { panic(fmt.Errorf("custom error")) })
}

// 6. NESTED DEFER RECOVER
func example6_NestedDefer() {
	fmt.Println("6. Nested Defer Recover:")

	inner := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Inner recovered: %v\n", r)
				panic(fmt.Sprintf("re-panicked: %v", r))
			}
		}()
		panic("inner panic")
	}

	outer := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Outer recovered: %v\n", r)
			}
		}()
		inner()
	}

	outer()
}

// 7. PANIC IN GOROUTINE
func example7_PanicInGoroutine() {
	fmt.Println("7. Panic in Goroutine:")

	ch := make(chan string, 3)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- fmt.Sprintf("goroutine panic: %v", r)
			}
		}()
		panic("goroutine error")
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- fmt.Sprintf("another goroutine panic: %v", r)
			}
		}()
		ch <- "success"
	}()

	fmt.Printf("  %s\n", <-ch)
	fmt.Printf("  %s\n", <-ch)
}

// 8. STACK TRACE WITH RECOVER
func example8_StackTrace() {
	fmt.Println("8. Stack Trace with Recover:")

	criticalOperation := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Panic: %v\n", r)
				stackTrace := string(debug.Stack())
				lines := strings.Split(stackTrace, "\n")
				// Show first few lines of stack
				for i := 0; i < 5 && i < len(lines); i++ {
					fmt.Printf("    %s\n", lines[i])
				}
			}
		}()
		panic("critical error occurred")
	}

	criticalOperation()
}

// 9. VALIDATE INPUTS TO AVOID PANIC
func example9_InputValidation() {
	fmt.Println("9. Input Validation to Prevent Panic:")

	parseInt := func(s string) (int, error) {
		if s == "" {
			return 0, fmt.Errorf("empty string")
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", s)
		}
		return val, nil
	}

	result, err := parseInt("123")
	if err == nil {
		fmt.Printf("  Parsed: %d\n", result)
	}

	_, err = parseInt("abc")
	if err != nil {
		fmt.Printf("  Error handled: %v\n", err)
	}
}

// 10. BEST PRACTICES
func example10_BestPractices() {
	fmt.Println("10. Best Practices:")
	fmt.Println("  ✓ Use panic for truly exceptional/unrecoverable errors")
	fmt.Println("  ✓ Use recover in defer blocks only")
	fmt.Println("  ✓ Each goroutine must have its own recover")
	fmt.Println("  ✓ Prefer returning errors for normal failures")
	fmt.Println("  ✓ Validate input to prevent panics")
	fmt.Println("  ✓ Don't panic in library code")
	fmt.Println("  ✓ Use panic for programmer errors (bugs)")
	fmt.Println("  ✓ Log panic information before recovery")
	fmt.Println("  ✓ Re-panic after recovery if necessary")
	fmt.Println("  ✓ Use panic in main() only as last resort")
}

func main() {
	fmt.Println("========== LEARNING GO PANIC & RECOVER ==========\n")

	// 1. Basic panic
	fmt.Println("--- 1. Basic Panic ---")
	example1_BasicPanic()

	// 2. Basic recover
	fmt.Println("\n--- 2. Basic Recover ---")
	example2_BasicRecover()

	// 3. Defer with recover
	fmt.Println("\n--- 3. Defer with Recover ---")
	example3_DeferWithRecover()

	// 4. Custom panic message
	fmt.Println("\n--- 4. Custom Panic Message ---")
	example4_CustomPanicMessage()

	// 5. Panic value type
	fmt.Println("\n--- 5. Detecting Panic Value Type ---")
	example5_PanicValueType()

	// 6. Nested defer
	fmt.Println("\n--- 6. Nested Defer Recover ---")
	example6_NestedDefer()

	// 7. Panic in goroutine
	fmt.Println("\n--- 7. Panic in Goroutine ---")
	example7_PanicInGoroutine()

	// 8. Stack trace
	fmt.Println("\n--- 8. Stack Trace with Recover ---")
	example8_StackTrace()

	// 9. Input validation
	fmt.Println("\n--- 9. Input Validation ---")
	example9_InputValidation()

	// 10. Best practices
	fmt.Println("\n--- 10. Best Practices ---")
	example10_BestPractices()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Safe Calculator with Panic Handling ==========\n")
	calculatorProject()
}

// ========== PROJECT: SAFE CALCULATOR WITH PANIC HANDLING ==========
// Real-world scenario: Math calculator that safely handles invalid operations
// Shows: Panic recovery, input validation, error handling
// Demonstrates: Safe division, type checking, comprehensive error reporting

type CalculatorError struct {
	Operation string
	Reason    string
	Input1    string
	Input2    string
}

func (e CalculatorError) Error() string {
	return fmt.Sprintf("calculator error [%s]: %s (inputs: %s, %s)", e.Operation, e.Reason, e.Input1, e.Input2)
}

func safeParseFloat(s string) (float64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty input")
	}
	s = strings.TrimSpace(s)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number", s)
	}
	return val, nil
}

func safeDivide(a, b float64) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic in safeDivide: %v\n", r)
		}
	}()

	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func safePower(base, exp float64) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic in safePower: %v\n", r)
		}
	}()

	if exp < 0 {
		return 0, fmt.Errorf("negative exponent not supported")
	}

	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}

	if result > 1e308 {
		return 0, fmt.Errorf("result too large (overflow)")
	}

	return result, nil
}

func safeSqrt(x float64) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic in safeSqrt: %v\n", r)
		}
	}()

	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number")
	}

	if x == 0 {
		return 0, nil
	}

	// Newton's method for square root
	guess := x / 2
	for i := 0; i < 100; i++ {
		guess = (guess + x/guess) / 2
	}

	return guess, nil
}

type CalculatorMemory struct {
	operations []string
	result     float64
	hasResult  bool
}

func (m *CalculatorMemory) addOperation(op string) {
	m.operations = append(m.operations, op)
}

func (m *CalculatorMemory) showHistory() {
	if len(m.operations) == 0 {
		fmt.Println("  📝 No operations yet")
		return
	}
	fmt.Println("  📝 History:")
	for i, op := range m.operations {
		fmt.Printf("    %d. %s\n", i+1, op)
	}
	if m.hasResult {
		fmt.Printf("  Last result: %.2f\n", m.result)
	}
}

func calculatorProject() {
	fmt.Println("🔢 Safe Calculator with Error Recovery\n")

	memory := &CalculatorMemory{}

	for {
		fmt.Print("Command (add/sub/mul/div/pow/sqrt/history/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "add":
			fmt.Print("Number 1: ")
			var s1 string
			fmt.Scanln(&s1)

			fmt.Print("Number 2: ")
			var s2 string
			fmt.Scanln(&s2)

			a, err1 := safeParseFloat(s1)
			b, err2 := safeParseFloat(s2)

			if err1 != nil || err2 != nil {
				fmt.Printf("❌ Invalid input\n")
			} else {
				result := a + b
				memory.hasResult = true
				operation := fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result)
				memory.addOperation(operation)
				fmt.Printf("✅ Result: %.2f\n", result)
			}

		case "sub":
			fmt.Print("Number 1: ")
			var s1 string
			fmt.Scanln(&s1)

			fmt.Print("Number 2: ")
			var s2 string
			fmt.Scanln(&s2)

			a, err1 := safeParseFloat(s1)
			b, err2 := safeParseFloat(s2)

			if err1 != nil || err2 != nil {
				fmt.Printf("❌ Invalid input\n")
			} else {
				result := a - b
				memory.hasResult = true
				operation := fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result)
				memory.addOperation(operation)
				fmt.Printf("✅ Result: %.2f\n", result)
			}

		case "mul":
			fmt.Print("Number 1: ")
			var s1 string
			fmt.Scanln(&s1)

			fmt.Print("Number 2: ")
			var s2 string
			fmt.Scanln(&s2)

			a, err1 := safeParseFloat(s1)
			b, err2 := safeParseFloat(s2)

			if err1 != nil || err2 != nil {
				fmt.Printf("❌ Invalid input\n")
			} else {
				result := a * b
				memory.hasResult = true
				operation := fmt.Sprintf("%.2f × %.2f = %.2f", a, b, result)
				memory.addOperation(operation)
				fmt.Printf("✅ Result: %.2f\n", result)
			}

		case "div":
			fmt.Print("Number 1: ")
			var s1 string
			fmt.Scanln(&s1)

			fmt.Print("Number 2: ")
			var s2 string
			fmt.Scanln(&s2)

			a, err1 := safeParseFloat(s1)
			b, err2 := safeParseFloat(s2)

			if err1 != nil || err2 != nil {
				fmt.Printf("❌ Invalid input\n")
			} else {
				result, err := safeDivide(a, b)
				if err != nil {
					fmt.Printf("❌ Division error: %v\n", err)
				} else {
					memory.hasResult = true
					operation := fmt.Sprintf("%.2f ÷ %.2f = %.2f", a, b, result)
					memory.addOperation(operation)
					fmt.Printf("✅ Result: %.2f\n", result)
				}
			}

		case "pow":
			fmt.Print("Base: ")
			var s1 string
			fmt.Scanln(&s1)

			fmt.Print("Exponent: ")
			var s2 string
			fmt.Scanln(&s2)

			base, err1 := safeParseFloat(s1)
			exp, err2 := safeParseFloat(s2)

			if err1 != nil || err2 != nil {
				fmt.Printf("❌ Invalid input\n")
			} else {
				result, err := safePower(base, exp)
				if err != nil {
					fmt.Printf("❌ Power error: %v\n", err)
				} else {
					memory.hasResult = true
					operation := fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exp, result)
					memory.addOperation(operation)
					fmt.Printf("✅ Result: %.2f\n", result)
				}
			}

		case "sqrt":
			fmt.Print("Number: ")
			var s string
			fmt.Scanln(&s)

			x, err := safeParseFloat(s)
			if err != nil {
				fmt.Printf("❌ Invalid input: %v\n", err)
			} else {
				result, err := safeSqrt(x)
				if err != nil {
					fmt.Printf("❌ Square root error: %v\n", err)
				} else {
					memory.hasResult = true
					operation := fmt.Sprintf("√%.2f = %.2f", x, result)
					memory.addOperation(operation)
					fmt.Printf("✅ Result: %.2f\n", result)
				}
			}

		case "history":
			memory.showHistory()

		case "quit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❓ Unknown command. Try: add, sub, mul, div, pow, sqrt, history, quit")
		}

		fmt.Println()
	}
}
