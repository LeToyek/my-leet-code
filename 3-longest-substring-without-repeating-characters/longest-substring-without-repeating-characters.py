class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        sb = ""
        longest_sb = ""
        for x in s:
            if x in sb:
                if len(longest_sb) < len(sb):
                    longest_sb = sb
                sb = sb[sb.index(x) + 1:]
            sb +=x
        if len(longest_sb) < len(sb):
            longest_sb = sb
        
        return max(len(longest_sb), len(sb))  # Return the length of the longest substring
            
            
        