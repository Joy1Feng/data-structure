# coding:utf-8

def insert_sort(l):
    """插入排序"""
    for i in range(1,len(l)):
        while i > 0:
            if l[i] < l[i-1]:
                l[i], l[i-1] = l[i-1], l[i]
                i -= 1
            else:
                break

l = [1,2,3,-1,0,-2]
# insert_sort(l)
# print(l)

def insert_sort1(l):
    for i in range(1,len(l)):
        for j in range(i,0,-1):
            if l[j] < l[j-1]:
                l[j], l[j-1] = l[j-1], l[j]
            else:
                break

