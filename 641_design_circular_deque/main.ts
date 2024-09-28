// Keep track using head and tail index
// to minimize calculating full or empty, use counter to store the current occupied indexes
// Time complexity: O(1)
// Space complexity: O(n) where n is the number of input
class MyCircularDeque {
  arr: Array<number>
  head: number
  tail: number
  counter: number

  constructor(k: number) {
    this.arr = new Array<number>(k)
    this.arr = this.arr.fill(-1)
    this.head = 0
    this.tail = 0
    this.counter = 0
  }

  insertFront(value: number): boolean {
    if (this.isFull()) {
      return false
    }

    if (this.arr[this.head] != -1) {
      this.head -= 1
      if (this.head == -1) {
        this.head = this.arr.length-1
      }
    }

    this.arr[this.head] = value
    this.counter += 1
    return true
  }

  insertLast(value: number): boolean {
    if (this.isFull()) {
      return false
    }

    if (this.arr[this.tail] != -1) {
      this.tail += 1
      if (this.tail == this.arr.length) {
        this.tail = 0
      }
    }

    this.arr[this.tail] = value
    this.counter += 1
    return true
  }

  deleteFront(): boolean {
    if (this.arr[this.head] == -1) {
      return false
    }

    this.arr[this.head] = -1
    if (this.head != this.tail) {
      this.head += 1
    }

    if (this.head == this.arr.length) { // if over the range, reset index
      this.head =  0
    }

    this.counter -= 1
    return true
  }

  deleteLast(): boolean {
    if (this.arr[this.tail] == -1) {
      return false
    }

    this.arr[this.tail] = -1
    if (this.tail != this.head) {
      this.tail -= 1
    }
    
    if (this.tail == -1) { // if over the range, reset index
      this.tail =  this.arr.length-1
    }

    this.counter -= 1
    return true
  }

  getFront(): number {
    return this.arr[this.head]
  }

  getRear(): number {
    return this.arr[this.tail]
  }

  isEmpty(): boolean {
    return this.counter == 0
  }

  isFull(): boolean {
    return this.counter == this.arr.length
  }
}

let k = 77
var obj = new MyCircularDeque(k)
obj.insertFront(89)
obj.getRear()
obj.deleteLast()
obj.getRear()
obj.insertFront(19)
obj.insertFront(23)
obj.insertFront(23)
obj.insertFront(82)
obj.isFull()
obj.insertFront(45)
obj.isFull()
obj.getRear()
obj.deleteLast()
obj.getFront()
obj.getFront()
obj.insertLast(74)
obj.deleteFront()
obj.getFront()
obj.insertLast(98)
obj.getRear()
obj.insertLast(99)
