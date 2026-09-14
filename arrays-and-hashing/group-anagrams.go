func groupAnagrams(strs []string) [][]string {
    group := make(map[string][]string)

    hash := func(s string) string {
        chars := strings.Split(s, "")
        sort.Strings(chars)
        return strings.Join(chars, "")
    }

    for _, s := range strs {
        key := hash(s)
        group[key] = append(group[key], s)
    }

    result := make([][]string, 0, len(group))
    for _, value := range group {
        result = append(result, value)
    }

    return result
}
