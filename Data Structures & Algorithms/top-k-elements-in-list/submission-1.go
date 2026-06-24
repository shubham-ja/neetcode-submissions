func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _,num := range nums {
		freqMap[num] += 1
	}
	//fmt.Println(freqMap)
	buckets := make([][]int,len(nums)+1)

	for num,freq := range freqMap {
		buckets[freq] = append(buckets[freq],num)
	}
	result := []int{}

	for i := len(buckets)-1; i >= 1; i-- {
		for _,val := range buckets[i] {
			result = append(result, val)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
