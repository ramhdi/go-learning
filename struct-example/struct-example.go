package main
import "fmt"

func main() {
	adek := Cat{"Adek", 4, 60};
	fmt.Println(adek);
	fmt.Printf("%T\n", adek);
}


type Cat struct {
	Name string;
	Weight int;
	Length int;
}