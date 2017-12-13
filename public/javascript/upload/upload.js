$("#upload-form").submit(function(e) {
    e.preventDefault();
    var formURL = $(this).attr("action");
    var formData = new FormData(this);
    $.ajax({
        url : formURL,
        type: "POST",
        data : formData,
        processData: false,
        contentType: false,
        success:function(data, textStatus, jqXHR) {
            var result = $.parseJSON(data);
            console.log("post success");
            //console.log(result.UploadStatus);
            /*
            if (result.UploadStatus == false) {

            } else {
                console.log(result.UploadStatus);            
            }
            */
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("ajax post error", textStatus);
        },
        always: function() {

        }
    });
});