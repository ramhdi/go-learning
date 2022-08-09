package main
import "fmt"

func main() {
	var anInt int = 42;
	var aFloat float64 = 3.14159;
	// res := anInt * aFloat; // error
	res := float64(anInt) * aFloat;
	fmt.Printf("Result = %v, Type = %T", res, res);
}