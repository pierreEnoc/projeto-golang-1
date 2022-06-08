package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Bomdia")
case t.Hour() < 18:
	fmt.Println("Boa tarde.")

default:
	fmt.Println("Boa noite")
}

}