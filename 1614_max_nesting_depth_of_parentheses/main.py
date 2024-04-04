from collections import deque

class Solution:
  def maxDepth(self, s: str) -> int:
    stack = deque()

    m = 0

    for c in s:
      if c == "(":
        stack.append(c)
      elif c == ")":
        m = max(m, len(stack))
        stack.pop()

    return m



sol = Solution()
print(sol.maxDepth("8*((1*(5+6))*(8/6))"))
