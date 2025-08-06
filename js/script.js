// script.js
document.getElementById("sync-btn").addEventListener("click", () => {
  const video = document.getElementById("vid-player");
  video.play(); // Later we'll sync this across devices
  alert("Syncing...");
});

const socket = new WebSocket("ws://localhost:8080/sync");
socket.onmessage = (e) => {
  if (e.data === "play") document.getElementById("vid-player").play();
};

document.getElementById("sync-btn").onclick = () => {
  socket.send("play"); // Send command to server
};
