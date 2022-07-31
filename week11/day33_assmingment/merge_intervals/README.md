# Problem Description

Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

###### Problem Constraints

```
0 <= |intervals| <= 10^5
```

###### input format

``` 
First argument is the vector of intervals

second argument is the new interval to be merged
```

###### Output Format

```
Return the vector of intervals after merging
```

###### Example Input

```
# input 1: 
    Given intervals [1, 3], [6, 9] insert and merge [2, 5] .
# output 1: 
   [ [1, 5], [6, 9] ]
# input 2: 
    Given intervals [1, 3], [6, 9] insert and merge [2, 6] .
# output 2: 
    [ [1, 9] ]
```
