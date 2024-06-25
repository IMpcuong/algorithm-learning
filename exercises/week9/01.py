from typing import List
from queue import PriorityQueue


class HeapEntry(object):
    def __init__(self, priority: int, pos: int):
        self.priority = priority
        self.pos = pos

    def __lt__(self, other):
        return self.priority > other.priority


def solve(jobs: int, pos: int, priorities: List[int]) -> int:
    if jobs == 1:
        return 1
    pq = PriorityQueue()
    for i, p in enumerate(priorities):
        pq.put(HeapEntry(priority=p, pos=i))
    print([(p.priority, p.pos) for p in pq.queue])
    counter = 0
    while not pq.empty():
        entry = pq.get()
        if entry.priority >= priorities[pos]:
            counter += 1
        if counter == jobs:
            counter -= 1
            break
    return counter


# pq = PriorityQueue()
# for i, p in enumerate([1, 1, 9, 1, 1, 1]):
#     pq.put(HeapEntry(p, i))
# tmp = [(p.priority, p.pos) for p in pq.queue]
# print(tmp)

n_queues = int(input())
n_jobs, n_pos = [], []
n_priorities = [[0] * 100 for _ in range(n_queues)]
for i in range(n_queues):
    jobs, pos = map(int, input().split())
    n_jobs.append(jobs)
    n_pos.append(pos)
    priorities = list(map(int, input().split()))
    n_priorities[i] = priorities

# print(n_jobs, n_pos)
for i in range(n_queues):
    print(solve(n_jobs[i], n_pos[i], n_priorities[i]))
