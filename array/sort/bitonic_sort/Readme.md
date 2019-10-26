# Bitonic Sort

Bitonic sort is a parallel sorting algorithm which performs O(n2 log n) comparisons. Although, the number of comparisons are more than that in any other popular sorting algorithm, It performs better for the parallel implementation because elements are compared in predefined sequence which must not be depended upon the data being sorted. The predefined sequence is called Bitonic sequence.

### Algorithm
1. Form a Bitonic sequence from the given random sequence which we have formed in the above step. We can consider it as the first step in our procedure. After this step, we get a sequence whose first half is sorted in ascending order while second step is sorted in descending order.
2. Compare first element of first half with the first element of the second half, then second element of the first half with the second element of the second half and so on. Swap the elements if any element in the second half is found to be smaller.
3. After the above step, we got all the elements in the first half smaller than all the elements in the second half. The compare and swap results into the two sequences of n/2 length each. Repeat the process performed in the second step recursively onto every sequence until we get single sorted sequence of length n.

### Complexity
```
Worst Case Time Complexity: O(log 2 n)
Best Case Time Complexity: O(log 2 n)
Average Time Complexity: O(log 2 n)
```
