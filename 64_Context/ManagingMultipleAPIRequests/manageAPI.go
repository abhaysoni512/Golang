// Consider a scenario where you need to fetch data from multiple APIs concurrently. By using context, you can ensure that all the API requests are canceled if any of them exceeds a specified timeout.

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	/*The NewRequestWithContext function is used to create a client-side (outgoing) HTTP request that is tightly bound to a context.Context.
	The http.Request object (often referred to simply as the Request object) is a fundamental structure in Go's net/http package. It represents an HTTP request—the message sent from a client (like a web browser or your Go program) to a server.
	*/
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s : %s", url, err.Error)
		return
	}

	/*
		http.DefaultClient is a pre-configured global instance of the http.Client type provided by the Go standard library. It is a convenience variable that provides sensible defaults for making outgoing HTTP requests, such as:

		Timeout: It has no default timeout, meaning it will wait indefinitely for a response unless you explicitly set a timeout on the context or the client itself.

		Transport: It uses the default http.Transport, which handles connection pooling, proxy settings, and basic network setup.

	*/
	client := http.DefaultClient

	/*
		🚀 Executing the Request (client.Do(req))
		Go

		resp, err := client.Do(req)
		What is client.Do(req)?
		The Do method is the heart of the Go HTTP client. It takes the Request object (req) you created and performs the following sequence of actions:

		Handles Context: It respects the ctx attached to the request, monitoring it for timeouts or cancellations.

		Network Connection: It establishes or reuses a network connection (TCP/TLS) to the server specified in the request's URL.

		Sending: It sends the request line, headers, and any body data across the connection.

		Receiving: It waits for and reads the server's response.

		Parsing: It parses the raw response data into the usable http.Response object.

		resp (*http.Response): This is the pointer to the Response object—the message received back from the server. It contains the status code, response headers, and the response body (which must be read).

		err (error): This is an error that occurs during the execution of the request (e.g., a network error, DNS lookup failure, or a context timeout). It is not an HTTP error status code (like 404 or 500).

	*/

	resp, err := client.Do(req)

	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	/*
			When you execute an HTTP request using client.Do(req), the server's response is returned as an http.Response object, which contains the response body as an io.ReadCloser (accessible via resp.Body).

		The data from the server (the response body) arrives on a network connection (a TCP socket).

		The connection remains open and busy until the entire response body is read or until the body is explicitly closed.

		If you forget to close the body, the Go runtime's http.Client cannot put that underlying network connection back into its connection pool for future use. Over time, making many requests without closing the body leads to a "leak" of connections. Eventually, the system runs out of available file descriptors (sockets), and your application will fail to make new requests.
	*/
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s : %d", url, resp.StatusCode)

}

func main() {
	/*
		WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
		WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

		Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this [Context] complete:

		func slowOperationWithTimeout(ctx context.Context) (Result, error) {
			ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
			defer cancel()  // releases resources if slowOperation completes before timeout elapses
			return slowOperation(ctx)
		}
	*/
	/*
			func context.Background() context.Context
			Background returns a non-nil, empty [Context]. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
		/*


	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://pkg.go.dev/",
		"https://go.dev/learn/",
		"https://go.dev/ref/spec",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}

}
