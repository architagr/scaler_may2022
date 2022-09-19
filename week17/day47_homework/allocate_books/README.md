# Problem Description

Given an array of integers A of size N and an integer B.

The College library has N books. The ith book has A[i] number of pages.

You have to allocate books to B number of students so that the maximum number of pages allocated to a student is minimum.

A book will be allocated to exactly one student.
Each student has to be allocated at least one book.
Allotment should be in contiguous order, for example: A student cannot be allocated book 1 and book 3, skipping book 2.
Calculate and return that minimum possible number.

NOTE: Return -1 if a valid assignment is not possible.

###### Problem Constraints

```
1 <= N <= 10^5
1 <= A[i], B <= 10^5
```

###### input format

``` 
The first argument given is the integer array A.
The second argument given is the integer B.
```

###### Output Format

```
Return that minimum possible number.
```

###### Example Input

```
# input 1: 
    A = [12, 34, 67, 90]
    B = 2
# output 1: 
    113
```
