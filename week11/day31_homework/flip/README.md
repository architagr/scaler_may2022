# Problem Description

You are given a binary string **A**(i.e., with characters **0** and **1**) consisting of characters **A1, A2, ..., AN.** In a single operation, you can choose two indices, **L** and **R**, such that **1 ≤ L ≤ R ≤ N** and flip the characters **AL, AL+1, ..., AR**. By flipping, we mean changing character **0** to **1** and vice-versa.

Your aim is to perform **ATMOST** one operation such that in the final string number of **1s** is maximized.

If you don't want to perform the operation, return an **empty** array. Else, return an array consisting of two elements denoting **L** and **R**. If there are multiple solutions, return the **lexicographically smallest** pair of **L** and **R**.

**NOTE:** Pair **(a, b)** is lexicographically smaller than pair **(c, d)** if **a < c** or, if **a == c** and **b < d.**

###### Problem Constraints

```
1 <= size of string <= 100000
```

###### input format

``` 
First and only argument is a string A.
```

###### Output Format

```
Return an array of integers denoting the answer.
```

###### Example Input

```
# input 1 : 
    A = "010"
# output 1: 
    [1, 1]
# input 2 : 
    A = "111"
# output 2: 
    []
```
