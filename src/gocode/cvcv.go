package main

func productExceptSelf(nums []int) []int {
	var n = len(nums)
	var res = make([]int, n)
	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i]
	}

	tmp := 1
	for i := n - 1; i > -1; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}
	return res

}
