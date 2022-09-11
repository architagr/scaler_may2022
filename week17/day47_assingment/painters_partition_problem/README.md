# Problem Description

Given 2 integers **A** and **B** and an array of integers **C** of size **N**. Element **C[i]** represents the length of **ith** board.
You have to paint all **N** boards **[C<sub>0</sub>, C<sub>1</sub>, C<sub>2</sub>, C<sub>3</sub> … C<sub>N-1</sub>]**. There are **A** painters available and each of them takes **B** units of time to paint **1 unit** of the board.

Calculate and return the minimum time required to paint all boards under the constraints that **any painter will only paint contiguous sections of the board.**
**NOTE:**
1. 2 painters cannot share a board to paint. That is to say, a board cannot be painted partially by one painter, and partially by another.
2. A painter will only paint contiguous boards. This means a configuration where painter 1 paints boards 1 and 3 but not 2 is invalid.

Return the **ans % 10000003.**

###### Problem Constraints

```
1 <= A <= 1000
1 <= B <= 10^6
1 <= N <= 10^5
1 <= C[i] <= 10^6
```

###### input format

``` 
The first argument given is the integer A.
The second argument given is the integer B.
The third argument given is the integer array C.
```

###### Output Format

```
Return minimum time required to paint all boards under the constraints that any painter will only paint contiguous sections of board % 10000003.
```

###### Example Input

```
# input 1: 
    A = 2
    B = 5
    C = [1, 10]
# output 1: 
    50
# input 2: 
    A = 10
    B = 1
    C = [1, 8, 11, 3]
# output 2: 
    11
```
