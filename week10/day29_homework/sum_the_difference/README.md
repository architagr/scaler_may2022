# Problem Description

Given an integer array, **A** of size **N.**
You have to find all possible non-empty subsequences of the array of numbers and then, for each subsequence, find the difference between the largest and smallest numbers in that subsequence. Then add up all the differences to get the number.

As the number may be large, output the number modulo 1e9 + 7 (1000000007).

**NOTE:** Subsequence can be non-contiguous.

###### Problem Constraints

```
1 <= N <= 10000
1<= A[i] <=1000
```

###### input format

``` 
First argument is an integer array A.
```

###### Output Format

```
Return an integer denoting the output.
```

###### Example Input

```
# input 1 : 
 A = [1, 2] 
# output 1: 
  1
# input 2 : 
  A = [1] 
# output 2: 
  0
```
