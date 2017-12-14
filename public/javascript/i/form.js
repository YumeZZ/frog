$( "#searchtype-organismname" ).click(function() {
    $('input[name=searchtype]').val('organismname');
});

$( "#searchtype-category" ).click(function() {
    $('input[name=searchtype]').val('category');
});

$( "#searchtype-location" ).click(function() {
    $('input[name=searchtype]').val('location');
});

$( "#searchtype-gps" ).click(function() {
    $('input[name=searchtype]').val('gps');
});

$( "#searchtype-season" ).click(function() {
    $('input[name=searchtype]').val('season');
});

$( "#searchtype-daterange" ).click(function() {
    $("#dateto-label").css("display","inline");
    $("#datefrom-label").css("display","inline");
    $('input[name=dateto]').css("display","inline-flex");
    $('input[name=datefrom]').css("display","inline-flex");
    $('input[name=searchtype]').val('daterange');
});


$("#search-form").submit(function(e) {
    e.preventDefault();
    var formURL = $(this).attr("action");
    var formData = $(this).serializeArray();

    
    //console.log(formData);
    $.ajax({
        url : formURL,
        type: "POST",
        data : formData,
        success:function(data, textStatus, jqXHR) {
            var result = $.parseJSON(data);
            //console.log(result.Records);

            $('.search-result').empty();

            searchResultHTML = '';
            jQuery.each(result.Records, function(i, val) {
                console.log(i, val);
                //console.log(result.Records[0].ID);
                //console.log(result.Records[0].Photo[0]);
                searchResultHTML += '<img ' + 'class="albumimg"' + 'id="' + result.Records[i].ID +'"' + ' ' + 'src="' + '/storage/photo/' + result.Records[i].Photo[0] + '" />';
            });
            $('.search-result').prepend(searchResultHTML);

            //$(location).attr('href', '/');
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("error");
        }
    });
});

