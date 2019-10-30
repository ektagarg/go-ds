# Lists

## Linked Lists

### Singly Linked List
### Doubly Linked List
### Circular Linked List


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
