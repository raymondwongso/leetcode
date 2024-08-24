def is_palindrome(n):
    return str(n) == str(n)[::-1]

# Start from the next number after 2976421
n = 2976421 - 1

# Increment and check until we find a palindrome
while not is_palindrome(n):
    n -= 1

print(n)
