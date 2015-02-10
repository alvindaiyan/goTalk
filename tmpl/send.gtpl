<html>
<head>
    <title>goTalk</title>

    <script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css">
    <!-- Latest compiled and minified JavaScript -->
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js"></script>

    <script type="text/javascript" src="js/sync.js"></script>
    <style>
        .bs-callout {
        padding: 20px;
        margin: 20px 0;
        border: 1px solid #eee;
        border-left-width: 5px;
        border-radius: 3px;
        }
        .bs-callout h4 {
        margin-top: 0;
        margin-bottom: 5px;
        }
        .bs-callout p:last-child {
        margin-bottom: 0;
        min-width: 500px;
        }
        .bs-callout code {
        border-radius: 3px;
        }
        .bs-callout+.bs-callout {
        margin-top: -5px;
        }
        .bs-callout-default {
        border-left-color: #777;
        }
        .bs-callout-default h4 {
        color: #777;
        }
        .bs-callout-primary {
        border-left-color: #428bca;
        }
        .bs-callout-primary h4 {
        color: #428bca;
        }
        .bs-callout-success {
        border-left-color: #5cb85c;
        }
        .bs-callout-success h4 {
        color: #5cb85c;
        }
        .bs-callout-danger {
        border-left-color: #d9534f;
        }
        .bs-callout-danger h4 {
        color: #d9534f;
        }
        .bs-callout-warning {
        border-left-color: #f0ad4e;
        }
        .bs-callout-warning h4 {
        color: #f0ad4e;
        }
        .bs-callout-info {
        border-left-color: #5bc0de;
        }
        .bs-callout-info h4 {
        color: #5bc0de;
        }

        .bs-callout-hs {
        padding: 20px;
        margin: 20px 0;
        border: 1px solid #eee;
        border-right-width: 5px;
        border-radius: 3px;
        }
        .bs-callout-hs h4 {
        margin-top: 0;
        margin-bottom: 5px;
        }
        .bs-callout-hs p:last-child {
        margin-bottom: 0;
        min-width: 500px;
        }
        .bs-callout-hs code {
        border-radius: 3px;
        }
        .bs-callout-hs+.bs-callout-hs {
        margin-top: -5px;
        }
        .bs-callout-hs-primary {
        border-right-color: #5cb85c;
        }
        .bs-callout-hs-primary h4 {
        color: #5cb85c;
        }
    </style>

</head>


<body>
<div class="container bs-doc-container">
    <div class="col-md-9" role="main">
        <div class="bs-docs-section">
            <!-- <form> -->
           
            <div class="form-group">
                <span class="label label-info" id="myNameTxt" > alvindaiyan</span>
                |
                <label for="myMsgNum">Message </label> 
                <span class="badge" id="myMsgNum">0</span>
            </div>
            <div class="form-group">
                <label for="myNameId">My ID</label>
                <span class="label label-info" id="myNameId" name="id"> alvindaiyan</span>
                <!-- <input type="text" class="form-control" id="myNameId" placeholder="e.g. 1234" name="id"> -->
            </div>
            <div class="form-group">
                <label for="reciverId">Reciver ID</label>
                <input type="text" class="form-control" id="reciverId" placeholder="e.g. 1234" name="sendToId">
            </div>
            <!--              <div class="form-group">
                             <label for="exampleInputFile">File input</label>
                             <input type="file" id="exampleInputFile">
         
                             <p class="help-block">send message</p>
                         </div> -->
            <!-- <div class="checkbox">
                <label>
                    <input type="checkbox"> Check me out
                </label>
            </div> -->
            <div class="jumbotron" style="overflow:scroll; overflow-x:hidden; height:50%; " id ="chat">
                <span style="float:left;">
                    <div class="bs-callout bs-callout-primary" style="width:50%; float:left;">
                        <h4 >receiver : </h4>
                        <p>test message 1</p>
                    </div>
                </span>               
                <span style="float:right;">
                    <div class="bs-callout-hs bs-callout-hs-primary" style="width:50%;  float:right;">
                        <h4 >sender (me): </h4>
                        <p>test message 2</p>
                    </div>
                </span>                  
            </div>


            <p>Test Msg:
                <output id="result"></output>
            </p>
            <button onclick="obj.startWorker()">Start Worker Tester</button>
            <button onclick="obj.stopWorker()">Stop Worker Tester</button>
            <br><br>

             <div class="form-group">
                <label for="enterTxt">Enter Text:</label>
                <input type="text" class="form-control" id="enterTxt" placeholder="Enter Text" name="content">
            </div>
            <button type="submit" class="btn btn-default" onclick="obj.sendMsg()">Submit</button>
            <!-- </form> -->
        </div>

    </div>
</div>


<script type="text/javascript">
            // setup the login page
            var obj = new sync();
</script>
</body>
</html>
