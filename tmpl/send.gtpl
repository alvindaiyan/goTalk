<html>
<head>
<title></title>
    <script type="text/javascript" src="js/sync.js"></script>
    <script type="text/javascript" src="js/demo_workers.js"></script>

</head>
<body>
<div id="chat"></div>
<form action="/send" method="post">
    txt:<input type="text" name="content">
    username:<input type="text" name="username">
    id:<input type="text" name="id">
    sendTo:<input type="text" name="sendToId">
    <input type="submit" value="submit">   
</form>

 <p>Count numbers: <output id="result"></output></p>
<button onclick="obj.startWorker()">Start Worker</button> 
<button onclick="obj.stopWorker()">Stop Worker</button>
<br><br>
 <script type="text/javascript">
            // setup the login page
            var obj = new sync();
    </script>
</body>
</html>
