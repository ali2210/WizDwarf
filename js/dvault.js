// credentilals object
const options = {
  key : "65993b3c66b5317411a5",
  cluster :"mt1",
};

// entries for web cache ... This is used for further processing..... 
const key_entries = [];
var values_entries = [];
//let count;

// pusher channels api object declaration 
const pusher = new Pusher(options.key, {cluster : options.cluster});

// weblocal "browser cache"
const weblocal = window.localStorage;


//@  Read all channels; if channel name is not specified then throw error 

pusher.allChannels().forEach((channels) =>{

  if (channels.name !== "keygen"){
    alert("The channel is not encrypted and not ready to be sent!", channel.name);
    return
  }
});

// Subcribe the channel and open for communication 
const channel = pusher.subscribe("keygen");


// Afterthat bind data for internal with specified channel 

channel.bind("pusher_internal:subscription_succeeded", (data) => {
  
  console.log("internal subscription executed ", data);

});

// Read that connection still open for further data exchange
pusher.connection.bind('state_change', (state) =>{
  
  if (state.current === "connected"){
    console.log("Connection connected ....", state.current);
  }

});

// @ If the channel is not connected then throw error   
channel.bind("pusher:subscription_error", (error) => {

  var { status } = error;
 
  if (status == 408 || status == 503) {
    alert("Channel disconnected", status);
    return
  }
});


// Dara bind with channel ; 

channel.bind("tnxs", (data) => {


        // message renderer on console
    console.log("Data: ... ", JSON.stringify(data));

      // read the data from the channel; data exist in map ....
    const stream = JSON.stringify(data,(_, value)=>
        (value instanceof Map? Array.from(value.entries()): value));

    const parser = JSON.parse(stream);

      // values , keys  store in web cache 
    const _ = JSON.stringify(parser, (_, value)=>{


        for (let i = 0; i < value["values"].length; i++){

            values_entries.push(value["values"][i]);
            weblocal.setItem(i, JSON.stringify(values_entries[i]));                
        }

        key_entries.push(value["keys"]);

        weblocal.setItem("key", JSON.stringify(key_entries[0]));
        //count = value["values"].length;    

    });


    console.log("entries: ...", key_entries[0], "values : ... ",...values_entries);  

});

let tObjects = [];


for (let i =0; i < weblocal.length; i++){
  
    if (weblocal.getItem(weblocal.length) !== "pusherTransportTLS" && weblocal.getItem(i) !== null){
      
      
         switch (i){

          case 0:

              JSON.stringify(JSON.parse(weblocal.getItem(i)), (_, values) => {
                
                if (JSON.parse(weblocal.getItem(i) !== values["CDR_LINK"])){
                  
                  document.getElementsByClassName('info-content-1')[0].children[0].innerHTML = values["CDR_LINK"];
                  document.getElementsByClassName('info-content-1')[0].children[1].innerHTML = values["SizeOf"];
                  document.getElementsByClassName('info-content-1')[0].children[2].innerHTML = values["Access"];
                  document.getElementsByClassName('info-content-1')[0].children[3].style.display = 'unset';
                  
                  if ((values["ImagePath"]).includes(".png")){
                    
                    document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = "/" +values["ImagePath"];

                  }else if ((values["ImagePath"]).includes(".jpeg")){
                    
                    document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = "/"+values["ImagePath"];
                    
                  }else if ((values["ImagePath"]).includes(".gif")){
                    
                    document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = "/"+values["ImagePath"];
                    
                  }

                  tObjects.push(values["Objects"]);            
                  
                  document.getElementsByClassName('info')[0].innerHTML = 'Peer ID : ' + weblocal.getItem("key"); 
                  
                }
                
            });
          break
          
          case 1:
            
              JSON.stringify(JSON.parse(weblocal.getItem(i)), (_, values) => {
            
                
                if (JSON.parse(weblocal.getItem(i) !== values["CDR_LINK"])){
                  
                  document.getElementsByClassName('info-content-2')[0].children[0].innerHTML = values["CDR_LINK"];
                  document.getElementsByClassName('info-content-2')[0].children[1].innerHTML = values["SizeOf"];
                  document.getElementsByClassName('info-content-2')[0].children[2].innerHTML = values["Access"];
                  document.getElementsByClassName('info-content-2')[0].children[3].style.display = 'unset';

                  if (!document.getElementsByClassName('info-content-2')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".png")){
                    
                    document.getElementsByClassName('info-content-2')[0].children[4].children[0].src = "/" +values["ImagePath"];
                    

                  }else if (! document.getElementsByClassName('info-content-2')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".jpeg")){
                    
                    document.getElementsByClassName('info-content-2')[0].children[4].children[0].src = "/"+values["ImagePath"];
                    
                  }else if (! document.getElementsByClassName('info-content-2')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".gif")){
                    
                    document.getElementsByClassName('info-content-2')[0].children[4].children[0].src = "/"+values["ImagePath"];
                    
                  }

                  tObjects.push(values["Objects"]);

                }
              
                
                
          });
          break
        
          case 2: 
        
          JSON.stringify(JSON.parse(weblocal.getItem(i)), (_, values) => {
        
            if (JSON.parse(weblocal.getItem(i) !== values["CDR_LINK"])){
            
              
              document.getElementsByClassName('info-content-3')[0].children[0].innerHTML = values["CDR_LINK"];
            
              document.getElementsByClassName('info-content-3')[0].children[1].innerHTML = values["SizeOf"];
            
              document.getElementsByClassName('info-content-3')[0].children[2].innerHTML = values["Access"];

              document.getElementsByClassName('info-content-3')[0].children[3].style.display = 'unset';
              
              if (! document.getElementsByClassName('info-content-3')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".png")){
                    
                document.getElementsByClassName('info-content-3')[0].children[4].children[0].src = "/" +values["ImagePath"];

              }else if (! document.getElementsByClassName('info-content-3')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".jpeg")){
                
                document.getElementsByClassName('info-content-3')[0].children[4].children[0].src = "/"+values["ImagePath"];
                
              }else if (! document.getElementsByClassName('info-content-3')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".gif")){
                
                document.getElementsByClassName('info-content-3')[0].children[4].children[0].src = "/"+values["ImagePath"];
                
              }

              tObjects.push(values["Objects"]);

            }
            
            
            
        });
        break  
        case 3: 
        
          JSON.stringify(JSON.parse(weblocal.getItem(i)), (_, values) => {
        
            if (JSON.parse(weblocal.getItem(i) !== values["CDR_LINK"])){
            
              
              document.getElementsByClassName('info-content-4')[0].children[0].innerHTML = values["CDR_LINK"];
            
              document.getElementsByClassName('info-content-4')[0].children[1].innerHTML = values["SizeOf"];
            
              document.getElementsByClassName('info-content-4')[0].children[2].innerHTML = values["Access"];
          
              document.getElementsByClassName('info-content-4')[0].children[3].style.display = 'unset';

              if (! document.getElementsByClassName('info-content-4')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".png")){
                    
                document.getElementsByClassName('info-content-4')[0].children[4].children[0].src = "/" +values["ImagePath"];

              }else if (! document.getElementsByClassName('info-content-4')[0].children[0].innerHTML === " " &&(values["ImagePath"]).includes(".jpeg")){
                
                document.getElementsByClassName('info-content-4')[0].children[4].children[0].src = "/"+values["ImagePath"];
                
              }else if (! document.getElementsByClassName('info-content-4')[0].children[0].innerHTML === " " &&(values["ImagePath"]).includes(".gif")){
                
                document.getElementsByClassName('info-content-4')[0].children[4].children[0].src = "/"+values["ImagePath"];
                
              }

              tObjects.push(values["Objects"]);
            }
            
        });
        break
        case 4: 
        
          JSON.stringify(JSON.parse(weblocal.getItem(i)), (_, values) => {
        
            if (JSON.parse(weblocal.getItem(i) !== values["CDR_LINK"])){
            
              document.getElementsByClassName('info-content-5')[0].children[0].innerHTML = values["CDR_LINK"];
            
              document.getElementsByClassName('info-content-5')[0].children[1].innerHTML = values["SizeOf"];
            
              document.getElementsByClassName('info-content-5')[0].children[2].innerHTML = values["Access"];
          
              document.getElementsByClassName('info-content-5')[0].children[3].style.display = 'unset';

              if (! document.getElementsByClassName('info-content-5')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".png")){
                    
                document.getElementsByClassName('info-content-5')[0].children[4].children[0].src = "/" +values["ImagePath"];

              }else if (!  document.getElementsByClassName('info-content-5')[0].children[0].innerHTML === " " &&(values["ImagePath"]).includes(".jpeg")){
                
                document.getElementsByClassName('info-content-5')[0].children[4].children[0].src = "/"+values["ImagePath"];
                
              }else if (! document.getElementsByClassName('info-content-5')[0].children[0].innerHTML === " " && (values["ImagePath"]).includes(".gif")){
                
                document.getElementsByClassName('info-content-5')[0].children[4].children[0].src = "/"+values["ImagePath"];
                
              }

              tObjects.push(values["Objects"]);
            }
            
            
        });
        break
      }
    }
        
  }
  
  
  document.getElementsByClassName('objects')[0].children[0].innerHTML = Math.max(...tObjects) + " Decentralized Objects";
  document.getElementsByClassName('objects')[0].children[2].innerHTML = Math.max(...tObjects) -5 <= 0 ? 0 : Math.max(...tObjects) -5;
