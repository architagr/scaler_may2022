# Problem Description

You are given a matrix **A**, you have to return another matrix which is the transpose of A.

**NOTE:** Transpose of a matrix A is defined as - AT[i][j] = A[j][i] ; Where 1 ≤ i ≤ col and 1 ≤ j ≤ row

###### Problem Constraints

```
1 <= A.size() <= 1000

1 <= A[i].size() <= 1000

1 <= A[i][j] <= 1000
```

###### input format

``` 
First argument is vector of vector of integers A representing matrix.
```

###### Output Format

```
You have to return a vector of vector of integers after doing required operations.
```

###### Example Input

```
# input 1 : 
    A = [[1, 2, 3],[4, 5, 6],[7, 8, 9]]
# output 1: 
    [[1, 4, 7], [2, 5, 8], [3, 6, 9]]
# input 2 : 
    A = [[1, 2],[1, 2],[1, 2]]
# output 2: 
    [[1, 1, 1], [2, 2, 2]]
```
