<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

class ListNode {

    public $val  = 0;
    public $next = null;

    function __construct($val = 0, $next = null) {

        $this->val  = $val;
        $this->next = $next;
    }
}