package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {

    s := "postgres://user:pass@host.com:5432/path?k=u&k=v&r=z&r=s#f"
	// s= scheme://user:password@host:port/path?query#fragment

	// fragment means: Jump to a section of a webpage.
	// ParseQuery turns the raw query into a map[string][]string

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}