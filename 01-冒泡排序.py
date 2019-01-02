# coding:utf-8

def bubble_sort(l):
    for j in range(len(l)-1):
        count = 0
        for i in range(len(l)-1-j):
            if l[i] > l[i+1]:
                l[i],l[i+1] = l[i+1],l[i]
                count += 1
        if count == 0:
            return l
    return l

print(bubble_sort([1,2,3]))
