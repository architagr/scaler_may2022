# Problem Description

Given an array of positive integers A and an integer B, find and return first continuous subarray which adds to B.

If the answer does not exist return an array with a single element "-1".

First sub-array means the sub-array for which starting index in minimum.

###### Problem Constraints

```
1 <= length of the array <= 100000
1 <= A[i] <= 10^9
1 <= B <= 10^9
```

###### input format

``` 
The first argument given is the integer array A.

The second argument given is integer B.
```

###### Output Format

```
Return the first continuous sub-array which adds to B and if the answer does not exist return an array with a single element "-1".
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3, 4, 5]
    B = 5
# output 1: 
    [2, 3]
# input 2: 
    A = [5, 10, 20, 100, 105]
    B = 110
# output 2: 
    [-1]
```
