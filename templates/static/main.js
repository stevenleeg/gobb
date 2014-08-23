var Sirjtaa = (function() {
    function quotePostClicked(e) {
        e.preventDefault();

        var pid = $(this).data("postid");
        var content = $("#p"+ pid +"-unparsed-content").text();
        var author  = $("#p"+ pid +"-author").text();

        // Remove any previous quotes from the content
        content = content.replace(/.* said:\s*(\>.*\n)*/g, "");
        content = content.replace(/>.*\n/g, "").trim();
        content = content.replace(/\n/g, "\n>");

        var reply = $("#reply-field");
        var prev = reply.text();
        var val  = author + " said: \n\n>" + content + "\n\n";
        reply.focus()
            .val('').val(val)
            .scrollTop(reply[0].scrollHeight);
    }

    function threadReplyClicked() {
        $("#reply-field").focus();
    }

    function clickConfirmDelete() {
        r = confirm("Really delete this? It can't be undone.");
        if(r == false) {
            return false;
        }
    }

    function clickModerate(e) {
        e.preventDefault();
        $(this).hide();
        $(this).parent().children(".mod_tools").show();
        $(this).parent().children(".prec_slash").hide();
    }

    $(document).ready(function() {
        $(".quote-post").on("click", quotePostClicked);
        $(".thread-reply-btn").on("click", threadReplyClicked);
        $(".delete").on("click", clickConfirmDelete);
        $(".moderate").on("click", clickModerate);
    });
})();
