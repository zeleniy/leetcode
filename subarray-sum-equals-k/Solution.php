<?php

namespace Zeleniy\Leetcode\SubarraySumEqualsK;

class Solution {

    function subarraySum($nums, $k) {

        $count       = 0;
        $current_sum = 0;
        $prefix_sums = [0 => 1];

        foreach ($nums as $num) {

            $current_sum += $num;

            if (isset($prefix_sums[$current_sum - $k])) {
                $count += $prefix_sums[$current_sum - $k];
            }

            if (isset($prefix_sums[$current_sum])) {
                $prefix_sums[$current_sum]++;
            } else {
                $prefix_sums[$current_sum] = 1;
            }
        }

        return $count;
    }
}
