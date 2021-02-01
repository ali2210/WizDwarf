function InputCardNumberEvent(){
  var text_in = document.getElementById('cardNo');
  var regular = '^d{4}([-]?)((d{6}(d{1})?d{5})|(d{4}d{1}?d{4}(d{1})?d{4}))$';
  var status = text_in.textContent.match(regular);
  if (!status){
    console.error('data is not valid', status);
  }else{
    console.log('Gotcha!', status);
  }

}

