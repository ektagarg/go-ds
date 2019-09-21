# Selection Sort
The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.

1) The subarray which is already sorted.
2) Remaining subarray which is unsorted.

In every iteration of selection sort, the minimum element (considering ascending order) from the unsorted subarray is picked and moved to the sorted subarray.


### Complexity
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n2)
Average Time Complexity: O(n2)
```

### Algorithm
1. Get a list of unsorted numbers
2. Set a marker for the unsorted section at the front of the list
3. Repeat steps 4 - 6 until one number remains in the unsorted section
4. Compare all unsorted numbers in order to select the smallest one
5. Swap this number with the first number in the unsorted section
6. Advance the marker to the right one position
7. Stop

### Example
```
array : [23, 45, 67, 27, 32, 87, 90, 12]
Starting from 0 to n, find the minimum element and insert at a[0]
[12, 45, 67, 27, 32, 87, 90, 23]
[12, 23, 67, 27, 32, 87, 90, 45]
[12, 23, 27, 67, 32, 87, 90, 45]
[12, 23, 27, 32, 67, 87, 90, 45]
[12, 23, 27, 32, 45, 87, 90, 67]
[12, 23, 27, 32, 45, 67, 90, 87]
[12, 23, 27, 32, 45, 67, 87, 90]
<<<<<<< HEAD
```
=======
```
>>>>>>> 8ab32ffe4b48bfb6de3bb5fd887ad31184a0c5ae
