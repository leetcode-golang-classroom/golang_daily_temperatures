package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	stack := []int{} // 存放沒處理過的 index
	answers := make([]int, len(temperatures))
	for idx, val := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < val {
			topIdx := len(stack) - 1
			ele := stack[topIdx]
			stack = stack[0:topIdx]
			answers[ele] = idx - ele
		}
		stack = append(stack, idx)
	}
	return answers
}
