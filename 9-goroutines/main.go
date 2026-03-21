package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== SIMPLE GOROUTINES LEARNING EXAMPLES ==========

// 1. BASIC GOROUTINE: Run function concurrently with 'go' keyword
func example1_BasicGoroutine() {
	fmt.Println("Main: Starting")

	// Launch goroutine (runs concurrently)
	go func() {
		fmt.Println("Goroutine: Running concurrently")
	}()

	fmt.Println("Main: Continuing without waiting")
	time.Sleep(100 * time.Millisecond) // Wait so goroutine can finish
}

// 2. MULTIPLE GOROUTINES: Launch many at once
func example2_MultipleGoroutines() {
	fmt.Println("Starting 5 goroutines:")

	for i := 1; i <= 5; i++ {
		go func(num int) {
			fmt.Printf("  Goroutine %d executing\n", num)
			time.Sleep(50 * time.Millisecond)
		}(i) // Important: pass 'i' as parameter to avoid closure issues
	}

	time.Sleep(200 * time.Millisecond) // Wait for all to complete
	fmt.Println("All goroutines done")
}

// 3. CHANNELS: Communication between goroutines
// Channel acts as a pipe - send data from one goroutine, receive in another
func example3_BasicChannels() {
	fmt.Println("Using channels:")

	messages := make(chan string) // Create string channel

	go func() {
		messages <- "Hello from goroutine" // Send data to channel
	}()

	msg := <-messages // Receive data from channel (blocking)
	fmt.Println(msg)
}

// 4. CHANNELS WITH LOOP: Send multiple values through channel
func example4_ChannelLoop() {
	fmt.Println("Channel with loop:")

	numbers := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i * 10
			time.Sleep(50 * time.Millisecond)
		}
		close(numbers) // Close channel when done sending
	}()

	// Receive until channel closes
	for num := range numbers {
		fmt.Printf("  Received: %d\n", num)
	}
}

// 5. WAITGROUP: Wait for multiple goroutines to complete
// Better than sleep() - ensures all goroutines finish
func example5_WaitGroup() {
	fmt.Println("Using WaitGroup:")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter

		go func(num int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("  Task %d executing\n", num)
			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	wg.Wait() // Wait until all Done() called
	fmt.Println("All tasks completed")
}

// 6. BUFFERED CHANNELS: Channel with capacity (can hold multiple values)
// Non-buffered: blocks if no receiver; Buffered: can hold N values
func example6_BufferedChannels() {
	fmt.Println("Buffered channels:")

	// Create channel that can hold 3 values
	results := make(chan string, 3)

	results <- "First"
	results <- "Second"
	results <- "Third"

	// Receive all without goroutines
	fmt.Printf("  %s, %s, %s\n", <-results, <-results, <-results)
}

// 7. RACE CONDITION: Problem without synchronization
func example7_RaceCondition() {
	fmt.Println("Race condition (UNSAFE):")

	counter := 0

	// Multiple goroutines incrementing same variable (DANGEROUS!)
	for i := 0; i < 5; i++ {
		go func() {
			counter++ // Race condition: simultaneous access
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  Counter: %d (expected 5, might be less!)\n", counter)
}

// 8. MUTEX: Prevent race conditions with lock
// Mutex ensures only one goroutine accesses code at a time
func example8_MutexLock() {
	fmt.Println("Using Mutex (SAFE):")

	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()   // Lock before accessing shared variable
			counter++   // Now safe
			mu.Unlock() // Release lock
		}()
	}

	wg.Wait()
	fmt.Printf("  Counter: %d (always 5 - SAFE!)\n", counter)
}

// 9. SELECT: Choose between multiple channel operations
// Like switch but for channels
func example9_Select() {
	fmt.Println("Using select:")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		chan1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		chan2 <- "from channel 2"
	}()

	// Wait for first channel to send (others ignored)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Printf("  Received: %s\n", msg1)
		case msg2 := <-chan2:
			fmt.Printf("  Received: %s\n", msg2)
		}
	}
}

// 10. WORKER POOL PATTERN: Manage many tasks efficiently
// Reuse goroutines instead of creating millions
func example10_WorkerPool() {
	fmt.Println("Worker pool pattern:")

	jobs := make(chan int, 10)
	results := make(chan string, 10)

	// Create 3 workers
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Queue jobs
	for j := 1; j <= 7; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	// Print results
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result := fmt.Sprintf("Worker %d processed job %d", id, job)
		results <- result
	}
}

func main() {
	fmt.Println("========== LEARNING GOROUTINES ==========\n")

	// 1. Basic goroutine
	fmt.Println("--- 1. Basic Goroutine ---")
	example1_BasicGoroutine()

	// 2. Multiple goroutines
	fmt.Println("\n--- 2. Multiple Goroutines ---")
	example2_MultipleGoroutines()

	// 3. Basic channels
	fmt.Println("\n--- 3. Basic Channels ---")
	example3_BasicChannels()

	// 4. Channel loop
	fmt.Println("\n--- 4. Channel with Loop ---")
	example4_ChannelLoop()

	// 5. WaitGroup
	fmt.Println("\n--- 5. WaitGroup ---")
	example5_WaitGroup()

	// 6. Buffered channels
	fmt.Println("\n--- 6. Buffered Channels ---")
	example6_BufferedChannels()

	// 7. Race condition
	fmt.Println("\n--- 7. Race Condition (UNSAFE) ---")
	example7_RaceCondition()

	// 8. Mutex
	fmt.Println("\n--- 8. Mutex Lock (SAFE) ---")
	example8_MutexLock()

	// 9. Select
	fmt.Println("\n--- 9. Select Statement ---")
	example9_Select()

	// 10. Worker pool
	fmt.Println("\n--- 10. Worker Pool Pattern ---")
	example10_WorkerPool()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Concurrent Download Manager ==========\n")
	downloadManagerProject()
}

// ========== PROJECT: CONCURRENT DOWNLOAD MANAGER ==========
// Real-world scenario: Download multiple files simultaneously
// Shows practical use of goroutines, channels, and synchronization

type DownloadTask struct {
	id   int
	url  string
	size int // simulated file size
}

type DownloadResult struct {
	id       int
	url      string
	status   string
	duration time.Duration
}

func simulateDownload(task DownloadTask, results chan DownloadResult) {
	start := time.Now()

	// Simulate download time based on file size
	downloadTime := time.Duration(task.size) * time.Millisecond

	fmt.Printf("📥 [%d] Starting download: %s (%.1fs)\n", task.id, task.url, downloadTime.Seconds())
	time.Sleep(downloadTime)

	duration := time.Since(start)
	results <- DownloadResult{
		id:       task.id,
		url:      task.url,
		status:   "✓ Success",
		duration: duration,
	}
}

func downloadManagerProject() {
	fmt.Println("🌐 Download Manager - Download multiple files concurrently\n")

	// Create tasks
	tasks := []DownloadTask{
		{1, "https://example.com/file1.zip", 100},
		{2, "https://example.com/file2.zip", 150},
		{3, "https://example.com/file3.zip", 80},
		{4, "https://example.com/file4.zip", 120},
		{5, "https://example.com/file5.zip", 90},
	}

	for {
		fmt.Print("\nCommand (start/status/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "start":
			fmt.Println("\n🚀 Starting concurrent downloads...")
			startTime := time.Now()

			// Create result channel
			results := make(chan DownloadResult, len(tasks))

			// Launch goroutine for each download
			for _, task := range tasks {
				go simulateDownload(task, results)
			}

			// Collect all results
			var completedDownloads []DownloadResult
			for i := 0; i < len(tasks); i++ {
				result := <-results
				completedDownloads = append(completedDownloads, result)
			}

			// Print summary
			fmt.Println("\n📊 Download Summary:")
			totalTime := time.Since(startTime)
			totalSize := 0

			for _, task := range tasks {
				totalSize += task.size
			}

			for _, result := range completedDownloads {
				fmt.Printf("  [%d] %s - %s (%.2fs)\n",
					result.id, result.url, result.status, result.duration.Seconds())
			}

			fmt.Printf("\n⏱️  Sequential time would be: %.2fs\n", time.Duration(totalSize)*time.Millisecond)
			fmt.Printf("⏱️  Concurrent time achieved: %.2fs\n", totalTime.Seconds())
			fmt.Printf("📈 Speed improvement: %.1fx faster\n",
				float64(totalSize*100/int(totalTime.Milliseconds()))/100)

		case "status":
			fmt.Println("\n📋 Available downloads:")
			for _, task := range tasks {
				fmt.Printf("  [%d] %s (size: %d)\n", task.id, task.url, task.size)
			}

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Use: start, status, quit")
		}
	}
}
