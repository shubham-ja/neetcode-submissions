func partition(s string) [][]string {
	cache := make(map[string]bool)
	return partitionAndCheckPalindrom(s, 0, len(s)-1, cache)
	
}

func partitionAndCheckPalindrom(s string, start, end int, cache map[string]bool) [][]string {
	result := [][]string{}
	limit := start
	for limit < end {
		if isPalindrome(s, start, limit, cache) {
			response := partitionAndCheckPalindrom(s, limit+1, end, cache)
			if len(response) > 0 {
				for _,row := range response {
					result = append(result,append([]string{s[start:limit+1]},row...))
				}
			}
		}
		limit += 1
	}
	if isPalindrome(s, start, end, cache) {
		result = append(result,[]string{s[start:end+1]})
	}
	return result
}

func isPalindrome(s string, start, end int, cache map[string]bool) bool {
	key := fmt.Sprintf("%d:%d",start,end)
	if val, exists := cache[key]; exists {
		return val
	} else {
		for start <= end {
			if s[start] != s[end] {
				cache[key] = false
				return cache[key]
			}
			start += 1
			end -= 1
		}
	}
	cache[key] = true
	return cache[key]
}
