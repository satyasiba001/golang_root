package main
import "fmt"

func greaterno(a int ,b int)int{
	var result int
	if(a>b){
		result = a
	}else{
		result = b
	}
	return result
}

func main(){
	fmt.Println(greaterno(5,6))
}

