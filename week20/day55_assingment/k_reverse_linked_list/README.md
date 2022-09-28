# Problem Description

Given a singly linked list A and an integer B, reverse the nodes of the list B at a time and return the modified linked list.

###### Problem Constraints

```
1 <= |A| <= 10^3

B always divides A
```

###### input format

``` 
The first argument of input contains a pointer to the head of the linked list.

The second arugment of input contains the integer, B.
```

###### Output Format

```
Return a pointer to the head of the modified linked list.
```

###### Example Input

```
# input 1 : 
    A = [1, 2, 3, 4, 5, 6]
    B = 2
# output 1: 
    [2, 1, 4, 3, 6, 5]
# input 2: 
    A = [1, 2, 3, 4, 5, 6]
    B = 3
# output 2: 
    [3, 2, 1, 6, 5, 4]
```
