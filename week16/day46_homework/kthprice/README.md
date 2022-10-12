# Problem Description

Given the price list at which tickets for a flight were purchased, figure out the kth smallest price for the flight. kth smallest price is the minimum possible n such that there are at least k price elements in the price list with value <= n. In other words, if the price list was sorted, then A[k-1] ( k is 1 based, while the array is 0 based ).

NOTE You are not allowed to modify the price list ( The price list is read only ). Try to do it using constant extra space.

###### Example Input

```
# input 1: 
    A : [2 1 4 3 2]
    k : 3
# output 1: 
    2
```
