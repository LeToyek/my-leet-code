func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    chars := make(map[rune]int)
    
    for _,ch := range s{
        chars[ch]++
    }
    for _,ch := range t{
        chars[ch]--
        if chars[ch]<0{
            return false
        }
    }

    return true
}