# Problem Description

Given a sorted array A of size N and a target value B, return the index (0-based indexing) if the target is found.
If not, return the index where it would be if it were inserted in order.

NOTE: You may assume no duplicates in the array. Users are expected to solve this in O(log(N)) time.

###### Problem Constraints

```
1 <= N <= 10^6
```

###### input format

``` 
The first argument is an integer array A of size N.
The second argument is an integer B.
```

###### Output Format

```
Return an integer denoting the index of target value.
```

###### Example Input

```
# input 1: 
    A = [1, 3, 5, 6]
    B = 5 
# output 1: 
    2
# input 2: 
    A = [1]
    B = 1
# output 2: 
    0
```
