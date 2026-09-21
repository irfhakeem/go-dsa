package leetcode

func moveZeroes(nums []int) {
	r := 0
	l := 0

	for r < len(nums) {
		if nums[l] == 0 && nums[r] != 0 {
			tmp := nums[l]
			nums[l] = nums[r]
			nums[r] = tmp
			l++
		} else if nums[l] != 0 {
			l++
		}
		r++
	}
}
