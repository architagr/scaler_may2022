# Problem Description

Given two BST's A and B, return the (sum of all common nodes in both A and B) % (10<sup>9</sup> +7) .

In case there is no common node, return 0.

NOTE:

Try to do it one pass through the trees.

###### Problem Constraints

```
1 <= Number of nodes in the tree A and B <= 10^5

1 <= Node values <= 10^6
```

###### input format

``` 
First argument represents the root of BST A.

Second argument represents the root of BST B.
```

###### Output Format

```
Return an integer denoting the (sum of all common nodes in both BST's A and B) % (10^9 +7) .
```

###### Example Input

```
# input 1 : 
    Tree A:
    5
   / \
  2   8
   \   \
    3   15
        /
        9

 Tree B:
    7
   / \
  1  10
   \   \
    2  15
       /
      11
# output 1: 
    17
# input 2: 
    Tree A:
    7
   / \
  1   10
   \   \
    2   15
        /
       11

 Tree B:
    7
   / \
  1  10
   \   \
    2  15
       /
      11
# output 2: 
      46
```
