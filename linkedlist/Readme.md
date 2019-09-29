# Inserting a node before the head(at beginning)
In this case, a new node is inserted before the current head node. Only one pointer(next) needs to be modified to new node's next pointer.

### Algorithm
1. Update the next pointer of new node, to point to the current node.
2. Update the head pointer to point to the new node.

### Example
head pointing to: 15
data ---> 15->7->40->NULL

head points to: data
data->15->7->40->NULL

### complexity
time complexity: o(1)