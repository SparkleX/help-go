package main
import (
	"fmt"
)
type Animal struct {
	Name string
	Age  int
}
func (a* Animal) GetAge() int {
	return a.Age
}
type Dog struct {
	Animal
}
func (a* Dog) GetAge() int {
	return a.Age+100
}
func main() {
	dog := Dog{Animal{"dog", 10}}
	fmt.Println(dog.Age)
	fmt.Println(dog.GetAge())
}