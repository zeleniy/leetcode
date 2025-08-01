<?php

namespace Zeleniy\Leetcode\LongestSubstringWithoutRepeatingCharacters;

class Solution {

    function lengthOfLongestSubstring($s) {

        $n = strlen($s);
        $charMap = [];
        $left = 0;
        $maxLen = 0;

        for ($right = 0; $right < $n; $right ++) {

            $char = $s[$right];

            if (isset($charMap[$char])) {
                $left = \max($left, $charMap[$char] + 1);
            }

            $charMap[$char] = $right;
            $maxLen = \max($maxLen, $right - $left + 1);
        }

        return $maxLen;
    }
}

