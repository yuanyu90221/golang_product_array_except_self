package product_array

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for idx, _ := range result {
		result[idx] = 1
	}
	lower := 1
	larger := 1
	for idx, _ := range result {
		if idx != 0 {
			lower *= nums[idx-1]
			larger *= nums[len(nums)-idx]
		}
		result[idx] *= lower
		result[len(nums)-1-idx] *= larger
	}
	return result
}
