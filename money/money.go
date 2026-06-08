package money

import (
	"errors"
	"fmt"
)

func withdraw(amount int) (int, error) {
	if amount > 0 {
		return 0, errors.New("withdraw: insufficient fund")
	}
	return amount, nil
}

func Money() {
	amount, err := withdraw(200)
	if err != nil {
		fmt.Println("Money:", err)
		return
	}
	fmt.Println("Please collect your money: ", amount)

}
