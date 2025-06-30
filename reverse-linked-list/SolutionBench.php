<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

class SolutionBench {

    private Solution $solution;

    public function __construct() {

        $this->solution = new Solution();
    }

    /**
     * @Revs(1000)
     * @Iterations(5)
     */
    public function benchReverseListRecursive() {

        foreach (DataSet::getData() as $data) {
            $this->solution->reverseListRecursive($data[1]);
        }
    }

    /**
     * @Revs(1000)
     * @Iterations(5)
     */
    public function benchReverseListIterative() {

        foreach (DataSet::getData() as $data) {
            $this->solution->reverseListIterative($data[1]);
        }
    }
}
