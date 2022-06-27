# Problem Description

Groot has N trees lined up in front of him where the height of the i'th tree is denoted by H[i].

He wants to select some trees to replace his broken branches.

But he wants uniformity in his selection of trees.

So he picks only those trees whose heights have frequency K.

He then sums up the heights that occur K times. (He adds the height only once to the sum and not K times).

But the sum he ended up getting was huge so he prints it modulo 10^9+7.

In case no such cluster exists, Groot becomes sad and prints -1.

###### Problem Constraints

```
1.   1<=N<=100000
2.   1<=K<=N
3.   0<=H[i]<=10^9
```

###### input format

``` 
 Integers N and K and array of size A
```

###### Output Format

```
 Sum of all numbers in the array that occur exactly K times.
```

###### Example Input

```
# input 1 : 
    N=5 ,K=2 ,A=[1 2 2 3 3]
# output 1: 
    5
# input 2 : 
    N=6 ,K=2 ,A=[1000000000, 1000000000, 999999999, 999999999, 999999998, 1]
# output 2: 
    999999992
```
