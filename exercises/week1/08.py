from typing import List


s: str = input()
t: str = input()

NT = False
RM = False
SH = False


def expose_char_freq_from(s: str) -> List[int]:
    char_freqs = [0] * 26
    for c in s:
        char_freqs[ord(c) - ord("a")] += 1
    return char_freqs


def get_first_matched_pos(s: List[int], ord_ch: int, pos: int) -> int:
    for i in range(pos, len(s) - 1):
        if s[i] == ord_ch:
            return i
    return -1


freq_s = expose_char_freq_from(s)
freq_t = expose_char_freq_from(t)

for i in range(0, 25):
    if freq_s[i] < freq_t[i]:
        NT = True
    elif freq_s[i] > freq_t[i]:
        RM = True

matched_pos = -1
for i in range(0, len(t) - 1):
    matched_pos = get_first_matched_pos(s, t[i], matched_pos + 1)
    if matched_pos == -1:
        SH = True
        break

if NT:
    print("need tree")
elif RM and SH:
    print("both")
elif RM:
    print("automaton")
else:
    print("array")
