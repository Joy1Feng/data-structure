# coding:utf-8


def quick_sort(l,first,last):
    if first >= last:
        return
    mid_value = l[first]
    low = first
    high = last
    while low < high:

        while (low < high) and (l[high] >= mid_value):
            high -= 1
        l[low] = l[high]
        # low += 1

        while low < high and l[low] < mid_value:
            low += 1
        l[high] = l[low]
        # high -= 1
    l[low] = mid_value
    quick_sort(l,first,low-1)
    quick_sort(l,low+1,last)
