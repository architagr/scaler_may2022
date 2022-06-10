# Problem Description

You have an array **A** with **N** elements. We have **two** types of operation available on this array :
1. We can split an element **B** into two elements, **C** and **D**, such that B = C + D.
2. We can merge two elements, **P** and **Q**, to one element, **R**, such that R = P ^ Q i.e., XOR of P and Q.
You have to determine whether it is possible to convert array A to size 1, **containing a single element equal to 0** after several splits and/or merge?

###### Problem Constraints

```
1 <= N <= 100000

1 <= A[i] <= 10^6
```

###### input format

``` 
The first argument is an integer array A of size N.
```

###### Output Format

```
Return "Yes" if it is possible otherwise return "No".
```

###### Example Input

```
# input 1 : 
    A = [9, 17]
# output 1: 
    Yes
# input 2 : 
    A = [1]
# output 2: 
    No
```
