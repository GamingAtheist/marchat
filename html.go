package main

import "log"
import "html/template"
import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := rootTemplate.Execute(w, config)
	if err != nil {
		log.Fatal("[!] serving root: ", err.Error())
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!doctype html>
<html>
  <head>
    <meta charset="utf8" />
    <title>MarCHat</title>
    <link href="//netdna.bootstrapcdn.com/twitter-bootstrap/2.3.1/css/bootstrap-combined.no-icons.min.css"
          rel="stylesheet">
    <script>
        var transmitter, receiver, input, output;
        
        function printMessages(ml) {
                if (ml.length === 0)
                        return;
                for (var i = 0; i < ml.length; i++)
                        output.innerHTML = '<p>' + ml[i] + '</p>' + output.innerHTML;
        };
        
        function onKey(e) {
        	if (e.keyCode == 13) {
        		sendMessage();
        	}
        };
        
        function sendMessage() {
                var m = input.value;
                input.value = "";
                transmitter.send(m + '\n');
        };
        
        function checkMessages() {
                console.log('checking for messages');
                var receiver = new WebSocket('ws://127.0.0.1:{{.Port}}/incoming')
                receiver.onmessage = printMessages
                setTimeout(receiver.close, 100);
        }
        
        function init() {
                output = document.getElementById('messages');
                transmitter = new WebSocket('ws://127.0.0.1:{{.Port}}/socket');
        
        	input = document.getElementById("input");
        	input.addEventListener("keyup", onKey, false);
                check = setInterval(checkMessages, 250);
        
        };
        window.addEventListener("load", init, false);
   
    </script>

    <style type="text/css">
        html,
        body {
            height: 100%; 
        }
        #wrap {
            min-height: 100%;
            height: auto !important;
            height: 100%; 
            margin: 0 auto -60px;
        }
        #push, #footer {
            height: 60px;
        }
        #footer {
            background-color: #f5f5f5;
        }
        @media (max-width: 767px) {
            #footer {
                margin-left: -20px;
                margin-right: -20px;
                padding-left: 20px;
                padding-right: 20px;
            }
        }
        .container {
            width: auto;
            max-width: 680px;
        }
        .container .credit {
            margin: 20px 0;
        }
    </style>
  </head>

  <body>
    <div id="wrap">
      <div class="container"> 
        <div class="row">
          <div class="span4"></div>
          <div class="span4">
            <h3 style="text-align:center">chatting as '{{.User}}'</h3>
            <p style="text-align:center"><input id='input' type = "text"></p>
            <small style="text-align:center">Hit enter to send a message.</small>
            <h3 style="text-align:center">Messages</h3>
            <div id="messages"></div>
          </div>
          <div class="span4"></div>
        </div>
      </div>
    </div>
  </body>
</html>
`))
