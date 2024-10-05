function checkInclusion(s1, s2) {
    if (s1.length > s2.length) {
        return false;
    }
    var s1arr = new Array(26);
    s1arr.fill(0);
    var s2arr = new Array(26);
    s2arr.fill(0);
    var offset = 97;
    for (var i = 0; i < s1.length; i++) {
        s1arr[s1.charCodeAt(i) - offset]++;
        s2arr[s2.charCodeAt(i) - offset]++;
    }
    if (s1arr == s2arr) {
        return true;
    }
    for (var i = s1.length; i < s2.length; i++) {
        var j = i - s1.length;
        s2arr[s2.charCodeAt(j) - offset]--;
        s2arr[s2.charCodeAt(i) - offset]++;
        if (s1arr == s2arr) {
            return true;
        }
    }
    return false;
}
;
checkInclusion("abc", "aadbcasext");
