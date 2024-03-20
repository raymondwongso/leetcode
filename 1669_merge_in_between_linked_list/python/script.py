class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def printlistnode(list: ListNode):
  if list is None:
    return
  
  print(list.val)
  if list.next is not None:
    printlistnode(list.next)

# Submission result: 193ms (54.05%) | 21.08MB (70.92%)
def mergeInBetween(list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
  i = 0
  root = list1

  while True:
      if i == a-1:
          diff = b-a
          counter = 0
          i += 1
          tail = None

          runningList = list1.next
          while True:
              if counter < diff:
                  counter += 1
                  i += 1
                  runningList = runningList.next
              else:
                  if runningList.next is not None:
                      tail = runningList.next
                  break
          
          rootList2 = list2
          while list2.next is not None:
              list2 = list2.next

          list2.next = tail
          list1.next = rootList2
          break
      
      i += 1
      list1 = list1.next
      
  printlistnode(root)
  return root

list1 = ListNode(10, ListNode(1, ListNode(13, ListNode(6, ListNode(9, ListNode(5))))))
list2 = ListNode(100, ListNode(101, ListNode(102)))
print(mergeInBetween(list1, 3, 4, list2))
# list1 = ListNode(0, ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5, ListNode(6)))))))
# list2 = ListNode(100, ListNode(101, ListNode(102, ListNode(103))))
# print(mergeInBetween(list1, 2, 5, list2))
