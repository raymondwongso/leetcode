# Naive solution
# We iterate all possibility
# Time complexity:  O(n2)
# Space complexity: O(1) 
#   excluding input size, additional space required is 
#   for i, n, j, n2 which is constant
# Submission result: 1596ms (34.21%) | 17.31MB (88.46%)
#
# class Solution:
#     def twoSum(self, nums: List[int], target: int) -> List[int]:
#         for i, n in enumerate(nums):
#             for j in range(i+1, len(nums)):
#                 n2 = nums[j]
#                 if n + n2 == target:
#                     return [i, j]
# =========================================================================
# Hashmap Solution
# We populate hashmap and then use lookup for faster search
# Time complexity:  O(n)
# Space complexity: O(n) 
#   determined by calculating space used by hashmap
# Submission result: 65ms (48.06%) | 17.82MB (38.30%)
#
# class Solution:
#     def twoSum(self, nums: List[int], target: int) -> List[int]:
#         hash_map = {}
#         for i, n in enumerate(nums):
#             hash_map[n] = i
#         for i, n in enumerate(nums):
#             lookup = target - n
#             if lookup in hash_map:
#                 if hash_map[lookup] != i:
#                     return [i, hash_map[lookup]]
# =========================================================================
# Improved hashmap solution (1 loop)
# We populate hashmap and then use lookup for faster search
# the search happens whilst populating the hashmap to reduce time
# Time complexity:  O(n)
# Space complexity: O(n) 
#   determined by calculating space used by hashmap
#
# Submission result: 55ms (83.05%) | 17.90MB (38.30%)
class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        hash_map = {}
        for i, n in enumerate(nums):
            lookup = target - n
            if lookup in hash_map:
                return [i, hash_map[lookup]]
            hash_map[n] = i
