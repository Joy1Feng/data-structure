# coding:utf-8


def shell_sort(l):
    """希尔排序"""
    n = len(l)
    gap = n // 2
    while gap >= 1:
        for j in range(gap,n):
            i = j
            while i > 0:
                if l[i] < l[i-gap]:
                    l[i], l[i-gap] = l[i-gap], l[i]
                    i -= gap
                else:
                    break
        gap //= 2

l = [1,2,3,-1,0]
shell_sort(l)
print(l)