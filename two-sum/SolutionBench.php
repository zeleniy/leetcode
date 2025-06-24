<?php

namespace Zeleniy\Leetcode\TwoSum;

/**
 *
 */
class SolutionBench
{
    private Solution $solution;

    public function __construct()
    {
        $this->solution = new Solution();
    }

    /**
     * @Revs(1000)
     * @Iterations(5)
     */
    function benchTwoSum()
    {
        foreach (DataSet::getData() as $data) {
            $this->solution->twoSum($data[2], $data[1]);
        }
    }
}
