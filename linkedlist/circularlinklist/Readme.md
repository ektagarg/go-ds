# Circular linked list
In circular linked list, there is no end. While traversing the circular linked list we should be careful otherwise we will be traversing the list indefinitely. 
In circular linked list each node has a successor. 
Note: Unlike singly linked lists, there is no node with NULL pointer in a circular linked list. 

### Use case
When several processes are using the same computer resource(CPU) for the same amount of time, we have to assure that no process accesses the resource before all other processes did(Round Robin Algorithm)

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