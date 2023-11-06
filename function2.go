package main

import(
	"fmt"
)
func fooo(bb ...int)int{
	sum := 0
	fmt.Println(bb)
	fmt.Printf("%T\n",bb)
	for _,v := range bb{
		sum += v
	}
	return sum
}
func main(){
	ans := fooo(1,2,3,4,5)
	fmt.Println(ans)
}