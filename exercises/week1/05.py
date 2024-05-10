from typing import List


sizes: List[int] = list(map(int, input().split()))
size_a, size_b = sizes[0], sizes[1]
pick_sizes: List[int] = list(map(int, input().split()))
pick_a, pick_b = pick_sizes[0], pick_sizes[1]

a: List[int] = list(map(int, input().split()))
b: List[int] = list(map(int, input().split()))

print("YES" if a[pick_a - 1] < b[size_b - pick_b] else "NO")
