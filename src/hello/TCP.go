package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handle(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		// removing ' '
		temp := strings.TrimSpace(string(netData))
		fmt.Println(temp)
		if temp == "STOP" {
			break
		}
		res := strconv.Itoa(random()) + "\n"
		// go equivalent of write
		c.Write([]byte(string(res)))
	}

	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Please Provide us Port No.")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Evaluated LIFO only when all's done
	defer l.Close()

	//rand method
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(c)
	}
}
