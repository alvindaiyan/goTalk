function sync(){
	var self = this;
	this.serviceUrl="/";
	var  w;
	

	this.startWorker = function() {
		if(typeof(Worker) !== "undefined") {
			if(typeof(w) == "undefined") {
				w = new Worker("js/demo_workers.js");
			}
			w.addEventListener('message', function(e) {
				document.getElementById("result").innerHTML = event.data;
			}, false)
			// w.onmessage = function(event) {
			// 	document.getElementById("result").innerHTML = event.data;
			// };
		} else {
			document.getElementById("result").innerHTML = "Sorry, your browser does not support Web Workers...";
		}
	}

	this.stopWorker = function() { 
		w.terminate();
		w = undefined;
	}



}