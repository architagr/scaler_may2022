# Problem Description

Given an array **A** of size **N**, Groot wants you to pick **2** indices **i** and **j** such that

1. **1 <= i, j <= N**
2. **A[i] % A[j]** is maximum among all possible pairs of **(i, j)**.
Help Groot in finding the maximum value of **A[i] % A[j]** for some **i, j**.

###### Problem Constraints

```
1 <= N <= 100000
0 <= A[i] <= 100000
```

###### input format

``` 
First and only argument in an integer array A.
```

###### Output Format

```
Return an integer denoting the maximum value of A[i] % A[j] for any valid i, j.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 44, 3]
# output 1: 
    3
# input 2: 
    A = [2, 6, 4]
# output 2: 
    4
```
