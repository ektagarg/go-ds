# Insertion sort
Insertion sort is based on the idea that one element from the input elements is consumed in each iteration to find its correct position i.e, the position to which it belongs in a sorted array.

It iterates the input elements by growing the sorted array at each iteration. It compares the current element with the largest value in the sorted array. If the current element is greater, then it leaves the element in its place and moves on to the next element else it finds its correct position in the sorted array and moves it to that position. This is done by shifting all the elements, which are larger than the current element, in the sorted array to one position ahead.


### Algorithm
1. Loop from i = 1 to n-1.
2. Pick element arr[i] and insert it into sorted sequence arr[0…i-1]

### Example
```
Array: 12, 11, 13, 5, 6

1. Let's loop for i = 1 (second element of the array) to 4 (last element of the array)

2. i = 1. Since 11 is smaller than 12, move 12 and insert 11 before 12
    11, 12, 13, 5, 6

3. i = 2. 13 will remain at its position as all elements in A[0..I-1] are smaller than 13
    11, 12, 13, 5, 6

4. i = 3. 5 will move to the beginning and all other elements from 11 to 13 will move one position ahead of their current position.
    5, 11, 12, 13, 6

5. i = 4. 6 will move to position after 5, and elements from 11 to 13 will move one position ahead of their current position.
    5, 6, 11, 12, 13
```

### Complexity
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n) when it's already sorted
Average Time Complexity: O(n2)
```