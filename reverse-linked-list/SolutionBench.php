<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

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
    function benchReverseList()
    {
        foreach (DataSet::getData() as $data) {
            $this->solution->reverseList($data[1]);
        }
    }
}
