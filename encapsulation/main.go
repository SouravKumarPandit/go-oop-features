package main

type Vehicle struct {
	brand string
	wheels int
}
func (v *Vehicle)GetBrand() string{
	return v.brand
}

func (v *Vehicle)GetSpeed() int{
	return v.wheels* 5
}

func main() {
/*
	Capital letter will expose your package struct to oter package a
	Same goes with method
	*/

}
