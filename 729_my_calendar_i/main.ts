class MyCalendar {
  root: Tree

  constructor() {}

  book(start: number, end: number): boolean {
    try {
      if (this.root == null) {
        this.root = new Tree(start, end)
      } else {
        this.root.add(start, end)
      }
    }catch(error) {
      return false
    }
    
    return true
  }
}

class Tree {
  start: number
  end: number
  left: Tree
  right: Tree

  constructor(start: number, end: number) {
    this.start = start
    this.end = end
  }

  add(start: number, end: number) {
    if (start < this.end && this.start < end) {
      throw new RangeError("new time interval overlap with current time interval")
    }

    if (end <= this.start) {
      if (this.left == null) {
        this.left = new Tree(start, end)
        return
      }
      return this.left.add(start, end)
    }

    if (this.right == null) {
      this.right = new Tree(start, end)
      return
    }

    return this.right.add(start, end)
  }
}

var obj = new MyCalendar()
console.log(obj.book(20, 29))
console.log(obj.book(13, 22))
console.log(obj.book(44, 50))
console.log(obj.book(1, 7))
console.log(obj.book(2, 10))
console.log(obj.book(14, 20))
console.log(obj.book(19, 25))
console.log(obj.book(36, 42))
console.log(obj.book(45, 50))
console.log(obj.book(47, 50))
console.log(obj.book(39, 45))
console.log(obj.book(44, 50))
console.log(obj.book(16, 25))
console.log(obj.book(45, 50))
console.log(obj.book(45, 50))
console.log(obj.book(12, 20))
console.log(obj.book(21, 29))
console.log(obj.book(11, 20))
console.log(obj.book(12, 17))
console.log(obj.book(34, 40))
console.log(obj.book(10, 18))
console.log(obj.book(38, 44))
console.log(obj.book(23, 32))
console.log(obj.book(38, 44))
console.log(obj.book(15, 20))
console.log(obj.book(27, 33))
console.log(obj.book(34, 42))
console.log(obj.book(44, 50))
console.log(obj.book(35, 40))
console.log(obj.book(24, 31))
