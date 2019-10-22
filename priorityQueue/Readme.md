# Priority queue
In some situations we need to find the minimum/maximum element among a collection of elements.
We can do this with the help of Priority Queue ADT.
A priority queue is a data structure that supports the operations like:
1. MinHeap(It returns the minimum element)
2. MaxHeap(It returns the maximum element)

These operations are equivalent to EnQueue and DeQueue operations of a queue. The difference is that, in priority queue, the order in which elements arrive, may not be the order in which elemnts get processed.

Ascending PQ - If the item with the smallest key has the highest priority. (delete smallest element first)

Descending PQ - If the item with the largest key has the smallest priority. (delete largest element first)

### Application
1. Job Scheduling - which is prioritized instead of first come first serve.
2. Data compression - Huffman Coding algorithm
3. Shortest path algorithm - Dijkstra's algorithm
4. Minimum spanning tree - Prim's algorithm
5. Even driven simulation
6. Selection problem - finding kth smallest element

### PQ Implementation

**Unordered array implementation** - Elements are inserted into the array without bothering about the order. But deletions are performed by searching the key and followed by deletion.
Complexity - Insertion(1), Deletion(n), Find(n)

**Unordered list implemenation** - Similar to array implementation instead using linkedlist.
Complexity - Insertion(1), Deletion(n), Find(n)

**Ordered array implementation** - Elements are inserted into the array in sorted order based on key field. Deletions are performed at one end.
Complexity - Insertion(n), Deletion(1), Find(1)

**Ordered list implemenation** - Similar to array implementation instead using linkedlist.
Complexity - Insertion(n), Deletion(1), Find(1)

**Binary heap imlementation** - A heap is a tree with special properties which is used for implementation.
[What is Heap?](../array/sort/heap_sort/Readme.md)
Complexity - Insertion(logn), Deletion(logn), Find(1)



