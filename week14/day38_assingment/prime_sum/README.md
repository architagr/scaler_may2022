# Problem Description

Given an even number **A** ( greater than 2 ), return two prime numbers whose sum will be equal to the given number.

If there is more than one solution possible, return the lexicographically smaller solution.
If [a, b] is one solution with a <= b, and [c,d] is another solution with c <= d, then 
[a, b] < [c, d], If a < c OR a==c AND b < d.
**NOTE:** A solution will always exist. Read Goldbach's conjecture.

###### Problem Constraints

```
4 <= A <= 2*107
```

###### input format

``` 
First and only argument of input is an even number A.
```

###### Output Format

```
Return a integer array of size 2 containing primes whose sum will be equal to given number.
```

###### Example Input

```
# input 1: 
    4
# output 1: 
    [2, 2]
# input 2: 
    10
# output 2: 
    [3, 7]
```
