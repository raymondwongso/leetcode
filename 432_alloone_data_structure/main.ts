class AllOne {
  counter: Map<string, number> // stores counter for each word
  words: Map<number, Map<string, boolean>>

  constructor() {
    this.counter = new Map<string, number>()
    this.words = new Map<number, Map<string, boolean>>()
  }

  inc(key: string): void {
    if (this.counter.get(key) == undefined) {
      this.counter.set(key, 0)
    }
    let oldCounter = this.counter.get(key)!
    if (this.words.get(oldCounter) == undefined) {
      this.words.set(oldCounter, new Map<string, boolean>())
    }

    this.words.get(oldCounter)!.delete(key)

    this.counter.set(key, oldCounter + 1)
    if (this.words.get(this.counter.get(key)!) == undefined) {
      this.words.set(this.counter.get(key)!, new Map<string, boolean>())
    }
    this.words.get(this.counter.get(key)!)!.set(key, true)
  }

  dec(key: string): void {
    let oldCounter = this.counter.get(key)!
    this.words.get(oldCounter)!.delete(key)

    this.counter.set(key, oldCounter - 1)
    this.words.get(this.counter.get(key)!)!.set(key, true)
  }

  getMaxKey(): string {
    let n = this.words.size
    let res = ""

    if (this.words.get(n-1)!.size > 0) {
      res = this.words.get(n-1)!.keys[0]
    }

    return res
  }

  getMinKey(): string {
    let n = this.words.size
    let res = ""

    for (let i = 0; i < n; i++) {
      if (this.words.get(i)!.size > 0) {
        res = this.words.get(i)!.keys[0]
        break
      }
    }

    return res
  }
}

var obj = new AllOne()
obj.inc("hello")
obj.inc("hello")
obj.getMaxKey()
obj.getMinKey()
obj.inc("leet")
obj.getMaxKey()
obj.getMinKey()
