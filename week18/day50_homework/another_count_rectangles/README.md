# Problem Description

Given a sorted array of distinct integers A and an integer B, find and return how many rectangles with distinct configurations can be created using elements of this array as length and breadth whose area is lesser than B.

(Note that a rectangle of 2 x 3 is different from 3 x 2 if we take configuration into view)

###### Problem Constraints

```
1 <= |A| <= 100000
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
Return the number of rectangles with distinct configurations with area less than B modulo (109 + 7).
```

###### Example Input

```
# input 1: 
    A = [1, 2]
    B = 5
# output 1: 
    4
# input 2: 
    A = [1, 2]
    B = 1
# output 2: 
    0
```
