type key [26]int

func genKey(s string) key {
	ans := key{}
	for _, c := range(s) {
		ans[c - 'a'] ++
	}
	return ans
}

func groupAnagrams(strs []string) [][]string {
	cache := make(map[key][]string)
	for _, s := range(strs) {
		k := genKey(s)
		cache[k] = append(cache[k], s)
	}

	ans := make([][]string, 0)
	for _, v := range(cache) {
		ans = append(ans, v)
	}
	return ans
}
