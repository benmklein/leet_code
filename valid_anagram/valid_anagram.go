package validAnagram

func isAnagram(s string, t string) bool {
  count := make(map[byte]int)
  
  for _, c := range []byte(s){
    count[c] += 1
  }
  for _, c := range []byte(t){
    count[c] -= 1
  }
  
  for _, v := range count {
    if v != 0 {return false}
  }
  return true
}
