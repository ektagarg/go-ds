# Quick Sort
Quick sort is highly efficient algorithm which is based on Divide and Conquer strategy. So, the technique is to divide the given array into smaller arrays.
A large array is partitioned into two arrays one of which holds values smaller than a specific value, say pivot, based on which the partition is made and another array holds values greater than the pivot value.
Quick sort partitions an array and then calls itself recursively twice to sort the two resulting subarrays. This algorithm is quite efficient for large-sized data sets as its average and worst case complexity are of Ο(n2), where n is the number of items.

### Complexity
```
Worst Case Time Complexity: O(n2) when an array is already sorted
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```

### Algorithm
1. Find a “pivot” item in the array. This item is the basis for comparison for a single round.
2. Start a pointer (the left pointer) as the first item in the array.
3. Start a pointer (the right pointer) as the last item in the array.
4. While the value at the left pointer in the array is less than the pivot value, move the left pointer to the right (add 1). Continue until the value at the left pointer is greater than or equal to the pivot value.
5. While the value at the right pointer in the array is greater than the pivot value, move the right pointer to the left (subtract 1). Continue until the value at the right pointer is less than or equal to the pivot value.
6. If the left pointer is less than or equal to the right pointer, then swap the values at these locations in the array.
7. Move the left pointer to the right by one and the right pointer to the left by one.
8. If the left pointer and right pointer don’t meet, go to step 1.

### Example
```
array : [58, 45, 67, 27, 32, 87, 90, 23]
indexes:[0,  1,  2,  3,  4,  5,  6,  7]
low = 1; high = 7; pivot = 58

Partition algorithm:

i=low, j=high
compare pivot with x[j] = (58 > 23)
stop j (j=7, i=1)
compare pivit with x[i] = (58 > 45)
since it's in correct order, i++ (j=7, i=2)
compare pivit with x[i] = (58 < 67)
stop i 
swap i with j, i++, j-- (j=6, i=2)

[58, 45, 23, 27, 32, 87, 90, 67]
compare pivot with x[j] = (58 < 90)
since it's in correct order, j-- (j=5, i=3)
compare pivot with x[j] = (58 < 87)
since it's in correct order, j-- (j=4, i=3)
compare pivot with x[j] = (58 > 32)
stop j 
compare pivit with x[i] = (58 > 27)
since it's in correct order, i++ (j=4, i=4)
i is now equals to j, we got correct position of pivot element i; swap a[i] with pivot
[32, 45, 23, 27] 58 [87, 90, 67]

