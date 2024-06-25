k, n, w = map(int, input().split())

paid = 0
for i in range(w):
  	paid += k * (i + 1)

paid = paid - n
if paid <= 0:
  	paid = 0
print(paid)