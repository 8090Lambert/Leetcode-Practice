<?php


function isPalindrome($x) {
    $x = (int) $x;
    if ($x < 0 || $x % 10 == 0) {

    }
    $target = 0;
    while ($x >= $target) {
        $target = $target * 10 + $x % 10;
        $x = $x / 10;
    }

    
    return $target == $x || $x == $target / 10;
}



function isPalindrome1($x)
{
    if ($x < 0 || ($x % 10 == 0 && $x != 0)) {
        return 0;
    }
    $digit = strlen($x);
    $times = $digit / 2;
    $target = 0;
    $begin = 0;

    while ($begin <= $times) {
        $target = $target * 10 + $x % 10;
        $x /= 10;
        $begin++;
    }

    
    //return $target == sprintf('%.1f', $x * 10);
}


echo isPalindrome1(121);