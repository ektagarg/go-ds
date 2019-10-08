# Trees
Tree is an example of non-linear data structures. It is a data structure similar to a linked list but instead of each node pointing simply to the next node in a linear fashion, each node points to a number of nodes. 
A tree structure is a way of representing the hierarchical nature of a structure in a graphical form.

In trees, order of the elements is not important. If we need ordering information linear data structures like linked lists, stacks, queues can be used.

```
                                3         <-- Level-0
                                /\
                               1  4        <-- Level-1
                                  /\
                                 2  6       <-- Level-2
```
### Binary trees
A tree is called binary tree if each node has zero child, one child or two children. Empty tree is also a valid binary tree. We can visualize a binary tree as consisting of a root and two disjoint binary trees, called the left and right subtrees of the root.

### Types of binary trees

**1. Strict binary tree:** A binary tree is called strict binary tree if each node has exactly two children or no children.
```
                                3         
                                /\
                               1  4        
                                  /\
                                 2  6
```

**2. Full binary tree:** A binary tree is called a full binary tree if each node has exactly two children and all leaf nodes are present at same level.
```
                                1         
                                /\
                               2  3        
                              /\   /\
                             4  5 6  7
```                                 
**2. Complete binary tree:** In complete binary trees, if we give numbering for the nodes by starting at root then we get a complete binary trees, then we get a complete sequence from 1 to number of nodes in the tree. While traversing we should give numbering for NULL pointers also. A binary tree is called Complete binary tree if all leaf nodes are at height h or h-1 and also without any missing number in the sequence.
```
                                1         
                                /\
                               2  3        
                              /\   
                             4  5  
``` 

### Applications of Binary Trees:
1. Expression trees are used in Compilers.
2. Huffman coding trees that are used in compression algorithms.
3. Binary search trees supports search, insertion, deletion on a collection of items in 0(logn)

### Binary Tree traversal
**1. PreOrder Traversal (Root-Left-Right):** In this traversal each node is processed before each of its subtree.

Process:
* Visit the Root
* Traverse the left subtree in Preorder
* Traverse the right subtree in Preorder

Time complexity: o(n)
Space complexity: o(n)

**2. InOrder Traversal (Left-Root-Right):** In this traversal the root is visited between its subtrees. It gives a sorted array as a result when applied on binary search tree

Process:
* Traverse the left subtree in InOrder
* Visit the Root
* Traverse the right subtree in InOrder

Time complexity: o(n)
Space complexity: o(n)

**3. PostOrder Traversal (Left-Right-Root):** In this traversal each node is processed after each of its subtree.

Process:
* Traverse the left subtree in PostOrder
* Traverse the right subtree in PostOrder
* Visit the Root

Time complexity: o(n)
Space complexity: o(n)
