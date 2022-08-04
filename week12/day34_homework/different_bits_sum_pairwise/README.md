# Problem Description

We define f(X, Y) as the number of different corresponding bits in the binary representation of X and Y.
For example, f(2, 7) = 2, since the binary representation of 2 and 7 are 010 and 111, respectively. The first and the third bit differ, so f(2, 7) = 2.

You are given an array of N positive integers, A1, A2,..., AN. Find sum of f(Ai, Aj) for all pairs (i, j) such that 1 ≤ i, j ≤ N. Return the answer modulo 10^9+7.

###### Problem Constraints

```
1 <= N <= 10^5

1 <= A[i] <= 2^31 - 1
```

###### input format

``` 
The first and only argument of input contains a single integer array A.
```

###### Output Format

```
Return a single integer denoting the sum.
```

###### Example Input

```
# input 1: 
    A = [1, 3, 5]
# output 1: 
    8
# input 2: 
    A = [2, 3]
# output 2: 
    2
```
