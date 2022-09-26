# Problem Description

Design and implement a Linked List data structure.
A node in a linked list should have the following attributes - an integer **value** and a **pointer** to the next node. It should support the following operations:

insert_node(position, value) - To insert the input value at the given position in the linked list.
delete_node(position) - Delete the value at the given position from the linked list.
print_ll() - Print the entire linked list, such that each element is followed by a single space.

**Note:**

1. If an input position does not satisfy the constraint, no action is required.
2. Each print query has to be executed in a new line.

###### Problem Constraints

```
1 <= position <= n where, n is the size of the linked-list.
```

###### input format

``` 
First line contains an integer denoting number of cases, let's say t.
Next t line denotes the cases.
```

###### Output Format

```
When there is a case of print_ll(),  Print the entire linked list, such that each element is followed by a single space.
NOTE: You don't need to return anything.
```

###### Example Input

```
# input 1 : 
    5
    i 1 23
    i 2 24
    p
    d 1
    p
# output 1: 
    23 24
    24
```
