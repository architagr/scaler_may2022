# Problem Description

You are given an integer array A denoting the Level Order Traversal of the Binary Tree.

You have to Deserialize the given Traversal in the Binary Tree and return the root of the Binary Tree.

NOTE:

In the array, the NULL/None child is denoted by -1.
For more clarification check the Example Input.


###### Problem Constraints

```
1 <= number of nodes <= 10^5

-1 <= A[i] <= 10^5
```

###### input format

``` 
Only argument is an integer array A denoting the Level Order Traversal of the Binary Tree.
```

###### Output Format

```
Return the root node of the Binary Tree.
```

###### Example Input

```
# input 1 : 
 A = [1, 2, 3, 4, 5, -1, -1, -1, -1, -1, -1]
# output 1: 
           1
         /   \
        2     3
       / \
      4   5
# input 2: 
A = [1, 2, 3, 4, 5, -1, 6, -1, -1, -1, -1, -1, -1]
# output 2: 
            1
          /   \
         2     3
        / \ .   \
       4   5 .   6
```
