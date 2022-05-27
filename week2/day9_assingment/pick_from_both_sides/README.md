# Problem Description

You are given an integer array **A** of size **N**.

You have to pick **B** elements in total. Some (possibly 0) elements from left end of array A and some (possibly 0) from the right end of array **A** to get the maximum sum.

Find and return this **maximum possible sum.**

**NOTE:** Suppose B = 4, and array A contains 10 elements, then

You can pick the first four elements or can pick the last four elements, or can pick 1 from front and 3 from the back, etc. You need to return the maximum possible sum of elements you can pick.

###### Problem Constraints

```
1 <= N <= 10^5
1 <= B <= N
-10^3 <= A[i] <= 10^3
```

###### input format

``` 
First argument is an integer array A.

Second argument is an integer B.
```

###### Output Format

```
Return an integer denoting the maximum possible sum of elements you picked.
```

###### Example Input

```
# input 1 : 
    A = [5, -2, 3 , 1, 2]
    B = 3
# output 1: 
    8
# input 2 : 
    A = [1, 2]
    B = 1
# output 2: 
    2
```
