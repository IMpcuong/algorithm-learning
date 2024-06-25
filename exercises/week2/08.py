n = int(input())
chocolates = list(map(int, input().split()))

a_consume = b_consume = 0
i, j = 0, n - 1
while i <= j:
    if a_consume + chocolates[i] <= b_consume + chocolates[j]:
        a_consume += chocolates[i]
        i += 1
    else:
        b_consume += chocolates[j]
        j -= 1

print(i, n - i)
