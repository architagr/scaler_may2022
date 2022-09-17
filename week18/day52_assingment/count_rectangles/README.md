# Problem Description

Given two arrays of integers **A** and **B** of size **N** each, where each pair **(A[i], B[i])** for **0 <= i < N** represents a unique point **(x, y)** in 2D Cartesian plane.

Find and return the number of unordered quadruplet **(i, j, k, l)** such that **(A[i], B[i]), (A[j], B[j]), (A[k], B[k]) and (A[l], B[l])** form a rectangle with the rectangle having all the sides parallel to either x-axis or y-axis.

###### Problem Constraints

```
1 <= N <= 2000
0 <= A[i], B[i] <= 10^9
```

###### input format

``` 
The first argument given is an integer array A.
The second argument given is the integer array B.
```

###### Output Format

```
Return the number of unordered quadruplets that form a rectangle.
```

###### Example Input

```
# input 1: 
    A = [1, 1, 2, 2]
    B = [1, 2, 1, 2]
# output 1: 
    1
# input 2: 
    A = [1, 1, 2, 2, 3, 3]
    B = [1, 2, 1, 2, 1, 2]
# output 2: 
    3
```
