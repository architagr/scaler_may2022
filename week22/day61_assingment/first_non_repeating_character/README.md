# Problem Description

Given a string A denoting a stream of lowercase alphabets, you have to make a new string B.
B is formed such that we have to find the first non-repeating character each time a character is inserted to the stream and append it at the end to B. If no non-repeating character is found, append '#' at the end of B.

###### Problem Constraints

```
1 <= |A| <= 100000
```

###### input format

``` 
The only argument given is string A.
```

###### Output Format

```
Return a string B after processing the stream of lowercase alphabets A.
```

###### Example Input

```
# input 1 : 
    A = "abadbc"
# output 1: 
    "aabbdd"
# input 2: 
    A = "abcabc"
# output 2: 
    "aaabc#"
```
