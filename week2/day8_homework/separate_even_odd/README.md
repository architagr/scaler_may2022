# Problem Description

You are given an integer **T** denoting the number of test cases. For each test case, you are given an integer array **A**.

You have to put the odd and even elements of array **A** in 2 separate arrays and print the new arrays.

NOTE: Array elements should have the same relative order as in **A**.

###### Problem Constraints

```
1 <= T <= 10

1 <= |A| <= 10^5

1 <= A[i] <= 10^9
```

###### input format

``` 
First line of the input contains a single integer T.

For each test case:

First line consists of a single integer |A| denoting the length of array.
Second line consists of |A| space separated integers denoting the elements of array A.
```

###### Output Format

```
For each test case:

First line should contain an array of space separated integers containing all the odd elements of array A
Second line should contain an array of space separated integers containing all the even elements of array A

```

###### Example Input

```
# input 1 : 
    2 
    5
    1 2 3 4 5
    3
    4 3 2
# output 1: 
    1 3 5
    2 4
    3
    4 2
# input 2 : 
    2 
    3
    2 2 2
    2
    1 1
# output 2: 
        
    2 2 2
    1 1
     
```
