const button_Mobile = document.getElementById('checkbox-menu');

var messages = document.getElementById("messages");
var mb = document.getElementById("messageBox");
var btn = document.getElementById("submit");


let ip 
fetch("http://"+ document.location.host + "/api/myip").then((response) => response.json()).then(data => ip = data);
function toggle_Menu(event) {
  const nav = document.getElementById('nav');
  nav.classList.toggle('active');  

}


function postMessage(msg, user) {
	var li = document.createElement("li");
	var span = document.createElement("span");
	var m = document.createTextNode(msg);
	var u = document.createTextNode(user);
	messages.appendChild(li);
	li.appendChild(span);
	span.appendChild(u);
	li.appendChild(m);
    return {li , span, m, u};
}

/*
btn.addEventListener("click", function() {
	if (mb.value == "") {
 		alert("you did't type a message!");
 	} else {
 		postMessage(mb.value, "Eduadro137");
 		mb.value = null;
	}
});
*/

/*
mb.addEventListener("keyup", function(event) {
	event.preventDefault();
	if (event.keyCode == 13) {
		btn.click();
	}
});
*/

function getIPandMSG (msg) {
    let ip , str;
    for (var count = 1; count < msg.length; count ++ ){
        if(msg[count] == ":" ) {
        ip = msg.substr(1, count - 1);
        str = msg.substr(count + 2, msg.length);
        return {ip , str};
        }
    }
}

window.onload = function () {
  var conn;
  var msg = document.getElementById("inputs");
  var log = document.getElementById("messages");
  function appendLog(item) {
    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
        if (!doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    mb.addEventListener("keyup", function(event) {
        event.preventDefault();
        if (event.keyCode == 13) {
            btn.click();
        }
    });

  if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        
        conn.onclose = function (evt) {
            
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };

        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                let mens = getIPandMSG(messages[i]);
                let user = function(){return (mens.ip == ip.ipv4 ) ? "Você" : mens.ip};
                let item = postMessage(mens.str, user(mens));
                //item.li.innerText = messages[i];
                appendLog(item.li);
            }
        };
        btn.addEventListener("click", function() {
            if (mb.value == "") {
                alert("Você não digitou a menssagem!");
             } else {
                conn.send(mb.value);              
                mb.value = null;
            }
        });
    }else{
            var item = document.createElement("div");
            item.innerHTML = "<b>Seu navegador não suporta WebSockets.</b>";
            appendLog(item);
        }
};

button_Mobile.addEventListener('click', toggle_Menu);
button_Mobile.addEventListener('touchstart', toggle_Menu);

