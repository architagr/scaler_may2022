# Problem Description

You are given an array of **N** integers, **A1, A2, .... AN.**

Return the maximum value of f(i, j) for all 1 ≤ i, j ≤ N. f(i, j) is defined as |A[i] - A[j]| + |i - j|, where |x| denotes absolute value of x.

###### Problem Constraints

```
1 <= N <= 100000
-109 <= A[i] <= 109
```

###### input format

``` 
First argument is an integer array A of size N.
```

###### Output Format

```
Return an integer denoting the maximum value of f(i, j).
```

###### Example Input

```
# input 1: 
    A = [1, 3, -1]
# output 1: 
    5
# input 2: 
    A = [2]
# output 2: 
    0
```
