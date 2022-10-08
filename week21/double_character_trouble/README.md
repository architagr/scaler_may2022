# Problem Description

You are given a string A.

An operation on the string is defined as follows:

Remove the first occurrence of the same consecutive characters. eg for a string "abbcd", the first occurrence of same consecutive characters is "bb".

Therefore the string after this operation will be "acd".

Keep performing this operation on the string until there are no more occurrences of the same consecutive characters and return the final string.

###### Problem Constraints

```
1 <= |A| <= 100000
```

###### input format

``` 
First and only argument is string A.
```

###### Output Format

```
Return the final string.
```

###### Example Input

```
# input 1 : 
    A = abccbc
# output 1: 
    "ac"
# input 2: 
    A = ab
# output 2: 
    "ab"
```
