# Problem Description

Shaggy has an array **A** consisting of N elements. We call a pair of distinct indices in that array a special if elements at those indices in the array are equal.

Shaggy wants you to find a special pair such that the distance between that pair is minimum. Distance between two indices is defined as |i-j|. If there is no special pair in the array, then return -1.

###### Problem Constraints

```
1 <= |A| <= 10^5
```

###### input format

``` 
The first and only argument is an integer array A.
```

###### Output Format

```
Return one integer corresponding to the minimum possible distance between a special pair.
```

###### Example Input

```
# input 1 : 
    A = [7, 1, 3, 4, 1, 7]
# output 1: 
    3
# input 2 : 
    A = [1, 1]
# output 2: 
    1
```
