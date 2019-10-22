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

### Example
```
12, 90, 23, 87, 32, 67, 27
h = h*3 + 1 where h=1; h=4

Pass 1:
i=1; j=4
12, 90, 23, 87, 32, 67, 27

i=2; j=5
12, 32, 23, 87, 90, 67, 27

i=3; j=6
12, 32, 23, 87, 90, 67, 27

i=4; j=7
12, 32, 23, 27, 90, 67, 87

Pass 2:
now h = h/2 = 2
i=1; j=3
12, 32, 23, 27, 90, 67, 87

i=2; j=4
12, 27, 23, 32, 90, 67, 87

i=3; j=5
12, 27, 23, 32, 90, 67, 87

i=4; j=6
12, 27, 23, 32, 90, 67, 87

i=5; j=7
12, 27, 23, 32, 87, 67, 90

Pass 3:
now h = h/2 = 1
simply apply insertion sort
[12, 23, 27, 67, 87, 90]
```

### Complexity
```
Worst complexity: Depends on gap sequence
Average complexity: n*log(n)^2 or n^(3/2)
Best complexity: n
```