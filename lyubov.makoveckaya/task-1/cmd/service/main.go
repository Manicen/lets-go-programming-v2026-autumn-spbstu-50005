package main

import "fmt"

func main() {
	var (
		a    int
		b    int
		pr   string
		err1 error
		err2 error
	)
	_, err1 = fmt.Scan(&a)
	_, err2 = fmt.Scan(&b)
	fmt.Scan(&pr)

	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	switch pr {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
