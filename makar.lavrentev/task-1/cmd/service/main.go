package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation byte
	_, err = fmt.Scanf("%c", &operation)

	if b == 0 && operation == '/' {
		fmt.Println("Division by zero")
		return
	}

	var ans int
	switch operation {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		ans = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(ans)
}
