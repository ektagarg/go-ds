# Stack
A stack is an ordered list in which insertion and deletion are done at one end, called top. The last element inserted is the first one to be deleted. Hence it is called LIFO(Last In First Out) or First In Last Out(FILO) list.

To understand it in a better way, we can take this example of pile of plates of a cafeteria, The plates are added to the stack as they are cleaned. They are placed on the top. When a plate is required it is taken from the top of the stack. The first plate placed on the satck os the last one to be used.

### Operations:
1. Push: When as elemnt is inserted in stack, the operation is called Push.
2. Pop: When an element is deleted from stack, the operation is called Pop.

```
                          push 3                        Pop: (Element after performing pop: 3)

                            3                             
12                          12                          12
43                          43                          43
33                          33                          33
34                          34                          34
```

### Applications:

**Direct applications:**
1. Balancing of symbols.
2. Infix to postfix conversion.
3. Evaluation of postfix Expression.
4. Implementing the function calls(including recursion)
5. Page visited history in web browser.
6. Undo sequence in a text editor.
7. Matching tags in HTML and XML.

**Indirect Applications:**
1. Auxiliary data structure for other algorithms. (Example: Tree traversal algorithm)
2. Component of other data structures (Example: Stimulating queues)