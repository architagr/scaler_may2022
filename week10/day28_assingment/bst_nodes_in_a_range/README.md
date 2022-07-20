# Problem Description

Given a binary search tree of integers. You are given a range B and C.

Return the count of the number of nodes that lie in the given range.

###### Problem Constraints

```
1 <= Number of nodes in binary tree <= 100000

0 <= B < = C <= 10^9
```

###### input format

``` 
First argument is a root node of the binary tree, A.

Second argument is an integer B.

Third argument is an integer C.
```

###### Output Format

```
Return the count of the number of nodes that lies in the given range.
```

###### Example Input

```
# input 1 : 
            15
          /    \
        12      20
        / \    /  \
       10  14  16  27
      /
     8

     B = 12
     C = 20
# output 1: 
    5
# input 2 : 
            8
           / \
          6  21
         / \
        1   7

     B = 2
     C = 20  
# output 2: 
    3
```
