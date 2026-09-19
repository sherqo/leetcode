func checkInclusion(s1 string, s2 string) bool {
    var freq1 [26]int
    for i := 0; i < len(s1); i++ {
        freq1[s1[i]-'a']++
    }

    var freq2 [26]int
    l := 0

    for r := 0; r < len(s2); r++ {
        char := s2[r] - 'a'
        freq2[char]++

        for freq2[char] > freq1[char] {
            freq2[s2[l]-'a']--
            l++
        }

        if r-l+1 == len(s1) {
            return true
        }
    }

    return false
}
