// credentilals object
const options = {
  key : "65993b3c66b5317411a5",
  cluster :"mt1",
};

// constant message 
const message = "File Editor added in upcomming build v-0.1.2.b ";

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

// @ declaration of an object array
let tObjects = [];

const arr = [];
let filterObj = [];


// @arr hold web cache after JSON object
for (let i =0; i < weblocal.length; i++){
  
  arr.push(JSON.parse(weblocal.getItem(i)));
  
}


// @ In case array value is neither null nor undefined, then apply filter method.
// @ Filter method retreive object states according to specific rule. 
// @ then apply map method which change into map 
arr.forEach(value =>{

  if (value !== null){ 
    filterObj = arr.filter(PopUnusedObjects).map(object => {
      return object
    })
  }
});



      // @ keygen manger  
for (let i =0; i < filterObj.length; i++){
      
  switch (i){

          case 0:

              // content details 
                    document.getElementsByClassName('info-content-1')[0].children[0].innerHTML = filterObj[i].CDR_LINK;
                  
                    document.getElementsByClassName('info-content-1')[0].children[1].innerHTML = filterObj[i].SizeOf;
                  
                    document.getElementsByClassName('info-content-1')[0].children[2].innerHTML = filterObj[i].Access;
                  
                    document.getElementsByClassName('info-content-1')[0].children[3].style.display = 'unset';
                  
              //  image case
                    if ((filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".png")){
                       
                      document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = "/" +filterObj[i].ImagePath;
                  
                    }else if ((filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".jpeg")){
                    
                      document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                    
                    }else if (((filterObj[i].ImagePath !== undefined) && filterObj[i].ImagePath).includes(".gif")){
                    
                      document.getElementsByClassName('info-content-1')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                  
                    }else{
                      document.getElementsByClassName('info-content-1')[0].children[3].style.display = "none";
                    }


                    tObjects.push(filterObj[i].Objects);                              
                    
                    // peer id 
                    document.getElementsByClassName('info')[0].innerHTML = 'Peer ID : ' + weblocal.getItem("key"); 

          break
          
          case 1:
            
  
                  document.getElementsByClassName('info-content-2')[0].children[0].innerHTML = filterObj[i].CDR_LINK;
                  document.getElementsByClassName('info-content-2')[0].children[1].innerHTML = filterObj[i].SizeOf;
                  document.getElementsByClassName('info-content-2')[0].children[2].innerHTML = filterObj[i].Access;
                  document.getElementsByClassName('info-content-2')[0].children[3].style.display = 'unset';

                  if (!document.getElementsByClassName('info-content-2')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".png")){
                    
                    document.getElementsByClassName('info-content-2')[0].children[4].children[0].src = "/" +filterObj[i].ImagePath;
                    

                  }else if (! document.getElementsByClassName('info-content-2')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".jpeg")){
                    
                    document.getElementsByClassName('info-content-2')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                    
                  }else if (! document.getElementsByClassName('info-content-2')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".gif")){
                    
                    document.getElementsByClassName('info-content-2')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                    
                  }else{

                    document.getElementsByClassName('info-content-2')[0].children[3].style.display = "none";
                  }

                  tObjects.push(filterObj[i].Objects);

            break
        
          case 2: 
        
  
              document.getElementsByClassName('info-content-3')[0].children[0].innerHTML = filterObj[i].CDR_LINK;
            
              document.getElementsByClassName('info-content-3')[0].children[1].innerHTML = filterObj[i].SizeOf;
            
              document.getElementsByClassName('info-content-3')[0].children[2].innerHTML = filterObj[i].Access;

              document.getElementsByClassName('info-content-3')[0].children[3].style.display = 'unset';
              
              if (! document.getElementsByClassName('info-content-3')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".png")){
                    
                document.getElementsByClassName('info-content-3')[0].children[4].children[0].src = "/" +filterObj[i].ImagePath;

              }else if (! document.getElementsByClassName('info-content-3')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".jpeg")){
                
                document.getElementsByClassName('info-content-3')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                
              }else if (! document.getElementsByClassName('info-content-3')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".gif")){
                
                document.getElementsByClassName('info-content-3')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                
              }else{

                document.getElementsByClassName('info-content-3')[0].children[3].style.display = "none";
              }

              tObjects.push(filterObj[i].Objects);

         break  
        case 3: 
        
              
              document.getElementsByClassName('info-content-4')[0].children[0].innerHTML = filterObj[i].CDR_LINK;
            
              document.getElementsByClassName('info-content-4')[0].children[1].innerHTML = filterObj[i].SizeOf;
            
              document.getElementsByClassName('info-content-4')[0].children[2].innerHTML = filterObj[i].Access;
          
              document.getElementsByClassName('info-content-4')[0].children[3].style.display = 'unset';

              

              if (! document.getElementsByClassName('info-content-4')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".png")){
                    
                document.getElementsByClassName('info-content-4')[0].children[4].children[0].src = "/" +filterObj[i].ImagePath;

              }else if (! document.getElementsByClassName('info-content-4')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) &&(filterObj[i].ImagePath).includes(".jpeg")){
                
                document.getElementsByClassName('info-content-4')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                
              }else if (! document.getElementsByClassName('info-content-4')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) &&(filterObj[i].ImagePath).includes(".gif")){
                
                document.getElementsByClassName('info-content-4')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                
              }else{
                      
                document.getElementsByClassName('info-content-4')[0].children[3].style.display = "none";
              }

              tObjects.push(filterObj[i].Objects);
  
         break
        case 4: 
            
              document.getElementsByClassName('info-content-5')[0].children[0].innerHTML = filterObj[i].CDR_LINK;
            
              document.getElementsByClassName('info-content-5')[0].children[1].innerHTML = filterObj[i].SizeOf;
            
              document.getElementsByClassName('info-content-5')[0].children[2].innerHTML = filterObj[i].Access;
          
              document.getElementsByClassName('info-content-5')[0].children[3].style.display = 'unset';

              if (! document.getElementsByClassName('info-content-5')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".png")){
                    
                document.getElementsByClassName('info-content-5')[0].children[4].children[0].src = "/" +filterObj[i].ImagePath;

              }else if (!  document.getElementsByClassName('info-content-5')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) &&(filterObj[i].ImagePath).includes(".jpeg")){
                
                document.getElementsByClassName('info-content-5')[0].children[4].children[0].src = "/"+filterObj[i].ImagePath;
                
              }else if (! document.getElementsByClassName('info-content-5')[0].children[0].innerHTML === " " && (filterObj[i].ImagePath !== undefined) && (filterObj[i].ImagePath).includes(".gif")){
                
                document.getElementsByClassName('info-content-5')[0].children[4].children[0].src = "/"+filterBj[i].ImagePath;
                
              }else{

                document.getElementsByClassName('info-content-5')[0].children[3].style.display = "none";
              }

              tObjects.push(filterObj[i].Objects);

         break
    }
}
  
  // segmented created 
  document.getElementsByClassName('objects')[0].children[0].innerHTML = Math.max(...tObjects) + " Decentralized Objects";
  document.getElementsByClassName('objects')[0].children[2].innerHTML = Math.max(...tObjects) -5 <= 0 ? 0 : Math.max(...tObjects) -5;
  
  
  // clear the objects states 
  pusher.unbind();
  weblocal.clear();
  console.log('local cache free now ...', weblocal.length)


  // annymous function 
  function PopUnusedObjects(object){

    return object;
  }