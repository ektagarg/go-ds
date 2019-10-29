# Prim's Alogrithm for Minimum Spanning Tree

A Minimum Spanning Tree is a tree that include all the vertex of a graph such that the total sum of all the edges are minimized. Prims alogorithm suggest a greedy approach for this where we select the minimum weighted edge of all the possible edges.

Let's take a look at the below example to understand it even better

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
First we will start from the vertex 0 and add it to our MST. Vertex 0 connects to 1 and 3 with weights 4 and 8 respectively. We need to select the minimum weighted edge. Since edge to the vertex 1 has lower weight we will add vertex 1 to out MST.
```

     0 -------[4]-------- 1 

     current MST. total weight 4
```

Now we have two vertex in our MST (0 and 1). We will first check all the edges going from 0 and select the minimum. Here 0 connects to vertex 3 with edge of weight 0 and with 1 but we will not consider 1 as its already in minimum spanning tree. We have one edge from 0 and this will be the minimum edge from vertex 0. We will repeat this for vertex 1 as well. Vertex 1 connects to vertex 0 and vertex 2. We won't consider 0 as its already in the minimum spanning tree, so let's go to the vertex 2 . the edge weight here is also 8. So now we have following scenario
```
    vertex A    vertex B        weight
    0           3               8
    1           2               8
```
 We need to select the minimum out of these edges. Since, these have same weights, we can take any. Let's choose (0-3)[8] and add it to our MST.

Lets have a look at our updated MST .
```

     0 -------[4]------- 1 
     |                                                                    
    [8]  
     |       
     3 
     Total Weight : 12
```

Let's Repeat the above process, we have now 3, 1 and 0 vertices in MST. We will check all the vertex that connects to them and are not already in the MST. We can see 0 is exhausted, we have visited all the vertex that it connects to. So, we have to do comparison between 1 and 3. After selecting minimum edges that connects both of the vertex to an unvisited vertex we will get following edges.
```
    vertex A    vertex B        weight
    1           2               8
    3           4               1

    see here 3 also connects with 5 but the weight is 7 so out of (3-5[7] and 3-4[1]) minimum is 3-4 which has weight 1.
```
Out of them the minimum edge is 3-4 having weight 1 so we add it to our MST

```
     0 -------[4]----------- 1 
     |        
     |        
    [8]       
     |     
     |   
     3 ------[1]-------- 4 
     Total Weight : 13
```

Similarly try implementing the above step on vertices and in the end when all the vertices are visited we will have the following Minimum Spanning Tree.
```
     0 -------[4]----------- 1                2
     |                                        |
     |                                        |
    [8]                       5 ------[2]-----|
     |                        |
     |                       [6]
     |                        |
     3 ------[1]------ 4 -----| 

    Total Weight : 21
```
Thus we have obtained out Minimum Spanning Tree.

 ## Algorithm
 The Prims algorithm suggests a greedy method to select the Minimum Spanning Tree. The concept is simple, choose the minimum weighted edge of the remaining edges that can connect to one of the vertices already added in MST.

 we can create a minimum spanning tree by following algorithm steps :

    1. Maintain Visited array and mark it as false for all the vertices intially. basically it is used to track the vertex that are already added to the MST.
    2. First We will select first vertex and make it visited. 
    3. While all the vertex are not added to the MST we do :
        a. Get already visited vertices(visited[i]==true) and select all the edges that connects it to the adjacent vertices that are not visited yet . out of all those edges select the minimum. Repeat this for all the visited vertices.
        b. After step 'a' we will have some edges, take the minimum edge out of them and add it to our MST.
        c. Mark the vertex that the edge connects to as visited.
    4. We have got ourself a MST.
