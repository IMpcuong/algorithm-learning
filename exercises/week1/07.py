from typing import Dict, List, Tuple
from collections import Counter


nums: Tuple[int, int] = tuple(map(int, input().split()))
retries = nums[0]
max_failures = nums[1]
input_pwds: List[str] = []
for i in range(0, retries):
    input_pwds.append(input())
input_pwds = sorted(input_pwds, key=lambda x: len(x))
len_freq_map: Dict[int, int] = dict(Counter(map(len, input_pwds)))
correct_pass: str = input()

best, worst = 0, 0
pwd_len = len(correct_pass)
for i in range(1, pwd_len + 1):
    if i not in len_freq_map:
        len_freq_map[i] = 0
    if i < pwd_len:
        best += len_freq_map[i]
    if i == pwd_len:
        best += 1
    worst += len_freq_map[i]

print(best, worst)
best += 5 * (best // max_failures)
worst += 5 * (worst // max_failures)
print(best, worst, sep=" ")
