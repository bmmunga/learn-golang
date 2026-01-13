package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)


	// help control the communication between different processes, allowing us to avoid a race condition bug
	// Channels are concurrent & allow for parallel execution of multiple goroutines
	// This creates an unbuffered channel of result structs
	// Sends & receives on this channel will block until the other side is ready
	resultChannel := make(chan result)


	// We normally wait for a func to return/process to finish(It blocking)
	// An op that doesn't block will run in a separate process called a "goroutine"
	// Only way to start a goroutine is to put go im front of a function call
	// Below uses an anonymous function

	for _, url := range urls {
		go func(u string) {
	// We send a result struct for each call to wc to resultChannel with a "send statement"
			resultChannel <- result{u, wc(u)}
		}(url) //passing url as an arg solves issue of goroutine closure
	}

	// We give each anony func a param cause it was taking a reference to the "url" var.
	// u is a copy of the value of url & fixed as value of url for iteration
	//this "()"makes goroutine execute the same time they are declared

	for i := 0; i < len(urls); i++ {
	// We use a receive expression (<-) which we assign a value received from the channel to a var
	// Receiving op blocks until a value is available on the channel
		r := <-resultChannel
		// Map happen in a single receive loop, avoiding concurrent map writes and therefore, race conditions
		results[r.string] = r.bool
	}
	return results
}

/* go test -race */ //Helps us debug problems with concurrent code


/*
Channels: Channels are the primary way that goroutines communicate safely
and synchronize their execution. You can send values into a channel and
receive those values through another goroutine. By default, sends and receives block until the other side is ready
allowing goroutines to synchronize without explicit locks or condition variables.
*/
