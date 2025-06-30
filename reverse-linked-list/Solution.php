<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

class Solution {

    public function reverseListRecursive($head) {

        if (null == $head || null == $head->next) {
            return $head;
        }

        $newHead = $this->reverseListRecursive($head->next);

        $head->next->next = $head;
        $head->next = null;

        return $newHead;
    }

    public function reverseListIterative($head) {

        $prev = null;
        $curr = $head;

        while ($curr) {
            $next       = $curr->next;
            $curr->next = $prev;
            $prev       = $curr;
            $curr       = $next;
        }

        return $prev;
    }
}
