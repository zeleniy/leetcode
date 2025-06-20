## [Duplicate Zeros](https://leetcode.com/problems/duplicate-zeros/)

Given a fixed-length integer array `arr`, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

### Example 1:

**Input**: `arr = [1,0,2,3,0,4,5,0]`<br />
**Output**: `[1,0,0,2,3,0,0,4]`<br />
**Explanation**: After calling your function, the input array is modified to: `[1,0,0,2,3,0,0,4]`

### Example 2:

**Input**: `arr = [1,2,3]`<br />
**Output**: `[1,2,3]`<br />
**Explanation**: After calling your function, the input array is modified to: `[1,2,3]`

### Constraints:

* `1 <= arr.length <= 10⁴`
* `0 <= arr[i] <= 9`
