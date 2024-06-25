n, m, x, y = list(map(int, input().split()))
soldiers = list(map(int, input().split()))
vests = list(map(int, input().split()))

# fit = [0 for _ in range(max(n, m) + 1)]
fit = []
i = 0
for j in range(m):
    while i < n and vests[j] > soldiers[i] + y:
        i += 1
    if i == n:
        break
    if soldiers[i] - x <= vests[j] <= soldiers[i] + y:
        fit.append((i + 1, j + 1))
        i += 1

print(len(fit))
[print(tup[0], tup[1]) for tup in fit]
