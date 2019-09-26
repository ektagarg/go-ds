# Shell sort
Shell sort is a highly efficient sorting algorithm and is based on insertion sort algorithm. This algorithm avoids large shifts as in case of insertion sort, if the smaller value is to the far right and has to be moved to the far left.

This algorithm uses insertion sort on a widely spread elements, first to sort them and then sorts the less widely spaced elements. This spacing is termed as interval. This interval is calculated based on Knuth's formula as −

```
Knuth's Formula
h = h * 3 + 1
where −
   h is interval with initial value 1
```

This algorithm is quite efficient for medium-sized data sets as its average and worst-case complexity of this algorithm depends on the gap sequence the best known is Ο(n), where n is the number of items. And the worst case space complexity is O(n).

### Algorithm
Step 1 − Initialize the value of h
Step 2 − Divide the list into smaller sub-list of equal interval h
Step 3 − Sort these sub-lists using insertion sort
Step 3 − Repeat until complete list is sorted

### Complexity
Worst complexity: Depends on gap sequence
Average complexity: n*log(n)^2 or n^(3/2)
Best complexity: n

### Example
12, 90, 23, 87, 32, 67, 27
h = h*3 + 1 where h=1; h=4

Pass 1:
i=1; j=4
12, 90, 23, 87

