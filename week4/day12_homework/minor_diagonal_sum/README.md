# Problem Description

You are given a **N X N** integer matrix. You have to find the sum of all the minor diagonal elements of **A**.

Minor diagonal of a **M X M** matrix **A** is a collection of elements **A[i, j]** such that **i + j = M + 1** (where i, j are 1-based).

###### Problem Constraints

```
1 <= N <= 10^3
-1000 <= A[i][j] <= 1000
```

###### input format

``` 
First and only argument is a 2D integer matrix A.
```

###### Output Format

```
Return an integer denoting the sum of minor diagonal elements.
```

###### Example Input

```
# input 1 : 
 A = [[1, -2, -3],
      [-4, 5, -6],
      [-7, -8, 9]]
# output 1: 
    -5
```
