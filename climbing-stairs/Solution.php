<?php

namespace Zeleniy\Leetcode\ClimbingStairs;

class Solution {

    function climbStairs($n) {

        if ($n <= 2) {
            return $n;
        }

        $a = 1;
        $b = 2;

        for ($i = 3; $i <= $n; $i ++) {

            $c = $a + $b;
            $a = $b;
            $b = $c;
        }

        return $b;
    }
}