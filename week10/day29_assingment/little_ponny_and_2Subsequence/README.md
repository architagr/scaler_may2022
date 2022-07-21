# Problem Description

Little Ponny has been given a string **A**, and he wants to find out the **lexicographically minimum subsequence** from it of **size >= 2**. Can you help him?

A string **a** is lexicographically smaller than string **b**, if the first different letter in **a** and **b** is smaller in **a**. For example, "abc" is lexicographically smaller than "acc" because the first different letter is 'b' and 'c' which is smaller in "abc".

###### Problem Constraints

```
1 <= |A| <= 10^5

A only contains lowercase alphabets.
```

###### input format

``` 
The first and the only argument of input contains the string, A.
```

###### Output Format

```
Return a string representing the answer.
```

###### Example Input

```
# input 1 : 
  A = "abcdsfhjagj" 
# output 1: 
  "aa" 
# input 2 : 
  A = "ksdjgha"    
# output 2: 
  "da"
```
