package main
// import "fmt"

// // func incr(p *int) int {
// // 	*p++
// // 	return *p
// // }
// func main() {
//     s := "apple"
//     r := []rune(s)

//     fmt.Println("Rune length:", len(r))  // 5
//     fmt.Println("Rune slice:", r)        // [97 112 112 108 101] → ASCII codes
//     fmt.Println("Original string:", s)

//     fmt.Println("\n--- Iterating over string (range s) ---")
//     for i, c := range s {
//         fmt.Printf("Character: %c, Byte Index: %d, ASCII/Unicode: %v\n", c, i, c)
//     }

//     fmt.Println("\n--- Iterating over rune slice (range r) ---")
//     for i, c := range r {
//         fmt.Printf("Character: %c, Rune Index: %d, ASCII/Unicode: %v\n", c, i, c)
//     }
// }
// package main
// import (
//     "errors"
//     "fmt"
//     "encoding/json"
// //     "time"
// //     // "sync"

// // )


// var ErrInvalidInput = errors.New("invalid input")

// func validateInput(input string) error {
//     if input == "" {
//         return fmt.Errorf("validation error: %w", ErrInvalidInput)
//     }
//     return nil
// }

// // func main() {
// //     err := validateInput("")
// //     if errors.Is(err, ErrInvalidInput) {
// //         fmt.Println("Detected invalid input")
// //     }
// // }
// // type Engine struct{
// //     HorsePower int
// //     Mileage    float32
// // }
// // type Car struct{
// //     Name  string
// //     Model string
// //     Engine
// // }
// // func main() {
// //     c := Car{
// //         Name:  "Toyota",
// //         Model: "Corolla",
// //         Engine: Engine{
// //             HorsePower: 132,
// //             Mileage:    30.0,
// //         },
// //     }
// //     fmt.Printf("Car: %+v\n", c)
// //     fmt.Printf("HorsePower: %d, Mileage: %.2f\n", c.HorsePower, c.Mileage)
// // }
// type User struct {
//     Name string
//     Age  int
// }
// func (u *User) HaveBirthday() {
//     u.Age++ // modifies the original struct
// }
// func (u User) Greet() {
//     fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.Name, u.Age)
// }

// // func main() {
// //     user := User{Name: "Alice", Age: 30}
// //     fmt.Printf("Before birthday: %+v\n", user)
// //     user.HaveBirthday()
// //     fmt.Printf("After birthday: %+v\n", user)
// //     user.Greet()
// // }
// func square(n int) int {
//     return n * n
// }

// func processAndPrint(nums []int, fn func(int) int) []int {
//     result := make([]int,len(nums))
//     for i, v := range nums {
//         result[i] = fn(v)
//     }
//     return result
// }
// // func main() {
// //     // numbers := []int{1, 2, 3, 4, 5}
// //     // squared := processAndPrint(numbers, square)
// //     // fmt.Println(squared)
// //     m:=make(map[string]int)
// //     m["a"]=1
// //     m["b"]=2
// //     fmt.Println(m)
// // }   

// func check(num int)(string,error){
//     if num <0{
//         return "",errors.New("negative number")
//     }
//     return "Num is Positive",nil
// }


// func read(filename string) error {
//     // Simulate an error
//     return fmt.Errorf("simulated read error")
// }

// func file() (string, error) {
//     e := read("")
//     if e != nil {
//         return "", fmt.Errorf("file read error: %w", e) // wrap here
//     }
//     return "file read successfully", nil
// }

// func divide(a, b int) (result int) {
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("Recovered from panic:", r)
//             result = 0
//         }
//     }()
//     if b == 0 {
//         panic("division by zero") 
//     }
//     return a / b
// }

// func main(){
//     fmt.Println(divide(10,2))
//     fmt.Println(divide(10,0))
// }
// package main

// import (
// 	"fmt"
// )

// Car interface: the "contract"
// type Car interface {
// 	StartEngine() string
// 	Drive() string
// }

// // Sedan struct
// type Sedan struct {
// 	Brand  string
// 	Length int // in feet
// }

// func (s Sedan) StartEngine() string {
// 	return fmt.Sprintf("Brand %s has started the engine silently.", s.Brand)
// }

// func (s Sedan) Drive() string {
// 	return fmt.Sprintf("Brand %s can drive smoothly with a length of %d feet.", s.Brand, s.Length)
// // }

// // SUV struct
// type SUV struct {
// 	Brand  string
// 	Height int // in meters
// }

// func (s SUV) StartEngine() string {
// 	return fmt.Sprintf("Brand %s has started the engine with a roar.", s.Brand)
// }

// func (s SUV) Drive() string {
// 	return fmt.Sprintf("Brand %s can drive on any terrain with a height of %d meters.", s.Brand, s.Height)
// }

// // Polymorphic function: works with ANY Car
// func testDrive(c Car) {
// 	fmt.Println(c.StartEngine())
// 	fmt.Println(c.Drive())
// 	fmt.Println("-------------")
// }

// func main() {
// 	// Use Sedan
// 	mySedan := Sedan{Brand: "Honda", Length: 15}
// 	// Use SUV
// 	mySUV := SUV{Brand: "Jeep", Height: 2}

// 	// Both Sedan and SUV satisfy Car interface
// 	var car Car

// 	car = mySedan
// 	testDrive(car)

// 	car = mySUV
// 	testDrive(car)
// }

// func main() {
//     msg, err := file()
//     if err != nil {
//         // Wrap again in main for additional context
//         err = fmt.Errorf("main encountered an issue: %w", err)
//         fmt.Println("Error:", err)

//         // Optional: unwrap to see the original error
//         orig := errors.Unwrap(err) // one level down
//         fmt.Println("Unwrapped error:", orig)
//         return
//     }
//     fmt.Println(msg)
// }
// func main() {
//     var interfaceValue interface{}

//     // Test with an integer
//     interfaceValue = 42
//     processValue(interfaceValue)

//     // Test with a string
//     interfaceValue = "Hello, Go!"
//     processValue(interfaceValue)

//     // Test with a float
//     interfaceValue = 3.14
//     processValue(interfaceValue)

//     // Test with a boolean
//     interfaceValue = true
//     processValue(interfaceValue)
// }

// func processValue(value interface{}) {
//     switch v := value.(type) {
//     case int:
//         fmt.Println("Type is int:", v)
//     case string:
//         fmt.Println("Type is string:", v)
//     default:
//         fmt.Println("Unsupported type:", v)
//     }
// }


// func main() {
// 	// Global recovery: catches any panic in main goroutine
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from:", r)
// 		}
// 	}()

// 	err := safeProcess()
// 	if err != nil {
// 		fmt.Println("Handled error:", err)
// 	}

// 	fmt.Println("Program continues running...")
// }

// func safeProcess() error {
// 	jsonResponse := `{bad json}` // broken JSON

// 	var data map[string]interface{}
// 	err := json.Unmarshal([]byte(jsonResponse), &data)
// 	if err != nil {
// 		// Option A: just return the error (clean handling)
// 		// return err

// 		// Option B: panic so defer+recover handles it
// 		panic(err)
// 	}

// 	fmt.Println("Data:", data)
// 	return nil
// }



// func sayHello() {
//     fmt.Println("Hello from Goroutine!")
// }

// func main() {
//     go sayHello() // Starts a new goroutine
//     time.Sleep(1 * time.Second) // Allow goroutine to complete
//     fmt.Println("Main function finished")
// }

// func sendData(channel chan int) {
//     channel <- 42 // Send data into the channel
// }

// func main() {
//     channel := make(chan int,3) // Create an unbuffered channel
//     // go sendData(channel)
//     channel <- 10
//     channel <- 20
//     channel <- 30
//     fmt.Println("Data sent to channel")
//     fmt.Println(<-channel)
//     channel <- 40 // This will block if the channel is unbuffered
//     fmt.Println(<-channel) 
//     fmt.Println(<-channel)
//     fmt.Println(<-channel) // Receive data from the channel
// }

// func printNumbers() {
//     for i := 1; i <= 10; i++ {
//         fmt.Println(i)
//         time.Sleep(time.Second) // Simulating work
//     }
// }

// func main() {
//     go printNumbers()
//     fmt.Println("Started printing numbers")
//     time.Sleep(11 * time.Second) // Wait for the goroutine to finish
//     fmt.Println("Finished printing numbers")
// }


// func main() {
//     var wg sync.WaitGroup
//     wg.Add(1) // Increment the WaitGroup counter
    
//     go func() {
//         defer wg.Done() // Decrement the counter when the goroutine completes
//         printNumbers()
//     }()
    
//     wg.Wait() // Wait for all goroutines to complete
//     fmt.Println("All goroutines completed.")
// }

// func printNumbers() {
//     for i := 1; i <= 10; i++ {
//         fmt.Println(i)
//     }
// }

// func worker(i int , wg *sync.WaitGroup) {
//     defer wg.Done()
//     fmt.Printf("Worker %d starting\n", i)
//     // Simulate work
//     time.Sleep(2*time.Second)
//     fmt.Printf("Worker %d done\n", i)
// }
// func main(){
//     var wg sync.WaitGroup
//     for i := 0; i < 3; i++ {
//         wg.Add(1)
//         go worker(i, &wg)
//     }
//     wg.Wait()
//     fmt.Println("All workers completed.")
// }
// func dynamicWorker(id int, wg *sync.WaitGroup, delay time.Duration) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d: Starting\n", id)
// 	time.Sleep(delay)
// 	fmt.Printf("Worker %d: Completed after %v\n", id, delay)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	delays := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second}

// 	// Add workers dynamically
// 	for i, delay := range delays {
// 		wg.Add(1)
// 		go dynamicWorker(i+1, &wg, delay)
// 	}
// 	// Wait for all workers to finish
// 	wg.Wait()

// 	fmt.Println("All workers with dynamic delays completed.")
// }
// var (
// 	counter  int
// 	rwMutex  sync.RWMutex
// 	wg       sync.WaitGroup
// )

// func readCounter() {
// 	defer wg.Done()
// 	rwMutex.RLock()
// 	fmt.Printf("Reading Counter: %d\n", counter)
// 	rwMutex.RUnlock()
// }

// func writeCounter() {
// 	defer wg.Done()
// 	rwMutex.Lock()
// 	counter++
//     fmt.Printf("Incremented Counter to: %d\n", counter)
// 	rwMutex.Unlock()
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go writeCounter()
// 	}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go readCounter()
// 	}

// 	wg.Wait()
// }

// package main

// import (
//     // "errors"
//     "fmt"
//     // "encoding/json"
//     "time"
//     "sync"

// )

// func main() {
// 	ch := make(chan int, 2)

// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// Sender goroutine
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Sending message 1")
// 		ch <- 1
// 		fmt.Println("Sending message 2")
// 		ch <- 2
// 		fmt.Println("Attempting to send message 3 (will wait)")
// 		ch <- 3 // blocks until space frees up
// 		fmt.Println("Message 3 sent")
// 	}()

// 	// Simulate receiver delay
// 	time.Sleep(3 * time.Second)

// 	// Receiver
// 	fmt.Println("Receiving message:", <-ch)
// 	fmt.Println("Receiving message:", <-ch)
// 	fmt.Println("Receiving message:", <-ch)

// 	// Wait for sender goroutine to finish printing
// 	wg.Wait()
// }
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type BufferedChan struct {
// 	mu       sync.Mutex
// 	notEmpty *sync.Cond
// 	notFull  *sync.Cond
// 	buffer   []int
// 	capacity int
// }

// func NewBufferedChan(cap int) *BufferedChan {
// 	bc := &BufferedChan{capacity: cap}
// 	bc.notEmpty = sync.NewCond(&bc.mu)
// 	bc.notFull = sync.NewCond(&bc.mu)
// 	return bc
// }

// func (bc *BufferedChan) Send(val int) {
// 	bc.mu.Lock()
// 	for len(bc.buffer) == bc.capacity {
// 		bc.notFull.Wait()
// 	}
// 	bc.buffer = append(bc.buffer, val)
// 	bc.notEmpty.Signal()
// 	bc.mu.Unlock()
// }

// func (bc *BufferedChan) Recv() int {
// 	bc.mu.Lock()
// 	for len(bc.buffer) == 0 {
// 		bc.notEmpty.Wait()
// 	}
// 	val := bc.buffer[0]
// 	bc.buffer = bc.buffer[1:]
// 	bc.notFull.Signal()
// 	bc.mu.Unlock()
// 	return val
// }

// func main() {
// 	ch := NewBufferedChan(2)

// 	go func() {
// 		fmt.Println("Sending 1")
// 		ch.Send(1)
// 		fmt.Println("Sending 2")
// 		ch.Send(2)
// 		fmt.Println("Sending 3 (waits until receiver consumes)")
// 		ch.Send(3)
// 		fmt.Println("Sent 3")
// 	}()

// 	fmt.Println("Received:", ch.Recv())
// 	fmt.Println("Received:", ch.Recv())
// 	fmt.Println("Received:", ch.Recv())
// }
// import (
//     "fmt"
//     "sync"
//     "time"
// )

// // Worker function that processes tasks from a task channel and sends results on a result channel
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
//     for task := range tasks {
//         result := processTask(task)
//         results <- result
//         fmt.Printf("Worker %d processed task %d\n", id, task)
//     }
//     wg.Done()
// }

// // Simulates task processing
// func processTask(task int) int {
//     time.Sleep(time.Second)  // Simulate time-consuming task
//     return task * 2
// }

// func main() {
//     const numWorkers = 3
//     const numTasks = 5

//     tasks := make(chan int, numTasks)   // Task channel (receive-only for workers)
//     results := make(chan int, numTasks) // Result channel (send-only for workers)
//     var wg sync.WaitGroup

//     // Start worker goroutines
//     for i := 1; i <= numWorkers; i++ {
//         wg.Add(1)
//         go worker(i, tasks, results, &wg)
//     }

//     // Send tasks to the worker pool
//     for j := 1; j <= numTasks; j++ {
//         tasks <- j
//     }
//     close(tasks)  // Close tasks channel when no more tasks are left to send

//     // Wait for all workers to finish
//     wg.Wait()
//     close(results)  // Close results channel after workers are done

//     // Print results
//     for result := range results {
//         fmt.Println("Result:", result)
//     }
// }

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     c1 := make(chan string)
//     c2 := make(chan string)
// 	c3 := make(chan string)

//     go func() {
//         time.Sleep(1 * time.Second)
//         c1 <- "one"
//     }()

//     go func() {
//         time.Sleep(2 * time.Second)
//         c2 <- "two"
//     }()
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		c3 <- "three"
// 	}()

//     for i := 0; i < 4; i++ {
//         select {
//         case msg1 := <-c1:
//             fmt.Println("Received on Channel 1", msg1)
//         case msg2 := <-c2:
//             fmt.Println("Received on Channel 2" , msg2)
// 		case msg3 := <-c3:
// 			fmt.Println("Received on Channel 3", msg3)
//         case <-time.After(1000 * time.Millisecond):
//             fmt.Println("timeout")
//         }
//     }
// }
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     messages := make(chan string,1)
//     signals := make(chan bool,1)

//     // Goroutine that sends a signal after 1 second
//     go func() {
//         time.Sleep(1 * time.Second)
//         signals <- true
//     }()

//     // First select: nothing in messages yet
//     select {
//     case msg := <-messages:
//         fmt.Println("received message:", msg)
//     default:
//         fmt.Println("no message received")
//     }

//     msg := "hi"
//     // Try to send a message (still no receiver)
//     select {
//     case messages <- msg:
//         fmt.Println("sent message:", msg)
//     default:
//         fmt.Println("no message sent")
//     }

//     // Give time for goroutine to send signal
//     time.Sleep(2 * time.Second)

//     // Now check channels again
//     select {
//     case msg := <-messages:
//         fmt.Println("received message:", msg)
//     case sig := <-signals:
//         fmt.Println("received signal:", sig)
//     default:
//         fmt.Println("no activity")
//     }
// }

import (
    "fmt"
    "sync"
    // "time"
	"net/http"
)

// Worker consumes jobs from `in` channel and sends results to `out` channel
// func worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
//     defer wg.Done() // always mark worker done at exit

//     for job := range in {
//         fmt.Printf("worker %d started job %d\n", id, job)
//         time.Sleep(time.Second) // simulate time-consuming task
//         fmt.Printf("worker %d finished job %d\n", id, job)
//         out <- job * 2 // send result to out channel
//     }
// }

// func main() {
//     const numJobs = 5
//     const numWorkers = 3

//     in := make(chan int, numJobs)   // channel carrying jobs
//     out := make(chan int, numJobs)  // channel carrying results

//     var wg sync.WaitGroup

//     // spin up worker goroutines
//     for w := 1; w <= numWorkers; w++ {
//         wg.Add(1)
//         go worker(w, in, out, &wg)
//     }

//     // push jobs into `in`
//     for j := 1; j <= numJobs; j++ {
//         in <- j
//     }
//     close(in) // no more jobs will be sent

//     // wait for all workers to finish
//     wg.Wait()
//     close(out) // no more results will be produced

//     // collect results
//     for result := range out {
//         fmt.Println("result:", result)
//     }
// }
// func worker(id int ,in <-chan int , out chan<-int , wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for job:=range in{
// 		fmt.Println("Worker",id,"started job",job)
// 		time.Sleep(time.Second)
// 		fmt.Println("Worker",id,"finished job",job)
// 		out <- job*2
// 	}
// }
// func main(){
// 	const numJobs=5
// 	const numWorkers=3
// 	in:=make(chan int , numJobs)
// 	out:=make(chan int , numJobs)
// 	var wg sync.WaitGroup
// 	for w:=1;w<=numWorkers;w++{
// 		wg.Add(1)
// 		go worker(w,in,out,&wg)
// 	}
// 	for j:=1;j<=numJobs;j++{
// 		in<-j
// 	}
// 	close(in)
// 	wg.Wait()
// 	close(out)
// 	for result:=range out{
// 		fmt.Println("Result:",result)
// 	}
// }
func fetchData(url string, wg *sync.WaitGroup, ch chan<- string) {
    defer wg.Done()  // Ensure Done is called when the worker finishes

    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
        return
    }
    ch <- fmt.Sprintf("Fetched %s with status: %s", url, resp.Status)
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan string, 2)  // Buffered channel to hold the results

    urls := []string{
        "https://api.example1.com/data",
        "https://api.example2.com/data",
    }

    // Launch goroutines to fetch data concurrently
    for _, url := range urls {
        wg.Add(1)
        go fetchData(url, &wg, ch)
    }

    wg.Wait()  // Wait for all goroutines to finish
    close(ch)  // Close the channel after all goroutines are done

    // Print the results
    for msg := range ch {
        fmt.Println(msg)
    }
}