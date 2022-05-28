package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// this sample demonstrates
// how to use a channel_demo to monitor the amount of time the program is running
// and terminate the program if it runs too long


const maxtime = 3 * time.Second

// use four channels : 3 unbuffered and 1 buffered of
var(
	// sigChan receives operating signals.
	// This will allow us to send a Ctrl-C to shut down our program cleanly.
	sigChan = make(chan os.Signal , 1)

	// timeout limits the amount of time the program has.
	// We really don't want to receive on this channel_demo because if we do, that means something bad
	// happens, we are timing out and we need to kill the program.
	timeout = time.After(maxtime)

	// complete is used to report processing is done.
	// This is the channel_demo we want to receive on. When the Goroutine finish the job, it will signal
	// to us on this complete channel_demo and tell us any error that occurred.
	complete = make(chan error)

	// shutdown provides system wide notification.
	shutdown = make(chan struct{})
)

func main() {

	log.Println("Starting Process")

	// receive all interrupt based signals
	// why we use buffered channel_demo of 1 ?
	signal.Notify(sigChan, os.Interrupt)

	// Launch the process.
	log.Println("Launching Processors")

	// this goroutine will do the processing job
	// 开启一个协程去完成相对应的工作
	go processor(complete)


	// main goroutine here is in the event loop
	// there are 3 cases in select: sigChan, timeout, and complete.
ControlLoop:
	for {
		select {
		case <-sigChan:
			// Interrupt event signaled by the operation system.
			log.Println("OS INTERRUPT")

			// Close the channel_demo to signal to the processor it needs to shutdown.
			close(shutdown)

			// Set the channel_demo to nil so we no longer process any more of these events.
			// If we try to send on a closed channel_demo, we are gonna panic. If we receive on a closed
			// channel_demo, that's gonna immediately return a signal without data. If we receive on a
			// nil channel_demo, we are blocked forever. Similar with send.
			// Why do we want to do that?
			// We don't want user to hold down Ctrl C or hit Ctrl C multiple times. If they do that
			// and we process the signal, we have to call close multiple time. When we call close
			// on a channel_demo that is already closed, the code will panic. Therefore, we cannot have
			// that.
			sigChan = nil

		case <-timeout:
			// We have taken too much time. Kill the app hard.
			log.Println("Timeout - Killing Program")

			// os.Exit will terminate the program immediately.
			os.Exit(1)

		case err := <-complete:
			// Everything completed within the time given.
			log.Printf("Task Completed: Error[%s]", err)

			// We are using a label break here.
			// We put one at the top of the for loop so the case has a break and the for has a
			// break.
			break ControlLoop
		}
	}

	// Program finished.
	log.Println("Process Ended")


}

// chan <-  means this channel_demo is a send-only channel_demo
// if we try to receive on this channel_demo, the compiler will give us an error
func processor(complete chan <- error) {

	log.Println("Processor - Starting")

	// Variable to store any error that occurs.
	// Passed into the defer function via closures.
	var err error

	// Defer the send on the channel_demo so it happens regardless of how this function terminates.
	// This is an anonymous function call like we saw with Goroutine. However, we are using the
	// keyword defer here.
	// We want to execute this function but after the processor function returns. This gives us an
	// guarantee that we can have certain things happen before control go back to the caller.
	// Also, defer is the only way to stop a panic. If something bad happens, say the image library
	// is blowing up, that can cause a panic situation throughout the code. In this case, we want
	// to recover from that panic, stop it and then control the shutdown.
	defer func() {
		// Capture any potential panic.
		if r := recover(); r != nil {
			log.Println("Processor - Panic", r)
		}

		// Signal the Goroutine we have shutdown.
		complete <- err
	}()

	// Perform the work.
	err = doWork()

	log.Println("Processor - Completed")

}


// doWork simulates task work.
// Between every single call, we call checkShutdown. After complete every tasks, we are asking:
// Have we been asked to shutdown? The only way we know is that shutdown channel_demo is closed. The
// only way to know if the shutdown channel_demo is closed is to try to receive. If we try to receive on
// a channel_demo that is not closed, it's gonna block. However, the default case is gonna save us here.
func doWork() error {
	log.Println("Processor - Task 1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}

// checkShutdown checks the shutdown flag to determine if we have been asked to interrupt processing.
func checkShutdown() bool {
	select {
	case <-shutdown:
		// We have been asked to shutdown cleanly.
		log.Println("checkShutdown - Shutdown Early")
		return true

	default:
		// If the shutdown channel_demo was not closed, presume with normal processing.
		return false
	}
}


// Output:
// -------
// - When we let the program run, since we configure the timeout to be 3 seconds, it will
// then timeout and be terminated.
// - When we hit Ctrl C while the program is running, we will see the OS INTERRUPT and the program
// is being shutdown early.
// - When we send a signal quit by hitting Ctrt \, we will get a full stack trace of all the
// Goroutines.