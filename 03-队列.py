# coding:utf-8

class Queue(object):
    """队列的实现"""
    def __init__(self):
        self.__list = []

    def enqueue(self,item):
        self.__list.append(item)

    def dequeue(self):
        return self.__list.pop(0)

    def is_empty(self):
        return not self.__list

    def size(self):
        return len(self.__list)

if __name__ == "__main__":
    q = Queue()