<?php

namespace Zeleniy\Leetcode\LongestSubstringWithoutRepeatingCharacters;

class Solution {

    function lengthOfLongestSubstring($s) {

        $n = strlen($s);

        if ($n < 2) {
            return $n;
        }

        $i = [];
        $l = 0;
        $m = 0;

        for ($r = 0; $r < $n; $r ++) {

            $c = $s[$r];

            if (isset($i[$c])) {
                $l = \max($l, $i[$c] + 1);
            }

            $i[$c] = $r;
            $m = \max($m, $r - $l + 1);
        }

        return $m;
    }
}