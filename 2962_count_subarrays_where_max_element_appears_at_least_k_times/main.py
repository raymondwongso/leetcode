from typing import Counter, List


class Solution:
  def countSubarrays(self, nums: List[int], k: int) -> int:
    m = max(nums)
    freq = Counter()
    ans = ii = 0 
    for idx, x in enumerate(nums): 
        freq[x] += 1
        while freq[m] >= k: 
            freq[nums[ii]] -= 1
            ii += 1
        ans += ii
    return ans 

sol = Solution()
print(sol.countSubarrays([61,23,38,23,56,40,82,56,82,82,82,70,8,69,8,7,19,14,58,42,82,10,82,78,15,82], 2))
