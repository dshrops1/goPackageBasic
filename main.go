package main

import(
	"fmt"
	"./packageExport"
)

func main(){
	fmt.Println("we are exporting a function that returns a value")
	fmt.Println(packageExport.DValue())
}
