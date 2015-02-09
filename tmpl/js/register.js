function register(){
	this.register = function() {
		var name = $("#registerName").val();
		var pwd = $("#registerPwd").val();
		$.post('/register',
			{
				username: name,
				password: pwd
			},
			function(data){				
				try{		
					console.log(JSON.stringify(data));
				} catch(err){
					alert("cannot load user" + err)
				}
			});		
	}
}