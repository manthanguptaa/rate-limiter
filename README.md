## Rate Limiting

**Description**

This project demonstrates two popular rate limiting algorithms in Go: the token bucket and the fixed window counter. It provides practical examples through HTTP handlers, allowing you to explore their behavior and configure them for different rate limiting needs.

**Key Features**

* **Dual Algorithms:** Choose between the flexible token bucket approach with gradual refills or the fixed window counter for burst protection within a specific timeframe.
* **IP-Based Tracking:** Monitor and limit requests based on client IP addresses.
* **Thread-Safe Design:** Utilizes mutexes for concurrent access to internal data structures, ensuring safe operation in multi-threaded environments.
* **Configurability:** Customize rate limits through adjustable constants for tokens, refill rate, window size, and maximum requests.

**Usage Instructions**

1. Clone this repository.
2. Run the code using `go run main.go`.
3. Access the functionalities via HTTP requests:
    * `/token-bucket`: Simulate token bucket behavior with adjustable refill rate and token capacity.
    * `/fixed-window`: Test the fixed window counter algorithm with configurable window size and maximum allowed requests within that timeframe.
4. Observe the server responses to understand whether requests are allowed or rejected based on the chosen algorithm and its current state.

**Configuration**

Modify the constants in `token_bucket.go` and `fixed_window_counter.go` to tailor the rate limiting behavior:

* **Token Bucket:**
    * `maxTokens`: Adjust the initial and maximum number of tokens available.
    * `refillRate`: Define the rate at which the bucket refills with tokens per second.
* **Fixed Window Counter:**
    * `windowSize`: Set the duration of the window in seconds, defining the timeframe for counting requests.
    * `maxRequest`: Indicate the maximum number of requests permitted within the specified window.

**Additional Information**

* This project serves as a basic example and can be further extended to integrate with real-world applications and authentication mechanisms.
