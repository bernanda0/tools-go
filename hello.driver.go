package main

import (
	"bernanda/learn/hello"
	"fmt"
	"time"
)

func runHello() {
	key := hello.GetKey()
	b := byte(32)

	for i := 0; i < len(key); {
		fmt.Print(string(b))
		time.Sleep(time.Millisecond)

		if b == key[i] {
			b = 32
			i++
			continue
		}

		fmt.Print("\b")
		b++
	}
}
