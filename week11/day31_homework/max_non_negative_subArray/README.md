# Problem Description

Given an array of integers, **A** of length **N**, find out the **maximum sum** sub-array of non negative numbers from **A.**

The sub-array should be contiguous i.e., a sub-array created by choosing the second and fourth element and skipping the third element is invalid.

Maximum sub-array is defined in terms of the sum of the elements in the sub-array.

Find and return the required subarray.

**NOTE:**

    1. If there is a tie, then compare with segment's length and return segment which has maximum length.
    2. If there is still a tie, then return the segment with minimum starting index.

###### Problem Constraints

```
1 <= N <= 1e5
-INT_MAX < A[i] <= INT_MAX
```

###### input format

``` 
The first and the only argument of input contains an integer array A, of length N.
```

###### Output Format

```
Return an array of integers, that is a subarray of A that satisfies the given conditions.
```

###### Example Input

```
# input 1 : 
    A = [1, 2, 5, -7, 2, 3]
# output 1: 
    [1, 2, 5]
# input 2 : 
    A = [10, -1, 2, 3, -4, 100]
# output 2: 
    [100]
```
