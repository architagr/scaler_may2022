# Problem Description

Given two sorted arrays of distinct integers, A and B, and an integer C, find and return the pair whose sum is closest to C and the pair has one element from each array.

More formally, find A[i] and B[j] such that abs((A[i] + B[j]) - C) has minimum value.

If there are multiple solutions find the one with minimum i and even if there are multiple values of j for the same i then return the one with minimum j.

Return an array with two elements {A[i], B[j]}.

###### Problem Constraints

```
1 <= length of both the arrays <= 100000

1 <= A[i], B[i] <= 10^9

1 <= C <= 10^9
```

###### input format

``` 
The first argument given is the integer array A.
The second argument given is the integer array B.
The third argument given is integer C.
```

###### Output Format

```
Return an array of size 2 denoting the pair which has sum closest to C.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3, 4, 5]
    B = [2, 4, 6, 8]
    C = 9
# output 1: 
    [1, 8]
# input 2: 
    A = [5, 10, 20]
    B = [1, 2, 30]
    C = 13
# output 2: 
    [10, 2]
```
