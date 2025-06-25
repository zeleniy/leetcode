<?php

namespace Zeleniy\Leetcode\ValidAnagram;

class Solution {

    function isAnagramWithCountChars($s, $t) {

        $sFrequencyMap = count_chars($s, 1);
        $tFrequencyMap = count_chars($t, 1);

        if (count($sFrequencyMap) != count($tFrequencyMap)) {
            return false;
        }

        foreach ($sFrequencyMap as $k => $v) {
            if (! isset($tFrequencyMap[$k]) || $tFrequencyMap[$k] != $v) {
                return false;
            }
        }

        return true;
    }

    function isAnagramWithHashMap($s, $t) {

        $n = \strlen($s);
        $m = \strlen($t);

        if ($n != $m) {
            return false;
        }

        $frequencyMap = [];

        for ($i = 0; $i < $n; $i ++) {

            $charS = $s[$i];
            $charT = $t[$i];

            if (isset($frequencyMap[$charS])) {
                $frequencyMap[$charS] ++;
            } else {
                $frequencyMap[$charS] = 1;
            }

            if (isset($frequencyMap[$charT])) {
                $frequencyMap[$charT] --;
            } else {
                $frequencyMap[$charT] = -1;
            }
        }

        foreach ($frequencyMap as $frequency) {
            if ($frequency !== 0) {
                return false;
            }
        }

        return true;
    }

    function isAnagramWithFixedArray($s, $t) {

        $n = \strlen($s);
        $m = \strlen($t);

        if ($n != $m) {
            return false;
        }

        $a = \ord('a');
        $frequencyMap = new \SplFixedArray(26);

        for ($i = 0; $i < $n; $i ++) {
            $frequencyMap[\ord($s[$i]) - $a] += 1;
            $frequencyMap[\ord($t[$i]) - $a] -= 1;
        }

        foreach ($frequencyMap as $frequency) {
            if ($frequency !== null && $frequency !== 0) {
                return false;
            }
        }

        return true;
    }
}