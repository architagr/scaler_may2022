# Problem Description

Given the root node of a Binary Tree denoted by A. You have to Serialize the given Binary Tree in the described format.

Serialize means encode it into a integer array denoting the Level Order Traversal of the given Binary Tree.

NOTE:

In the array, the NULL/None child is denoted by -1.
For more clarification check the Example Input.

###### Problem Constraints

```
1 <= number of nodes <= 10^5
```

###### input format

``` 
Only argument is a A denoting the root node of a Binary Tree.
```

###### Output Format

```
Return an integer array denoting the Level Order Traversal of the given Binary Tree.
```

###### Example Input

```
# input 1 : 
           1
         /   \
        2     3
       / \
      4   5
# output 1: 
    [1, 2, 3, 4, 5, -1, -1, -1, -1, -1, -1]
# input 2: 
            1
          /   \
         2     3
        / \     \
       4   5     6
# output 2: 
    [1, 2, 3, 4, 5, -1, 6, -1, -1, -1, -1, -1, -1]
```
