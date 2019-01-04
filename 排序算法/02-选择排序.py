# coding:utf-8

def chose_sort(l):
    n = len(l)
    for j in range(n-1):
        min_index = j
        for i in range(j+1,n):
            if l[min_index] > l[i]:
                min_index = i
        l[j],l[min_index] = l[min_index],l[j]
    return l

s = chose_sort([1,2,3,-1,0])
print(s)

