func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stack := []int{}

	for _, token := range tokens {
		switch token {
			case "+":
				a, b := stack[len(stack) - 2], stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, a+b)
			case "-":
				a, b := stack[len(stack) - 2], stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, a-b)
			case "*":
				a, b := stack[len(stack) - 2], stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, a*b)
			case "/":
				a, b := stack[len(stack) - 2], stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, a/b)
			case "%":
				a, b := stack[len(stack) - 2], stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, a&b)
			default:
				num, _ := strconv.Atoi(token)
				stack = append(stack, num)
		}
	}

	return stack[0]
}
