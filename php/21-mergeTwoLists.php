<?php

class ListNode {
public $val = 0;
public $next = null;
function __construct($val) { $this->val = $val; }
}


/**
 * @param ListNode $l1
 * @param ListNode $l2
 * @return ListNode
 */
function mergeTwoLists($l1, $l2)
{
    if ($l1 == null) {
        return $l2;
    }
    if ($l2 == null) {
        return $l1;
    }
    $new = new ListNode(0);
    $current = $new;
    while ($l1 != null && $l2 != null) {
        if ($l1->val < $l2->val) {
            $current->next = new ListNode($l1->val);
            $l1 = $l1->next != null ? $l1->next : null;
        } else {
            $current->next = new ListNode($l2->val);
            $l2 = $l2->next != null ? $l2->next : null;
        }
        $current = $current->next;
    }

    if ($l1 != null) {
        $current->next = $l1;
    } elseif ($l2 != null) {
        $current->next = $l2;
    }
    
    return $new->next;
}

$n1 = new ListNode(1);
$n1->next = new ListNode(2);
$n1->next->next = new ListNode(3);

$n2 = new ListNode(2);
$n2->next = new ListNode(3);
$n2->next->next = new ListNode(4);

mergeTwoLists($n1, $n2);



