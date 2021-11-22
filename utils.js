module.exports = {
    sleep: function (seconds) {
        var e = new Date().getTime() + (seconds * 1000);
        while (new Date().getTime() <= e) { }
    },
    currentTimeSeconds: function () {
        return new Date().getTime() / 1000
    }
}