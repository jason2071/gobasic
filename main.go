package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("hello world")
	// hr.Profile()
	// jsonDemo.TodoList()
	// jsonDemo.TodoListJsonPlaceholder()
	// users.User()
	// jsonDemo.JsonEncoderDecoder()
	// greet.HiPersonTextTemplate()
	// greet.HtmlTemplate()
	// numbers.NumberAvg()
	// outlineNodes.NodeOutline()

	Money()
}

func withdraw(amount int) (int, error) {
	if amount > 0 {
		return 0, errors.New("insufficient fund")
	}
	return amount, nil
}

func Money() {
	amount, err := withdraw(200)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please collect your money: ", amount)

}
