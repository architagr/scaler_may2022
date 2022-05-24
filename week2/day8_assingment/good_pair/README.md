# Problem Description

Given an array **A** and an integer **B**. A **pair(i, j)** in the array is a good pair if **i != j** and **(A[i] + A[j] == B)**. Check if any good pair exist or not.

###### Problem Constraints

```
1 <= A.size() <= 104

1 <= A[i] <= 109

1 <= B <= 109
```

###### input format

``` 
First argument is an integer array A.

Second argument is an integer B.
```

###### Output Format

```
Return 1 if good pair exist otherwise return 0.
```

###### Example Input

```
# input 1 : 
    A = [1,2,3,4]
    B = 7
# output 1: 
    1 
# input 2 : 
    A = [1,2,4]
    B = 4
# output 2: 
    0
```
