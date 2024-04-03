from typing import List


class Solution:
  def removeDuplicates(self, nums: List[int]) -> int:
    l, r, n, prev, offset, count = 0, 0, len(nums), nums[0], 0, 0
    
    if n <= 2:
      return n
    
    while r < n:
      if prev != nums[r]:
        offset = r - l
        count = 0
        
      if offset > 0:
        nums[l] = nums[r]
        offset -= 1
        
      count += 1
      if count < 3:
        l += 1

      prev = nums[r]
      r += 1
    return l
    
sol = Solution()
print(sol.removeDuplicates([1,1,1,2,2,2]))
