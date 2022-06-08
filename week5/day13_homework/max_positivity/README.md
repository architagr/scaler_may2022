# Problem Description

Given an array of integers **A**, of size **N**.

Return the maximum size subarray of **A** having only non-negative elements. If there are more than one such subarray, return the one having the smallest starting index in **A**.

**NOTE:** It is guaranteed that an answer always exists.

###### Problem Constraints

```
1 <= N <= 10^5
-10^9 <= A[i] <= 10^9
```

###### input format

``` 
The first and only argument given is the integer array A.
```

###### Output Format

```
Return maximum size subarray of A having only non-negative elements. If there are more than one such subarrays, return the one having earliest starting index in A.
```

###### Example Input

```
# input 1 : 
    A = [5, 6, -1, 7, 8]
# output 1: 
    [5, 6]
# input 2 : 
    A = [1, 2, 3, 4, 5, 6]
# output 2: 
    [1, 2, 3, 4, 5, 6]
```
