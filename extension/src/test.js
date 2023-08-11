let token = '';
let tuid = '';

const twitch = window.Twitch.ext;

function requestVideo(token, tuid) {
    fetch("http://localhost:3333/video?" + new URLSearchParams({
        streamId: tuid
    }), {
        method: "GET",
        mode: "cors",
        credentials: "same-origin"
    })
    .then(resp => resp.json())
    .then(resp => {
        console.log(resp.videoUrl)
        let videoUrl = "https://www.youtube.com/embed/uGGQGoht6ic"
        document.getElementById("video").src = videoUrl
    })
}

twitch.onAuthorized(function (auth) {
    // save our credentials
    token = auth.token;
    tuid = auth.userId;
  
    requestVideo(token, tuid)
});
