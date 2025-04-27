<?php

namespace Zeleniy\Leetcode\MergeIntervals;

class Solution {

    function merge($intervals) {

        \usort($intervals, function($a, $b) {
            return $a[0] - $b[0];
        });

        $n      = \count($intervals);
        $result = [];

        for ($i = 0; $i < $n; $i = $j) {

            $current = $intervals[$i];
            $j = $i + 1;

            while ($j < $n && $current[1] >= $intervals[$j][0]) {
                $current[1] = \max($current[1], $intervals[$j][1]);
                $j ++;
            }

            $result[] = $current;
        }

        return $result;
    }
}