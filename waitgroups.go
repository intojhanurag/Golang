// package main

// import (
//     "fmt"
//     "sync"
//     "time"
// )

// func worker(id int) {
//     fmt.Printf("Worker %d starting\n", id)
//     time.Sleep(time.Second)
//     fmt.Printf("Worker %d done\n", id)
// }

// func main() {
//     var wg sync.WaitGroup

//     for i := 1; i <= 5; i++ {
//         wg.Add(1)

//         go func(id int) {
//             defer wg.Done()
//             worker(id)
//         }(i) // 👈 pass current value of i to goroutine
//     }

//     wg.Wait()
// }


package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    fmt.Println("Main: Starting 2 students")

    wg.Add(2) // Waiting for 2 students

    go func() {
        fmt.Println("Student 1: Doing homework")
        time.Sleep(2 * time.Second)
        fmt.Println("Student 1: Done")
        wg.Done() // 1 student done
    }()

    go func() {
        fmt.Println("Student 2: Doing homework")
        time.Sleep(1 * time.Second)
        fmt.Println("Student 2: Done")
        wg.Done() // 2nd student done
    }()

    fmt.Println("Main: Waiting for students to finish...")
    wg.Wait() // ⏳ Wait until both students call Done()

    fmt.Println("Main: All students finished, leaving class.")
}
