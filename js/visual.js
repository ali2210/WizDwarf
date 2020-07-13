


var circlesNum = 100;
var circlesArr = [];

for (var i = 0; i < circlesNum; i++) {
		var circlePath = new Path.Circle({
			fillColor : 'red',
			center : [i * circlesNum , circlesNum],
			radius : 50
	});

	circlesArr.push(circlePath);
}	

function onMouseMove(event){

	for (var i = 0; i < circlesArr.length; i++) {
		circlesArr[i].position = event.point;
		circlesArr[i].fullySelected = true;
	}

}

function onMouseDown(event){
	
	for (var i = 0; i < circlesArr.length; i++) {
		circlesArr[i].fullySelected = false;
		circlesArr[i].position = event.point;
	}

}

function onDoubleClick(event){
			/*var text = new PointText(new Point(30, 30));
			text.fillColor = 'blue';
			text.content = 'Press shift... Clear Screen';
	*/
	console.log("Double:" + event.point);
}


