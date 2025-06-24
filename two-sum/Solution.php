<?php

namespace Zeleniy\Leetcode\TwoSum;

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {

        $map = [];

        foreach ($nums as $i => $num) {

            if (isset($map[$target - $num])) {
                return [$map[$target - $num], $i];
            }

            $map[$num] = $i;
        }
    }
}
