var sub = new WebSocket("ws://localhost:8000/subscribe");
var pub = new WebSocket("ws://localhost:8000/publish");

listen()

function appendHistory( equation ){

    var p = document.createElement("p");
    var node = document.createTextNode(equation);
    p.appendChild(node);
    document.getElementById("serverhistory").appendChild(p);
}

function clearHistory(){
    document.getElementById("serverhistory").innerHTML = "";
}

function listen(){

    //expects json like: 
    // calculations { 
    //      numCalculations : int  
    //      calculations : [
    //          {
    //              equation : string 
    //              user : string 
    //              time : time string 
    //          }
    //      ]
    // }

    sub.onmessage = function( msg ){

        clearHistory();
        var response = JSON.parse( msg.data );
        for( let i = response.calculations.length -1; i >= 0; i-- ){
            appendHistory( response.calculations[i]['equation']);
        }
    }
}


function publish(equation) {
    var req = { equation: equation, user: ""};
    pub.send( JSON.stringify(req));
}

function getHistory(){
	return document.getElementById("history-value").innerText;
}
function printHistory(num){
	document.getElementById("history-value").innerText=num;
}
function getOutput(){
	return document.getElementById("output-value").innerText;
}
function printOutput(num){
	if(num==""){
		document.getElementById("output-value").innerText=num;
	}
	else{
		document.getElementById("output-value").innerText=getFormattedNumber(num);
	}	
}
function getFormattedNumber(num){

	return num;
}
function reverseNumberFormat(num){
	//return Number(num.replace(/,/g,''));
    return num
}
var operator = document.getElementsByClassName("operator");
for(var i =0;i<operator.length;i++){
	operator[i].addEventListener('click',function(){
		if(this.id=="clear"){
			printHistory("");
			printOutput("");
		}
		else if(this.id=="backspace"){
			var output=reverseNumberFormat(getOutput()).toString();
			if(output){//if output has a value
				output= output.substr(0,output.length-1);
				printOutput(output);
			}
		}
        else if( this.id=='-' && getOutput() == '' ){
            printOutput('-')
        }
        else if( this.id=='-' && getOutput() == '-') {
            return 
        }
		else{
			var output=getOutput();
			var history=getHistory();
			if(output==""&&history!=""){
				if(isNaN(history[history.length-1])){
					history= history.substr(0,history.length-1);
				}
			}
			if(output!="" || history!=""){
				output= output==""?output:reverseNumberFormat(output);
				history=history+output;
				if(this.id=="="){
					var result= math.format( math.evaluate(history), 14 );
					printOutput(result);
					printHistory("");
                    publish(history + "=" + result); 
				}
				else{
					history=history+this.id;
					printHistory(history);
					printOutput("");
				}
			}
		}
		
	});
}
var number = document.getElementsByClassName("number");
for(var i =0;i<number.length;i++){
	number[i].addEventListener('click',function(){

        var output=getOutput();
		
        if( !isNaN(output + this.id)){
            console.log( output + this.id )
            output = output+this.id;
            printOutput(output)

        } 
	});
}
