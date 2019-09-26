# Bubble Sort
It is a comparision-based sorting algorithm where each pair of adjacent elements is compared and elements get swapped if they are not in order.

### Example 
```
Array: 23,98,67,45,12,24
preocedure: Take first pair and swap if it's not in order and repeat until last element gets sorted

###### iteration 1:
step 1: 23,98,67,45,12,24
step 2: 23,67,98,45,12,24
step 3: 23,67,45,98,12,24
step 4: 23,67,45,12,98,24
step 5: 23,67,45,12,24,98

result: we got last element sorted!!

###### iteration 2:
step 1: 23,67,45,12,24,98
step 2: 23,45,67,12,24,98
step 3: 23,45,12,67,24,98
step 4: 23,45,12,24,67,98

result: we got second last element sorted!!

###### iteration 3:
step 1: 23,45,12,24,67,98
step 2: 23,12,45,24,67,98
step 3: 23,12,24,45,67,98

result: we got third last element sorted!!

###### iteration 4:
step 1: 23,12,24,45,67,98
step 2: 12,23,24,45,67,98

result: we got array sorted!!
```

### Complexity
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n) when it's already sorted
Average Time Complexity: O(n2)
```