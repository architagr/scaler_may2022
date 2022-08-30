# Problem Description

Given an array **A** of non-negative integers of size **N**. Find the minimum sub-array **A<sub>l</sub>, A<sub>l+1</sub> ,..., A<sub>r</sub>** such that if we sort(in ascending order) that sub-array, then the whole array should get sorted. If **A** is already sorted, output **-1.**

###### Problem Constraints

```
1 <= N <= 1000000
1 <= A[i] <= 1000000
```

###### input format

``` 
First and only argument is an array of non-negative integers of size N.
```

###### Output Format

```
Return an array of length two where the first element denotes the starting index(0-based) and the second element denotes the ending index(0-based) of the sub-array. If the array is already sorted, return an array containing only one element i.e. -1.
```

###### Example Input

```
# input 1: 
    A = [1, 3, 2, 4, 5]
# output 1: 
    [1, 2]
# input 2: 
    A = [1, 2, 3, 4, 5]
# output 2: 
    [-1]
```
