from typing import List


def calculate_adjacent_gap(arr: List[int]) -> int:
    if arr[0] > 15:
        return 15

    for i in range(len(arr) - 1):
        intermed_gap = arr[i + 1] - arr[i]
        if intermed_gap > 15:
            return arr[i] + 15

    if arr[-1] + 15 < 90:
        return arr[-1] + 15
    return 90


thrilled_freq = int(input())
thrilled_minute_arr: List[int] = list(map(int, input().split()))

print(calculate_adjacent_gap(thrilled_minute_arr))
