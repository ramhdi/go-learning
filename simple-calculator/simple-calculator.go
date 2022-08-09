package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin);

	// input number 1
	fmt.Print("Enter number 1: ");
	numInput1, _ := reader.ReadString('\n');
	aFloat1, err1 := strconv.ParseFloat(strings.TrimSpace(numInput1), 64);

	if err1 != nil {
		panic(err1);
	}

	// input number 2
	fmt.Print("Enter number 2: ");
	numInput2, _ := reader.ReadString('\n');
	aFloat2, err2 := strconv.ParseFloat(strings.TrimSpace(numInput2), 64);

	if err2 != nil {
		panic(err2);
	}

	// add two numbers
	res := aFloat1 + aFloat2;
	fmt.Printf("Sum: %v + %v = %v\n", aFloat1, aFloat2, res);
}