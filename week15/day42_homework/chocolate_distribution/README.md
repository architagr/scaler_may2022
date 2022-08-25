# Problem Description

Given an array **A** of **N** integers where the i-th element represent the number of chocolates in the i-th packet.

There are **B** number of students, the task is to distribute chocolate packets following below conditions:

1. Each student gets one packet.
2. The difference between the number of chocolates given to any two students is minimum.
Return the minimum difference (that can be achieved) between the student who gets minimum number of chocolates and the student who gets maximum number of chocolates.

###### Problem Constraints

```
0 <= N <= 10^5
1 <= A[i] <= 10^7
0 <= B <= 10^5
```

###### input format

``` 
The first argument contains an integer array A.

The second argument contains an integer B.
```

###### Output Format

```
Return the minimum difference (that can be achieved) between maximum and minimum number of chocolates distributed.
```

###### Example Input

```
# input 1: 
    A : [3, 4, 1, 9, 56, 7, 9, 12]
    B : 5
# output 1: 
    6

```
