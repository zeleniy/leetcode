<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

class Solution {

    function reverseList($head) {

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
