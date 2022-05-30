# Problem Description

Given an integer array **A** containing **N** distinct integers, you have to find all the **leaders** in array **A**.

An element is a leader if it is strictly greater than all the elements to its right side.

**NOTE:** The rightmost element is always a leader.

###### Problem Constraints

```
1 <= N <= 10^5

1 <= A[i] <= 10^8
```

###### input format

``` 
First and only argument is an integer array A.
```

###### Output Format

```
Return an integer array denoting all the leader elements of the array.

NOTE: Ordering in the output doesn't matter.
```

###### Example Input

```
# input 1 : 
A = [16, 17, 4, 3, 5, 2]
# output 1: 
    [17, 2, 5]
# input 2 : 
    A = [1,2]
# output 2: 
    [2]
```
