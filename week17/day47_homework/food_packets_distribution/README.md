# Problem Description

The government wants to set up B distribution offices across N cities for the distribution of food
packets. The population of the ith city is A[i]. Each city must have at least 1 office, and people can go to an office of their own city. We want to maximize the minimum number of people who can get food in any single office.

###### Problem Constraints

```
1 <= N <= 10^5

1 <= A[i] <= 10^6

1 <= B <= 5 x 10^5
```

###### input format

``` 
The first line of input contains an integer array A. 

The second line of input contains an integer B.
```

###### Output Format

```
Return one integer representing the maximum number of people who can get food in any single office.
```

###### Example Input

```
# input 1: 
    A = [10000, 20000, 30000]
    B = 6
# output 1: 
    1000
# input 2: 
    A = [1, 1, 1]
    B = 4
# output 2: 
    0
```
