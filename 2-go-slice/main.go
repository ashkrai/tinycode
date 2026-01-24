package main

import "fmt"

func main() {
	// 1. Declare and create a slice
	var fruits []string
	fruits = append(fruits, "apple", "banana", "orange")
	fmt.Println("Slice:", fruits)

	// 2. Slice literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// 3. Length and Capacity
	fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))

	// 4. Slicing operation (start:end)
	subset := numbers[1:4] // elements at index 1, 2, 3
	fmt.Println("Subset [1:4]:", subset)

	// 5. Range iteration
	fmt.Println("Iterate with range:")
	for i, val := range fruits {
		fmt.Printf("  Index %d: %s\n", i, val)
	}

	// 6. Append
	fruits = append(fruits, "grape")
	fmt.Println("After append:", fruits)

	// 7. Modify element
	fruits[0] = "avocado"
	fmt.Println("After modify:", fruits)

	// 8. Make with capacity
	reserved := make([]int, 2, 5) // length 2, capacity 5
	reserved = append(reserved, 10)
	fmt.Println("Reserved slice:", reserved, "Cap:", cap(reserved))

	// 9. Copy slice
	copied := make([]int, len(numbers))
	copy(copied, numbers)
	fmt.Println("Copied:", copied)

	// PROJECT: Student Grade Manager
	fmt.Println("\n--- PROJECT: Student Grade Manager ---")
	studentGradeTracker()
}

// PROJECT: Simple grade management system using slices
func studentGradeTracker() {
	// Initialize grades slice
	var grades []float64

	// Add grades
	grades = append(grades, 85.5, 90.0, 78.5, 92.0, 88.0)
	fmt.Println("Grades:", grades)

	// Calculate average
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	average := sum / float64(len(grades))
	fmt.Printf("Average: %.2f\n", average)

	// Filter passing grades (>= 80)
	var passingGrades []float64
	for _, grade := range grades {
		if grade >= 80 {
			passingGrades = append(passingGrades, grade)
		}
	}
	fmt.Println("Passing grades (>=80):", passingGrades)

	// Find highest and lowest
	maxGrade := grades[0]
	minGrade := grades[0]
	for _, grade := range grades {
		if grade > maxGrade {
			maxGrade = grade
		}
		if grade < minGrade {
			minGrade = grade
		}
	}
	fmt.Printf("Highest: %.1f, Lowest: %.1f\n", maxGrade, minGrade)
}
