from typing import List


def isValidFashion(size: int, buttons: List[int]) -> bool:
    if size != len(buttons):
        return False
    if size == 1:
        return buttons[0] == 1

    f_counter = 0
    for b in buttons:
        if b == 1:
            f_counter += 1
    return f_counter == size - 1


# if __name__ == "main":
size: int = int(input())
buttons: List = list(map(int, input().split()))
if isValidFashion(size, buttons):
    print("YES")
else:
    print("NO")
