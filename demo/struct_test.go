package demo

import (
	"testing"
	"fmt"
)

type Dog struct {
	name string 
	// friends map[int]string  //若加上这个属性，则会报错 invalid operation: dog1 == dog2 (struct containing map[int]string cannot be compared)
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}


func TestStructJunior(t *testing.T){
	dog1 := Dog{name:"one"}
	var dog2 Dog = dog1
	// dog1.SetName("two")  //若加上这个语句， 下面的值就不会相等

	if dog1 == dog2 {
		fmt.Println("dog1==dog2")
	}
}
