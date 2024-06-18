package twoSum

func twoSum(nums []int, target int) []int {
  m := make(map[int][]int)
  for i := range nums {
    if len(m[nums[i]]) != 0 {
      return []int{m[nums[i]][0], i}
    }
    targetDiff := target - nums[i]
    m[targetDiff] = append(m[targetDiff], i)
  }
  return []int{}
}
