from typing import List, Tuple


size: int = int(input())
segments: List[Tuple[int, int]] = []
lhs: List[int] = []
rhs: List[int] = []
for i in range(0, size):
    seg = tuple(map(int, input().split()))
    segments.append(seg)
    lhs.append(seg[0])
    rhs.append(seg[1])


def is_exist_covered_segment() -> int:
    desired_segment = (min(lhs), max(rhs))
    for i in range(0, len(segments)):
        if segments[i] == desired_segment:
            return i + 1
    return -1


print(is_exist_covered_segment())
