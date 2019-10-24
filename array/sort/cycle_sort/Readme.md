# Cycle Sort
Cycle sort is an in-place, unstable sorting algorithm, a comparison sort that is theoretically optimal in terms of the total number of writes to the original array, unlike any other in-place sorting algorithm. It is based on the idea that the permutation to be sorted can be factored into cycles, which can individually be rotated to give a sorted result.
Unlike nearly every other sort, items are never written elsewhere in the array simply to push them out of the way of the action. Each value is either written zero times, if it's already in its correct position, or written one time to its correct position. This matches the minimal number of overwrites required for a completed in-place sort.

### Algorithm
1. Get a list of unsorted numbers
2. Remember current CycleStartPosition as 0
3. If CycleStartPosition > lenght of list go to step 9. Else remember element on CycleStartPosition as A.
4. Iterate through array, find number of elements lesser than current. Find by that correct position for current element. 
5. If the element is already at the correct position, increase CycleStartPosition by one and go to step 3. If it is not, write it to its intended position. That position is inhabited by a different element B. 
6. Find correct position for B.
7. If correct position for B == CycleStartPosition, write it, increase CycleStartPosition by one and go to step 3.
8. If it is not, write it to its intended position. That position is inhabited by a different element C. Remember C as B and go to step 6.
9. Stop

### Example
```
array : [55 52 79 20 65 26 36 93 97 22]
[55 52 79 20 65 55 36 93 97 22]
[55 52 26 20 65 55 36 93 97 22]
[55 52 26 20 65 55 36 79 97 22]
[55 52 26 20 65 55 36 79 93 22]
[55 52 26 20 65 55 36 79 93 97]
[55 22 26 20 65 55 36 79 93 97]
[55 22 26 20 52 55 36 79 93 97]
[55 22 26 20 52 55 65 79 93 97]
[55 22 26 36 52 55 65 79 93 97]
[20 22 26 36 52 55 65 79 93 97]
```

### Complexity
```
Worst-case performance	Θ(n2)
Best-case performance	Θ(n2)
Average performance	Θ(n2)
Worst-case space complexity	Θ(n) total, Θ(1) auxiliary
```
