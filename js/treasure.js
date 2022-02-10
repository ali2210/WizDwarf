const _m01 = document.getElementsByClassName('menu-div')[0];

const _sp01 = document.getElementsByClassName('span-item')[0];

const options = {
  // app_id : "1265511",
  key : "65993b3c66b5317411a5",
  // secret : "4f8bf3faf121d9c8dadf",
  cluster : "mt1",
};

 // channel created and get data for futher processing
 const pusher = new Pusher(options.key,{cluster: options.cluster});

  // craete channel log event
  // Pusher.logToConsole = true;
  pusher.allChannels().forEach((channels) =>{
    if (channels.name !== "protein"){

      alert("The channel is not encrypted and not ready to be sent!", channel.name);
    }
  });

  // channel bind with subcriber wait for trigger event 
  const channel = pusher.subscribe("protein");
  channel.bind("pusher_internal:subscription_succeeded", (data) => {
      console.log("internal subscription_succeeded", data);
  });   

  // channel state changed
  pusher.connection.bind('state_change', (state) =>{
    console.log("current state", state.current);
  });

  // channel subcribe error 
  channel.bind("pusher:subscription_error", (error) => {
    var { status } = error;
    if (status == 408 || status == 503) {
      alert("Channel disconnected", status);
    }
  });

  // read stream from channel
  channel.bind("molecule", (data) => {

    // debugger mod
    console.log(JSON.stringify(data));
    
    // Map object 
    const dataTensor = new Map();

    // object intialization
    for (let i in data) {
           dataTensor.set(data[i].key, data[i].values);  
    }

    console.log("tensor:", dataTensor.entries().next().value);

    // y values against input 
    let pi_streaming = [];

    for (let index in data) {
     if(dataTensor.has(data[index].key)) {
        pi_streaming.push(dataTensor.get(data[index].key));
      }
    }
    
    // if y tensor hold 0 value then throw exception
    if(pi_streaming.length<=0){
      alert("data stream should not empty");
    }

    alert("data stream");
    for (let index in pi_streaming){
      console.log("Y value:", pi_streaming[index]);
    }

    const web_local = window.localStorage;

    for(let index in pi_streaming){
      web_local.setItem(index, JSON.stringify(pi_streaming[index]));
    }

    let getY_value = [];

    for (let index in pi_streaming){
      getY_value = JSON.parse(web_local.getItem(index));
      console.log(getY_value);
    }
  
  }); 

_m01.addEventListener('click', (event) => {
  
  //  data values 
  // let pi_streaming = [];
 
  // chart metadata 
  let context2d = document.getElementById('chart').getContext('2d');


  // create new chart with channel data 
//  new Chart(context2d, {
//     type: 'bar',
//     data: {
//       labels: ['F', 'L', 'I', 'M', 'V', 'S', 'P', 'T','A', 'Y','!','!*','H','Q','N','K','D','E','C','!**','R','G',],
//       datasets: [{
//           label : 'protein per aminochain',
//           data: [...pi_streaming],
//           backgroundColor: [
//               'rgba(255, 99, 132, 0.2)',
//               'rgba(54, 162, 235, 0.2)',
//               'rgba(255, 206, 86, 0.2)',
//               'rgba(75, 192, 192, 0.2)',
//               'rgba(153, 102, 255, 0.2)',
//               'rgba(255, 159, 64, 0.2)',
//               'rgba(94, 222, 275, 0.2)',
//               'rgba(255, 220, 42, 0.2)',
//               'rgba(105, 82, 188, 0.2)',
//               'rgba(19, 235, 230, 0.2)',
//               'rgba(255, 4, 0, 0.2)',
//               'rgba(155, 314, 335, 0.2)',
//               'rgba(255,117,218, 0.2)',
//               'rgba(7,58,140,0.2)',
//               'rgba(255, 61, 3, 0.2)',
//               'rgba(0, 255, 0, 0.2)',
//               'rgba(255, 97, 25, 0.2)',
//               'rgba(0, 128, 255, 0.2)',
//               'rgba(255, 255, 250, 0.2)',
//               'rgba(102, 54, 128, 0.2)',
//               'rgba(123, 251, 236, 0.2)',
//               'rgba(102, 221, 128, 0.2)',
              
//           ],
//           borderColor: [
//               'rgba(255, 99, 132, 1)',
//               'rgba(54, 162, 235, 1)',
//               'rgba(255, 206, 86, 1)',
//               'rgba(75, 192, 192, 1)',
//               'rgba(153, 102, 255, 1)',
//               'rgba(255, 159, 64, 1)',
//               'rgba(94, 222, 275, 1)',
//               'rgba(255, 220, 42, 1)',
//               'rgba(105, 82, 188, 1)',
//               'rgba(19, 235, 230, 1)',
//               'rgba(255, 4, 0, 1)',
//               'rgba(155, 314, 335, 1)',
//               'rgba(255, 61, 3, 1)',
//               'rgba(7,58,140,1)',
//               'rgba(255,117,218, 1)',
//               'rgba(0, 255, 0, 1)',
//               'rgba(255, 97, 25, 1)',
//               'rgba(0, 128, 255, 1)',
//               'rgba(255, 255, 250,1)',
//               'rgba(102, 54, 128, 1)',
//               'rgba(123, 251, 236, 1)',
//               'rgba(102, 221, 128, 1)',
//           ],
//           borderWidth: 1
//       }]
//   },
//   options: {
//       scales: {
//           y: {
//               beginAtZero: true
//           }
//       }
//   }
// });
  

});

  // console.log(max_scale(model));
  // console.log(min_scale(model));
// });

// function max_scale(data){return Math.max(...data);}
// function min_scale(data){return Math.min(...data);}