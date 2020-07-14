


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
	var text = new PointText(new Point(500, 30));
	text.fillColor = 'blue';
	text.content = '';
	for (var i = 0; i < circlesArr.length; i++) {
		circlesArr[i].fullySelected = false;
		circlesArr[i].position = event.point;
	}

}

function onMouseUp(event){
	var text = new PointText(new Point(500, 30));
	text.fillColor = 'blue';
	text.content = 'Press shift...';
	
	if (Key.isDown('shift')) {

		for (var i = 0; i < circlesArr.length; i++) {
			circlesArr[i].position = [0,0];
			circlesArr[i].remove();
			text.content = '';
		}
		console.log("Fine..." );
		setTimeout(reDraw(),7000);
	}
	

}

function reDraw() {
	// body...
	for (var i = 0; i < circlesNum; i++) {
			var circlePath = new Path.Circle({
				fillColor : 'red',
				center : [i * circlesNum , circlesNum],
				radius : 50
			});

			circlesArr.push(circlePath);
			
		}
}


