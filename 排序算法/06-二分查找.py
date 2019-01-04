# coding: utf-8



#一定要先排序
def binary_search(l, item):
    """二分查找, 递归"""
    l.sort()
    n = len(l)
    if n > 0:
        mid = n // 2
        if l[mid] == item:
            return True
        elif item < l[mid]:
            return binary_search(l[:mid], item)
        else:
            return binary_search(l[mid+1:], item)
    return False
# res = binary_search([1, 2, 3, 0], 0)
# print(res)

def binary_search1(l,item):
    """二分查找, 非递归"""

    l.sort()
    n = len(l)
    first = 0
    last = n - 1
    while first <= last:
        mid = (first + last) // 2
        if l[mid] == item:
            return True
        elif item < l[mid]:
            last = mid - 1
        else:
            first = mid + 1
    return False

res = binary_search1([1, 2 , 3, 0], 0)
print(res)
