# Problem Description

We have a list **A** of points **(x, y)** on the plane. Find the **B** closest points to the origin **(0, 0)**.

Here, the distance between two points on a plane is the **Euclidean distance**.

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in.)

**NOTE:** Euclidean distance between two points **P1(x1, y1)** and **P2(x2, y2)** is **sqrt( (x1-x2)<sup>2</sup> + (y1-y2)<sup>2</sup> ).**

###### Problem Constraints

```
1 <= B <= length of the list A <= 10^5
-10^5 <= A[i][0] <= 10^5
-10^5 <= A[i][1] <= 10^5
```

###### input format

``` 
The argument given is list A and an integer B.
```

###### Output Format

```
Return the B closest points to the origin (0, 0) in any order.
```

###### Example Input

```
# input 1: 
     A = [ 
       [1, 3],
       [-2, 2] 
     ]
    B = 1
# output 1: 
    [ [-2, 2] ]
# input 2: 
    A = [
       [1, -1],
       [2, -1]
     ] 
    B = 1
# output 2: 
    [ [1, -11] ]
```
