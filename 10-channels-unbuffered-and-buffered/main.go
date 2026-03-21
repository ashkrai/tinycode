package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== SIMPLE CHANNELS LEARNING EXAMPLES ==========

// 1. UNBUFFERED CHANNEL: Send blocks until receiver is ready
// Synchronous communication - sender waits for receiver
func example1_UnbufferedBasic() {
	fmt.Println("1. Unbuffered Channel - Blocking behavior:")

	ch := make(chan string) // No capacity = unbuffered

	go func() {
		fmt.Println("  Goroutine: About to send message")
		ch <- "Hello" // BLOCKS here until receiver is ready
		fmt.Println("  Goroutine: Message sent!")
	}()

	time.Sleep(500 * time.Millisecond) // Delay to show blocking
	fmt.Println("  Main: About to receive")
	msg := <-ch // Receive unblocks the send
	fmt.Println("  Main: Received -", msg)
}

// 2. BUFFERED CHANNEL: Send doesn't block until buffer is full
// Asynchronous communication - sender can continue
func example2_BufferedBasic() {
	fmt.Println("2. Buffered Channel - Buffer capacity:")

	ch := make(chan string, 2) // Capacity = 2

	fmt.Println("  Sending first message...")
	ch <- "Message 1" // Doesn't block - buffer has space
	fmt.Println("  Sent! (no wait)")

	fmt.Println("  Sending second message...")
	ch <- "Message 2" // Doesn't block - buffer still has space
	fmt.Println("  Sent! (no wait)")

	fmt.Println("  Receiving messages...")
	fmt.Println("  Received:", <-ch)
	fmt.Println("  Received:", <-ch)
}

// 3. UNBUFFERED DEADLOCK: Common mistake!
func example3_UnbufferedDeadlock() {
	fmt.Println("3. Unbuffered Channel - Deadlock example:")
	fmt.Println("  (Demonstrating with recover)")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  ✓ Caught: All goroutines are asleep - DEADLOCK!")
		}
	}()

	// Example: This would deadlock:
	//   ch := make(chan string)
	//   ch <- "test" // Blocks - no receiver!
	// Main goroutine would freeze forever waiting
	fmt.Println("  ✓ Lesson: Unbuffered channels BLOCK until both send and receive")
}

// 4. BUFFERED vs UNBUFFERED: Performance difference
func example4_BufferPerformance() {
	fmt.Println("4. Performance difference:")

	// Unbuffered - each send waits
	fmt.Println("  Testing unbuffered channel...")
	startUnbuf := time.Now()
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch // Receive all values
		}
	}()

	for i := 0; i < 100; i++ {
		ch <- i // Each send waits for receive
	}
	wg.Wait()
	unbufTime := time.Since(startUnbuf)
	fmt.Printf("    Unbuffered time: %v\n", unbufTime)

	// Buffered - fills buffer then continues
	fmt.Println("  Testing buffered channel...")
	startBuf := time.Now()
	chBuf := make(chan int, 100) // Large buffer

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-chBuf // Receive all values
		}
	}()

	for i := 0; i < 100; i++ {
		chBuf <- i // Fills buffer quickly, doesn't wait
	}
	wg.Wait()
	bufTime := time.Since(startBuf)
	fmt.Printf("    Buffered time: %v\n", bufTime)
}

// 5. CHANNEL LEN and CAP: Check buffer state
func example5_LenAndCap() {
	fmt.Println("5. Channel len() and cap():")

	ch := make(chan int, 5)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Printf("  Capacity: %d (max values can hold)\n", cap(ch))
	fmt.Printf("  Length: %d (current values in buffer)\n", len(ch))

	<-ch // Remove one value

	fmt.Printf("  After receiving one: len=%d, cap=%d\n", len(ch), cap(ch))
}

// 6. MULTIPLE SENDERS: Multiple goroutines sending on same channel
func example6_MultipleSenders() {
	fmt.Println("6. Multiple senders (buffered):")

	ch := make(chan string, 10) // Buffer to handle multiple sends

	// Launch 3 sender goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 2; j++ {
				ch <- fmt.Sprintf("Sender %d - Message %d", id, j)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(500 * time.Millisecond)

	// Receive all messages
	fmt.Println("  Received messages:")
	for i := 0; i < 6; i++ {
		fmt.Printf("    %s\n", <-ch)
	}
}

// 7. MULTIPLE RECEIVERS: Multiple goroutines receiving from same channel
func example7_MultipleReceivers() {
	fmt.Println("7. Multiple receivers (buffered):")

	ch := make(chan int, 10)

	// Fill channel with values
	for i := 1; i <= 10; i++ {
		ch <- i * 10
	}

	// Launch 3 receiver goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for val := range ch {
				fmt.Printf("  Receiver %d got: %d\n", id, val)
			}
		}(i)
	}

	// Close channel when done (important!)
	// Wait a bit then close
	time.Sleep(100 * time.Millisecond)
	close(ch)

	time.Sleep(100 * time.Millisecond) // Let receivers finish
}

// 8. CLOSE CHANNEL: Signal that no more values will be sent
func example8_CloseChannel() {
	fmt.Println("8. Closing channels:")

	ch := make(chan string, 3)

	ch <- "First"
	ch <- "Second"
	close(ch) // No more values will be sent

	// Receiving after close still works
	fmt.Println("  After close:")
	fmt.Println("    Received:", <-ch) // "First"
	fmt.Println("    Received:", <-ch) // "Second"
	fmt.Println("    Received:", <-ch) // "" (zero value for string)

	// Range detects close and stops
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	fmt.Println("  Range over closed channel:")
	for val := range ch2 {
		fmt.Printf("    %d\n", val)
	}
	fmt.Println("    (range stops at close)")
}

// 9. SELECT with channels: Choose between multiple channel operations
func example9_SelectChannels() {
	fmt.Println("9. Select with multiple channels:")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	fmt.Println("  Receiving from first ready channel:")
	select {
	case msg := <-ch1:
		fmt.Printf("    Got ch1: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("    Got ch2: %s\n", msg)
	}

	time.Sleep(100 * time.Millisecond)
	select {
	case msg := <-ch1:
		fmt.Printf("    Got ch1: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("    Got ch2: %s\n", msg)
	}
}

// 10. PRACTICAL COMPARISON: Unbuffered vs Buffered for work queue
func example10_WorkQueueComparison() {
	fmt.Println("10. Work Queue - Unbuffered vs Buffered:")

	// Unbuffered - worker processes immediately
	fmt.Println("  Unbuffered (immediate processing):")
	workChanUnbuf := make(chan int) // 0 buffer

	go func() {
		for work := range workChanUnbuf {
			fmt.Printf("    Processing: %d\n", work)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	start := time.Now()
	for i := 1; i <= 3; i++ {
		workChanUnbuf <- i // Blocks until worker receives
	}
	close(workChanUnbuf)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("    Time: %v\n", time.Since(start))

	// Buffered - queues jobs
	fmt.Println("  Buffered (queue jobs, then process):")
	workChanBuf := make(chan int, 10) // Buffer 10 items

	go func() {
		for work := range workChanBuf {
			fmt.Printf("    Processing: %d\n", work)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	start = time.Now()
	for i := 1; i <= 3; i++ {
		workChanBuf <- i // Doesn't block - goes to buffer
	}
	fmt.Println("    (all jobs queued immediately)")
	close(workChanBuf)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("    Time: %v\n", time.Since(start))
}

func main() {
	fmt.Println("========== LEARNING CHANNELS: UNBUFFERED vs BUFFERED ==========\n")

	// 1. Unbuffered basic
	fmt.Println("--- 1. Unbuffered Channel ---")
	example1_UnbufferedBasic()

	// 2. Buffered basic
	fmt.Println("\n--- 2. Buffered Channel ---")
	example2_BufferedBasic()

	// 3. Deadlock
	fmt.Println("\n--- 3. Deadlock Example ---")
	example3_UnbufferedDeadlock()

	// 4. Performance
	fmt.Println("\n--- 4. Performance Comparison ---")
	example4_BufferPerformance()

	// 5. Len and cap
	fmt.Println("\n--- 5. Channel len() and cap() ---")
	example5_LenAndCap()

	// 6. Multiple senders
	fmt.Println("\n--- 6. Multiple Senders ---")
	example6_MultipleSenders()

	// 7. Multiple receivers
	fmt.Println("\n--- 7. Multiple Receivers ---")
	example7_MultipleReceivers()

	// 8. Close channel
	fmt.Println("\n--- 8. Closing Channels ---")
	example8_CloseChannel()

	// 9. Select
	fmt.Println("\n--- 9. Select Statement ---")
	example9_SelectChannels()

	// 10. Work queue
	fmt.Println("\n--- 10. Work Queue Comparison ---")
	example10_WorkQueueComparison()

	// ========== PROJECT ==========
	fmt.Println("\n========== PROJECT: Email Notification Queue System ==========\n")
	emailQueueProject()
}

// ========== PROJECT: EMAIL NOTIFICATION QUEUE SYSTEM ==========
// Real-world scenario: Send emails through a queue system
// Shows practical difference between unbuffered (immediate) and buffered (queue) channels

type Email struct {
	id        int
	recipient string
	subject   string
	body      string
}

func sendEmail(email Email) {
	// Simulate email sending delay
	time.Sleep(time.Duration(50+email.id*10) * time.Millisecond)
	fmt.Printf("  📧 Sent to %s - Subject: %s\n", email.recipient, email.subject)
}

func emailQueueProject() {
	fmt.Println("📮 Email Notification Queue System\n")

	emails := []Email{
		{1, "alice@example.com", "Welcome", "Welcome to our service!"},
		{2, "bob@example.com", "Update", "New features available"},
		{3, "charlie@example.com", "Alert", "Security update needed"},
		{4, "diana@example.com", "Reminder", "Subscription expires soon"},
		{5, "eve@example.com", "Notification", "New message received"},
	}

	for {
		fmt.Print("\nCommand (send-unbuffered/send-buffered/list/quit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "send-unbuffered":
			fmt.Println("\n🚀 Sending emails UNBUFFERED (immediate, waits for each):")
			startTime := time.Now()

			emailChan := make(chan Email) // UNBUFFERED - no capacity

			go func() {
				for email := range emailChan {
					sendEmail(email)
				}
			}()

			// Each send BLOCKS until receiver processes it
			for _, email := range emails {
				fmt.Printf("[%d] Queuing email to %s...\n", email.id, email.recipient)
				emailChan <- email // Blocks until processed
				fmt.Println("    ✓ Email processed immediately")
			}

			close(emailChan)
			time.Sleep(500 * time.Millisecond)
			totalTime := time.Since(startTime)
			fmt.Printf("\n⏱️  Total time (unbuffered): %.2fs\n", totalTime.Seconds())

		case "send-buffered":
			fmt.Println("\n🚀 Sending emails BUFFERED (queue all, process in background):")
			startTime := time.Now()

			emailChan := make(chan Email, len(emails)) // BUFFERED - capacity = number of emails

			go func() {
				for email := range emailChan {
					sendEmail(email)
				}
			}()

			// All sends are fast - just fill the buffer
			for _, email := range emails {
				fmt.Printf("[%d] Queuing email to %s...\n", email.id, email.recipient)
				emailChan <- email // Doesn't block - goes to buffer
				fmt.Println("    ✓ Email queued instantly")
			}

			fmt.Println("\n✓ All emails queued! Processing in background...")
			close(emailChan)
			time.Sleep(500 * time.Millisecond)
			totalTime := time.Since(startTime)
			fmt.Printf("\n⏱️  Total time (buffered): %.2fs\n", totalTime.Seconds())

		case "list":
			fmt.Println("\n📋 Email queue contents:")
			for _, email := range emails {
				fmt.Printf("  [%d] To: %s\n      Subject: %s\n", email.id, email.recipient, email.subject)
			}

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Use: send-unbuffered, send-buffered, list, quit")
		}
	}
}
