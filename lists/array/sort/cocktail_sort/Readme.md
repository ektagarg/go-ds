# Cocktail Sort
Cocktail Sort is a variation of Bubble sort. The Bubble sort algorithm always traverses elements from left and moves the largest element to its correct position in first iteration and second largest in second iteration and so on.

### Example 
```
Array: [1 5 4 2 8 0 2]

###### first iteration :
 gets biggest element to the last position:
 gets smallest element to the first position and so on:

First Forward Pass

[1 5 4 2 8 0 2]
[1 4 5 2 8 0 2]
[1 4 2 5 8 0 2]
[1 4 2 5 8 0 2]
[1 4 2 5 0 8 2]
[1 4 2 5 0 2 8]

First Backward Pass:

[1 4 2 5 0 2 8]
[1 4 2 0 5 2 8]
[1 4 0 2 5 2 8]
[1 0 4 2 5 2 8]
[0 1 4 2 5 2 8]

###### iteration 2:
 Forward Pass

[0 1 4 2 5 2 8]
[0 1 4 2 5 2 8]
[0 1 2 4 5 2 8]
[0 1 2 4 5 2 8]
[0 1 2 4 2 5 8]

 Backward Pass:
[0 1 2 2 4 5 8]
[0 1 2 2 4 5 8]
[0 1 2 2 4 5 8]
[0 1 2 2 4 5 8]

###### iteration 3:
[0 1 2 2 4 5 8]
[0 1 2 2 4 5 8]
[0 1 2 2 4 5 8]
[0 1 2 2 4 5 8]

Since array is already sorted, Backward Pass is not conducted

```

### Complexity
```
Worst Case Time Complexity: O(n^2)
Best Case Time Complexity: O(n) when it's already sorted
Average Time Complexity: O(n^2)
```