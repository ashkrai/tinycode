package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ========== CUSTOM ERROR TYPES DEFINITIONS ==========

// 1. BASIC ERROR INTERFACE: error is a built-in interface
// type error interface {
//     Error() string
// }

// 2. CUSTOM ERROR STRUCT: Implement the error interface
type ValidationError struct {
	field   string
	message string
	code    int
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("[%d] Field '%s': %s", e.code, e.field, e.message)
}

// 3. SENTINEL ERRORS: Known errors that can be compared with == or errors.Is()
var (
	ErrEmptyInput       = errors.New("input cannot be empty")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
	ErrUserNotFound     = errors.New("user not found in database")
)

// 4. CUSTOM ERROR WITH FIELDS: Store additional context
type FileError struct {
	Filename string
	Op       string // operation (read, write, delete)
	Err      error  // underlying error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file error: %s on '%s': %v", e.Op, e.Filename, e.Err)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

// 5. ERROR WRAPPING: errors.Is() and errors.As()
type DatabaseError struct {
	operation string
	table     string
	err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error [%s on %s]: %w", e.operation, e.table, e.err)
}

func (e *DatabaseError) Unwrap() error {
	return e.err
}

// 6. AUTHENTICATION ERROR: Custom error for auth failures
type AuthError struct {
	username string
	reason   string
	attempt  int
}

func (e AuthError) Error() string {
	return fmt.Sprintf("authentication failed for user '%s': %s (attempt %d)", e.username, e.reason, e.attempt)
}

// 7. CONFIGURATION ERROR: Multiple field validation errors
type ConfigError struct {
	errors []string
}

func (e ConfigError) Error() string {
	return fmt.Sprintf("configuration error: %d issues:\n  - %s", len(e.errors), strings.Join(e.errors, "\n  - "))
}

// ========== LEARNING EXAMPLES ==========

// 1. BASIC ERROR HANDLING: Return error as second value
func example1_BasicErrorHandling() {
	fmt.Println("1. Basic Error Handling:")

	parseAs := func(s string) (int, error) {
		for _, c := range s {
			if c < '0' || c > '9' {
				return 0, fmt.Errorf("invalid character '%c' in string", c)
			}
		}
		return len(s), nil
	}

	result, err := parseAs("12345")
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Result: %d\n", result)
	}

	_, err = parseAs("123abc")
	if err != nil {
		fmt.Printf("  Error caught: %v\n", err)
	}
}

// 2. CREATING CUSTOM ERROR TYPES
func example2_CustomErrorTypes() {
	fmt.Println("2. Creating Custom Error Types:")

	validateEmail := func(email string) error {
		if !strings.Contains(email, "@") {
			return ValidationError{
				field:   "email",
				message: "must contain @ symbol",
				code:    400,
			}
		}
		return nil
	}

	err := validateEmail("invalid-email")
	if err != nil {
		fmt.Printf("  %v\n", err)
	}

	err = validateEmail("user@example.com")
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Println("  ✓ Email valid")
	}
}

// 3. SENTINEL ERRORS: Compare known error values
func example3_SentinelErrors() {
	fmt.Println("3. Sentinel Errors:")

	validate := func(password string) error {
		if password == "" {
			return ErrEmptyInput
		}
		if len(password) < 8 {
			return ErrPasswordTooShort
		}
		return nil
	}

	err := validate("")
	if errors.Is(err, ErrEmptyInput) {
		fmt.Println("  Error: Empty input detected")
	}

	err = validate("short")
	if errors.Is(err, ErrPasswordTooShort) {
		fmt.Println("  Error: Password too short")
	}
}

// 4. ERROR WRAPPING & UNWRAPPING: Preserve error chain
func example4_ErrorWrapping() {
	fmt.Println("4. Error Wrapping & Unwrapping:")

	baseErr := errors.New("connection refused")
	wrappedErr := fmt.Errorf("failed to connect: %w", baseErr)

	fmt.Printf("  Wrapped error: %v\n", wrappedErr)
	fmt.Printf("  Unwrapped: %v\n", errors.Unwrap(wrappedErr))

	// Check if underlying error matches
	if errors.Is(wrappedErr, baseErr) {
		fmt.Println("  ✓ Found base error in chain")
	}
}

// 5. ERRORS.AS() FOR TYPE ASSERTION
func example5_ErrorsAs() {
	fmt.Println("5. Errors.As() - Type Assertion:")

	fileErr := &FileError{
		Filename: "config.json",
		Op:       "read",
		Err:      errors.New("permission denied"),
	}

	var target *FileError
	if errors.As(fileErr, &target) {
		fmt.Printf("  Caught FileError: %v\n", target)
		fmt.Printf("  Operation: %s on %s\n", target.Op, target.Filename)
	}
}

// 6. MULTIPLE ERRORS IN VALIDATION
func example6_MultipleErrors() {
	fmt.Println("6. Multiple Errors in Validation:")

	validateForm := func(name, email, password string) error {
		var errs []string

		if name == "" {
			errs = append(errs, "name is required")
		}
		if !strings.Contains(email, "@") {
			errs = append(errs, "invalid email format")
		}
		if len(password) < 8 {
			errs = append(errs, "password must be 8+ characters")
		}

		if len(errs) > 0 {
			return ConfigError{errors: errs}
		}
		return nil
	}

	err := validateForm("", "bad-email", "123")
	if err != nil {
		fmt.Printf("  %v\n", err)
	}
}

// 7. ERROR WITH DEFER: Defer error handling
func example7_DeferError() {
	fmt.Println("7. Error Handling with Defer:")

	processFile := func(filename string) (err error) {
		defer func() {
			if err != nil {
				fmt.Printf("  Cleanup: Failed to process %s\n", filename)
			}
		}()

		if filename == "" {
			return errors.New("filename cannot be empty")
		}
		fmt.Printf("  ✓ Processing %s\n", filename)
		return nil
	}

	processFile("data.txt")
	processFile("")
}

// 8. CUSTOM ERROR WITH RECOVERY HINT
func example8_ErrorWithHint() {
	fmt.Println("8. Error With Recovery Hint:")

	err := errors.New("failed to connect to database")
	if err != nil {
		fmt.Printf("  Problem: %v\n", err)
		fmt.Printf("  💡 Hint: Check if database server is running and connection string is correct\n")
	}
}

// 9. ERROR CHAINING IN OPERATIONS
func example9_ErrorChaining() {
	fmt.Println("9. Error Chaining:")

	readFile := func(filename string) (string, error) {
		if filename == "" {
			return "", errors.New("filename empty")
		}
		return "file contents", nil
	}

	parseContent := func(content string) error {
		if content == "" {
			return errors.New("empty content")
		}
		return nil
	}

	process := func(filename string) error {
		content, err := readFile(filename)
		if err != nil {
			return fmt.Errorf("read failed: %w", err)
		}

		if err := parseContent(content); err != nil {
			return fmt.Errorf("parse failed: %w", err)
		}

		return nil
	}

	if err := process("data.txt"); err != nil {
		fmt.Printf("  %v\n", err)
	}
}

// 10. BEST PRACTICES: Error handling patterns
func example10_BestPractices() {
	fmt.Println("10. Best Practices:")
	fmt.Println("  ✓ Always check error returns")
	fmt.Println("  ✓ Wrap errors with context using %w in fmt.Errorf")
	fmt.Println("  ✓ Use sentinel errors for known conditions")
	fmt.Println("  ✓ Create custom error types for specific errors")
	fmt.Println("  ✓ Use errors.Is() and errors.As() to inspect errors")
	fmt.Println("  ✓ Don't ignore errors silently")
	fmt.Println("  ✓ Include context: what, where, why")
	fmt.Println("  ✓ Log errors with sufficient information")
	fmt.Println("  ✓ Don't use panic for common errors")
	fmt.Println("  ✓ Keep error messages lowercase and actionable")
}

func main() {
	fmt.Println("========== LEARNING GO ERRORS & CUSTOM ERRORS ==========\n")

	// 1. Basic error handling
	fmt.Println("--- 1. Basic Error Handling ---")
	example1_BasicErrorHandling()

	// 2. Custom error types
	fmt.Println("\n--- 2. Custom Error Types ---")
	example2_CustomErrorTypes()

	// 3. Sentinel errors
	fmt.Println("\n--- 3. Sentinel Errors ---")
	example3_SentinelErrors()

	// 4. Error wrapping
	fmt.Println("\n--- 4. Error Wrapping ---")
	example4_ErrorWrapping()

	// 5. Errors.As()
	fmt.Println("\n--- 5. Errors.As() ---")
	example5_ErrorsAs()

	// 6. Multiple errors
	fmt.Println("\n--- 6. Multiple Errors ---")
	example6_MultipleErrors()

	// 7. Defer error
	fmt.Println("\n--- 7. Defer Error Handling ---")
	example7_DeferError()

	// 8. Error with hint
	fmt.Println("\n--- 8. Error With Hint ---")
	example8_ErrorWithHint()

	// 9. Error chaining
	fmt.Println("\n--- 9. Error Chaining ---")
	example9_ErrorChaining()

	// 10. Best practices
	fmt.Println("\n--- 10. Best Practices ---")
	example10_BestPractices()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: User Registration with Custom Errors ==========\n")
	userRegistrationProject()
}

// ========== PROJECT: USER REGISTRATION SYSTEM WITH CUSTOM ERRORS ==========
// Real-world scenario: Web service user registration with validation
// Shows: Custom errors, error wrapping, validation patterns
// Demonstrates: Type assertions, error inspection, user-friendly messages

type RegistrationError struct {
	Field   string
	Message string
	Code    string // error code for API responses
}

func (e RegistrationError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.Field, e.Message)
}

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

var database = make(map[string]*User) // Simple in-memory database

func validateUsername(username string) error {
	if username == "" {
		return RegistrationError{"username", "cannot be empty", "EMPTY"}
	}
	if len(username) < 3 {
		return RegistrationError{"username", "must be at least 3 characters", "SHORT"}
	}
	if len(username) > 20 {
		return RegistrationError{"username", "must be less than 20 characters", "LONG"}
	}
	if _, exists := database[username]; exists {
		return RegistrationError{"username", "already taken", "DUPLICATE"}
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return RegistrationError{"email", "cannot be empty", "EMPTY"}
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return RegistrationError{"email", "invalid email format", "INVALID"}
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return RegistrationError{"password", "cannot be empty", "EMPTY"}
	}
	if len(password) < 8 {
		return RegistrationError{"password", "must be at least 8 characters", "SHORT"}
	}
	if !strings.ContainsAny(password, "0123456789") {
		return RegistrationError{"password", "must contain at least one digit", "NO_DIGIT"}
	}
	return nil
}

func registerUser(username, email, password string) (*User, error) {
	// Validate each field
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// Create user
	user := &User{
		ID:        len(database) + 1,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	database[username] = user
	return user, nil
}

func loginUser(username, password string) (*User, error) {
	if username == "" {
		return nil, RegistrationError{"username", "cannot be empty", "EMPTY"}
	}

	user, exists := database[username]
	if !exists {
		return nil, RegistrationError{"username", "user not found", "NOT_FOUND"}
	}

	if user.Password != password {
		return nil, RegistrationError{"password", "incorrect password", "INVALID"}
	}

	return user, nil
}

func userRegistrationProject() {
	fmt.Println("👤 User Registration System with Error Handling\n")

	for {
		fmt.Print("Command (register/login/list/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "register":
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)

			fmt.Print("Email: ")
			var email string
			fmt.Scanln(&email)

			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			user, err := registerUser(username, email, password)
			if err != nil {
				var regErr RegistrationError
				if errors.As(err, &regErr) {
					fmt.Printf("❌ Registration failed:\n")
					fmt.Printf("   Field: %s\n", regErr.Field)
					fmt.Printf("   Error: %s\n", regErr.Message)
					fmt.Printf("   Code: %s\n", regErr.Code)
				} else {
					fmt.Printf("❌ Unexpected error: %v\n", err)
				}
			} else {
				fmt.Printf("✅ User registered successfully!\n")
				fmt.Printf("   ID: %d\n", user.ID)
				fmt.Printf("   Username: %s\n", user.Username)
				fmt.Printf("   Email: %s\n", user.Email)
			}

		case "login":
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)

			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			user, err := loginUser(username, password)
			if err != nil {
				var regErr RegistrationError
				if errors.As(err, &regErr) {
					fmt.Printf("❌ Login failed: %s\n", regErr.Message)
				} else {
					fmt.Printf("❌ Unexpected error: %v\n", err)
				}
			} else {
				fmt.Printf("✅ Login successful!\n")
				fmt.Printf("   Welcome, %s!\n", user.Username)
				fmt.Printf("   Account created: %s\n", user.CreatedAt.Format("2006-01-02 15:04:05"))
			}

		case "list":
			if len(database) == 0 {
				fmt.Println("📋 No users registered yet")
			} else {
				fmt.Println("📋 Registered Users:")
				for username, user := range database {
					fmt.Printf("  • %s (%s) - Joined: %s\n", username, user.Email, user.CreatedAt.Format("2006-01-02"))
				}
			}

		case "quit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Use: register, login, list, quit")
		}

		fmt.Println()
	}
}
