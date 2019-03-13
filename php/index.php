<?php

 class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
 }

function addTwoNumbers($l1, $l2){
    $num1 = 0;
    $num2 = 0;
    $location = 1;
    while ($l1->val != null && $l2->val != null) {
        $num1 += $l1->val * $location;
        $num2 += $l2->val * $location;
        $l1 = $l1->next;
        $l2 = $l2->next;
        $location *= 10;
    }
    $location = 1;


    $num1 = (string)$num1;
    $num2 = (string)$num2;
    $target = 0;
    $locationNum = strlen($num1);

    echo $num1;
    echo PHP_EOL;
    echo $num2;
    echo PHP_EOL;
    
    for ($i=$locationNum-1; $i>= 0; $i--) {
        # code...
        $target += ($num1[$i] + $num2[$i]) * $location;
        $location *= 10;
    }


    echo $target;
}


function TwoAdd($l1, $l2) {
    $target = 0;    // 目标值
    $carry = 0; // 进位
    $index = 1; // 位数
    $current = new ListNode(0);

    while ($l1 != null|| $l2 != null) {
        $sum = $l1->val + $l2->val + $carry;
        $carry = $sum >= 10 ? 1 : 0;
        $target = new ListNode(($sum % 10) * pow(10, $index - 1));

        $l1 = $l1->next != null ? $l1->next : null;
        $l2 = $l2->next != null ? $l2->next : null;
        $index++;
    }

    if ($carry > 0) {
        $current->next = new ListNode(1);
    }

    
    return $target;
}




function main() {
    $l11 = new ListNode(2);
    $l12 = new ListNode(4);
    $l13 = new ListNode(3);
    //$l11->next = $l12;
    //$l12->next = $l13;
    $l13->next = $l12;
    $l12->next = $l11;

    $l21 = new ListNode(5);
    $l22 = new ListNode(6);
    $l23 = new ListNode(4);
    // $l21->next = $l22;
    // $l22->next = $l23;
    $l23->next = $l22;
    $l22->next = $l21;

    //addTwoNumbers($l11, $l21);
    echo TwoAdd($l13, $l23);
}

main();
