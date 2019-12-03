/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	j := 0
	i := 0
	if n == 0 {
		return
	}
	for ; i < m; i++ {
		if nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]

			for ;j < n; j++ {
				k := j+1
				if k >= n {
					j = 0
					break
				}
				if nums2[j] > nums2[k] {
			       nums2[j], nums2[k] = nums2[k], nums2[j]
				}
			}
		}
	}

	for j = 0; j < n; j++ {
		nums1[i] = nums2[j]	
		i++
	}
}
// @lc code=end

