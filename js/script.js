// script.js
document.getElementById("sync-btn").addEventListener("click", () => {
  const video = document.getElementById("vid-player");
  video.play(); // Later we'll sync this across devices
  alert("Syncing... (WebSocket coming next!)");
});
