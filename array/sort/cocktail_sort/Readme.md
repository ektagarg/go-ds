# Cocktail Sort
It is a sorting algorithm that iterates over an array in both directions, leaving the largest elements to the right, and the lowest elements to the left. It is a variation of the Bubble Sort Algorithm.

### Example
[1, 4, 5, 3, 2]

1) <b>1</b> 4 5 3 2
2) 1 <b>4</b> 5 3 2
3) 1 4 <b>3</b> 5 2
4) 1 4 3 <b>2</b> 5
5) 1 4 3 2 <b>5</b>
5) 1 4 2 <b>3</b> 5
6) 1 2 <b>4</b> 3 5
6) 1 <b>2</b> 4 3 5
7) <b>1</b> 2 4 3 5
8) 1 <b>2</b> 4 3 5
9) 1 2 <b>3</b> 4 5
10) 1 2 3 <b>4</b> 5
11) 1 2 3 4 <b>5</b>

And then, the array is sorted.

```
Best case (When it is already sorted): O(n)
Worst case: O(n^2)
Average time complexity: O(n^2)
```
