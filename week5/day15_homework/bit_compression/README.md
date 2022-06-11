# Problem Description

Richard Hendricks, a mastermind in compression algorithms, is an employee of Hooli in Silicon Valley.
One day, he finally decided to quit and work on his new idea of the middle-out compression algorithm.

He needed to work at the bit - level to compress data. He, eventually, encountered this problem.
There is an array **A** of **N** integers. He has to perform certain operations on the elements.
In any one operation, two indices **i** and **j** (i < j) are chosen, and **A[i]** is replaced with **A[i] & A[j]**,
and **A[j]** is replaced with **A[i] | A[j]**, where **&** represents the Bitwise AND operation and **|** represents the Bitwise OR operation.
This operation is performed over all the pairs of integers in the array.

Help Richard find the Bitwise XOR of all the elements after performing the operations.

###### Problem Constraints

```
1 <= N <= 10^5
1 <= Ai <= 10^9
```

###### input format

``` 
The first argument is the integer array A.
```

###### Output Format

```
Return a single integer denoting the XOR of the elements after performing the operations.
```

###### Example Input

```
# input 1 : 
    A = [1, 3, 5]
# output 1: 
    7
# input 2 : 
    A = [1, 1, 1]
# output 2: 
    1
```
