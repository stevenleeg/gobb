var Sirjtaa = (function() {
    function clickThreadQuote() {
        var thread_id = $(this).attr("thread");
        var content = $("#p"+ thread_id +"_content").text();
        var author  = $("#p"+ thread_id +"_author").text();

        // Remove any previous quotes from the content
        content = content.replace(/.* said:\s*(\>.*\n)*/g, "");
        content = content.replace(/>.*\n/g, "").trim();
        content = content.replace(/\n/g, "\n>");

        var prev = $("#reply_content").text();
        var val  = author + " said: \n\n>" + content + "\n\n";
        $("#reply_content")
            .focus()
            .val('').val(val)
            .scrollTop($('#reply_content')[0].scrollHeight);
    }

    function clickThreadReply() {
        $("#reply_content").focus();
    }

    function clickConfirmDelete() {
        r = confirm("Really delete this? It can't be undone.");
        if(r == false) {
            return false;
        }
    }

    $(document).ready(function() {
        $(".thread_quote").on("click", clickThreadQuote);
        $(".thread_reply_btn").on("click", clickThreadReply);
        $(".delete").on("click", clickConfirmDelete);
    });
})();
