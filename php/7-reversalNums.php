<?php

/**
 * LeetCode第七题
 * 数字反转
 */

function reverse($x)
{
    $min = -1 * pow(2, 31);
    $max = pow(2, 31) - 1;
    $rev = 0;
    $temp = 0;
    if ($x > 0) {
        while ($x >= 1) {
            $temp = $x % 10;
            $rev = $rev * 10 + $temp;
            if ($rev > $max) {
                return 0;
            }
            $x /= 10;
        }
    } else {
        while ($x <= -1) {
            $temp = $x % 10;
            $rev = $rev * 10 + $temp;
            if ($rev < $min) {
                return 0;
            }
            $x /= 10;

        }
    }
    
    return $rev;
}


/**
 * Second Way.
 */
function reverseWithAnswer($x)
{
    $min = -1 * pow(2, 31);
    $max = pow(2, 31) - 1;
    $target = 0;
    if ($x > 0) {
        while ($x >= 1) {
            $pop = $x % 10;
            if ($target > $max / 10 || ($target == $max / 10 && $pop > 7)) {
                return 0;
            }
            $target = $target * 10 + $pop;
            $x = $x / 10;
        }
    } else {
        while ($x <= -1) {
            $pop = $x % 10;
            if ($target < $min / 10 || ($target == $min / 10 && $pop < -8)) {
                return 0;
            }
            $target = $target * 10 + $pop;
            $x = $x / 10;
        }
    }
    

    return $target;
}