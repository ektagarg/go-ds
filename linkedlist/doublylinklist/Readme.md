# Doubly link list
Doubly linked list(also called two-way linked list) can be traversed in both directions(forward and backward).

### Advantages
A node in a single link list cannot be removed unless we have the pointer to its predecessor.
But in doubly link list, we can delete a node even if we don't have previous node address(since, each node has left pointer pointing to previous node and can move backward).

### Disadvantages
1. Each node requires an extra pointer, requiring more space.
2. This insertion or deletion of a node takes a bit longer(more pointer operations)

### Operations
###### Insertion
1. Inserting a new node before the head.
2. Inserting a new node after the tail.
3. Inserting a new node at given position.
```
Time Complexity: O(n), In the worst case when we need to insert node at the end of the list.
Space Complexity: O(1), For creating a temp variable.
```

###### Deletion
1. Deleting the first node.
2. Deleting the last node.
3. Deleting an intermediate node.
```
Time Complexity: O(n), for scanning the complete list of size n.
Space Complexity: O(1), For creating a temp variable.
```