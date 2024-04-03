from collections import defaultdict
from typing import List


class Solution:
  def maxProfit(self, prices: List[int], fee: int) -> int:
    n = len(prices)
    arr = [[0 for i in range(n)] for j in range(n)]
    hmaxRow = defaultdict(self.defaultzero)
    hmaxCol = defaultdict(self.defaultzero)
    
    row = n-1
    while row >= 0:
      col = n-1
      while col >= 0:
        if col > row:
          arr[row][col] = prices[col] - prices[row] - fee
          hmaxRow[row] = max(hmaxRow[row], prices[col] - prices[row] - fee)
          hmaxCol[col] = max(hmaxCol[col], prices[col] - prices[row] - fee)
        col -= 1
      row -= 1
    
    m = 0
    for i in range(n):
      maxc, maxr = hmaxCol[i], hmaxRow[i+1]
      m = max(m, maxc + maxr)
    
    return m

  def defaultzero(self):
    return 0

  def printArr(self, arr):
    for col in arr:
      for row in col:
        print(row)
      print("\n")

sol = Solution()
sol.maxProfit([1,4,6,2,8,3,10,14], 3)
