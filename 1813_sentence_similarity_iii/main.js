function areSentencesSimilar(sentence1, sentence2) {
    if (sentence1 == sentence2) {
        return true;
    }
    if (sentence2.length < sentence1.length) {
        var oldsentence1 = sentence1;
        sentence1 = sentence2;
        sentence2 = oldsentence1;
    }
    if (sentence2.indexOf(sentence1 + " ") == 0) {
        return true;
    }
    if (sentence2.substring(sentence2.length - sentence1.length - 1) == " " + sentence1) {
        return true;
    }
    // sentence1 is found but not at the edge of string
    if (sentence2.indexOf(sentence1) != -1) {
        return false;
    }
    var i = sentence1.indexOf(" ");
    var word1 = "";
    var word2 = "";
    while (i != -1) {
        word1 = sentence1.substring(0, i);
        word2 = sentence1.substring(i + 1);
        if (sentence2.indexOf(word1 + " ") == 0 && sentence2.indexOf(" " + word2) == sentence2.length - (word2.length + 1)) {
            return true;
        }
        i = sentence1.indexOf(" ", i + 1);
    }
    return false;
}
;
console.log(areSentencesSimilar("A", "a A b A"));
