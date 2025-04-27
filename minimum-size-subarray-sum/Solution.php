<?php

namespace Zeleniy\Leetcode\MinimumSizeSubarraySum;

class Solution {

    function minSubArrayLen($target, $nums) {

        $length    = \count($nums);
        $minLength = \PHP_INT_MAX;

        for ($left = 0, $right = 0, $sum = 0; $right < $length; $right ++) {

            $sum += $nums[$right];

            for (; $sum >= $target; $left ++) {

                $minLength  = \min($minLength, $right - $left + 1);
                $sum       -= $nums[$left];
            }
        }

        return $minLength == \PHP_INT_MAX ? 0 : $minLength;
    }
}