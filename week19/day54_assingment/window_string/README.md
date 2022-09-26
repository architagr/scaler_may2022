# Problem Description

Given a string A and a string B, find the window with minimum length in A, which will contain all the characters in B in linear time complexity.
Note that when the count of a character c in B is x, then the count of c in the minimum window in A should be at least x.

Note:

If there is no such window in A that covers all characters in B, return the empty string.
If there are multiple such windows, return the first occurring minimum window ( with minimum start index and length )

###### Problem Constraints

```
1 <= size(A), size(B) <= 10^6
```

###### input format

``` 
The first argument is a string A.
The second argument is a string B.
```

###### Output Format

```
Return a string denoting the minimum window.
```

###### Example Input

```
# input 1: 
    A = "ADOBECODEBANC"
    B = "ABC"
# output 1: 
    "BANC"
# input 2: 
    A = "Aa91b"
    B = "ab"
# output 2: 
    "a91b"
```
