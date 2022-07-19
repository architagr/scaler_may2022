# Problem Description

Given a Binary Tree **A** containing **N** nodes, you need to find the path from **Root** to a given node **B**.

**NOTE:**

No two nodes in the tree have the same data values.
You can assume that **B** is present in tree **A** and a path always exists.

###### Problem Constraints

```
1 <= N <= 10^5

1 <= Data Values of Each Node <= N

1 <= B <= N
```

###### input format

``` 
First Argument represents pointer to the root of binary tree A.

Second Argument is an integer B denoting the node number.
```

###### Output Format

```
Return an one-dimensional array denoting the path from Root to the node B in order.
```

###### Example Input

```
# input 1 : 
 A =     
           1
         /   \
        2     3
       / \   / \
      4   5 6   7 

 B = 5 
# output 1: 
  [1, 2, 5]
# input 2 : 
 A = 
            1
          /   \
         2     3
        / \ .   \
       4   5 .   6

 B = 1      
# output 2: 
  [1]
```
