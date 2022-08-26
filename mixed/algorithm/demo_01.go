package algorithm

import "fmt"

/*
 * 输入n，打印出其中连续自然数的和为n的数
 */
func Sun(n int) []int {
	for i := 1; i < n/2; i++ {
		for j := i + 1; j < n; j++ {
			//fmt.Println("i:", i,"j:", j)
			var total []int
			a := j + 1
			if i+j+a == n && i == j-1 && j == a-1 {
				total = append(total, i, j, a)
				return total
				fmt.Println(total)
			} else if j+a == n && j == a-1 && i == j-1 {
				total = append(total, j, a)
				fmt.Println(total)
				return total
			}
		}
	}
	return nil
}
