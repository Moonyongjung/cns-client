function charactercheck() {
    $(function(){
        $("#alert-success").hide();
        $("#alert-danger").hide();
        $("#alert-danger-characters").hide();

        $("input").keyup(function(){
            var domain=$("#domain").val();
            
            if(domain.length > 20) {
                $("#alert-success").hide();
                $("#alert-danger").hide();
                $("#alert-danger-characters").show();
                $("#login-button").attr("disabled", "disabled");
            } else if(domain.length == 0) {
                $("#alert-success").hide();
                $("#alert-danger").show();
                $("#alert-danger-characters").hide();
                $("#login-button").attr("disabled", "disabled");            
            } else {
                $("#alert-success").show();
                $("#alert-danger").hide();
                $("#alert-danger-characters").hide();
                $("#login-button").removeAttr("disabled");
            }
            
        });
    });
}

charactercheck()