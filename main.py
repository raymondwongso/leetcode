from typing import List


class Solution:
    def haveConflict(self, event1: List[str], event2: List[str]) -> bool:
        h1, m1 = event1
        h2, m2 = event2
        
        return h1 <= h2 <= m1 or h2 <= h1 <= m2

      
sol = Solution()
print(sol.haveConflict(["10:13", "13:02"], ["13:17", "21:38"]))
