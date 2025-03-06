package main
import (
	"fmt"
	"math/rand"
)
func main() {
	num := rand.Intn(10)
	if num > 5 {
		fmt.Println("The number is greater than 5.")
	} else if num < 5 {
		fmt.Println("The number is less than 5.")
	} else {
		fmt.Println("The number is equal to 5.")
	}
}
