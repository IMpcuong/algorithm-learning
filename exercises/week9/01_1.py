from typing import List


def solve(jobs: int, pos: int, priorities: List[int]) -> int:
    if jobs == 1:
        return 1

    minutes = 0
    arr = [0] * 10
    for p in priorities:
        arr[p] += 1
    target_priority = priorities[pos]
    for i in range(target_priority, 9 + 1):
        minutes += arr[i]
    return minutes


n_queues = int(input())
n_jobs, n_pos = [], []
n_priorities = [[0] * 100 for _ in range(n_queues)]
for i in range(n_queues):
    jobs, pos = map(int, input().split())
    n_jobs.append(jobs)
    n_pos.append(pos)
    priorities = list(map(int, input().split()))
    n_priorities[i] = priorities

for i in range(n_queues):
    print(solve(n_jobs[i], n_pos[i], n_priorities[i]))
