function sync(){
	var self = this;
	this.serviceUrl="/";
	var  w;
	

	this.startWorker = function() {
		if(typeof(Worker) !== "undefined") {
			if(typeof(w) == "undefined") {
				w = new Worker("js/sync_worker.js");
			}
			// addEventListener is bettern than onmessage()
			w.addEventListener('message', function() {				
				var strArr = event.data.split(" ||| ");
				var json = strArr[0];
				var count = strArr[1];
				var msgcount = strArr[2];

					

				// Content       string
				// UserIdSend    int
				// UserIdReceive int

				// update the chat area
				var msgObjs = jQuery.parseJSON(json);
				for( var i = 0; i < msgObjs.length; i++){				
					appendMsg(msgObjs[i].UserIdSend, msgObjs[i].Content, true);	
				}

				// update the msg count
				$("#myMsgNum").text(msgObjs.length);	
				
				document.getElementById("result").innerHTML = "data: " + json + ", count: " + count + ", msgs: " + msgcount + " mgs";
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

	this.displayReceivedMsg = function(userId, msg){
		appendMsg()
	}

	this.sendMsg = function(){
		var username = $("#myNameTxt").val();
		var id = $("#myNameId").val();
		var sendToId = $("#reciverId").val();
		var content = $("#enterTxt").val();
		if(content == '') {
			alert("no empty content -_-!!!")
		} else{
			$.post('/send', 
				{
					username: username,
					id: id,
					sendToId: sendToId,
					content: content
				}, function(data) {
					console.log(JSON.stringify(data));
					appendMsg("me (success)", content, false);
					$("#chat").scrollTop($("#chat")[0].scrollHeight);
				});
		}
		$("#enterTxt").val("");

	}


	var appendMsg = function (title, msg, isSender) {
		if (isSender){
			$("#chat").append('<span style="float:left;"><div class="bs-callout bs-callout-primary" style="width:50%; float: left"><h4 >' + title +': </h4><p>' + msg +'</p></div></span>');
		} else {
			$("#chat").append('<span style="float:right;"><div class="bs-callout-hs bs-callout-hs-primary" style="width:50%; float: right"><h4 >' + title +': </h4><p>' + msg +'</p></div></span>');
		}

	}

}