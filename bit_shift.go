package main

import "fmt"

func main(){
	x := 2
	fmt.Printf("%d\t\t%b\n",x,x)
	
	y := x << 1
	fmt.Printf("%d\t\t%b\n",y,y)

	z := y << 2
        fmt.Printf("%d\t\t%b\n",z,z)
}
