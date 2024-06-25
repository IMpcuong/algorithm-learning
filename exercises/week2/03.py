n, t = map(int, input().split())
arr = list(map(int, input().split()))

max_book = 0
read_time = 0
start_book = 0
for end_book in range(n):
   read_time += arr[end_book]
   if read_time > t:
     read_time -= arr[start_book]
     start_book += 1
   max_book = max(max_book, end_book - start_book + 1)

print(max_book)