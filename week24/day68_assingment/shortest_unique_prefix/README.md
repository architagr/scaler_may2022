# Problem Description

Given a list of N words, find the shortest unique prefix to represent each word in the list.

NOTE: Assume that no word is the prefix of another. In other words, the representation is always possible


###### Problem Constraints

```
1 <= Sum of length of all words <= 10^6
```

###### input format

``` 
First and only argument is a string array of size N.
```

###### Output Format

```
Return a string array B where B[i] denotes the shortest unique prefix to represent the ith word.
```

###### Example Input

```
# input 1 : 
  A = ["zebra", "dog", "duck", "dove"]
# output 1: 
  ["z", "dog", "du", "dov"]
# input 2: 
A = ["apple", "ball", "cat"]
# output 2: 
  ["a", "b", "c"]
```
