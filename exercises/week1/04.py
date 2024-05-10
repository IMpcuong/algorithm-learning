from typing import List


def convert_to_str_from(l: List[int]) -> str:
    # NOTE: All `l` members are validated ASCII codes.
    return "".join(map(chr, l))


# def eval_qualified_str_from(l1: List[int], l2: List[int]) -> str:
#     # NOTE: "abcde" versus "abcdg".
#     matched_chars: List[int] = []
#     for i, t in enumerate(zip(l1, l2)):
#         if t[1] - t[0] == 0:
#             if i == len(l1) - 1:
#                 matched_chars.append(t[0] + 1)
#             else:
#                 matched_chars.append(t[0])
#         if t[1] - t[0] == 1:
#             matched_chars.append(t[0])
#         if t[1] - t[0] > 1:
#             matched_chars.append(t[0] + 1)
#     r = convert_to_str_from(matched_chars)
#     if r == convert_to_str_from(l1):
#         return "No such string"
#     return r


def convert_to_adjacent_lexico_str(l: List[int]) -> str:
    for i in range(len(l) - 1, -1, -1):
        if l[i] == ord("z"):
            l[i] = ord("a")
        else:
            l[i] += 1
            break
    return convert_to_str_from(l)


s: List[int] = [ord(c) for c in input()]
t: List[int] = input()

# if t[-1] - s[-1] > 1 or t[-1] - s[-1] == 0:
#     m = [chr(c) for c in s[0:-1]]
#     r = "".join(m) + chr(s[-1] + 1)
#     print(r)
# elif t[-1] - s[-1] == 1:
#     print("No such string")

# s = eval_qualified_str_from(l1=s, l2=t)
r = convert_to_adjacent_lexico_str(s)
if r == t:
    print("No such string")
print(r)
