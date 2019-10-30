# Kruskal Algorithm for Minimum Spanning Tree

A Minimum Spanning Tree is a tree that include all the vertex of a graph such that the total sum of all the edges are minimized. Kruskal alogorithm suggest a greedy approach for this, We sorts all the edges in non decreasing way and keep selecting the minimum edges and adding it to our MST. When we add an Edge to existing MST it should not form a cycle.

Let's look at it with example.
Suppose we have a graph as described below
```

     0 -------[4]----------- 1 -----[8]------ 2
     |                                        |
     |                                        |
    [8]       -- [7] -------- 5 ------[2]-----|
     |      /                 |
     |   /                   [6]
     | /                      |
     3 ------[1]-------- 4 ---|
```

it can be summarized as 
```
    Vertex From   Vertex To         Weight
        0           1                 4
        0           3                 8
        1           2                 8
        2           5                 2
        3           4                 1
        4           5                 6
        5           3                 7

```

Our Aim is to find the Minimum Spanning Tree From this Graph. 
First we will Sort the Edges in a non decreasing order. After sorting in non decreasing order we would have following order of edges.
```
  Vertex From   Vertex To         Weight
        3           4                 1
        2           5                 2
        0           1                 4
        4           5                 6
        5           3                 7
        0           3                 8
        1           2                 8

```
Now we will Select the minimum edges one by one try to add it to the MST if it does not form a cycle 
then we will add it to the MST else we will skip that edge.

The first edge we select is
```

     3-------[1]-------- 4 

     current MST. total weight 1
```

Then we move to next minimum edge. Similarly If we keep repeating at some point after adding 4 edges we will have our MST like this
```
     0 -------[4]----------- 1                2
                                              |
                                              |
                              5 ------[2]-----|
                              |
                             [6]
                              |
     3 ------[1]-------- 4 ---|
```

now the next minim edge is 5-3 with weight 7 but addding it will make a cycle so we will skip this edge. We will then try to add the edge 0-3 [weight=8]

```
     0 -------[4]----------- 1 -----[8]------ 2
                                              |
                                              |
                              5 ------[2]-----|
                              |
                             [6]
                              |
     3 ------[1]-------- 4 ---|

    Total Weight : 21
```
Thus we have obtained out Minimum Spanning Tree.



# Algorithm

The algorithm for kruskal works following way
1. Sort all the edges in non-decreasing order of their weight.
2. Pick the smallest edge. Check if it forms a cycle with the spanning tree formed so far. If cycle is not formed, include this edge. Else, discard it.
3. Repeat step#2 until there are (V-1) edges in the spanning tree.

