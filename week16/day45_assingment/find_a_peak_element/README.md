# Problem Description

Given an array of integers **A**, find and return the peak element in it. An array element is peak if it is NOT smaller than its neighbors.

For corner elements, we need to consider only one neighbor. We ensure that answer will be unique.

**NOTE:** Users are expected to solve this in O(log(N)) time. The array may have duplicate elements.

###### Problem Constraints

```
1 <= |A| <= 100000

1 <= A[i] <= 10^9
```

###### input format

``` 
The only argument given is the integer array A.
```

###### Output Format

```
Return the peak element.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3, 4, 5]
# output 1: 
    5
# input 2: 
    A = [5, 17, 100, 11]
# output 2: 
    100
```
