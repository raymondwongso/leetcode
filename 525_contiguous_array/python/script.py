from typing import List

# Submission result: 604ms (89.21%) | 22.22MB (25.49%)
class Solution:
  def findMaxLength(self, nums: List[int]) -> int:
    h = {0: -1}
    res = 0

    prefixSum = 0
    for i, n in enumerate(nums):
        if n == 0:
            n = -1
        prefixSum += n

        if h.get(prefixSum, None) is None:
            h[prefixSum] = i
        else:
            res = max(res, i-h[prefixSum])
    
    return res

sol = Solution()
print(sol.findMaxLength([0, 1, 0]))
