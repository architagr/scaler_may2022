# Problem Description

Given the root of a tree A with each node having a certain value, find the count of nodes with more value than all its ancestor.

###### Problem Constraints

```
1 <= Number of Nodes <= 200000

1 <= Value of Nodes <= 2000000000
```

###### input format

``` 
The first and only argument of input is a tree node.
```

###### Output Format

```
Return a single integer denoting the count of nodes that have more value than all of its ancestors.
```

###### Example Input

```
# input 1 : 
  3 
# output 1: 
    1
# input 2 : 
    4
   / \
  5   2
     / \
    3   6
# output 2: 
    3
```
