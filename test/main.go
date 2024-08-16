package main

import "fmt"


type Car struct {
	name string
}

func (car * Car) Mycar() {
	fmt.Print(car.name)
}

func NewInstance(name string) Car {
	return Car {
		name,
	}
}

func main(){
	srujan:= NewInstance("audi")
	
	srujan.Mycar()
}