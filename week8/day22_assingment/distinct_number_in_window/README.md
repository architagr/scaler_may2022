# Problem Description

You are given an array of **N** integers, **A1, A2 ,..., AN** and an integer **B**. Return the of count of distinct numbers in all windows of size **B**.

Formally, return an array of size **N-B+1** where **i'th** element in this array contains number of distinct elements in sequence **Ai, Ai+1 ,..., Ai+B-1.**

**NOTE:** if **B > N**, return an empty array.


###### input format

``` 
First argument is an integer array A
Second argument is an integer B.
```

###### Output Format

```
Return an integer array.
```

###### Example Input

```
# input 1 : 
    A = [1, 2, 1, 3, 4, 3]
    B = 3
# output 1: 
    [2, 3, 3, 2]
# input 2 : 
    A = [1, 1, 2, 2]
    B = 1
# output 2: 
    [1, 1, 1, 1]
```
