package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ========== SIMPLE WORKER POOL LEARNING EXAMPLES ==========

// 1. BASIC WORKER: Single worker processing jobs
// Worker: Function that runs in its own goroutine
// Job: Work that needs to be done
func example1_BasicWorker() {
	fmt.Println("1. Basic Worker:")

	jobs := make(chan int, 5)
	results := make(chan string, 5)

	// Single worker goroutine
	go func() {
		for job := range jobs {
			result := fmt.Sprintf("Processed job %d", job)
			results <- result
		}
	}()

	// Queue jobs
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	for i := 0; i < 3; i++ {
		fmt.Printf("  %s\n", <-results)
	}
}

// 2. MULTIPLE WORKERS: Several workers sharing same job queue
// Each worker pulls from job queue independently
// Solves: Efficiently process many jobs with limited resources
func example2_MultipleWorkers() {
	fmt.Println("2. Multiple Workers (2 workers):")

	jobs := make(chan int, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// Launch 2 workers
	numWorkers := 2
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("  Worker %d: processing job %d\n", workerID, job)
				results <- fmt.Sprintf("Job %d (Worker %d)", job, workerID)
				time.Sleep(50 * time.Millisecond)
			}
		}(w)
	}

	// Queue jobs
	for i := 1; i <= 6; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for workers
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("  Results:")
	for result := range results {
		fmt.Printf("    %s\n", result)
	}
}

// 3. WORKER POOL PATTERN CONCEPT: Reusable worker template
// Components: jobs channel, results channel, worker function, WaitGroup
func example3_WorkerPoolStructure() {
	fmt.Println("3. Worker Pool Structure:")
	fmt.Println("  Components:")
	fmt.Println("  - Jobs channel: Queue of work to do")
	fmt.Println("  - Results channel: Processed results")
	fmt.Println("  - Workers: Fixed number of goroutines")
	fmt.Println("  - WaitGroup: Synchronization")
	fmt.Println("  Benefits: Resource efficiency, load balancing, easy scaling")
}

// 4. BOUNDED WORKER POOL: Limit total goroutines with semaphore
// Prevents creating too many goroutines (expensive resource)
func example4_BoundedWorkerPool() {
	fmt.Println("4. Bounded Worker Pool (max 3 concurrent):")

	semaphore := make(chan struct{}, 3) // Only 3 can run simultaneously

	for i := 1; i <= 8; i++ {
		job := i
		// Each job acquires semaphore slot
		go func() {
			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			fmt.Printf("  Job %d executing\n", job)
			time.Sleep(100 * time.Millisecond)
		}()
	}

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("  All jobs completed")
}

// 5. ERROR HANDLING IN WORKERS: Process errors from jobs
func example5_ErrorHandling() {
	fmt.Println("5. Error Handling in Workers:")

	type Job struct {
		id    int
		value int
	}

	type Result struct {
		id    int
		value int
		err   error
	}

	jobs := make(chan Job, 5)
	results := make(chan Result, 5)
	var wg sync.WaitGroup

	// Worker with error handling
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				result := Result{id: job.id}

				// Simulate error for certain jobs
				if job.value < 0 {
					result.err = fmt.Errorf("negative value")
				} else {
					result.value = job.value * 2
				}

				results <- result
			}
		}()
	}

	// Queue jobs (some invalid)
	jobs <- Job{1, 5}
	jobs <- Job{2, -1} // Error!
	jobs <- Job{3, 10}
	close(jobs)

	wg.Wait()
	close(results)

	// Check results
	for result := range results {
		if result.err != nil {
			fmt.Printf("  Job %d: ERROR - %v\n", result.id, result.err)
		} else {
			fmt.Printf("  Job %d: value=%d\n", result.id, result.value)
		}
	}
}

// 6. WORKER METRICS: Track worker statistics
func example6_WorkerMetrics() {
	fmt.Println("6. Worker Metrics:")

	jobs := make(chan int, 20)
	var processed int64

	var wg sync.WaitGroup

	// Workers with atomic counter
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
				atomic.AddInt64(&processed, 1)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	// Queue jobs
	for i := 1; i <= 15; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	fmt.Printf("  Processed %d jobs\n", atomic.LoadInt64(&processed))
}

// 7. PRIORITY QUEUE WORKERS: Process high-priority jobs first
func example7_PriorityQueue() {
	fmt.Println("7. Priority Queue Workers:")

	type Job struct {
		id       int
		priority int // 1=high, 2=normal, 3=low
		name     string
	}

	// For simplicity, just demonstrate concept
	jobs := []Job{
		{1, 2, "normal task"},
		{2, 1, "urgent task"},
		{3, 3, "background task"},
		{4, 1, "critical task"},
	}

	// In real code, use a heap or priority queue
	fmt.Println("  Jobs sorted by priority:")
	for _, job := range jobs {
		priority := []string{"", "🔴 HIGH", "🟡 NORMAL", "🟢 LOW"}
		fmt.Printf("  [%d] %s - %s\n", job.id, priority[job.priority], job.name)
	}
}

// 8. WORKER TIMEOUT: Cancel slow jobs
func example8_WorkerTimeout() {
	fmt.Println("8. Worker Timeout:")

	type Job struct {
		id       int
		duration time.Duration
	}

	jobs := make(chan Job, 5)
	results := make(chan string, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for job := range jobs {
			// Create timeout context
			done := make(chan bool, 1)

			go func() {
				time.Sleep(job.duration)
				done <- true
			}()

			select {
			case <-done:
				results <- fmt.Sprintf("Job %d completed", job.id)
			case <-time.After(100 * time.Millisecond):
				results <- fmt.Sprintf("Job %d TIMEOUT", job.id)
			}
		}
	}()

	jobs <- Job{1, 50 * time.Millisecond}  // Fast
	jobs <- Job{2, 150 * time.Millisecond} // Slow - timeout
	jobs <- Job{3, 50 * time.Millisecond}  // Fast
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

// 9. DYNAMIC WORKER POOL: Adjust workers based on load
func example9_DynamicWorkerPool() {
	fmt.Println("9. Dynamic Worker Pool (adjust workers):")
	fmt.Println("  Concept: Increase workers when queue grows, decrease when idle")
	fmt.Println("  - Monitor queue length (len(jobs))")
	fmt.Println("  - Spawn new workers if queue > threshold")
	fmt.Println("  - Remove workers if idle > timeout")
	fmt.Println("  Benefit: Auto-scale to workload")
}

// 10. WORKER POOL BEST PRACTICES: Summary of patterns
func example10_BestPractices() {
	fmt.Println("10. Worker Pool Best Practices:")
	fmt.Println("  ✓ Use fixed pool size for predictable resources")
	fmt.Println("  ✓ Buffer job channel appropriately")
	fmt.Println("  ✓ Always close job channel when done")
	fmt.Println("  ✓ Use WaitGroup to wait for completion")
	fmt.Println("  ✓ Handle errors gracefully")
	fmt.Println("  ✓ Monitor metrics (processed jobs, errors, latency)")
	fmt.Println("  ✓ Consider timeouts for long operations")
	fmt.Println("  ✓ Don't spawn unlimited goroutines")
}

func main() {
	fmt.Println("========== LEARNING WORKER POOLS ==========\n")

	// 1. Basic worker
	fmt.Println("--- 1. Basic Worker ---")
	example1_BasicWorker()

	// 2. Multiple workers
	fmt.Println("\n--- 2. Multiple Workers ---")
	example2_MultipleWorkers()

	// 3. Structure
	fmt.Println("\n--- 3. Worker Pool Structure ---")
	example3_WorkerPoolStructure()

	// 4. Bounded pool
	fmt.Println("\n--- 4. Bounded Worker Pool ---")
	example4_BoundedWorkerPool()

	// 5. Error handling
	fmt.Println("\n--- 5. Error Handling ---")
	example5_ErrorHandling()

	// 6. Metrics
	fmt.Println("\n--- 6. Worker Metrics ---")
	example6_WorkerMetrics()

	// 7. Priority queue
	fmt.Println("\n--- 7. Priority Queue ---")
	example7_PriorityQueue()

	// 8. Timeout
	fmt.Println("\n--- 8. Worker Timeout ---")
	example8_WorkerTimeout()

	// 9. Dynamic pool
	fmt.Println("\n--- 9. Dynamic Worker Pool ---")
	example9_DynamicWorkerPool()

	// 10. Best practices
	fmt.Println("\n--- 10. Best Practices ---")
	example10_BestPractices()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: API Request Processor ==========\n")
	apiRequestProcessorProject()
}

// ========== PROJECT: API REQUEST PROCESSOR ==========
// Real-world scenario: API server with fixed worker pool processing requests
// Shows: Worker pool for request handling, load distribution, timeouts

type Request struct {
	id       int
	url      string
	duration time.Duration // Simulated processing time
}

type RequestResult struct {
	id     int
	url    string
	status string
	time   time.Duration
}

func processRequest(req Request) RequestResult {
	start := time.Now()
	fmt.Printf("  🔄 [%d] Processing: %s\n", req.id, req.url)

	// Simulate processing
	time.Sleep(req.duration)

	result := RequestResult{
		id:     req.id,
		url:    req.url,
		status: "✓ 200 OK",
		time:   time.Since(start),
	}

	return result
}

func apiRequestProcessorProject() {
	fmt.Println("🌐 API Request Processor - Fixed worker pool\n")

	// Sample requests
	requests := []Request{
		{1, "GET /users", 100 * time.Millisecond},
		{2, "POST /login", 150 * time.Millisecond},
		{3, "GET /profile", 80 * time.Millisecond},
		{4, "PUT /settings", 120 * time.Millisecond},
		{5, "DELETE /account", 90 * time.Millisecond},
		{6, "GET /data", 110 * time.Millisecond},
	}

	for {
		fmt.Print("\nCommand (process/info/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "process":
			fmt.Print("Number of workers (1-4): ")
			var numWorkers int
			fmt.Scanln(&numWorkers)

			if numWorkers < 1 || numWorkers > 4 {
				fmt.Println("Invalid number. Using 2.")
				numWorkers = 2
			}

			fmt.Printf("\n⚙️  Starting server with %d workers...\n\n", numWorkers)

			// Create worker pool
			jobsChan := make(chan Request, 10)
			resultsChan := make(chan RequestResult, 10)
			var wg sync.WaitGroup

			startTime := time.Now()

			// Launch workers
			for w := 1; w <= numWorkers; w++ {
				wg.Add(1)
				go func(workerID int) {
					defer wg.Done()
					for req := range jobsChan {
						result := processRequest(req)
						resultsChan <- result
					}
				}(w)
			}

			// Send requests to queue
			fmt.Println("📤 Queuing incoming requests...")
			go func() {
				for _, req := range requests {
					jobsChan <- req
					fmt.Printf("  [📨] Request %d queued\n", req.id)
					time.Sleep(20 * time.Millisecond)
				}
				close(jobsChan)
			}()

			// Wait for workers to finish
			wg.Wait()
			close(resultsChan)

			// Print results
			fmt.Println("\n📊 Results:")
			totalTime := 0.0
			for result := range resultsChan {
				fmt.Printf("  [%d] %s - %s (%.0fms)\n",
					result.id, result.url, result.status, result.time.Seconds()*1000)
				totalTime += result.time.Seconds()
			}

			totalElapsed := time.Since(startTime)
			fmt.Printf("\n📈 Stats:\n")
			fmt.Printf("  Total requests: %d\n", len(requests))
			fmt.Printf("  Workers: %d\n", numWorkers)
			fmt.Printf("  Total time: %.2fs\n", totalElapsed.Seconds())
			fmt.Printf("  Avg time/request: %.2fms\n", (totalTime/float64(len(requests)))*1000)
			fmt.Printf("  Throughput: %.1f req/s\n", float64(len(requests))/totalElapsed.Seconds())

		case "info":
			fmt.Println("\n📋 Worker Pool Benefits:")
			fmt.Println("  • Fixed resource usage (no unlimited goroutines)")
			fmt.Println("  • Fair load distribution across workers")
			fmt.Println("  • Easy to monitor and control")
			fmt.Println("  • Prevents resource exhaustion")
			fmt.Println("  • Better than 1 goroutine per task for high load")
			fmt.Println("\n📝 Requests to process:")
			for _, req := range requests {
				fmt.Printf("  [%d] %s (%.0fms)\n", req.id, req.url, float64(req.duration.Milliseconds()))
			}

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Use: process, info, quit")
		}
	}
}
