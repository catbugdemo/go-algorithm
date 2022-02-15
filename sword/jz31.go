package sword

// 用一个队列模拟入栈出栈
//
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	index := 0
	length := len(pushed)
	for i := 0; i < length; i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && index < len(popped) && popped[index] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	return len(stack) == 0
}
