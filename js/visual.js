window.requestAnimationFrame(meter);

function meter() {
  // read canvas
  var canvas = document.getElementById('canvasPaper');
  var ctx = canvas.getContext('2d');
  ctx.save();

  ctx.clearRect(0, 0, 150, 150);
  ctx.translate(75, 75);
  ctx.scale(0.4, 0.4);
  ctx.rotate(-Math.PI / 2);
  ctx.strokeStyle = 'black';
  ctx.fillStyle = 'white';
  ctx.lineWidth = 8;
  ctx.lineCap = 'round';

  ctx.save();
  for (var i = 0; i < 100; i++) {
    ctx.beginPath();
    ctx.rotate(Math.PI / 50);
    ctx.moveTo(100, 0);
    ctx.lineTo(120, 0);
    ctx.stroke();
  }
  ctx.restore();



  var res = document.getElementById('result').textContent;;
  var valueRes = parseFloat(res);

  // console.log(valueRes);

  ctx.fillStyle = 'green';
  ctx.save();
  ctx.rotate((Math.PI * valueRes) / (Math.PI));
  ctx.lineWidth = 15;
  ctx.beginPath();
  ctx.moveTo(-20, 0);
  ctx.lineTo(80, 0);
  ctx.strokeStyle = 'red';
  ctx.arc(0, 0, 10, 0, 2 * Math.PI, true);
  ctx.lineTo(100, 0);
  ctx.stroke();
  //ctx.restore();

  //second hand
  // ctx.save();
  // ctx.rotate((Math.PI/ 360) * valueRes);
  // ctx.lineWidth = 10;
  // ctx.beginPath();
  // ctx.moveTo(-28,0);
  // ctx.lineTo(112,0);
  // ctx.stroke();
  // ctx.restore();


  ctx.beginPath();
  ctx.lineWidth = 14;
  ctx.strokeStyle = '#325FA2';
  ctx.arc(0, 0, 142, 0, 2 * Math.PI, true);
  ctx.stroke();
  ctx.restore();

  //window.requestAnimationFrame(meter);
}

var t = 0;
var c = document.querySelector("canvas");
var $ = c.getContext('2d');
c.width = window.innerWidth;
c.height = window.innerHeight;
$.fillStyle = 'hsla(0,0%,0%,1)';

window.addEventListener('resize', function () {
  c.width = window.innerWidth;
  c.height = window.innerHeight;
}, false);

function draw() {
  $.globalCompositeOperation = 'source-over';
  $.fillStyle = 'hsla(0,0%,0%,.1)';
  $.fillRect(0, 0, c.width, c.height);
  var foo, i, j, r;
  foo = Math.sin(t) * 2 * Math.PI;
  for (i = 0; i < 400; ++i) {
    r = 400 * Math.sin(i * foo);
    $.globalCompositeOperation = '';
    $.fillStyle = 'hsla(' + i + 12 + ',100%, 60%,1)';
    $.beginPath();
    $.arc(Math.sin(i) * r + (c.width / 2),
      Math.cos(i) * r + (c.height / 2),
      1.5, 0, Math.PI * 2);
    $.fill();

  }
  t += 0.000005;
  return t %= 2 * Math.PI;

};

function run() {
  window.requestAnimationFrame(run);
  x = draw();

  const process = document.querySelector('.btn-success');
  const div = document.querySelector('.canvasdiv');
  const transactBtn = document.querySelector('.lnkBtn');
  setInterval(() => {
    if (x == 0.0017950000000000043) {
      process.style.visibility = "visible";
      div.style.visibility = "visible";
      transactBtn.style.visibility = "visible";
    }
  }, 500);




}
run();

const progress = document.querySelector('.progress-done');
setTimeout(() => {
  progress.style.opacity = 1;
  progress.style.width = progress.getAttribute('data-done') + '%';
}, 200);


