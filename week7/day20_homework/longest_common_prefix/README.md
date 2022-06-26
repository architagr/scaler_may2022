# Problem Description

Given the array of strings **A**, you need to find the longest string **S**, which is the prefix of **ALL** the strings in the array.

The longest common prefix for a pair of strings **S1** and **S2** is the longest string **S** which is the prefix of both **S1** and **S2**.

**Example:** the longest common prefix of "abcdefgh" and "abcefgh" is "abc".

###### Problem Constraints

```
0 <= sum of length of all strings <= 1000000
```

###### input format

``` 
The only argument given is an array of strings A.
```

###### Output Format

```
Return the longest common prefix of all strings in A.
```

###### Example Input

```
# input 1 : 
    A = ["abcdefgh", "aefghijk", "abcefgh"]
# output 1: 
    "a"
# input 2 : 
    A = ["abab", "ab", "abcd"]
# output 2: 
    "ab"
```
