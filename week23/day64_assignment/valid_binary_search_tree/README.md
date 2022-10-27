# Problem Description

You are given a binary tree represented by root A.

Assume a BST is defined as follows:

1) The left subtree of a node contains only nodes with keys less than the node's key.

2) The right subtree of a node contains only nodes with keys greater than the node's key.

3) Both the left and right subtrees must also be binary search trees.

###### Problem Constraints

```
1 <= Number of nodes in binary tree <= 10^5

0 <= node values <= (2^32)-1
```

###### input format

``` 
First and only argument is head of the binary tree A.
```

###### Output Format

```
Return 0 if false and 1 if true.
```

###### Example Input

```
# input 1 : 
    1
  /  \
 2    3
# output 1: 
     0
# input 2: 
  2
 / \
1   3
# output 2: 
      1
```
