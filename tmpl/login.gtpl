<html>
<head>
	 <script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css">
    <!-- Latest compiled and minified JavaScript -->
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js"></script>

    <script type="text/javascript" src="js/login.js"></script>
<title>login</title>
</head>
<body>
  <div>
    user name:<input type="text" name="username" id = "unameLogin">
    password:<input type="password" name="password" id = "pwdLogin">
    <!-- <button type="button" class="btn btn-primary btn-lg" onclick="obj.register()" value = "login"/> -->
    <button type="button" data-loading-text="Loading..." class="btn btn-primary" autocomplete="off" onclick="obj.login()">
      login
    </button>
    <!-- Button trigger modal -->
    <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#myModalRegister" onclick="obj.register()">
      Register
    </button>
  </div>
<!-- <form action="/login" method="post">
    user name:<input type="text" name="username" id = "unameLogin">
    password:<input type="text" name="password" id = "pwdLogin">
    <input type="submit" value="submit">    
</form>
 -->



<!-- Modal -->
<div class="modal fade" id="myModalRegister" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">Register</h4>
      </div>
      <div class="modal-body" id = "result">        
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div>





<script type="text/javascript">
            // setup the login page
            var obj = new login();
</script>
</body>
</html>
