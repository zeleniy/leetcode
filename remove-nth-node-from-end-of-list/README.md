## [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.

### Example 1:
<img src="img/remove_ex1.png" width="542px"/><br/>

**Input**: `head = [1,2,3,4,5]`, `n = 2`<br/>
**Output**: `[1,2,3,5]`<br/>

### Example 2:

**Input**: `head = [1]`, `n = 1`<br/>
**Output**: `[]`<br/>

### Example 3:

**Input**: `head = [1,2]`, `n = 1`<br/>
**Output**: `[1]`<br/>

### Constraints:

* The number of nodes in the list is `sz`.
* `1 <= sz <= 30`
* `0 <= Node.val <= 100`
* `1 <= n <= sz`

### Follow-up:

Could you do this in one pass?
