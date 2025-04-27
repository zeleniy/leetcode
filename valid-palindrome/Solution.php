<?php

namespace Zeleniy\Leetcode\ValidPalindrome;

class Solution {

    function isPalindrome($s) {

        $n = \strlen($s);

        for ($i = 0, $j = $n - 1; $i < $j; $i ++, $j --) {

            while (! \ctype_alnum($s[$i]) && $i < $j) {
                $i ++;
            }

            while (! \ctype_alnum($s[$j]) && $i < $j) {
                $j --;
            }

            if (\lcfirst($s[$i]) != \lcfirst($s[$j])) {
                return false;
            }
        }

        return true;
    }
}
