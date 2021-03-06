# coding:utf-8

class Stack(object):

    def __init__(self):
        self.__list = []

    def push(self,item):
        """添加一个元素到栈顶"""
        self.__list.append(item)

    def pop(self):
        return self.__list.pop()

    def peek(self):
        if self.__list:
            return self.__list[-1]
        return None

    def is_empty(self):
        return not self.__list

    def size(self):
        return len(self.__list)

if __name__ == "__main__":
    s = Stack()
    s.push(1)
    s.push(2)
    s.push(3)
    s.push(4)
    s.push(5)
    print(s.pop())
    print(s.pop())
    print(s.pop())
    print(s.pop())
    print(s.pop())
