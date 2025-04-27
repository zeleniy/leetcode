<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

class Solution {

    function reverseList($head) {

        if ($head === null || $head->next === null) {
            return $head;
        }

        $prev = null;
        $curr = $head;

        while ($curr !=- null) {

            $next       = $curr->next;
            $curr->next = $prev;
            $prev       = $curr;
            $curr       = $next;
        }

        return $prev;
    }
}
