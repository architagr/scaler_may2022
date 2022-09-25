# Problem Description

You are given a string **A** of length **N** consisting of lowercase alphabets. Find the **period** of the string.

Period of the string is the **minimum** value of **k (k >= 1)**, that satisfies **A[i] = A[i % k]** for all valid i.

###### Problem Constraints

```
1 <= N <= 10^6
```

###### input format

``` 
First and only argument is a string A of length N.
```

###### Output Format

```
Return an integer, denoting the period of the string.
```

###### Example Input

```
# input 1: 
    A = "abababab"
# output 1: 
    2
# input 2: 
    A = "aaaa"
# output 2: 
    1
```
