# Heap Sort
Heap sort in comparision based, in-memory sorting technique based on binary heap data structures. It involves building a Heap data structure from the given array and then utilizing the Heap to sort the array.

### What is a Heap ?
Heap is a special tree-based data structure with the following properties:

1. Shape Property: Heap data structure is always a Complete Binary Tree, which means all levels of the tree are fully filled.

2. Heap Property: All nodes are either greater than or equal to or less than or equal to each of its children. If the parent nodes are greater than their child nodes, heap is called a Max-Heap, and if the parent nodes are smaller than their child nodes, heap is called Min-Heap.

###### Min Heap
In min-heap, first element is the smallest. So, when we want to sort a list in ascending order, we create a min-heap from that list, and picks the first element as smallest and then repeat the following procedure with all the elements.

###### Max Heap
In max-heap (or reverse of min-heap), the first element is the largest, hence it is used when we need to sort the list in descending order.

### Algorithm
1. Creating a Heap of the unsorted list/array.
2. Then a sorted array is created by repeatedly removing the largest/smallest element from the heap, and inserting it into the array. The heap is reconstructed after each removal.

### Example
```
Input data: 41, 98, 32, 56, 11
         41(0)
        /   \
     98(1)   32(2)
    /   \
 56(3)    11(4)

note: The numbers in bracket represent the indices in the array 
representation of data.

Applying heapify procedure to index 1:
         41(0)
        /   \
    98(1)    32(2)
    /   \
56(3)    11(4)

Applying heapify procedure to index 0:
        98(0)
        /  \
     56(1)  32(2)
    /   \
 41(3)    11(4)
The heapify procedure calls itself recursively to build heap
 in top down manner.
```

### Complexity
Best-case: o(nlogn)
Worst-case: o(nlogn)
Avg-case: o(nlogn)
Space: o(n)