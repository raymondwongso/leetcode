from typing import List


class Solution:
  def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
    if k == 0:
        return 0

    ans, ii, ln, p = 0, 0, len(nums), 1

    for i, x in enumerate(nums):
      p *= x

      if p >= k:
        n = i-ii
        ans += (n/2)*(2+n-1)-1 # minus 1 in the end to delete overlap
        while ii < i-1:
          p /= nums[ii]
          ii += 1

    if ii < ln:
      n = ln-ii
      ans += (n/2)*(2+n-1)

    return int(ans)

sol = Solution()

print(sol.numSubarrayProductLessThanK([1,1,1], 1))
