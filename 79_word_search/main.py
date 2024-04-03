from collections import defaultdict
from typing import List

class Solution:
  def default_value(self):
    return False

  def exist(self, board: List[List[str]], word: str) -> bool:
    self.board = board
    self.word = word
    self.m = len(board)
    self.n = len(board[0])

    wi = 0
    res = False

    for i, row in enumerate(board):
      for j, col in enumerate(row):
        res = self.do(defaultdict(self.default_value), i, j, wi)
        if res:
          return True

    return res
  
  def do(self, h, i, j, wi) -> bool:
    res = False

    if self.board[i][j] == self.word[wi]:
      h[f"{i}{j}"] = True
      if wi == len(self.word)-1:
        return True
  
      if i > 0:
        if h[f"{i-1}{j}"] == False:
          res = self.do(h, i-1, j, wi+1)
      if res:
        return True

      if j > 0:
        if h[f"{i}{j-1}"] == False:
          res = self.do(h, i, j-1, wi+1)
      if res:
        return True

      if i < self.m-1:
        if h[f"{i+1}{j}"] == False:
          res = self.do(h, i+1, j, wi+1)
      if res:
        return True

      if j < self.n-1:
        if h[f"{i}{j+1}"] == False:
          res = self.do(h, i, j+1, wi+1)
      if res:
        return True

    # at this point, no match found, delete from hash
    h[f"{i}{j}"] = False
    return res

sol = Solution()
print(sol.exist([["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "ABCCED"))
