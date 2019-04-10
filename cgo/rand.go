package main

import (
	"fmt"
	"time"
)

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed((int)(time.Now().Unix()))
	fmt.Println("random:", Random())
}
