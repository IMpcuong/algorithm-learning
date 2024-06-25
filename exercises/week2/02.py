n, m = map(int, input().split())
required = list(map(int, input().split())) # len(required) == n
required.sort()
prepared = list(map(int, input().split())) # len(prepared) == m
prepared.sort()


i = 0
counter = 0
for j in range(m):
    if i >= n:
      break
    if prepared[j] >= required[i]:
      counter += 1
      i += 1

print(n - counter)
