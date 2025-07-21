package main

import (
    "fmt"
    "time"
)


// Unix time is just a number — it counts how many seconds have passed since:

// 📅 1 Jan 1970 00:00:00 UTC

// This date is called the Unix Epoch.


func main() {

    now := time.Now()
    fmt.Println(now)

    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}