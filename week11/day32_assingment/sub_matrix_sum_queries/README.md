# Problem Description

Given a matrix of integers **A** of size **N x M** and multiple queries **Q**, for each query, find and return the submatrix sum.

Inputs to queries are **top left (b, c)** and bottom right **(d, e)** indexes of submatrix whose sum is to find out.

**NOTE:**

1. Rows are numbered from top to bottom, and columns are numbered from left to right.
2. Sum may be large, so return the answer mod 10^9 + 7.

###### Problem Constraints

```
1 <= N, M <= 1000
-100000 <= A[i] <= 100000
1 <= Q <= 100000
1 <= B[i] <= D[i] <= N
1 <= C[i] <= E[i] <= M
```

###### input format

``` 
The first argument given is the integer matrix A.
The second argument given is the integer array B.
The third argument given is the integer array C.
The fourth argument given is the integer array D.
The fifth argument given is the integer array E.
(B[i], C[i]) represents the top left corner of the i'th query.
(D[i], E[i]) represents the bottom right corner of the i'th query.
```

###### Output Format

```
Return an integer array containing the submatrix sum for each query.
```

###### Example Input

```
# input 1 : 
    A = [   [1, 2, 3]
         [4, 5, 6]
         [7, 8, 9]   ]
    B = [1, 2]
    C = [1, 2]
    D = [2, 3]
    E = [2, 3]
# output 1: 
    [12, 28]

# input 2 :
    A = [   [5, 17, 100, 11]
            [0, 0,  2,   8]    ]
    B = [1, 1]
    C = [1, 4]
    D = [2, 2]
    E = [2, 4]
 # output 2: 
    [22, 19]
```
