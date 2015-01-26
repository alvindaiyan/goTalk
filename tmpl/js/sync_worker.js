var i = 0;

function timedCount() {
    i = i + 1;
    postMessage(i + ".test");
    setTimeout("timedCount()",500);
}

timedCount();