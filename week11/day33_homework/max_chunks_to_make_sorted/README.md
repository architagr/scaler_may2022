# Problem Description

Given an array of integers **A** of size **N** that is a permutation of **[0, 1, 2, ..., (N-1)]**, if we split the array into some number of "chunks" (partitions), and individually sort each chunk. After concatenating them in order of splitting, the result equals the sorted array.

What is the **most** number of chunks we could have made?

###### Problem Constraints

```
1 <= N <= 100000
0 <= A[i] < N
```

###### input format

``` 
The only argument given is the integer array A.
```

###### Output Format

```
Return the maximum number of chunks that we could have made.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3, 4, 0]
# output 1: 
    1
# input 2: 
    A = [2, 0, 1, 3]
# output 2: 
    2
```
