package linkedlist_test

type LinkedList interface {
    length() int
    insertAt(item any, index uint)
    remove(item any) (any, error)
    removeAy(index uint) (any, error)
    append(item any)
    prepend(item any)
    get(index uint) (any, error)
}

