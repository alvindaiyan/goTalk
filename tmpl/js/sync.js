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
			// $.ajax({
			// 	type: "POST",
			// 	url: "/send",
			// 	async:true,
			// 	contentType:false,
			// 	cache: false,
			// 	processData: false,
			// 	data:{
			// 		username: username,
			// 		id: id,
			// 		sendToId: sendToId,
			// 		content: content
			// 	},
			// 	dataFilter: function(data) {
			// 		console.log(JSON.stringify(data));
			// 		appendMsg("me (success)", content, false);
			// 		$("#chat").scrollTop($("#chat")[0].scrollHeight);

			// 	}, 
			// 	traditional:true,
			// 	error :  function(err) {
			// 		var result = jQuery.parseJSON(err.responseText);
			// 		if(result.error)
			// 		{
			// 			alert(decodeURIComponent(result.error));                  
			// 		}                 
			// 	}
			// });		

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