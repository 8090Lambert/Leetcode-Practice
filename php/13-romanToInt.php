<?php

/**
 * LeetCode 第十三题
 */
function romanToInt($s)
{
    $mapper = ['I' => 1, 'V' => 5, 'X' => 10, 'L' => 50, 'C' => 100, 'D' => 500,'M' => 1000];
    $len = strlen($s);
    $cur = 0;
    $pre = 0;
    $target = 0;
    for ($i = 0; $i < $len; $i++) {
        $cur = $mapper[$s[$i]];
        if ($cur <= $pre) {
            $target = $target + $cur;
        } else {
            $target = $target + $cur - 2 * $pre;
        }
        $pre = $mapper[$s[$i]];
    }

    return $target; 
}

/**
 * WithMine
 */
function romanToIntWithMine($s)
{
    $mapper = ['I' => 1, 'V' => 5, 'X' => 10, 'L' => 50, 'C' => 100, 'D' => 500,'M' => 1000];
    $len = strlen($s);
    $target = 0;
    for ($i = 0; $i < $len; $i++) {
        $target += $mapper[$s[$i]];
        if ($i > 0) {
            if (in_array($s[$i], ['V', 'X'])) {
                if ($s[$i-1] == 'I') {
                    $target = $target - 2 * $mapper[$s[$i-1]];
                }
            } elseif (in_array($s[$i], ['L', 'C'])) {
                if ($s[$i-1] == 'X') {
                    $target = $target - 2 * $mapper[$s[$i-1]];
                }
            } elseif (in_array($s[$i], ['D', 'M'])) {
                if ($s[$i-1] == 'C') {
                    $target = $target - 2 * $mapper[$s[$i-1]];
                }
            }
        }
    }

    return $target;
}