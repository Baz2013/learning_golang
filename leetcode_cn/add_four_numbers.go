package leetcode_cn

// leetcode 454

// AddFourNumbers1  该方法虽然能测试通过，但是时间复杂度未 O(n^4), 不可取
func AddFourNumbers1(A [3]int, B [3]int, C [3]int, D [3]int) (length int) {

	//myMap := make(map[int]int)
	var result [][]int = make([][]int, 0)

	for i0, v0 := range A {
		for i1, v1 := range B {
			for i2, v2 := range C {
				for i3, v3 := range D {
					if v0+v1+v2+v3 == 0 {
						result = append(result, []int{i0, i1, i2, i3})
					}
				}
			}
		}
	}

	length = len(result)

	return
}

// AddFourNumbers 该方法的时间复杂度为 O(n^2)
func AddFourNumbers(A [3]int, B [3]int, C [3]int, D [3]int) (count int) {
	myMap := make(map[int]int)
	for _, v0 := range A {
		for _, v1 := range B {
			myMap[v0+v1]++
		}
	}

	for _, v0 := range C {
		for _, v1 := range D {
			count += myMap[0-(v0+v1)]
		}
	}

	return
}
