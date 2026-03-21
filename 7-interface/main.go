package main

import "fmt"

// ========== SIMPLE INTERFACE EXAMPLES ==========

// 1. INTERFACE: Contract - "Any type must have these methods"
type Employee interface {
	GetRole() string
	GetSalary() float64
	DoWork() string
}

// 2. DEVELOPER: Implements Employee interface
type Developer struct {
	name   string
	salary float64
}

func (d Developer) GetRole() string {
	return "Developer"
}

func (d Developer) GetSalary() float64 {
	return d.salary
}

func (d Developer) DoWork() string {
	return "💻 Writing code and fixing bugs"
}

// 3. DESIGNER: Also implements Employee interface
type Designer struct {
	name   string
	salary float64
}

func (d Designer) GetRole() string {
	return "Designer"
}

func (d Designer) GetSalary() float64 {
	return d.salary
}

func (d Designer) DoWork() string {
	return "🎨 Creating beautiful UI/UX designs"
}

// 4. MANAGER: Another Employee type
type Manager struct {
	name   string
	salary float64
}

func (m Manager) GetRole() string {
	return "Manager"
}

func (m Manager) GetSalary() float64 {
	return m.salary
}

func (m Manager) DoWork() string {
	return "📋 Managing team and organizing tasks"
}

// 5. POLYMORPHISM: One function works with ANY Employee type
func ShowEmployeeInfo(e Employee) {
	fmt.Printf("Role: %s | Salary: $%.2f | Task: %s\n",
		e.GetRole(), e.GetSalary(), e.DoWork())
}

// 6. EMPTY INTERFACE: Can hold ANY type
func PrintAnything(val interface{}) {
	fmt.Printf("Value: %v (Type: %T)\n", val, val)
}

// 7. TYPE ASSERTION: Check if interface holds specific type
func CheckEmployeeType(e Employee) {
	if dev, ok := e.(Developer); ok {
		fmt.Printf("✓ Yes! It's a Developer earning: $%.2f\n", dev.salary)
	} else {
		fmt.Printf("❌ Not a Developer\n")
	}
}

// 8. TYPE SWITCH: Check multiple types
func CheckType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Printf("It's an integer: %d\n", v)
	case string:
		fmt.Printf("It's a string: %s\n", v)
	case bool:
		fmt.Printf("It's a boolean: %v\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	fmt.Println("========== LEARNING INTERFACE BASICS ==========\n")

	// Create employees
	dev := Developer{"Alice", 85000}
	designer := Designer{"Bob", 75000}
	manager := Manager{"Charlie", 95000}

	// 1. Call methods directly
	fmt.Println("--- 1. Direct Method Calls ---")
	fmt.Printf("%s: $%.2f - %s\n", dev.GetRole(), dev.GetSalary(), dev.DoWork())
	fmt.Printf("%s: $%.2f - %s\n", designer.GetRole(), designer.GetSalary(), designer.DoWork())
	fmt.Printf("%s: $%.2f - %s\n", manager.GetRole(), manager.GetSalary(), manager.DoWork())

	// 2. Use interface parameter (POLYMORPHISM)
	// Same function, different behavior for each type!
	fmt.Println("\n--- 2. Using Interface Parameter (Polymorphism) ---")
	ShowEmployeeInfo(dev)
	ShowEmployeeInfo(designer)
	ShowEmployeeInfo(manager)

	// 3. Slice of interfaces (store different types in one collection)
	fmt.Println("\n--- 3. Slice of Interface (Different Types Together) ---")
	employees := []Employee{dev, designer, manager}
	for i, emp := range employees {
		fmt.Printf("Employee %d: %s\n", i+1, emp.GetRole())
	}

	// 4. Empty interface (holds ANY type)
	fmt.Println("\n--- 4. Empty Interface (Any Type) ---")
	PrintAnything(42)
	PrintAnything("Hello Go")
	PrintAnything(3.14)
	PrintAnything(true)

	// 5. Type assertion
	fmt.Println("\n--- 5. Type Assertion ---")
	var emp Employee = Developer{"David", 90000}
	if dev, ok := emp.(Developer); ok {
		fmt.Printf("✓ Yes! It's a Developer named: %s\n", dev.name)
	}

	// 6. Type switch
	fmt.Println("\n--- 6. Type Switch ---")
	CheckType(100)
	CheckType("Golang")
	CheckType(false)

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Vehicle Management ==========\n")
	vehicleProject()
}

// ========== PROJECT: VEHICLE MANAGEMENT ==========

// Interface: Vehicle contract
type Vehicle interface {
	Start() string
	Stop() string
	GetType() string
}

// Car implements Vehicle
type Car struct {
	brand string
}

func (c Car) Start() string {
	return "🚗 Car engine started: Vroom!"
}

func (c Car) Stop() string {
	return "🚗 Car engine stopped"
}

func (c Car) GetType() string {
	return fmt.Sprintf("Car (%s)", c.brand)
}

// Bike implements Vehicle
type Bike struct {
	brand string
}

func (b Bike) Start() string {
	return "🏍️ Bike engine started: Brrr!"
}

func (b Bike) Stop() string {
	return "🏍️ Bike engine stopped"
}

func (b Bike) GetType() string {
	return fmt.Sprintf("Bike (%s)", b.brand)
}

// Truck implements Vehicle
type Truck struct {
	brand string
}

func (t Truck) Start() string {
	return "🚚 Truck engine started: Rumble!"
}

func (t Truck) Stop() string {
	return "🚚 Truck engine stopped"
}

func (t Truck) GetType() string {
	return fmt.Sprintf("Truck (%s)", t.brand)
}

// Start any vehicle
func StartVehicle(v Vehicle) {
	fmt.Printf("[%s] %s\n", v.GetType(), v.Start())
}

// Stop any vehicle
func StopVehicle(v Vehicle) {
	fmt.Printf("[%s] %s\n", v.GetType(), v.Stop())
}

func vehicleProject() {
	// Create vehicles
	car := Car{"Toyota"}
	bike := Bike{"Harley"}
	truck := Truck{"Volvo"}

	// Collection of different vehicles
	vehicles := []Vehicle{car, bike, truck}

	// Start all vehicles
	fmt.Println("--- Starting All Vehicles ---")
	for _, v := range vehicles {
		StartVehicle(v)
	}

	// Stop all vehicles
	fmt.Println("\n--- Stopping All Vehicles ---")
	for _, v := range vehicles {
		StopVehicle(v)
	}

	// Interactive demo
	fmt.Println("\n--- Interactive Demo ---")
	for {
		fmt.Print("\nSelect vehicle (1-car, 2-bike, 3-truck, 0-exit): ")
		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			fmt.Println("Goodbye!")
			break
		}

		if choice < 1 || choice > 3 {
			fmt.Println("Invalid choice")
			continue
		}

		selectedVehicle := vehicles[choice-1]
		fmt.Printf("Starting: %s\n", selectedVehicle.Start())

		fmt.Print("Press Enter to stop: ")
		fmt.Scanln()

		fmt.Printf("Stopping: %s\n", selectedVehicle.Stop())
	}
}
