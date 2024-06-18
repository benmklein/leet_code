package validAnagram

import (
  "testing"
  "fmt"
)

func TestIsAnagramTable(t *testing.T){
  var tests = []struct {
    str1, str2 string
    want bool
  }{
    {"anagram", "nagaram", true},
    {"rat", "car", false},
  }
  for _, tt := range tests {
    testname := fmt.Sprintf("%s, %s", tt.str1, tt.str2)
    t.Run(testname, func(t *testing.T) {
      ans := isAnagram(tt.str1, tt.str2)
      if ans != tt.want {
        t.Errorf("Got %t, want %t", ans, tt.want)
      }
    })
  }
}
