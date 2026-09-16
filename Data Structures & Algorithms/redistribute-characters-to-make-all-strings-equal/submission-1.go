func makeEqual(words []string) bool {
    data := make(map[byte]int)
    for i := range words {
        for j := range words[i] {
            data[words[i][j]]++
        }
    }

    for _, count := range data {
        if count % len(words) != 0 {
            return false
        }
    }
    return true
}