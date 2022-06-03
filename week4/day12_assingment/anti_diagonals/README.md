# Problem Description

Give a **N * N** square matrix **A**, return an array of its anti-diagonals. Look at the example for more details.

###### Problem Constraints

```
1<= N <= 1000
1<= A[i][j] <= 1e9
```

###### input format

``` 
First argument is an integer N, denoting the size of square 2D matrix.
Second argument is a 2D array A of size N * N.
```

###### Output Format

```
Return a 2D integer array of size (2 * N-1) * N, representing the anti-diagonals of input array A.
The vacant spaces in the grid should be assigned to 0.
```

###### Example Input

```
# input 1 : 
3
1 2 3
4 5 6
7 8 9
# output 1: 
1 0 0
2 4 0
3 5 7
6 8 0
9 0 0

# input 2 : 
2
1 2
3 4
# output 2: 
1 0
2 3
4 0
```
