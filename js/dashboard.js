
    $("#file").on("change", () => {
      if( $("#file")[0].files.length !== 0 ){
        $(".selected-file")[0].innerHTML = `Seleected File: ${$("#file")[0].files[0].name}`;
        $(".selected-file").addClass("file-uploaded")
      }
    })
