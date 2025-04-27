<?php

namespace Zeleniy\Leetcode\ValidAnagram;

class Solution {

    function isAnagram($s, $t) {

        $frequencyMap = [];
        $strLength = \strlen($s);

        for ($i = 0; $i < $strLength; $i ++) {

            $char = $s[$i];

            if (isset($frequencyMap[$char])) {
                $frequencyMap[$char] ++;
            } else {
                $frequencyMap[$char] = 1;
            }
        }

        $strLength = \strlen($t);

        for ($i = 0; $i < $strLength; $i ++) {

            $char = $t[$i];

            if (! isset($frequencyMap[$char])) {
                return false;
            }

            $frequencyMap[$char] --;

            if ($frequencyMap[$char] < 0) {
                return false;
            }
        }

        return \array_sum($frequencyMap) == 0;
    }
}