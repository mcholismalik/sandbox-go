package rotate

import "fmt"

func New(nums []int, k int) {
	// Without changing the function contract, try to solve the problem.
	for i := 0; i < k; i++ {
		tmp := make(map[int]int)
		e := nums[len(nums)-1]

		for j := 0; j < len(nums); j++ {
			if j == 0 {
				tmp[j] = nums[j]
				nums[j] = e
			} else {
				tmp[j] = nums[j]
				nums[j] = tmp[j-1]
			}
			delete(tmp, j-1)
		}
	}

	fmt.Println(nums)
}
