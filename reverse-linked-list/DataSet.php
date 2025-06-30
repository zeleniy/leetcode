<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

class DataSet {

    public static function getData() {

        return [
            [[5, 4, 3, 2, 1], new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5)))))],
            [[2, 1], new ListNode(1, new ListNode(2))],
            [[], null]
        ];
    }
}
