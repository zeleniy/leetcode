<?php

namespace Zeleniy\Leetcode\ContainsDuplicate;

class Solution {

    function containsDuplicate($nums) {

        $frequencyMap = [];

        foreach ($nums as $num) {
            if (isset($frequencyMap[$num])) {
                return true;
            } else {
                $frequencyMap[$num] = 1;
            }
        }

        return false;
    }
}