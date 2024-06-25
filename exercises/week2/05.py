n = int(input())
cards = list(map(int, input().split()))

i, j = 0, n - 1
res = [0, 0]
player = 0
while i <= j:
    if cards[i] > cards[j]:
        res[player] += cards[i]
        i += 1
    else:
        res[player] += cards[j]
        j -= 1
    player = 1 - player

print(" ".join(map(str, res)))
