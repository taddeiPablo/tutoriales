package main

import(
	"fmt"
)

//TUTORIAL 7 CICLOS EN GO
func main() {
	//primer ejemplo de for
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world")
	}

	//for tipo while
	for i < 10 {
		
	} 
	i := 0
	for {

		if i ==2{
			i++
			continue
		}
		fmt.Println(i)
		i++
		if i > 10{
			break
		}
	}
}