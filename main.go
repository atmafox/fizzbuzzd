package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		var out string

		msg := strings.TrimSpace(string(netData))
		switch strings.ToLower(msg) {
		case "stop":
			return
		case "help":
			out = "Send a number to get fizzbuzz status, send stop to close connection.\n"
		default:
			i, err := strconv.Atoi(msg)
			if err != nil {
				out = "Not a valid number\n"
			} else {
				// only hit here if a valid number

				if i%3 == 0 && i%5 == 0 {
					out = strconv.Itoa(i) + " FizzBuzz\n"
				} else if i%3 == 0 {
					out = strconv.Itoa(i) + " Fizz\n"
				} else if i%5 == 0 {
					out = strconv.Itoa(i) + " Buzz\n"
				} else {
					out = strconv.Itoa(i) + "\n"
				}
			}
		}

		c.Write([]byte(out))
	}
}

func main() {
	flag.Parse()

	PORT := ":3666"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
