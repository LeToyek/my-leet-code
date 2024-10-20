# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = ListNode(0)
        current = head
        carry = 0
        while l1 or l2 or carry:
            a = l1.val if l1 is not None else 0
            b = l2.val if l2 is not None else 0 

            res = a + b + carry
            carry = res // 10
            current.next = ListNode(res%10)

            current = current.next
            if l1 is not None : l1 = l1.next
            if l2 is not None : l2 = l2.next

        result = head.next
        head.next = None
        return result

            
            
        

            

        