function login(){
	this.register = function(){
		$.get( "/register", function( data ) {
			$( "#result" ).html( data );			
		});
	}

}