<?php

namespace Zeleniy\Leetcode\ValidParentheses;

class Solution {

    private $forwardMap = [
        '(' => ')',
        '[' => ']',
        '{' => '}',
    ];
    private $backwardMap = [
        ')' => '(',
        ']' => '[',
        '}' => '{',
    ];

    function isValid($s) {

        $length = \strlen($s);
        $stack  = [];

        for ($i = 0; $i < $length; $i ++) {

            $char = $s[$i];

            if (isset($this->forwardMap[$char])) {
                \array_push($stack, $char);
                continue;
            }

            $closingChar = \array_pop($stack);

            if ($closingChar != $this->backwardMap[$char]) {
                return false;
            }
        }

        return \count($stack) == 0;
    }
}
