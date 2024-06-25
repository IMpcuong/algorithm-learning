n = int(input())
claws = list(map(int, input().split()))

victims = 0
j = n - 1
for i in range(n - 1, -1, -1):
    j = min(i, j)
    last_victim_pos = max(0, i - claws[i])

    if j > last_victim_pos:
        victims += j - last_victim_pos
        j = last_victim_pos

print(n - victims)
