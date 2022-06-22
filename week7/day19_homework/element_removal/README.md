# Problem Description

Given an integer array **A** of size **N**. You can **remove** any element from the array in one operation.
The cost of this operation is the **sum of all elements** in the array present **before this operation.**

Find the **minimum cost** to remove all elements from the array.

###### Problem Constraints

```
0 <= N <= 1000
1 <= A[i] <= 10^3
```

###### input format

``` 
First and only argument is an integer array A.
```

###### Output Format

```
Return an integer denoting the total cost of removing all elements from the array.
```

###### Example Input

```
# input 1 : 
    A = [2, 1]
# output 1: 
    4
# input 2 : 
    A = [5]
# output 2: 
    5
```
