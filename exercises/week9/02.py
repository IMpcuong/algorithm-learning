n = int(input())
str = input()

arr = [0] * 132
for c in str.lower():
    arr[ord(c)] += 1


def solve(lower: int, upper: int) -> bool:
    for c in range(lower, upper + 1):
        if arr[c] < 1:
            return False
    return True


if solve(ord('a'), ord('z')):
    print("YES")
else:
    print("NO")
