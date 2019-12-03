/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	for i := len(nums)-1 ; i >= 0; i-- {
        j := (i+k) % len(nums)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
// @lc code=end

