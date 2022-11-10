# Problem Description

Given an array of integers A, find and return the maximum result of A[i] XOR A[j], where i, j are the indexes Given an integer array A of size N.

You have to find the product of the three largest integers in array A from range 1 to i, where i goes from 1 to N.

Return an array B where B[i] is the product of the largest 3 integers in range 1 to i in array A. If i < 3, then the integer at index i in B should be -1.

###### Problem Constraints

```
1 <= N <= 10^5
0 <= A[i] <= 10^3
```

###### input format

``` 
First and only argument is an integer array A.
```

###### Output Format

```
Return an integer array B. B[i] denotes the product of the largest 3 integers in range 1 to i in array A.
```

###### Example Input

```
# input 1 : 
   A = [1, 2, 3, 4, 5]
# output 1: 
   [-1, -1, 6, 24, 60]
# input 2: 
  A = [10, 2, 13, 4]
# output 2: 
  [-1, -1, 260, 520]
```
