<?php

namespace Zeleniy\Leetcode\FirstUniqueCharacterInString;

class Solution {

    function firstUniqChar($s) {

        $length   = \strlen($s);
        $freqMap  = [];
        $indexMap = [];

        for ($i = 0; $i < $length; $i ++) {

            $char = $s[$i];

            if (isset($freqMap[$char])) {
                $freqMap[$char] ++;
                continue;
            } else {
                $freqMap[$char] = 1;
                $indexMap[$char] = $i;
            }
        }

        foreach ($freqMap as $char => $freq) {
            if ($freq == 1) {
                return $indexMap[$char];
            }
        }

        return -1;
    }
}