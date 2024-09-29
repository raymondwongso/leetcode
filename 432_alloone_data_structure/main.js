var AllOne = /** @class */ (function () {
    function AllOne() {
        this.counter = new Map();
        this.words = new Map();
    }
    AllOne.prototype.inc = function (key) {
        if (this.counter.get(key) == undefined) {
            this.counter.set(key, 0);
        }
        var oldCounter = this.counter.get(key);
        if (this.words.get(oldCounter) == undefined) {
            this.words.set(oldCounter, new Map());
        }
        this.words.get(oldCounter).delete(key);
        this.counter.set(key, oldCounter + 1);
        if (this.words.get(this.counter.get(key)) == undefined) {
            this.words.set(this.counter.get(key), new Map());
        }
        this.words.get(this.counter.get(key)).set(key, true);
    };
    AllOne.prototype.dec = function (key) {
        var oldCounter = this.counter.get(key);
        this.words.get(oldCounter).delete(key);
        this.counter.set(key, oldCounter - 1);
        this.words.get(this.counter.get(key)).set(key, true);
    };
    AllOne.prototype.getMaxKey = function () {
        var n = this.words.size;
        var res = "";
        if (this.words.get(n - 1).size > 0) {
            res = this.words.get(n - 1).keys[0];
        }
        return res;
    };
    AllOne.prototype.getMinKey = function () {
        var n = this.words.size;
        var res = "";
        for (var i = 0; i < n; i++) {
            if (this.words.get(i).size > 0) {
                res = this.words.get(i).keys[0];
                break;
            }
        }
        return res;
    };
    return AllOne;
}());
var obj = new AllOne();
obj.inc("hello");
obj.inc("hello");
obj.getMaxKey();
obj.getMinKey();
obj.inc("leet");
obj.getMaxKey();
obj.getMinKey();
