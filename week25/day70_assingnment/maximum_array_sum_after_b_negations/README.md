# Problem Description

Given an array of integers A and an integer B. You must modify the array exactly B number of times. In a single modification, we can replace any one array element A[i] by -A[i].

You need to perform these modifications in such a way that after exactly B modifications, sum of the array must be maximum.

###### Problem Constraints

```
1 <= length of the array <= 5*10^5
1 <= B <= 5 * 10^6
-100 <= A[i] <= 100
```

###### input format

``` 
The first argument given is an integer array A.
The second argument given is an integer B.
```

###### Output Format

```
Return an integer denoting the maximum array sum after B modifications.
```

###### Example Input

```
# input 1 : 
  A = [24, -68, -29, -9, 84]
  B = 4
# output 1: 
   196
# input 2: 
  A = [57, 3, -14, -87, 42, 38, 31, -7, -28, -61]
  B = 10
# output 2: 
  362
```
