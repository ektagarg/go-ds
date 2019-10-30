# Tim-sort
Tim-sort is a sorting algorithm derived from insertion sort and merge sort. It was designed to perform in an optimal way on different kind of real world data. It is an adaptive sorting algorithm which needs O(n log n) comparisons to sort an array of n elements. It was designed and implemented by Tim Peters in 2002 in python programming language. It has been python's standard sorting algorithm since version 2.3.

### Algorithm
1. Divide the array into the number of blocks known as run.
2. Consider size of run either 32 or 64(in the below implementation, size of run is 32.)
3. Sort the individual elements of every run one by one using insertion sort.
4. Merge the sorted runs one by one using merge function of merge sort.
5. Double the size of merged sub-arrays after every iteration.

### Complexity
```
Worst Case Time Complexity: O(n log n)
Best Case Time Complexity: O(n)
Average Time Complexity: O(n log n)
```
