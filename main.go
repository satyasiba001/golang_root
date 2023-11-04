package main

import "fmt"

func main(){
	//var s string = "Hello satya"
	var g string
	fmt.Println(g)
	g = "satya"
	//k = 10
	//fmt.Println(k)
	k := 12.432
	fmt.Printf("%v\t%T\n",k,k)
	fmt.Println(g)
	a,b,c := 0,1,2
	fmt.Println(a,b,c,"\v")
	var d,e,f = 4, true ,"satya"
	fmt.Println(d,e,f)
	fmt.Printf("%T %T %T \n",d,e,f)
	j := int(k)
	//n := bool(k)
	fmt.Printf("%v\t%T\n",j,j)
}