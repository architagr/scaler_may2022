# Problem Description

Given two arrays of integers **A** and **B** describing a pair of **(A[i], B[i])** coordinates in a 2D plane. **A[i]** describe x coordinates of the ith point in the 2D plane, whereas **B[i]** describes the y-coordinate of the ith point in the 2D plane.

Find and return the **maximum** number of points that lie on the same line.

###### Problem Constraints

```
1 <= (length of the array A = length of array B) <= 1000
-10^5 <= A[i], B[i] <= 10^5
```

###### input format

``` 
The first argument is an integer array A.
The second argument is an integer array B.
```

###### Output Format

```
Return the maximum number of points which lie on the same line.
```

###### Example Input

```
# input 1: 
    A = [-1, 0, 1, 2, 3, 3]
    B = [1, 0, 1, 2, 3, 4]
# output 1: 
    4
# input 2: 
    A = [3, 1, 4, 5, 7, -9, -8, 6]
    B = [4, -8, -3, -2, -1, 5, 7, -4]
# output 2: 
    2
```
