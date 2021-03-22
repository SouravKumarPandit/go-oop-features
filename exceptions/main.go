package main

import "fmt"

func SomeErrorThrowMethod(i int) (string, error) {
	if i < 0 {
		return "", fmt.Errorf("Number Is Invalid number. ")
	}

	s := fmt.Sprintf("Your Binary input Value is %b", i)

	return s, nil
}
func main() {
	value, err := SomeErrorThrowMethod(-45)
	//value, err := SomeErrorThrowMethod(45)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)

}
