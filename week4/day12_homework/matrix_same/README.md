# Problem Description

You are given two matrices A & B of equal dimensions and you have to check whether two matrices are equal or not.

NOTE: Both matrices are equal if A[i][j] == B[i][j] for all i and j in the given range.

###### Problem Constraints

```
1 <= A.size(), B.size() <= 1000

1 <= A[i].size(), B[i].size() <= 1000

1 <= A[i][j], B[i][j] <= 1000
```

###### input format

``` 
First argument is vector of vector of integers representing matrix A.

Second argument is vector of vector of integers representing matrix B.
```

###### Output Format

```
Return 1 if both matrices are equal or return 0.
```

###### Example Input

```
# input 1 : 
    A = [[1, 2, 3],[4, 5, 6],[7, 8, 9]]
    B = [[1, 2, 3],[4, 5, 6],[7, 8, 9]]
# output 1: 
    1
# input 2 : 
    A = [[1, 2, 3],[4, 5, 6],[7, 8, 9]]
    B = [[1, 2, 3],[7, 8, 9],[4, 5, 6]]
# output 2: 
    0
```
