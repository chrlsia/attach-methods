package main

import "fmt"

type car struct {
	color string
	body  string
}

func main() {
	porsche := car{color: "blue", body: "coupe"}
	bentley := car{color: "green", body: "saloon"}

	fmt.Printf("porsche is :%v\n", porsche)
	fmt.Printf("bentley is :%v\n", bentley)

	fmt.Printf("porsche color is : %v\n", porsche.color)
	fmt.Printf("porsche style is : %v\n", porsche.body)

	fmt.Printf("bentley color is : %v\n", bentley.color)
	fmt.Printf("bentley style is : %v\n", bentley.body)

}

func (c car) accelerate() string {
	return "accelerating ..."
}
