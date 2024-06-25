from typing import List


_n = int(input())
rates = list(map(int, input().split()))
rate_freq = [0] * 2001
for v in rates:
    rate_freq[v] += 1


def get_rank(num: int, rates: List[int]) -> int:
    desc_order_list = sorted(list(set(rates)), reverse=True)
    return [i + 1 for i, v in enumerate(desc_order_list) if v == num][0]


def solve(rates: List[int]) -> List[int]:
    for i in range(_n):
        rates[i] = get_rank(rates[i], rates)
    return rates

print(solve(rates=rates))
