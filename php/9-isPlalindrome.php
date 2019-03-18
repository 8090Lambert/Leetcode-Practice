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
    

}


class A 
{
    public function __construct()
    {

    }
}
echo isPalindrome(121);