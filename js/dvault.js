const options = {
    key : "65993b3c66b5317411a5",
    cluster :"mt1",
 };

 var key_entries = [];
 var values_entries = [];
 var count = 0;
  
  const pusher = new Pusher(options.key, {cluster : options.cluster});
  const weblocal = window.localStorage;


  pusher.allChannels().forEach((channels) =>{

    if (channels.name !== "keygen"){
      alert("The channel is not encrypted and not ready to be sent!", channel.name);
      return
    }
  });
  
  const channel = pusher.subscribe("keygen");

  channel.bind("pusher_internal:subscription_succeeded", (data) => {
    
    console.log("internal subscription executed ", data);
  
  });

  // channel state changed
  pusher.connection.bind('state_change', (state) =>{
    
    if (state.current === "connected"){
      console.log("Connection connected ....", state.current);
    }
  
  });

  // channel subcribe error 
  channel.bind("pusher:subscription_error", (error) => {

    var { status } = error;
    
    if (status == 408 || status == 503) {
      alert("Channel disconnected", status);
      return
    }
});

channel.bind("tnxs", (data) => {
  
  console.log("Data:", JSON.stringify(data));
  const stream = JSON.stringify(data,(key, value)=>
      (value instanceof Map? Array.from(value.entries()): value));
  
  const parser = JSON.parse(stream);

  for (let i = 0; i < 6; i++){
   const parse_arr = parser["values"][i] instanceof Array ? Array.from(JSON.parse(parser["values"][i])): parser["values"][i];
   if (parse_arr !== null){
      Object.entries(parse_arr).forEach(([key, value])=>{
            
        if (value !== " "){
          key_entries.push(key);
          values_entries.push(value);
        }
      });

      console.log("Output:", ...key_entries, "Value:", ...values_entries);
      weblocal.setItem(count, JSON.stringify(values_entries));
   }
  }
});

count = 0;
console.log("Web cache:", weblocal.getItem(count));
const obj = [];

obj.push(weblocal.getItem(count));

if (Array.from(JSON.parse(obj[0]))[0] !== ''){
    document.getElementsByClassName('info')[0].innerHTML = 'Peer ID :' + Array.from(JSON.parse(obj[0]))[0];
    document.getElementsByClassName('info-content-1')[0].children[0].innerHTML = Array.from(JSON.parse(obj[0]))[1];
    document.getElementsByClassName('info-content-1')[0].children[1].innerHTML = Array.from(JSON.parse(obj[0]))[4];
    document.getElementsByClassName('info-content-1')[0].children[2].innerHTML = Array.from(JSON.parse(obj[0]))[5];
    if (Array.from(obj[0])[3]){
      document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = JSON.parse(obj[0])[2];
    }else{
      document.getElementsByClassName('info-content-1')[0].children[4].children[1].innerHTML = '';
    }
}
