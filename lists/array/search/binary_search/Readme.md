# Binary Search
A binary search or half-interval search aims to find the element in a sorted array by first comparing target element with middle element to see whether the element is in first half or second half and keep repeating this until it finds the target element.

### Algorithm
1. Find middle element in sorted array: (first+last)/2
2. If x matches with middle element, we return the mid index.
3. Else If x is greater than the mid element, then x can only lie in right half subarray after the mid element. So we recur for right half.
4. Else (x is smaller) recur for the left half.

### Example
```
[1,3,12,23,34,56,62,98,99]
x = 12

Iteration 1:
1. mid = 1+9/2 = 5
2. x<mid
3. recur in left half

Iteration 2:
1. mid = 1+5/2 = 3
2. x==mid
3. return mid

Element is at position 3
```

### Complexity
Worst-case complexity: o(logn)
Best-case complexity: o(1)