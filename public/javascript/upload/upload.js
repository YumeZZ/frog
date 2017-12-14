$("#upload-form").submit(function(e) {
    e.preventDefault();
    var formURL = $(this).attr("action");
    var formData = new FormData(this);
    location.reload();
    $('input[name=organismname]').val('');

    $.ajax({
        url : formURL,
        type: "POST",
        data : formData,
        processData: false,
        contentType: false,
        success:function(data, textStatus, jqXHR) {
            //var result = $.parseJSON(data);
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("ajax post error", textStatus);
        },
        always: function() {

        }
    });
});