# go-ds
A library which aims to explain and implement data structures in Golang.

## Types of data structures
# Arrays
An array is a data structure that contains fixed-sized collection of elements of same data type, such as an integer or string. Arrays are commonly used in computer programs to organize data so that a related set of values can be easily sorted or searched.

Operations on Arrays:
### Searching
**1. [Linear Search](array/search/binary_search/Readme.md):** A linear search or sequential search is a method for finding an element within a list. It sequentially checks each element of the list until a match is found or the whole list has been searched. 
```
Worst-case complexity: o(n)
Best-case complexity: o(1)
```

**2. [Binary Search](array/search/binary_search/Readme.md):** A binary search or half-interval search aims to find the element in a sorted array by first comparing target element with middle element to see whether the element is in first half or second half and keep repeating this until it finds the target element.
```
Worst-case complexity: o(logn)
Best-case complexity: o(1)
```

### Sorting

**1. [Bubble Sort](array/sort/bubble_sort/Readme.md):** It is a comparision-based sorting algorithm where each pair of adjacent elements is compared and elements get swapped if they are not in order.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n)
```

**2. [Insertion Sort](array/sort/insertion_sort/Readme.md):** Insertion sort is based on the idea that one element from the input elements is consumed in each iteration to find its correct position i.e, the position to which it belongs in a sorted array.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n)
```

**3. [Selection Sort](array/sort/selection_sort/Readme.md):** The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(n2)
```

**4. [Quick Sort](array/sort/quick_sort/Readme.md):** Quick sort is highly efficient algorithm which is based on Divide and Conquer strategy. So, the technique is to divide the given array into smaller arrays.
A large array is partitioned into two arrays one of which holds values smaller than a specific value, say pivot, based on which the partition is made and another array holds values greater than the pivot value.
```
Worst Case Time Complexity: O(n2)
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```

**5. [Merge Sort](array/sort/merge_sort/Readme.md):** Like quick sort, merge sort is based on divide and conquer methodology. It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.
```
Worst Case Time Complexity: O(nlog(n))
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```

**6. [Heap Sort](array/sort/heap_sort/Readme.md):** Heap sort in comparision based, in-memory sorting technique based on binary heap data structures. It involves building a Heap data structure from the given array and then utilizing the Heap to sort the array.
```
Worst Case Time Complexity: O(nlog(n))
Best Case Time Complexity: O(nlog(n)) 
Average Time Complexity: O(nlog(n))
```

**7. [Radix Sort](array/sort/radix_sort/Readme.md):** It is not based on comparision-based algorithms, instead it is an integer sorting algorithm that sorts data with integer keys by grouping the keys by individual digits that share the same significant position and value.
```
time complexity: o(kn) 
where k is total number of digits in the largest number bcs that is number of iterations
```

**8. [Shell Sort](array/sort/shell_sort/Readme.md):** Shell sort is a highly efficient sorting algorithm and is based on insertion sort algorithm. This algorithm avoids large shifts as in case of insertion sort, if the smaller value is to the far right and has to be moved to the far left.

This algorithm uses insertion sort on a widely spread elements, first to sort them and then sorts the less widely spaced elements. This spacing is termed as interval. This interval is calculated based on Knuth's formula(h = h * 3 + 1).
```
Worst complexity: Depends on gap sequence
Average complexity: n*log(n)^2 or n^(3/2)
Best complexity: n
```
# [Linked lists](linkedlist/Readme.md)
It is a data structure for storing collections of data with the following properties:
* Successive elements are connected via pointers
* Last element point to NULL
* Can grow or shrink in size during the execution of the program
* Can be made just as long as required(until system memory exhausts)

### Operations:
Insertion:
1. Insertion at the beginning
2. Insertion at the end
3. Insertion in the middle

Deletion:
1. Deletion from the beginning
2. Deletion from the end
3. Deletion from the middle

# [Stack](stack/Readme.md)
A stack is an ordered list in which insertion and deletion are done at one end, called top. The last element inserted is the first one to be deleted. Hence it is called LIFO(Last In First Out) or First In Last Out(FILO) list.

### Operations:
1. Push
2. Pop