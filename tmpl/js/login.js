function login(){


	this.login = function(){
		var name = $("#unameLogin").val();
		var pwd = $("#pwdLogin").val();
		$.post('/login',
		{
			username: name,
			password: pwd
		},
		function(data){				
			try{		
				console.log(JSON.stringify(data));
				window.location.href = "/send?username=" + name + "&token=" + data.SessionID;
			} catch(err){
				alert("cannot load user" + err)
			}
		})
		.fail(function() {
    		alert( "fail to login, please try again" );
  		});		
	}

	this.register = function(){
		$.get( "/register", function( data ) {
			$( "#result" ).html( data );			
		});
	}

}