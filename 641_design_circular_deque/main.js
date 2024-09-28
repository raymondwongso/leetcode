var MyCircularDeque = /** @class */ (function () {
    function MyCircularDeque(k) {
        this.arr = new Array(k);
        this.arr = this.arr.fill(-1);
        this.head = 0;
        this.tail = 0;
        this.counter = 0;
    }
    MyCircularDeque.prototype.insertFront = function (value) {
        if (this.counter == this.arr.length) {
            return false;
        }
        if (this.arr[this.head] != -1) {
            this.head -= 1;
            if (this.head == -1) {
                this.head = this.arr.length - 1;
            }
        }
        this.arr[this.head] = value;
        this.counter += 1;
        return true;
    };
    MyCircularDeque.prototype.insertLast = function (value) {
        if (this.counter == this.arr.length) {
            return false;
        }
        if (this.arr[this.tail] != -1) {
            this.tail += 1;
            if (this.tail == this.arr.length) {
                this.tail = 0;
            }
        }
        this.arr[this.tail] = value;
        this.counter += 1;
        return true;
    };
    MyCircularDeque.prototype.deleteFront = function () {
        if (this.arr[this.head] == -1) {
            return false;
        }
        this.arr[this.head] = -1;
        this.head += 1;
        if (this.head > this.tail) { // head can't overtaking tail
            this.head = this.tail;
        }
        else if (this.head == this.arr.length) { // if over the range, reset index
            this.head = 0;
        }
        this.counter -= 1;
        return true;
    };
    MyCircularDeque.prototype.deleteLast = function () {
        if (this.arr[this.tail] == -1) {
            return false;
        }
        this.arr[this.tail] = -1;
        this.tail -= 1;
        if (this.tail < this.head) { // tail can't be behind of head
            this.tail = this.head;
        }
        else if (this.tail == -1) { // if over the range, reset index
            this.tail = this.arr.length - 1;
        }
        this.counter -= 1;
        return true;
    };
    MyCircularDeque.prototype.getFront = function () {
        return this.arr[this.head];
    };
    MyCircularDeque.prototype.getRear = function () {
        return this.arr[this.tail];
    };
    MyCircularDeque.prototype.isEmpty = function () {
        return this.counter == 0;
    };
    MyCircularDeque.prototype.isFull = function () {
        return this.counter == this.arr.length;
    };
    return MyCircularDeque;
}());
var k = 77;
var obj = new MyCircularDeque(k);
obj.insertFront(89);
obj.getRear();
obj.deleteLast();
obj.getRear();
obj.insertFront(19);
obj.insertFront(23);
obj.insertFront(23);
obj.insertFront(82);
obj.isFull();
obj.insertFront(45);
obj.isFull();
obj.getRear();
obj.deleteLast();
obj.getFront();
obj.getFront();
obj.insertLast(74);
obj.deleteFront();
obj.getFront();
obj.insertLast(98);
obj.getRear();
obj.insertLast(99);
