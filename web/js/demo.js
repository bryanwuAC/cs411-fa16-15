

function appenditem(item){
    $("#body").append(
		"<div>" + item.username + "</div>" 
  	)
}




$("#demo").on('click', function () {
    $.get("ajax/getAdmins", function(data){
        $.each(data.MyData_array, function(i, item) {
            appenditem(item)
        });
    })
})