# Problem Description

You are given a function to_lower() consisting of a character array **A.**

Convert each character of A into lowercase characters if it exists. If the lowercase of a character does not exist, it remains unmodified.
The uppercase letters from A to Z are converted to lowercase letters from a to z respectively.

Return the **lowercase ** version of the given character array.

###### Problem Constraints

```
1 <= |A| <= 10^5
```

###### input format

``` 
Only argument is a character array A.
```

###### Output Format

```
Return the uppercase version of the given character array.
```

###### Example Input

```
# input 1 : 
    A = ['S', 'c', 'A', 'L', 'E', 'r', 'A', 'c', 'a', 'D', 'e', 'm', 'y']
# output 1: 
    ['s', 'c', 'a', 'l', 'e', 'r', 'a', 'c', 'a', 'd', 'e', 'm', 'y']
# input 2 : 
    A = ['S', 'c', 'a', 'L', 'e', 'R', '#', '2', '0', '2', '0']
# output 2: 
    ['s', 'c', 'a', 'l', 'e', 'r', '#', '2', '0', '2', '0']
```
