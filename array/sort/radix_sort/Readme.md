# Radix Sort
It is not based on comparision-based algorithms, instead it is an integer sorting algorithm that sorts data with integer keys by grouping the keys by individual digits that share the same significant position and value.

Radix sort takes in a list of n integers which are in base b (the radix) and so each number has at most d digits where d = |log_b(k) +1| and k is the largest number in the list. 

### Algorithm
1) For each digit i where i varies from least significant digit to the most significant digit.
2) Sort input array using counting sort according to the i’th digit.

### Example

```
Let's take a slice
213, 003, 045, 067, 078, 001, 034, 456, 044

Create buckets based on the digit, on base 10
Iteration 1:
0-
1- 1
2- 
3- 213, 3
4- 34, 44
5- 45
6- 456
7- 67
8- 78
9-

push numbers to new slice sequentially from bucket
001, 213, 003, 034, 044, 045, 456, 067, 078

Create buckets based on the digit, on base 100
Iteration 2:
0- 1, 3
1- 213
2- 
3-34, 
4- 44, 45, 
5- 456
6- 67
7- 78
8-
9-

001, 003, 213, 034, 044, 045, 456, 067, 078

Create buckets based on the digit, on base 1000
Iteration 3:
0- 1, 3, 33, 34, 45, 67, 78
1- 213
4 - 456

1, 3, 33, 34, 45, 67, 78, 213, 456
```

### Complexity
```
Time Complexity: o(kn) 
where k is total number of digits in the largest number bcs that is number of iterations
```
