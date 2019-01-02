# coding:utf-8

class Node(object):
    """节点"""
    def __init__(self,elem):
        self.elem = elem
        self.next = None

class SingleLinkList(object):
    """单链表"""
    def __init__(self,node=None):
        self.__head = node

    def is_empty(self):
        return self.__head == None

    def lenth(self):
        # cur游标,用来移动遍历节点
        cur = self.__head
        # count记录数量
        count = 0
        while cur != None:
           count += 1
           cur = cur.next
        return count

    def travel(self):
        cur = self.__head
        while cur != None:
            print(cur.elem)
            cur = cur.next

    def append(self,item):
        node = Node(item)
        cur = self.__head
        if self.is_empty():
            self.__head = node
        else:
            while cur.next != None:
                cur = cur.next
            cur.next = node

    def insert(self,pos,item):
        if pos <= 0:
            self.add(item)
        elif pos > (self.lenth()-1):
            self.append(item)
        else:
            node = Node(item)
            pre = self.__head
            count = 0

            while count < (pos-1):
                count += 1
                pre = pre.next
            #当循环退出后,pre = pos-1
            node.next = pre.next
            pre.next = node

    def search(self,item):
        cur = self.__head
        count = 0
        while cur != None:
            count += 1
            cur = cur.next
            if cur.elem == item:
                return count
        return -1

    def add(self,item):
        node = Node(item)
        node.next = self.__head
        self.__head = node

    def remove(self,item):
        pre = None
        cur = self.__head
        while cur != None:
            if cur.elem == item:
                #先判断此节点是否是头结点,如果是头结点,
                if cur == self.__head:
                    self.__head = cur.next
                else:
                    pre.next = cur.next
                break
            else:
                pre = cur
                cur = cur.next





# sll = SingleLinkList()
# print(sll.lenth())
# sll.travel()
# sll.append(124)
# print(sll.lenth())


if __name__ == "__main__":
    sll = SingleLinkList()
    print(sll.is_empty())
    print(sll.lenth())
    sll.append(1)
    print(sll.is_empty())
    print(sll.lenth())
    sll.add(8)
    sll.append(2)

    sll.append(3)
    sll.append(4)
    sll.append(5)
    sll.append(6)
    sll.insert(1, 123)
    sll.travel()
    print("================")
    # print(sll.search(6))
    sll.remove(6)
    sll.travel()