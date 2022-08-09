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
	fmt.Print("Enter number: ");
	numInput, _ := reader.ReadString('\n');
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64);

	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println("Input float: ", aFloat);
		fmt.Printf("Input type: %T\n", aFloat);
	}
}