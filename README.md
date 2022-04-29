# golang_daily_temperatures

Given an array of integers `temperatures` represents the daily temperatures, return *an array* `answer` *such that* `answer[i]` *is the number of days you have to wait after the* `ith` *day to get a warmer temperature*. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

## Examples

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100

# 解析

給定一個整數陣列 temperatures , 陣列中每一個值代表該天的溫度, temperature[i] 代表第 i 天的溫度,題目想求的是對於每天為起點經過幾天後天氣會變高

所以回傳是一個整數陣列 answer ,每個值代表從該天經過幾天後會變暖, answer[i] 代表第 i 天溫度變暖需要經過 answer[i] 天

最直覺的作法是輪詢每次以第i天起點找到第一個大於 temperature[i] 的 index 做相減

這種作法需要 O(n^2) 因為每次都需要 從 i 找到 n-1所以需要 找 n-i 次有 n 個

要降低每次尋找的次數必須要把每次經過的值沒有處理過紀錄下來

這樣就不會需要再重新找

並且知道每個點只關心在他之後的點 只關心第一個大的點

所以可以透過 stack 來做這件事

這樣因為每次找到大於 element 都只需要花費 O(1) , 可以將複雜度降低到 O(1)

## 透過 stack 來降低複雜度的解法

step 1 初始化一個 index = 0 , answer 為長度 temperatures.length 的整數陣列, stack 用來儲存當下沒處理過的 index

step 2 當 index < temperatures.length  檢查 stack 是否為空

step 3 當stack 不為空則 把小於 top 元素都 pop 出來 並且 跟新  answer[pop_Element] = index - pop_Element, 把 index 放入 stack 

step 4  當 stack 為空, 把 index 放入

step 5 index = index + 1 , 回到 step 2

step 6 回傳 answers

## 程式碼

```go
func dailyTemperatures(temperatures []int) []int {
  tLen := len(temperatures)
  stack := make([]int, 0, tLen) // last value is capacity not len
  answers := make([]int, tLen)
  for idx, value := range temperatures {
    for len(stack) > 0 && temperatures[stack[len(stack)-1]] < value {
       topIdx := len(stack) - 1
       ele := stack[topIdx]
       stack = stack[0: topIdx]
       answers[ele] = idx - ele
    } 
    stack = append(stack, idx)
  }
  return answers
}
```

# Solve Points

- [x]  Understand What the problem wants
- [x]  Choose suitable Data structure
- [x]  Analysis Complexity