# Problem Description

Given an integer array **A** containing **N** distinct integers.

Find the number of unique pairs of integers in the array whose XOR is equal to **B**.

**NOTE:**

Pair (a, b) and (b, a) is considered to be the same and should be counted once.

###### Problem Constraints

```
1 <= N <= 10^5

1 <= A[i], B <= 10^7
```

###### input format

``` 
The first argument is an integer array A.

The second argument is an integer B.
```

###### Output Format

```
Return a single integer denoting the number of unique pairs of integers in the array A whose XOR is equal to B.
```

###### Example Input

```
# input 1 : 
    A = [5, 4, 10, 15, 7, 6]
    B = 5
# output 1: 
    1
# input 2 : 
    A = [3, 6, 8, 10, 15, 50]
    B = 5
# output 2: 
    2
```
