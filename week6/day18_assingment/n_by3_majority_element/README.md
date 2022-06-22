# Problem Description

You're given a read-only array of **N** integers. Find out if any integer occurs more than N/3 times in the array in linear time and constant additional space.
If so, return the integer. If not, return **-1.**

If there are multiple solutions, return any one.

###### Problem Constraints

```
1 <= N <= 7*10^5
1 <= A[i] <= 10^9
```

###### input format

``` 
The only argument is an integer array A.
```

###### Output Format

```
Return an integer.
```

###### Example Input

```
# input 1 : 
   [1 2 3 1 1]
# output 1: 
    1
# input 2 : 
    [3, 2, 3, 1, 2, 8, 9, 3, 2, 3, 2]
# output 2: 
    2 
# input 3 : 
    [1]
# output 3: 
    1 
```
