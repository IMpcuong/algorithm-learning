from typing import List


def calculate_adjacent_distances(arr: List[int]) -> List[int]:
    distances: List[int] = []
    for i in range(len(arr) - 1):
        intermed_distance = abs(arr[i] - arr[i + 1])
        if intermed_distance > 13:
            distances.append(26 - intermed_distance)
        else:
            distances.append(intermed_distance)
    return distances


input_str: str = str(input())
ord_arr: List[int] = [ord("a")] + [ord(c) for c in input_str]

rune_arr: List[int] = list(range(97, 122))
adjacent_distances: List[int] = calculate_adjacent_distances(ord_arr)
optimal_rotations: int = sum(adjacent_distances, start=0)

print(optimal_rotations)
