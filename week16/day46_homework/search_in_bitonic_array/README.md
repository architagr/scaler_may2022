# Problem Description

Given a bitonic sequence **A** of **N** distinct elements, write a program to find a given element **B** in the bitonic sequence in **O(logN)** time.

**NOTE:**

A Bitonic Sequence is a sequence of numbers which is first strictly increasing then after a point strictly decreasing.

###### Problem Constraints

```
3 <= N <= 10^5

1 <= A[i], B <= 10^8

Given array always contain a bitonic point.

Array A always contain distinct elements.
```

###### input format

``` 
First argument is an integer array A denoting the bitonic sequence.

Second argument is an integer B.
```

###### Output Format

```
Return a single integer denoting the position (0 index based) of the element B in the array A if B doesn't exist in A return -1.
```

###### Example Input

```
# input 1: 
    A = [3, 9, 10, 20, 17, 5, 1]
    B = 20
# output 1: 
    3
# input 2: 
    A = [5, 6, 7, 8, 9, 10, 3, 2, 1]
    B = 30
# output 2: 
    -1
```
