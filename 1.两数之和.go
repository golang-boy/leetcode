/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	var i, j int
    for i = 0; i < len(nums); i++ {
	   for j = i+1; j < len(nums); j++ {
		   if j < len(nums) && nums[j] == target - nums[i] {
			   goto out
		   }
	   }
	}
out:
    return []int{i, j}
}
// @lc code=end

