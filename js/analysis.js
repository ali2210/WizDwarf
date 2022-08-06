const _b01 = document.getElementsByClassName('matices')[0];
const _d01 = document.getElementsByClassName('visual')[0];
const _mn01 = document.getElementsByClassName('menu-div')[0];

// const _sp01 = document.getElementsByClassName('span-item')[0];

const options = {
  // app_id : "1265511",
  key : "65993b3c66b5317411a5",
  // secret : "4f8bf3faf121d9c8dadf",
  cluster : "mt1",
};

 // channel created and get data for futher processing
 const pusher = new Pusher(options.key,{cluster: options.cluster});

//  create localStorage Object
 const web_local = window.localStorage;

 let pi_streaming = [];
 var onetime = false;

  // craete channel log event
  // Pusher.logToConsole = true;
  pusher.allChannels().forEach((channels) =>{
    if (channels.name !== "protein"){
      alert("The channel is not encrypted and not ready to be sent!", channel.name);
      return
    }
  });

  // channel bind with subcriber wait for trigger event 
  const channel = pusher.subscribe("protein");
  channel.bind("pusher_internal:subscription_succeeded", (data) => {
      console.log("internal subscription executed ", data);
  });   

  // channel state changed
  pusher.connection.bind('state_change', (state) =>{
    _d01.style.backgroundColor = "transparent";

    if (state.current === "connected" && !onetime){
      reloadPage();
      console.log("pusher channel status:", state.current, "flag:", onetime);
    }

    onetime = true;
    
    // _sp01.innerHTML = "Reload the webpage " + (new Date().getTime() / 1000000000000 );
  });

  // channel subcribe error 
  channel.bind("pusher:subscription_error", (error) => {
    var { status } = error;
    if (status == 408 || status == 503) {
      alert("Channel disconnected", status);
      return
    }
  });

  function reloadPage() {
    var currentDocumentTimestamp = new Date(performance.timing.domLoading).getTime();
    // Current Time //
    var now = Date.now();
    // Total Process Lenght as Minutes //
    var tenSec = 10 * 1000;
    // End Time of Process //
    var plusTenSec = currentDocumentTimestamp + tenSec;
    if (now > plusTenSec) {
        window.location.reload();
    }

    
  }

  // read stream from channel
  channel.bind("molecule", (data) => {

    //alert("Event initating:" );
    // debugger mod
    console.log("data :", JSON.stringify(data));
    
    // Map object 
    const dataTensor = new Map();

    // object intialization
    for (let i in data) {
           dataTensor.set(data[i].key, data[i].values);  
    }

    console.log("tensor:", dataTensor.entries().next().value);

    // y values against input 

    for (let index in data) {
     if(dataTensor.has(data[index].key)) {
        pi_streaming.push(dataTensor.get(data[index].key));
      }
    }
    
    // if y tensor hold 0 value then throw exception
    if(pi_streaming.length<=0){
      alert("data stream should not empty");
      return
    }

    for(let index in pi_streaming){
      web_local.setItem(index, JSON.stringify(pi_streaming[index]));
    }
  
  }); 


  // const progressCircle = document.getElementsByClassName('progress-container')[0];
_b01.addEventListener('click', (event) => {

  if (_mn01.style.visibility === 'hidden') {
    _mn01.style.visibility = 'visible';
  }

  // chart metadata 
  let context2d = document.getElementById('chart').getContext('2d');

  let getY_value = [];
    
  // read local-cache and sore back 
    for (let index = 0; index < 300; index++) {
      if(!JSON.parse(web_local.getItem(index) <= 0)){
        getY_value.push(JSON.parse(web_local.getItem(index)));
      }
    }

    // _sp01.style.visibility = 'hidden';

    
  // create new chart with channel data 
  new Chart(context2d, {
    type: 'bar',
    data: {
      labels: ['F', 'L', 'I', 'M', 'V', 'S', 'P', 'T','A', 'Y','!','!*','H','Q','N','K','D','E','C','!**','R','G',],
      datasets: [{
          label : 'Proteins Occurrence',
          data: [...getY_value],
          backgroundColor: [
              'rgba(255, 99, 132, 0.2)',
              'rgba(54, 162, 235, 0.2)',
              'rgba(255, 206, 86, 0.2)',
              'rgba(75, 192, 192, 0.2)',
              'rgba(153, 102, 255, 0.2)',
              'rgba(255, 159, 64, 0.2)',
              'rgba(94, 222, 275, 0.2)',
              'rgba(255, 220, 42, 0.2)',
              'rgba(105, 82, 188, 0.2)',
              'rgba(19, 235, 230, 0.2)',
              'rgba(255, 4, 0, 0.2)',
              'rgba(155, 314, 335, 0.2)',
              'rgba(255,117,218, 0.2)',
              'rgba(7,58,140,0.2)',
              'rgba(255, 61, 3, 0.2)',
              'rgba(0, 255, 0, 0.2)',
              'rgba(255, 97, 25, 0.2)',
              'rgba(0, 128, 255, 0.2)',
              'rgba(255, 255, 250, 0.2)',
              'rgba(102, 54, 128, 0.2)',
              'rgba(123, 251, 236, 0.2)',
              'rgba(102, 221, 128, 0.2)',
              'rgba(255, 99, 132, 0.2)',
          ],
          borderColor: [
              'rgba(255, 99, 132, 1)',
              'rgba(54, 162, 235, 1)',
              'rgba(255, 206, 86, 1)',
              'rgba(75, 192, 192, 1)',
              'rgba(153, 102, 255, 1)',
              'rgba(255, 159, 64, 1)',
              'rgba(94, 222, 275, 1)',
              'rgba(255, 220, 42, 1)',
              'rgba(105, 82, 188, 1)',
              'rgba(19, 235, 230, 1)',
              'rgba(255, 4, 0, 1)',
              'rgba(155, 314, 335, 1)',
              'rgba(255, 61, 3, 1)',
              'rgba(7,58,140,1)',
              'rgba(255,117,218, 1)',
              'rgba(0, 255, 0, 1)',
              'rgba(255, 97, 25, 1)',
              'rgba(0, 128, 255, 1)',
              'rgba(255, 255, 250,1)',
              'rgba(102, 54, 128, 1)',
              'rgba(123, 251, 236, 1)',
              'rgba(102, 221, 128, 1)',
              'rgba(255, 99, 132, 1)',
              'rgba(54, 162, 235, 1)',
          ],
          borderWidth: 1
      }]
  },
  options: {
      scales: {
          y: {
              beginAtZero: true
          }
      }
  }
});

  // delete cache 
  web_local.clear();
  pusher.unbind();
  console.log("Special data stream closed & cache now free ");
  reloadPage();
});

// generate visual graph cubic Interpolation (line with points). 
document.getElementsByClassName('line-chart')[0].addEventListener('click', (event)=>{
  
  if (_mn01.style.visibility === 'hidden') {
    _mn01.style.visibility = 'visible';
  }

  // read canvas over web canvas
  let ctx = document.getElementById('chart').getContext('2d');
  
  // data stream hold channel data;
  let data_stream = [];
  
  // read local-cache and sore back 
  for (let index = 0; index < 300; index++) {
    
    if(!JSON.parse(web_local.getItem(index) <= 0)){
      
      data_stream.push(JSON.parse(web_local.getItem(index)));
    
    }
  }

  // data processing
  const data = {
    // labels contains proteins name 
    labels : ['F', 'L', 'I', 'M', 'V', 'S', 'P', 'T','A', 'Y','!','!*','H','Q','N','K','D','E','C','!**','R','G',],
   
    // datasets contain data about proteins
    datasets : [
      { 
        label : 'Proteins Occurrence',
        data : [...data_stream],
        borderColor : 'rgba(0, 255, 255,1)',
        fill: false,
        cubicInterpolationMode: 'monotone',
        tension: 0.4,
      }
    ],
  }

  // create new chart with default settings
  new Chart(ctx, {
    type : 'line',
    data: data,
    option : {
      responsive : true,
      plugins :{
        title : {
          display : true,
          text : 'Proteins Occurrence Line Visualization',
        },
      },
      intersection : {
        intersect : false,
      },
      scales :{
        x : {
          display : true,
          title : {
            display : true,
          }
        },

        y : {
          display : true,
          title : {
            display : true,
            text : 'Occurrence',
          }
        },

        suggestedMin: 1,
        suggestedMax: 100,
      },
    }
  })


  // clear the storage  
  web_local.clear();
  pusher.unbind();
  console.log("Special data stream closed & cache now free ");
  reloadPage();
});


  