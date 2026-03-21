package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== SIMPLE WAITGROUP LEARNING EXAMPLES ==========

// 1. BASIC WAITGROUP: Wait for goroutines to complete
// WaitGroup: Counter that tracks running goroutines
// Add(): Increment counter
// Done(): Decrement counter
// Wait(): Block until counter reaches 0
func example1_BasicWaitGroup() {
	fmt.Println("1. Basic WaitGroup:")

	var wg sync.WaitGroup

	// Add 2 goroutines to wait group
	wg.Add(2)

	go func() {
		defer wg.Done() // Decrement counter when done
		fmt.Println("  Goroutine 1 executing")
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		defer wg.Done() // Decrement counter when done
		fmt.Println("  Goroutine 2 executing")
		time.Sleep(100 * time.Millisecond)
	}()

	// Wait until all Done() calls complete
	wg.Wait()
	fmt.Println("  All goroutines finished!")
}

// 2. ADD IN LOOP: Add goroutines in a loop
func example2_AddInLoop() {
	fmt.Println("2. Add in Loop:")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Worker %d starting\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("  Worker %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("  All workers completed")
}

// 3. PRE-ADD: Add total count upfront
func example3_PreAdd() {
	fmt.Println("3. Pre-Add Total Count:")

	var wg sync.WaitGroup
	numWorkers := 4

	wg.Add(numWorkers) // Add total upfront

	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Task %d\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("  All tasks done")
}

// 4. ERROR HANDLING WITH WAITGROUP: Collect errors from goroutines
func example4_ErrorHandling() {
	fmt.Println("4. Error Handling:")

	var wg sync.WaitGroup
	errorsChan := make(chan error, 3)

	jobs := []struct {
		id    int
		valid bool
	}{
		{1, true},
		{2, false}, // This will error
		{3, true},
	}

	for _, job := range jobs {
		wg.Add(1)

		go func(id int, valid bool) {
			defer wg.Done()

			if !valid {
				errorsChan <- fmt.Errorf("job %d failed", id)
				return
			}

			fmt.Printf("  Job %d completed\n", id)
		}(job.id, job.valid)
	}

	wg.Wait()
	close(errorsChan)

	// Check for errors
	fmt.Println("  Errors:")
	errCount := 0
	for err := range errorsChan {
		fmt.Printf("    %v\n", err)
		errCount++
	}
	if errCount == 0 {
		fmt.Println("    None")
	}
}

// 5. CHANNEL CLOSE PATTERN: Close channel after all goroutines finish
func example5_ChannelClose() {
	fmt.Println("5. Channel Close Pattern:")

	var wg sync.WaitGroup
	results := make(chan string)

	// Launch goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			results <- fmt.Sprintf("Result from goroutine %d", id)
		}(i)
	}

	// Close channel in separate goroutine when all done
	go func() {
		wg.Wait()
		close(results) // Signal no more values
	}()

	// Receive until channel closes
	fmt.Println("  Receiving results:")
	for result := range results {
		fmt.Printf("    %s\n", result)
	}
	fmt.Println("  Channel closed")
}

// 6. WAITGROUP GOTCHA: Pass by pointer not value!
func example6_PointerGotcha() {
	fmt.Println("6. WaitGroup Gotcha (must pass pointer):")

	// WRONG: var wg sync.WaitGroup - if passed by value, counter lost!
	wg := &sync.WaitGroup{} // Correct: use pointer

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 2")
	}()

	wg.Wait()
	fmt.Println("  ✓ Using pointer ensures counter is shared")
}

// 7. COLLECTING RESULTS: Gather results from multiple goroutines
func example7_CollectResults() {
	fmt.Println("7. Collecting Results:")

	var wg sync.WaitGroup
	results := make([]int, 3)

	jobs := []struct {
		id    int
		index int
		value int
	}{
		{1, 0, 10},
		{2, 1, 20},
		{3, 2, 30},
	}

	for _, job := range jobs {
		wg.Add(1)

		go func(id, idx, val int) {
			defer wg.Done()
			results[idx] = val * 2 // Process
			fmt.Printf("  Job %d: %d * 2 = %d\n", id, val, val*2)
		}(job.id, job.index, job.value)
	}

	wg.Wait()
	fmt.Printf("  Final results: %v\n", results)
}

// 8. NESTED WAITGROUPS: WaitGroup inside functions
func example8_NestedWaitGroups() {
	fmt.Println("8. Nested WaitGroups:")

	var outerWg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		outerWg.Add(1)

		go func(batchID int) {
			defer outerWg.Done()

			var innerWg sync.WaitGroup

			// Inner goroutines
			for j := 1; j <= 2; j++ {
				innerWg.Add(1)

				go func(taskID int) {
					defer innerWg.Done()
					fmt.Printf("  Batch %d, Task %d\n", batchID, taskID)
				}(j)
			}

			innerWg.Wait()
			fmt.Printf("  Batch %d completed\n", batchID)
		}(i)
	}

	outerWg.Wait()
	fmt.Println("  All batches done")
}

// 9. TIMING WITH WAITGROUP: Measure concurrent operations
func example9_Timing() {
	fmt.Println("9. Timing Concurrent Operations:")

	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Task %d starting\n", id)
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			fmt.Printf("  Task %d done\n", id)
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Printf("  Total time: %v (parallel not sequential)\n", elapsed)
}

// 10. WAITGROUP BEST PRACTICES: Summary
func example10_BestPractices() {
	fmt.Println("10. Best Practices:")
	fmt.Println("  ✓ Always use defer wg.Done()")
	fmt.Println("  ✓ Never pass WaitGroup by value")
	fmt.Println("  ✓ Add before launching goroutines")
	fmt.Println("  ✓ Never call Wait() inside goroutine")
	fmt.Println("  ✓ Use WaitGroup for coordination")
	fmt.Println("  ✓ Close channels after Wait() completes")
	fmt.Println("  ✓ Consider using context for timeouts")
	fmt.Println("  ✓ Combine with channels for result collection")
}

func main() {
	fmt.Println("========== LEARNING WAITGROUP ==========\n")

	// 1. Basic WaitGroup
	fmt.Println("--- 1. Basic WaitGroup ---")
	example1_BasicWaitGroup()

	// 2. Add in loop
	fmt.Println("\n--- 2. Add in Loop ---")
	example2_AddInLoop()

	// 3. Pre-add
	fmt.Println("\n--- 3. Pre-Add Count ---")
	example3_PreAdd()

	// 4. Error handling
	fmt.Println("\n--- 4. Error Handling ---")
	example4_ErrorHandling()

	// 5. Channel close
	fmt.Println("\n--- 5. Channel Close Pattern ---")
	example5_ChannelClose()

	// 6. Pointer gotcha
	fmt.Println("\n--- 6. Pointer Gotcha ---")
	example6_PointerGotcha()

	// 7. Collect results
	fmt.Println("\n--- 7. Collecting Results ---")
	example7_CollectResults()

	// 8. Nested
	fmt.Println("\n--- 8. Nested WaitGroups ---")
	example8_NestedWaitGroups()

	// 9. Timing
	fmt.Println("\n--- 9. Timing ---")
	example9_Timing()

	// 10. Best practices
	fmt.Println("\n--- 10. Best Practices ---")
	example10_BestPractices()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Concurrent File Processor ==========\n")
	fileProcessorProject()
}

// ========== PROJECT: CONCURRENT FILE PROCESSOR ==========
// Real-world scenario: Process multiple files concurrently using WaitGroup
// Shows: Coordinating goroutines, collecting results, progress tracking

type FileTask struct {
	id       int
	filename string
	size     int // Simulated file size
}

type ProcessResult struct {
	id       int
	filename string
	lines    int
	status   string
	duration time.Duration
}

func processFile(task FileTask) ProcessResult {
	start := time.Now()
	fmt.Printf("  📄 [%d] Processing: %s (%d bytes)\n", task.id, task.filename, task.size)

	// Simulate file processing
	time.Sleep(time.Duration(task.size) * time.Millisecond)

	result := ProcessResult{
		id:       task.id,
		filename: task.filename,
		lines:    task.size / 100,
		status:   "✓ Complete",
		duration: time.Since(start),
	}

	return result
}

func fileProcessorProject() {
	fmt.Println("📂 Concurrent File Processor - Using WaitGroup for coordination\n")

	// Sample files
	files := []FileTask{
		{1, "log_2024_01.txt", 150},
		{2, "log_2024_02.txt", 200},
		{3, "log_2024_03.txt", 100},
		{4, "log_2024_04.txt", 180},
		{5, "log_2024_05.txt", 120},
	}

	for {
		fmt.Print("\nCommand (process/info/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "process":
			fmt.Print("Number of concurrent workers (1-5): ")
			var numWorkers int
			fmt.Scanln(&numWorkers)

			if numWorkers < 1 || numWorkers > 5 {
				fmt.Println("Invalid number. Using 2.")
				numWorkers = 2
			}

			fmt.Printf("\n⚙️  Processing %d files with %d workers...\n\n", len(files), numWorkers)

			// Use channels for work distribution
			fileChan := make(chan FileTask, len(files))
			resultsChan := make(chan ProcessResult, len(files))

			startTime := time.Now()

			// Create WaitGroup for coordination
			var wg sync.WaitGroup

			// Launch workers
			for w := 1; w <= numWorkers; w++ {
				wg.Add(1)

				go func(workerID int) {
					defer wg.Done()

					// Each worker processes files from channel
					for file := range fileChan {
						result := processFile(file)
						resultsChan <- result
					}
				}(w)
			}

			// Send files to workers
			go func() {
				for _, file := range files {
					fileChan <- file
				}
				close(fileChan)
			}()

			// Wait for all workers to finish
			wg.Wait()
			close(resultsChan)

			// Collect results
			fmt.Println("\n📊 Processing Results:")
			totalLines := 0
			for result := range resultsChan {
				fmt.Printf("  [%d] %s - %s (%.2fs)\n",
					result.id, result.filename, result.status, result.duration.Seconds())
				totalLines += result.lines
			}

			totalTime := time.Since(startTime)
			fmt.Printf("\n📈 Stats:\n")
			fmt.Printf("  Files processed: %d\n", len(files))
			fmt.Printf("  Workers used: %d\n", numWorkers)
			fmt.Printf("  Total time: %.2fs\n", totalTime.Seconds())
			fmt.Printf("  Total lines: %d\n", totalLines)
			fmt.Printf("  Throughput: %.1f files/sec\n", float64(len(files))/totalTime.Seconds())

		case "info":
			fmt.Println("\n🏗️  How WaitGroup Works:")
			fmt.Println("  1. Create WaitGroup")
			fmt.Println("  2. Call wg.Add(n) for each goroutine")
			fmt.Println("  3. Launch goroutines")
			fmt.Println("  4. Call wg.Done() at end of each goroutine")
			fmt.Println("  5. Call wg.Wait() to block until all done")
			fmt.Println("\n📝 Files to process:")
			for _, file := range files {
				fmt.Printf("  [%d] %s (%d bytes)\n", file.id, file.filename, file.size)
			}

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Use: process, info, quit")
		}
	}
}
