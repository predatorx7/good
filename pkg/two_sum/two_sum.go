package two_sum

func TwoSum(nums []int, target int) []int {
	previousNumbers := make(map[int]int, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		it := nums[i]
		requiredNumber := target - it
		if indexOfRequiredNum, ok := previousNumbers[requiredNumber]; ok {
			return []int{indexOfRequiredNum, i}
		}
		previousNumbers[it] = i
	}
	return []int{}
}
