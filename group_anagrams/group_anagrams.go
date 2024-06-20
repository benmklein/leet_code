package main
import ("fmt")

func groupAnagrams(strs []string) [][]string {
  result := make(map[[26]uint8][]string)

  for _, s := range strs {
    char_counts := [26]uint8{}
    for _, c := range []byte(s) {
      c_num := c - "a"[0]
      char_counts[c_num] += 1
    }
    result[char_counts] = append(result[char_counts], s)
  }

  anagrams := [][]string{}
  for _, v := range result {
    anagrams = append(anagrams, v)
  }
  return anagrams
}

func main(){
  fmt.Println( groupAnagrams([]string{ "cat", "bat", "tab" }) )
}
