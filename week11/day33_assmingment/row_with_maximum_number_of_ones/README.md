# Problem Description

Given a binary sorted matrix **A** of size **N x N.** Find the row with the **maximum** number of **1.**

**NOTE:**

1. If two rows have the maximum number of 1 then return the row which has a **lower index.**
2. Rows are numbered from top to bottom and columns are numbered from left to right.
3. Assume **0-based** indexing.
4. Assume each row to be sorted by values.
5. Expected time complexity is O(rows).

###### Problem Constraints

```
1 <= N <= 1000

0 <= A[i] <= 1
```

###### input format

``` 
The only argument given is the integer matrix A.
```

###### Output Format

```
Return the row with the maximum number of 1.
```

###### Example Input

```
# input 1: 
     A = [   [0, 1, 1]
         [0, 0, 1]
         [0, 1, 1]   ]
# output 1: 
    0
# input 2: 
     A = [   [0, 0, 0, 0]
         [0, 1, 1, 1]    ]
# output 2: 
    1
```
