Given a string of length n consisting of digits [0-9], count the number of ways the given string can be split into prime numbers.
Since the answer can be large, return the answer modulo 109 + 7. 
Note: A partition that contains numbers with leading zeroes will be invalid and the initial string does not contain leading zeroes. 

Example 1:
 the input string to be s = "3175", then this string can be split into 3 different ways as (31, 7,5), (3, 17, 5), (317,5) where each one of them contains only prime numbers.

Example 2:
s = "11373", then this string can be split into 6 different ways as [11, 37, 3), [113, 7, 3), [11, 3, 73), [11, 37, 3), (113, 73) and [11, 373) where each one of them contains only prime numbers.