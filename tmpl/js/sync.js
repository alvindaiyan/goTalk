function sync(){
	var self = this;
	this.serviceUrl="/";
	var  w;
	var token = null;
	
	var getUrlParameter = function(sParam) {
		var sPageURL = window.location.search.substring(1);
		var sURLVariables = sPageURL.split('&');
		for (var i = 0; i < sURLVariables.length; i++) 
		{
			var sParameterName = sURLVariables[i].split('=');
			if (sParameterName[0] == sParam) 
			{
				return sParameterName[1];
			}
		}
	} 


	var startWorker = function() {
		if(typeof(Worker) !== "undefined") {
			if(typeof(w) == "undefined") {
				w = new Worker("js/sync_worker.js");
				w.postMessage($("#myNameId").text());
			}
			// addEventListener is bettern than onmessage()
			w.addEventListener('message', function() {				
				var strArr = event.data.split(" ||| "); // firefox not work!!!
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
				$("#chat").scrollTop($("#chat")[0].scrollHeight);
			}, false);
		} else {
			document.getElementById("result").innerHTML = "Sorry, your browser does not support Web Workers...";
		}
	}



	$( document ).ready(function(){
		var name = getUrlParameter('username');
		token = getUrlParameter('token');
		var id = 0;
		$.post('/getuseridbyname',
			{username: name},
			function(data){
				console.log(JSON.stringify(data));
				try{					
					id = data.Id;
					$("#myNameId").text(id);
					$("#myNameTxt").text(name);
					startWorker();		
				} catch(err){
					alert("cannot load user" + err)
				}
			});		
	})



	this.stopWorker = function() { 
		w.terminate();
		w = undefined;
	}

	this.startWorker = function(){
		startWorker();
	}

	this.displayReceivedMsg = function(userId, msg){
		appendMsg()
	}

	this.sendMsg = function(){
		var username = $("#myNameTxt").text();
		var id = $("#myNameId").text();
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