n = int(input())
arr = list(map(int, input().split()))

freq = [0] * (10 ** 5 + 5)
longest_range = 0
diff = 0
j = 0
for i in range(n):
    if freq[arr[i]] == 0:
      diff += 1
    freq[arr[i]] += 1

    if diff > 2 and j < n:
      if freq[arr[j]] == 1:
        diff -= 1
      freq[arr[j]] -= 1
      j += 1

    longest_range = max(longest_range, i - j + 1)

print(longest_range)