var Sirjtaa = (function() {
    function clickThreadQuote() {
        var thread_id = $(this).attr("thread");
        var content = $("#p"+ thread_id +"_content").text();
        var author  = $("#p"+ thread_id +"_author").text();

        var prev = $("#reply_content").text();
        var val  = author + " said: \n\n>" + content.replace('\n', '\n>') + "\n\n";
        $("#reply_content").focus().val('').val(val);
    }

    function clickThreadReply() {
        $("#reply_content").focus();
    }

    $(document).ready(function() {
        $(".thread_quote").on("click", clickThreadQuote);
        $(".thread_reply_btn").on("click", clickThreadReply);
    });
})();
