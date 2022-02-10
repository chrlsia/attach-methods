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

}

func (c car) accelerate() string {
	return "accelerating ..."
}
