# Problem Description

Given two arrays of integers **A** and **B**, Sort **A** in such a way that the relative order among the elements will be the same as those are in **B**.
For the elements not present in **B**, append them at last in sorted order.

Return the array **A** after sorting from the above method.

**NOTE:** Elements of B are unique.

###### Problem Constraints

```
1 <= length of the array A <= 100000
1 <= length of the array B <= 100000
-10^9 <= A[i] <= 10^9
```

###### input format

``` 
The first argument given is the integer array A.

The second argument given is the integer array B.
```

###### Output Format

```
Return the array A after sorting as described.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3, 4, 5]
    B = [5, 4, 2]
# output 1: 
    [5, 4, 2, 1, 3]
# input 2: 
    A = [5, 17, 100, 11]
    B = [1, 100]
# output 2: 
    [100, 5, 11, 17]
```
