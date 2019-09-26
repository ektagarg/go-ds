# Merge sort
Like quick sort, merge sort is based on divide and conquer methodology. It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.

Conceptually, a merge sort works as follows:

Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.

### Algorithm
MergeSort(arr[], l,  r)

If r > l

1. Find the middle point to divide the array into two halves:  
         middle m = (l+r)/2
2. Call mergeSort for first half:   
         Call mergeSort(arr, l, m)
3. Call mergeSort for second half:
         Call mergeSort(arr, m+1, r)
4. Merge the two halves sorted in step 2 and 3:
         Call merge(arr, l, m, r)

### Example
```
                                [35,26,47,3,8,89,11]
                                /                 \      
                        [35,26,47,3]            [8,89,11]
                        /       \               /       \
                [35,26]        [47,3]       [8,89]     [11]   
                /    \          /   \         /  \        \
            [35]    [26]    [47]     [3]     [8]  [89]    [11]
                \    /          \     /        \   /    /
                [26,35]         [3,47]        [8,89]  [11]
                        \       /                 \    /
                        [3,26,35,47]              [8,11,89]
                                \                  /
                                [3,8,11,26,35,47,89]
```

### Complexity
```
Worst Case Time Complexity: O(nlog(n))
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```
