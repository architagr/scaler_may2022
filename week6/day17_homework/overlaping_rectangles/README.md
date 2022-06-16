# Problem Description

Eight integers **A, B, C, D, E, F, G, and H** represent two rectangles in a **2D plane**s.
For the first rectangle, its bottom left corner is (A, B), and the top right corner is (C, D), and for the second rectangle, its bottom left corner is (E, F), and the top right corner is (G, H).

Find and return whether the two rectangles overlap or not.

###### Problem Constraints

```
-10000 <= A < C <= 10000

-10000 <= B < D <= 10000

-10000 <= E < G <= 10000

-10000 <= F < H <= 10000
```

###### input format

``` 
The eight arguments are integers A, B, C, D, E, F, G, and H.
```

###### Output Format

```
Return 1 if the two rectangles overlap else, return 0.
```

###### Example Input

```
# input 1 : 
    A = 0   B = 0
    C = 4   D = 4
    E = 2   F = 2
    G = 6   H = 6
# output 1: 
    1
# input 2 : 
    A = 0   B = 0
    C = 1   D = 1
    E = 1   F = 1
    G = 6   H = 6
# output 2: 
    0
# input 3 : 
    A = 60   B = 65
    C = 99   D = 84
    E = 85   F = 24
    G = 99   H = 84
# output 3: 
    1
```
