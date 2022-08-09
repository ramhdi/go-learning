package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input number 1
	fmt.Print("Enter number 1: ");
	numInput1 := readFloat();

	// input number 1
	fmt.Print("Enter number 2: ");
	numInput2 := readFloat();

	// input operator
	fmt.Print("Enter operator (+ - * /): ");
	op := readString();

	// operate two numbers
	var res float64;
	switch op {
	case "+":
		res = numInput1 + numInput2;
	case "-":
		res = numInput1 - numInput2;
	case "*":
		res = numInput1 * numInput2;
	case "/":
		res = numInput1 / numInput2;
	default:
		panic("Operation not supported");
	}

	// print result
	fmt.Printf("Sum: %v %s %v = %v\n", numInput1, op, numInput2, res);
}

func readFloat() float64 {
	//fmt.Print("Enter number 1: ");
	reader := bufio.NewReader(os.Stdin);
	numInput, _ := reader.ReadString('\n');
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64);

	if err != nil {
		panic(err);
	}
	return aFloat;
}

func readString() string {
	//fmt.Print("Enter number 1: ");
	reader := bufio.NewReader(os.Stdin);
	strInput, _ := reader.ReadString('\n');
	aOp := strings.TrimSpace(strInput)
	
	return aOp;
}