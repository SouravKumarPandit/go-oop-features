package apackage

import "fmt"
//https://github.com/go101/golds
//some text to generate go doc
func PrintSomeThing() {
	fmt.Print("Can You Print me")
	//can't access package name with internal
	//apack.Bar()
	//bpack.Bar()
	//somepack.Bar()
}