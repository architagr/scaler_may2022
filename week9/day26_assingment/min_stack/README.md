# Problem Description

Design a stack that supports **push, pop, top,** and **retrieve the minimum element** in constant time.
**push(x)** -- Push element x onto stack.
**pop()** -- Removes the element on top of the stack.
**top()** -- Get the top element.
**getMin()** -- Retrieve the minimum element in the stack.
**NOTE:**
All the operations have to be constant time operations.
getMin() should return **-1** if the stack is empty.
pop() should return **nothing** if the stack is empty.
top() should return **-1** if the stack is empty.

###### Problem Constraints

```
1 <= Number of Function calls <= 10^7
```

###### input format

``` 
Functions will be called by the checker code automatically.
```

###### Output Format

```
Each function should return the values as defined by the problem statement.
```

###### Example Input

```
# input 1 : 
    push(1)
    push(2)
    push(-2)
    getMin()
    pop()
    getMin()
    top()
# output 1: 
    -2 1 2
# input 2 : 
    getMin()
    pop()
    top()
# output 2: 
    -1 -1
```
