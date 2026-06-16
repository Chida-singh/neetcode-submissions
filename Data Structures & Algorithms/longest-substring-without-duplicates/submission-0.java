class Solution {
    public int lengthOfLongestSubstring(String s) {
        HashSet<Character> set = new HashSet<>();
        int l = 0, res = 0;
        for(int i =0; i < s.length();i++){
            char c = s.charAt(i);
            while(set.contains(c)){
                set.remove(s.charAt(l++));
            }
            set.add(s.charAt(i));
            res = Math.max(res, i - l + 1);
            
        }
        return res;
    }
}
