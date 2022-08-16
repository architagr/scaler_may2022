# Problem Description

Given an array **A** having **N** positive numbers. You have to find the number of **Prime subsequences** of **A**.

A Prime subsequence is one that has **only prime** numbers, for example [2, 3], [5] are the Prime subsequences where [2, 4] and [1, 2, 3, 4] are not.

###### Problem Constraints

```
1 <= N <= 1e3
1 <= A[i] <= 1e6
```

###### input format

``` 
The first argument given is an Array A, having N integers.
```

###### Output Format

```
Return an integer X, i.e number of Prime subsequences. 
As X can be very large print X % (1000000007), here % is modulus operator.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3]
# output 1: 
    3
# input 2: 
    A = [1, 2, 3, 5]
# output 2: 
    7
```
