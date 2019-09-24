# go-ds
A library which aims to explain and implement data structures in Golang.

## Types of data structures
### Arrays
An array is a data structure that contains fixed-sized collection of elements of same data type, such as an integer or string. Arrays are commonly used in computer programs to organize data so that a related set of values can be easily sorted or searched.

Operations on Arrays:
### Searching
**1. Linear Search:** A linear search or sequential search is a method for finding an element within a list. It sequentially checks each element of the list until a match is found or the whole list has been searched. 
```
Worst-case complexity: o(n)
Best-case complexity: o(1)
```

**2. Binary Search:** A binary search or half-interval search aims to find the element in a sorted array by first comparing target element with middle element to see whether the element is in first half or second half and keep repeating this until it finds the target element.
```
Worst-case complexity: o(logn)
Best-case complexity: o(1)
```

### Sorting

**1. Bubble Sort:** It is a comparision-based sorting algorithm where each pair of adjacent elements is compared and elements get swapped if they are not in order.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n)
```
For more details, you can read this: https://github.com/ektagarg/go-ds/blob/master/Arrays/Sort/bubble_sort/Readme.md

**2. Insertion Sort:** Insertion sort is based on the idea that one element from the input elements is consumed in each iteration to find its correct position i.e, the position to which it belongs in a sorted array.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n)
```
For more details, you can read this: https://github.com/ektagarg/go-ds/blob/master/Arrays/Sort/insertion_sort/Readme.md

**3. Selection Sort:** The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n2)
```
For more details, you can read this: https://github.com/ektagarg/go-ds/blob/master/Arrays/Sort/selection_sort/Readme.md

**4. Quick Sort:** Quick sort is highly efficient algorithm which is based on Divide and Conquer strategy. So, the technique is to divide the given array into smaller arrays.
A large array is partitioned into two arrays one of which holds values smaller than a specific value, say pivot, based on which the partition is made and another array holds values greater than the pivot value.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```
For more details, you can read this: https://github.com/ektagarg/go-ds/blob/master/Arrays/Sort/quick_sort/Readme.md

**5. Merge Sort:** Like quick sort, merge sort is based on divide and conquer methodology. It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.
```
Worst Case Time Complexity: O(nlog(n))
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```
For more details, you can read this: https://github.com/ektagarg/go-ds/blob/master/Arrays/Sort/merge_sort/Readme.md

**6. Heap Sort:** Heap sort in comparision based, in-memory sorting technique based on binary heap data structures. It involves building a Heap data structure from the given array and then utilizing the Heap to sort the array.
```
Worst Case Time Complexity: O(nlog(n))
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```
For more details, you can read this: https://github.com/ektagarg/go-ds/blob/master/Arrays/Sort/heap_sort/Readme.md
