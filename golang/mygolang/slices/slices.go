package main
import ("fmt")

func main(){
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}
	
	fmt.Println(arr1)
	fmt.Println(arr2)


	prices := [...]int{10,20,30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	var newarr= [...]int{1:1,2:2}
	fmt.Print(newarr) //prints [0,1,2]
}