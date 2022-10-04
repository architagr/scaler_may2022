# Problem Description

Given a linked list A, remove the B-th node from the end of the list and return its head.

For example, Given linked list: 1->2->3->4->5, and B = 2. After removing the second node from the end, the linked list becomes 1->2->3->5.

NOTE: If B is greater than the size of the list, remove the first node of the list.

NOTE: Try doing it using constant additional space.

###### Problem Constraints

```
1 <= |A| <= 10^6
```

###### input format

``` 
The first argument of input contains a pointer to the head of the linked list.

The second argument of input contains the integer B.
```

###### Output Format

```
Return the head of the linked list after deleting the B-th element from the end.
```

###### Example Input

```
# input 1: 
    A = [1, 2, 3, 4, 5]
    B = 2
# output 1: 
    [1, 2, 3, 5]
# input 2: 
    A = [1]
    B = 1
# output 2: 
    []
```
