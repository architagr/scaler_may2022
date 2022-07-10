# Problem Description

Given an expression string A, examine whether the pairs and the orders of “{“,”}”, ”(“,”)”, ”[“,”]” are correct in A.

Refer to the examples for more clarity.

###### Problem Constraints

```
1 <= |A| <= 100
```

###### input format

``` 
The first and the only argument of input contains the string A having the parenthesis sequence.
```

###### Output Format

```
Return 0 if the parenthesis sequence is not balanced.

Return 1 if the parenthesis sequence is balanced.
```

###### Example Input

```
# input 1 : 
     A = {([])}
# output 1: 
    1
# input 2 : 
     A = {([)}
# output 2: 
    0
```
