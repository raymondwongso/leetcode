# Definition for singly-linked list.
from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Time complexity: O(n)
# Space complexity: O(n)
# Submission result: 39ms (46.24%) | 17.75MB (33.99%)
# class Solution:
#   def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
#     if head is None:
#       return []
    
#     root, tail = self.recursive(head)
    
#     return root

#   def recursive(self, head: Optional[ListNode]):
#     if head.next is not None:
#       root, tail = self.recursive(head.next)
#       head.next = None
#       tail.next = head
#       tail = tail.next
#       return root, tail
    
#     return head, head
      
# Iterative solution
# Time complexity: O(n)
# Space complexity: O(n)
# Submission result: 45ms (9.69%) | 17.68MB (77.25%)
class Solution:
  def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
    stack = [head]
    while head.next is not None:
        stack.append(head.next)
        head = head.next

    res = None
    root = None
    for n in reversed(stack):
        if res is None:
            res = n
            root = res
        else:
            n.next = None
            res.next = n
            res = n

    return root      

# Improved iteration solution
# Time complexity: O(n)
# Space complexity: O(1)
# Submission result: 
class Solution:
  def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
    prevNode = None
    currentNode = head
    
    while currentNode is not None:
        nextNode = currentNode.next
        currentNode.next = prevNode
        prevNode = currentNode
        currentNode = nextNode
        
    return prevNode




# Driver code
sol = Solution()
def printListNode(head: Optional[ListNode]):
  print(head.val)
  
  if head.next is not None:
    printListNode(head.next)

ls = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
res = sol.reverseList(ls)
printListNode(res)
