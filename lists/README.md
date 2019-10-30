# Lists

## Linked List

It is a data structure for storing collections of data with the following properties:
* Successive elements are connected via pointers
* Last element point to NULL
* Can grow or shrink in size during the execution of the program
* Can be made just as long as required(until system memory exhausts)

### Operations
###### 1. Inserting a node before the head(at beginning)
In this case, a new node is inserted before the current head node. Only one pointer(next) needs to be modified to new node's next pointer.

###### Algorithm
1. Update the next pointer of new node, to point to the current node.
2. Update the head pointer to point to the new node.

###### Example
```
new node : 12
head pointing to: 15
data ---> 15->7->40->NULL

head points to: data
data->15->7->40->NULL
12->15->7->40->NULL
```

### complexity
time complexity: o(1)

### 2. Inserting a node in the end
In this we traverse the list till end, and then insert the new node in the list.

###### Algorithm
1. New node right pointer points to NULL and left pointer points to the end of list.
2. Update right of pointer of last node to point to new node.

###### Example
```
new node : 12
head pointing to: 15
15->7->40->NULL

15->7->40->12->NULL
```

### 3. Inserting node in the Middle
Traverse the list till that position and insert new node.

###### Algorithm
1. New node right pointer points to the next node of the position node, where we want to insert the new node. Also, new node left pointer points to the position node.
2. Position node right pointer points to the new node and the next node of position nodes left pointer points to new node

###### Example
```
new node : 12, position:3
head pointing to: 15
15->7->40->NULL

15->7->12->40->NULL
```

### 4. Deleting the first node
Current head node to remove from the list.

###### Algorithm
1. Create a temp node which will point to same node as that of head.
2. Now, move the head nodes pointer to the next node and change the heads left pointer to NULL.

###### Example
```
before deletion: 15->7->12->40->NULL
delete
after deletion: 7->12->40->NULL
```

### 5. Deleting the last node
Traverse till last node and delete that node.

###### Algorithm
1. Traverse the list and while traversing maintain the previous node address also. By the time, we reach the end of list, we will have two pointers one pointing to the tail node and other pointing to the node before tail node.
2. Update the previous node's next pointer with null.
3. Delete the tail node.

###### Example
```
before deletion: 15->7->12->40->NULL
delete
after deletion: 15->7->12->NULL
```

### 6. Deleting the Middle node
Delete a node located between two nodes. Head and tails are not updated, so the removal can be done is two steps:

###### Algorithm
1. Maintain previous node while traversing the list. Once we found the node to be deleted, change the previous node's next pointer to the next pointer of the node to be deleted.
2. Delete the current node to be deleted.

###### Example
```
before deletion: 15->7->12->40->NULL
delete node at 2
after deletion: 15->12->40->NULL
```


## Doubly Linked List

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




## Circular Linked List

In circular linked list, there is no end. While traversing the circular linked list we should be careful otherwise we will be traversing the list indefinitely.
In circular linked list each node has a successor.
**Note:** Unlike singly linked lists, there is no node with NULL pointer in a circular linked list.

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




## Skip List

If you ever needed to use a subway to go around you probably know that there are different ways of going from one station to another.

```
                            A------B------C------D------E------F------G------H------I------J

                    Line 1  X------X------X------X------X------X------X------X------X------X

                    Line 2  X-------------X-------------X-------------X-------------X------X

                    Line 3  X---------------------------X----------------------------------X

                    Line 4  X--------------------------------------------------------------X
```
Imagine that you need to go from station A to station J. The fastest line is 4 because it goes directly from A to J. However,if we wanted to go from A to B then we'd go with line 1. But what if we wanted to go from A to H? Well, we'd take the line 3 to go from A to E, then we'd switch lines at station E and use line 2 to go from E to G, then switch lines again and take line 1 to go from G to H.

Skip lists work just like subways. In a skip list each station is a leaf node containing data and each
