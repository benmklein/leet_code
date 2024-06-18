package twoSum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum( t *testing.T ) {
  tests := []struct {
    nums []int
    target int
    want []int
  }{
    {[]int{ 2,7,11,15 }, 9, []int{0,1}},
    {[]int{3,2,4}, 6, []int{1,2}},
    {[]int{3,3}, 6, []int{0,1}},
  }
  for _, tt := range tests {
    testname := fmt.Sprintf("%v, %v", tt.nums, tt.target)
    t.Run(testname, func(t *testing.T){
      ans := twoSum(tt.nums, tt.target)
      if !reflect.DeepEqual(ans, tt.want){ t.Errorf("want: %v, got: %v", tt.want, ans) }
    })

  }
}
