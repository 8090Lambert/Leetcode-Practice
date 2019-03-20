<?php
/**
 * Leetcode 第九题
 */
function isPalindrome($x)
{
    if ($x < 0 || ($x % 10 == 0 && $x != 0)) {
        return 0;
    }
    $digit = strlen($x);
    $times = $digit / 2;
    $target = 0;
    $begin = 0;

    while ($begin < floor($times)) {
        $target = $target * 10 + $x % 10;
        $x /= 10;
        $begin++;
    }

    return $target == floor($x) || $target == floor($x / 10);
}