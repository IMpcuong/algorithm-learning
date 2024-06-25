# a1 a1 a3 a4 a5
# 5 3 -> a1 a1 a3 a4 -> 1 4

n, k = map(int, input().split())
arr = list(map(int, input().split()))

freq = [0] * (10**5 + 1)
unique = 0
start_pos = 0
for end_pos in range(n):
    num = arr[end_pos]
    if freq[num] == 0:
        unique += 1
    freq[num] += 1
    if unique == k:
        while True:
            if freq[arr[start_pos]] > 1: # first non-zero number occurrence
                freq[arr[start_pos]] -= 1
                start_pos += 1
            elif freq[num] == 1:
                print(start_pos + 1, end_pos + 1, sep=" ")
                exit()

print(-1, -1)
