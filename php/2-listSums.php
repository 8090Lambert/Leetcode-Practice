<?php

/**
 * LeetCode 第二题
 * 
 * 两链表之和
 */

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}


function listSums($l1, $l2) {
    $response = new ListNode(0);
    $carry = 0;
    $current = $response;
    while ($l1 || $l2) {
        $sum = $l1->val + $l2->val + $carry;
        $current->next = new ListNode($sum % 10);
        $current = $current->next;
        $carry = $sum >= 10 ? 1 : 0;
        $l1 = $l1->next == null ? null : $l1->next;
        $l2 = $l2->next == null ? null : $l2->next;
    }

    if ($carry != 0) {
        $current->next = new ListNode(1);
    }

    return $response->next;
}


$l11 = new ListNode(2);
$l12 = new ListNode(4);
$l13 = new ListNode(3);
$l13->next = $l12;
$l12->next = $l11;

$l21 = new ListNode(5);
$l22 = new ListNode(6);
$l23 = new ListNode(4);
$l23->next = $l22;
$l22->next = $l21;

listSums($l13, $l23);
