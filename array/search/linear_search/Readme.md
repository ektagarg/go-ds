# Linear search 
A linear search (or sequential search) is a method for finding an element within a list. It sequentially checks each element of the list until a match is found or the whole list has been searched. 

### Algorithm
1. Start from the leftmost element of arr[] and one by one compare x with each element of arr[]
2. If x matches with an element, return the index.
3. If x doesn’t match with any of elements, return -1.

### Example
```
[1,3,12,23,34,56,62,98,99]
x = 34

Iteration 1:
compare x with element at position 0
NOT SAME

Iteration 2:
compare x with element at position 1
NOT SAME

Iteration 3:
compare x with element at position 2
NOT SAME

Iteration 4:
compare x with element at position 3
NOT SAME

Iteration 5:
compare x with element at position 4
SAME
Element is at position 4
```

### Complexity
Worst-case complexity: o(n)
Best-case complexity: o(1)
