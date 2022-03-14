package main

import (
	"fmt"
	"net"
)

func worker(ports, results chan int) {
	url := "https://scanme.nmap.org:"

	for i := range ports {
		str := fmt.Sprintf("%s%d", url, i)
		conn, err := net.Dial("tcp", str)

		if err != nil {
			results <- 0
			continue
		}

		results <- i
		conn.Close()
	}
}

func main() {
	var ports = make(chan int, 100)
	var results = make(chan int)
	defer close(ports)
	defer close(results)

	for i := 

	for i := 1; i <= 1024; i++ {
		ports <- i
	}

}
