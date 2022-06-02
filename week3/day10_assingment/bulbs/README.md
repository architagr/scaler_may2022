# Problem Description

A wire connects **N** light bulbs.

Each bulb has a switch associated with it; however, due to faulty wiring, a switch also changes the state of all the bulbs to the right of the current bulb.

Given an initial state of all bulbs, find the **minimum number** of switches you have to press to turn on all the bulbs.

You can press the same switch multiple times.

**Note: 0** represents the bulb is off and **1** represents the bulb is on.

###### Problem Constraints

```
1 <= N <= 5×10^5
0 <= A[i] <= 1
```

###### input format

``` 
The first and the only argument contains an integer array A, of size N.
```

###### Output Format

```
Return an integer representing the minimum number of switches required.
```

###### Example Input

```
# input 1 : 
    A = [0, 1, 0, 1]
# output 1: 
    4
# input 2 : 
    A = [1, 1, 1, 1]
# output 2: 
    0
```
