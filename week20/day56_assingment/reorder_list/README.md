# Problem Description

Given a singly linked list A

 A: A0 → A1 → … → An-1 → An 
reorder it to:

 A0 → An → A1 → An-1 → A2 → An-2 → … 
You must do this in-place without altering the nodes' values.

###### Problem Constraints

```
1 <= |A| <= 10^6
```

###### input format

``` 
The first and the only argument of input contains a pointer to the head of the linked list A.
```

###### Output Format

```
Return a pointer to the head of the modified linked list.
```

###### Example Input

```
# input 1 : 
    A = [1, 2, 3, 4, 5]  
# output 1: 
    [1, 5, 2, 4, 3] 
```
