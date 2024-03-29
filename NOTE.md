0. Binary search:

- Binary search is an algorithm; its input is a sorted list of elements.
- Return the location or index of the needed element or else return -1.

```go
package main

import "fmt"

// Verbose implementation:
func binary_search_v1(num int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == num {
			return mid
		} else if guess > num {
			high = mid - 1
		} else if guess < num {
			low = mid + 1
		}
	}
	return -1
}

func binary_search_v2(num int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess > num {
			high = mid - 1
		} else if guess < num {
			low = mid + 1
		}
		return mid
	}
	return -1
}

func binary_search_recursive(num, low, high int, arr []int) int {
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess > num {
			binary_search_recursive(num, low, mid-1, arr)
		}
		if guess < num {
			binary_search_recursive(num, mid+1, high, arr)
		}
		return mid
	}
	return -1
}

func main() {
	arr := []int{0, 2, 4, 5, 7, 9, 11, 12, 13, 18}
	fmt.Println(binary_search_v1(7, arr), arr[4])
	fmt.Println(binary_search_v2(7, arr), arr[4])
	fmt.Println(binary_search_recursive(7, 0, len(arr)-1, arr), arr[4])
}
```

```c
#include <stdint.h>
#include <stdio.h>

uint32_t binary_search(uint32_t num, uint32_t *arr)
{
	uint32_t low = 0;
	uint32_t high = (uint32_t)((sizeof arr) / (sizeof arr[0])) - 1; // FIXME: @@@
	while (low < high)
	{
		uint32_t mid = (low + high) / 2;
		uint32_t guess = *(arr + mid);
		if (guess > num)
			high = mid - 1;
		if (guess < num)
			low = mid + 1;
		return mid;
	}
	return 0;
}

int main(void)
{
	uint32_t arr[8] = {0, 2, 6, 7, 8, 9, 10, 12};
	printf("%d: %d", binary_search(8, arr), arr[4]);
	return 0;
}
```

- In this book: `log_x(n) -> x := 2`.
- Exercise:

  - 1.1: log_2(128) = log(128) = 7
  - 1.2: log(128 \* 2) = 8

- BigO notation lets you compare the number of operations. Some examples:

  | BigO           | Stands for        | Example                                       |
  | -------------- | ----------------- | --------------------------------------------- |
  | `O(log n)`     | **_log time_**    | Binary search                                 |
  | `O(n)`         | **_linear time_** | Simple search                                 |
  | `O(n * log n)` |                   | Quicksort (Fast sorting algorithm)            |
  | `O(n^2)`       |                   | Selection sort (Slow sorting algorithm)       |
  | `O(n!)`        |                   | Traveling salesperson (Really slow algorithm) |

- Algorithm speed isn't measured in seconds, but in growth of the number of operation.
- Instead, we talk about how quickly the runtime of an algorithm increases as the size of the input increases.
- Runtime of algorithm is expressed in BigO notation.
- `O(log n)` is faster than `O(n)`, but it gets a lot faster as the list of items you're searching grows.

- Exercise:

  - 1.3: `O(log n)`.
  - 1.4: `O(n)`.
  - 1.5: `O(n)`.
  - 1.6: `O(log (n/26))`.

1. Selection sort:

- Runtime for common operation between arrays and lists:

  - Array: allow random access.
  - List (Linked lists): allow sequential access.

  |           | Arrays | Lists  |
  | --------- | ------ | ------ |
  | Reading   | `O(1)` | `O(n)` |
  | Insertion | `O(n)` | `O(1)` |
  | Deletion  | `O(n)` | `O(1)` |

- Exercise:

  - 2.1: List.
  - 2.2: List.
  - 2.3: Array.
  - 2.4: Downside ~ array easy to append, an inappropriate for in-between insertion
    -> hard to sort -> to serve for binary-search algorithm.
  - 2.5: Let’s consider a hybrid data structure: an array of linked lists.
    You have an array with 26 slots. Each slot points to a linked list.
    For example, the first slot in the array points to a linked list containing all the usernames starting with `a`.
    The second slot points to a linked list containing all the usernames starting with `b`, and so on.
    My guess:

    | Hybrid    | Array  | Linked list |
    | --------- | :----: | :---------: |
    | Searching | Faster |   Slower    |
    | Insertion | Slower |   Faster    |

2. Recursion:

- Quotes:

  - Loops may achieve a performance gain for your program.
  - Recursion may achieve a performance gain for your programmer.
  - Choose which is more important in your situation!”

- Using the stack is convenient because you don't have to keep track of a pile of boxes yourself - the stack does that for you.
- Using the stack comes at the cost of over-allocating memory. When the stack to tall,
  that means your computer is saving information for many functions call.

- Recap:

  - Recursion is when a function calls itself.
  - Every recursive function has `2` cases: the base case and the recursive case.
  - A stack has `2` operations: `push` and `pop`.
  - All functions call go onto the `call stack`.
  - The call stack can get very large, which takes up a lot of memory.

3. Divide-and-conquer:

- Recap:

  - Step 1: Figure out simple case as the base case.
  - Step 2: Figure out how to reduce your problem and get to the base case (recursive case).

- Final recap:

  - `D&C` works by breaking problem down into smaller and smaller pieces. If you are using `D&C` on a list, the base case is
    probably an empty array or an array with one elements.
  - If you're implementing quick-sort, choose a random element as the pivot. The average runtime of quick-sort is `O(n * log n)`!
  - The constant in BigO notation can matter some time. That's why quick-sort is faster than merge-sort.
  - The constant is almost never matter for simple search versus binary search, because `O(log n)` is so much faster than `O(n)`
    when your list gets big.

4. Hash tables:

- Definition: A `hash-function` is a function where you put in a string and you get back a number.
- Mathematics reflection: In mathematics, a function (injective function) is a procedure that maps a value from an infinite set with a correspondent value that belongs to another infinite set.

- Arrays and lists map straight to the memory, but `hash-tables` are smarter (combination between: `hash-function + array`). They use a `hash-function` to intelligently figure out where to store elements.

- Recap for hashes:

  - Speed up your looks up process.
  - Modelling the relationship from one thing to another thing.
  - Filtering out duplicates.
  - Caching/memorizing data instead of making the server do all the work.

- Collisions: To avoid key collisions, the implementation method must takes care of the two most important parts:

  - Choosing the hash function wisely.
  - Use a compound data structure such as hashmap + linked-list (not so sure).
  - A low load factor := `Number of items in hash tables / Total number of slots`.
    Once the load factor variant surpass `0.7`, it's time to resize your hash table.

- Performance:

  | Operation | Average Case | Worst Case |
  | --------- | :----------: | :--------: |
  | Search    |    `O(1)`    |   `O(n)`   |
  | Insert    |    `O(1)`    |   `O(n)`   |
  | Delete    |    `O(1)`    |   `O(n)`   |

- Time: `Chapter 6: Breadth-first Search`; `99 (118 of 258)`.

5. Breath-first search:

- Recap for BFS:

  - BFS tells you if there's ac path from A to B.
  - If there's path, BFS will find the shortest path.
  - If you have a problem like "find the shortest X", try modeling your problem as a graph,
    and use BFS to solve.
  - A directed graph has arrows, and the relationship follows the direction of arrow
    (A -> B means "A owes B money").
  - Undirected graph don't have arrows, the relationship goes both ways
    (Joey - Phoebe means "Joey dated Phoebe and Phoebe dated Joey").
  - Queues are FIFO (First In, First Out).
  - Stacks are LIFO (Last In, First Out).
  - You need to check people in the order they were added to the search list, so the search list needs to be a queue.
    Otherwise, you won't get the shortest path.
  - Once you check someone, mark them as a tainted person and do not repeat to check them again.
    Otherwise, you might end up in an infinite loop.