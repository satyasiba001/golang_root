package main

import(
	"fmt"
)

func foo(){
	fmt.Println("I am from foo")
}

func faa(s string){
	fmt.Println(s,"is here")
}

func bar(s string) string{
	return fmt.Sprint("They call me : ", s)
}

func person(name string, age int)(string,int){
	return fmt.Sprint("The square of age of ",name," is "), age*age
}

func main(){
	foo()
	faa("satya")
	var k = bar("xerxes")
	fmt.Println(k)
	a,b := person("papu", 20)
	fmt.Println(a,b)
}