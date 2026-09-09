func genVersions(word string) []string {
	ans := []string{}
	for i := range word {
		for j := range 26 {
			nb := byte('a' + j)
			if nb != word[i] {
				b := []byte(word)
				b[i] = nb
				ans = append(ans, string(b))
			}
		}
	}
	return ans
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    words := make(map[string]struct{})
	for i := range wordList {
		words[wordList[i]] = struct{}{}
	}
	words[beginWord] = struct{}{}

	neighbours := make(map[string][]string)
	for word := range words {
		neighbours[word] = []string{}
		variants := genVersions(word)

		for _, variant := range variants {
			if _, found := words[variant]; found {
				neighbours[word] = append(neighbours[word], variant)
			}
		}
	}

	queue := []string{beginWord}
	visited := make(map[string]struct{})
	hops := 1

	for len(queue) > 0 {
		stash := []string{}
		for _, word := range queue {
			if word == endWord {
				return hops
			}
			visited[word] = struct{}{}
			for _, neighbour := range neighbours[word] {
				if _, vis := visited[neighbour]; !vis {
					stash = append(stash, neighbour)
				}
			}
		}
		queue = stash
		hops++
	}

	return 0
}
