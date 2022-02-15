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

  // const process = document.querySelector('.btn-success');
  setInterval(() => {
    if (x == 0.0017950000000000043) {

    }
  }, 500);

}
run();
const processBtn = document.getElementsByClassName('process')[0];
const progress = document.querySelector('.progress-done');
setTimeout(() => {
  progress.style.opacity = 1;
  progress.style.width = progress.getAttribute('data-done') + '%';
  if (progress.style.width !== '100%') {
    processBtn.style.visibility = "hidden";
  }
}, 200);

setTimeout(() => {
  processBtn.style.visibility = "visible";
}, 10000 * 2.5);



const uvindex_div = document.getElementsByClassName("uv-div")[0];
const linediv = document.getElementsByClassName("risk-div")[0];
const uv_field = document.getElementsByClassName("uvindex")[0];
const sunburn = document.getElementsByClassName("fa-sun-o")[0];
uvindex_div.children[0].addEventListener('click', function(){
  uv_field.style.visibility = 'visible';
  console.log("uvindex:", uv_field);
   if (uvindex_div.children[0].children[1].innerHTML.includes("[0 2.9]")){
     uvindex_div.children[0].children[0].style.color = "green";
     linediv.children[0].style.backgroundColor = "green";
     uvindex_div.children[0].innerHTML = "Risk_Level :Low";
     uvindex_div.children[0].className = "btn btn-success btn-feature";
     uvindex_div.children[0].style.marginLeft = "-46px";
     sunburn.style.color = 'green';
     uvindex_div.children[0].disabled = true;
   }else if(uvindex_div.children[0].children[1].innerHTML.includes("[3 5.9]")){
     uvindex_div.children[0].children[0].style.color = "yellow";
     linediv.children[0].style.backgroundColor = "yellow";
     uvindex_div.children[0].innerHTML = "Risk_Level :Moderate";
     uvindex_div.children[0].className = "btn btn-warning btn-feature";
     uvindex_div.children[0].style.marginLeft = "2px";
     sunburn.style.color = 'yellow';
     uvindex_div.children[0].disabled = true;
   }else if (uvindex_div.children[0].children[1].innerHTML.includes("[6 7.9]")) {
    uvindex_div.children[0].children[0].style.color = "orange";
    linediv.children[0].style.backgroundColor = "orange";
    uvindex_div.children[0].innerHTML = "Risk_Level :High";
    uvindex_div.children[0].className = "btn btn-warning btn-feature";
    uvindex_div.children[0].style.marginLeft = "2px";
    sunburn.style.color = 'orange';
    uvindex_div.children[0].disabled = true;
   }else if (uvindex_div.children[0].children[1].innerHTML.includes("[6 7.9]")) {
    uvindex_div.children[0].children[0].style.color = "red";
    linediv.children[0].style.backgroundColor = "red";
    uvindex_div.children[0].innerHTML = "Risk_Level :Very high";
    uvindex_div.children[0].className = "btn btn-danger btn-feature";
    uvindex_div.children[0].style.marginLeft = "2px";
    sunburn.style.color = 'red';
    uvindex_div.children[0].disabled = true;
   }else if (uvindex_div.children[0].children[1].innerHTML.includes("[11]")) {
    uvindex_div.children[0].children[0].style.color = "violet";
    linediv.children[0].style.backgroundColor = "violet";
    uvindex_div.children[0].innerHTML = "Risk_Level :Extreme";
    uvindex_div.children[0].className = "btn btn-danger btn-feature";
    uvindex_div.children[0].style.marginLeft = "2px";
    sunburn.style.color = 'violet';
    uvindex_div.children[0].disabled = true;
   }
   
});