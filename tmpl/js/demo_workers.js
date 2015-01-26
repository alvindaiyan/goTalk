var i = 0;

function timedCount() {
    i = i + 1;
    postMessage(i + ".test");
    setTimeout(function(){timedCount()},1000);

}
timedCount(i);