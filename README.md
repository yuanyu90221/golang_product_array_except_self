# product_array_except_self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

## 分析

需要把所有自己之外的 element 相乘的新 array

原本會是 先 loop 一遍乘出所有 product 在利用除法去算出來

然而這邊不能這樣做並且需要在 O(n) 時間複雜度

## 初步解法

定義 product [i] 為除了 i 之外值的 product

可以拆解成兩個部份 

1 大於 i 的 product 

2 小於 i 的 product

所以可以用兩個 array 依次把 大於 index 0,..,n 的 product 及 小於 index 0,..., n 的 product 算出來

然後只要 把這兩個部份 array 部份相乘就是答案

## 優化空間複雜度

察看每個 product 其實都與前一項有關係

所以其實不需要用 array 去紀錄

只需透過關係式每次推演當下的 lower 與 upper 即可


## 程式碼

```golang
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	larger := make([]int, len(nums))
	lower := make([]int, len(nums))
	for idx, _ := range result {
		result[idx] = 1
    if idx == 0 {
			lower[idx] = 1
			larger[len(nums)-idx-1] = 1
		} else {
			lower[idx] = lower[idx-1] * nums[idx-1]
			larger[len(nums)-1-idx] = nums[len(nums)-idx] * larger[len(nums)-idx]
		}
	}
	for idx, _ := range result {
		result[idx] = lower[idx] * larger[idx]
	}
	return result
}
```

## 解法2

```golang
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
```