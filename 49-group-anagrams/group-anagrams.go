func groupAnagrams(strs []string) [][]string {
    var allStrs [][]string
    used := make([]bool, len(strs)) // to avoid grouping the same word multiple times

    for i := 0; i < len(strs); i++ {
        if used[i] {
            continue
        }

        var tempGroup []string
        tempGroup = append(tempGroup, strs[i])
        used[i] = true

        for j := i + 1; j < len(strs); j++ {
            if !used[j] && isAnagram(strs[i], strs[j]) {
                tempGroup = append(tempGroup, strs[j])
                used[j] = true
            }
        }

        allStrs = append(allStrs, tempGroup)
    }

    return allStrs
}

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := make(map[rune]int)

    for _, ch := range s {
        count[ch]++
    }

    for _, ch := range t {
        count[ch]--
        if count[ch] < 0 {
            return false
        }
    }

    return true
}
