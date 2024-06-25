n = int(input()) # NOTE: n is a odd number.
arr = list(map(int, input().split()))


def find_median(arr, length):
    arr.sort()
    if length % 2 == 0:
        return (arr[length // 2] + arr[length // 2 - 1]) / 2
    return arr[n // 2]


print(find_median(arr=arr, length=n))

