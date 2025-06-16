import "sort"

func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

    for _, str := range strs {
        key := sortedString(str)
        groups[key] = append(groups[key], str)
    }

    result := [][]string{}
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}

// Helper to sort a string
func sortedString(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}
