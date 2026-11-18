const articleElement = document.querySelector("article");
const loadingElement = document.querySelector("#loading");
const statusElement = document.querySelector("#status");
const errorElement = document.querySelector("#error");
const errorMessageElement = document.querySelector("#error-message");
const finishedElement = document.querySelector("#finished");
const canvasElement = document.querySelector("canvas");

const onWindowResize = () => {
  canvasElement.width = window.innerWidth;
  canvasElement.height = window.innerHeight;
};
window.addEventListener("resize", onWindowResize);
onWindowResize();

canvasElement.addEventListener("contextmenu", (event) => {
  event.preventDefault();
});

const toMB = (bytes) => (bytes / (1024 * 1024)).toFixed(1);

const showProgress = (loaded, total) => {
  statusElement.textContent = total > 0
    ? `Downloading... ${toMB(loaded)} / ${toMB(total)} MB`
    : `Downloading... ${toMB(loaded)} MB`;
};

const hideLoading = () => {
  articleElement.style.display = "none";
  canvasElement.style.display = "block";
  canvasElement.focus();
};

const showError = (error) => {
  console.error(error);
  canvasElement.style.display = "none";
  loadingElement.style.display = "none";
  articleElement.style.display = "flex";
  errorElement.style.display = "block";
  errorMessageElement.textContent = String(error.message || error);
};

const showFinished = () => {
  canvasElement.style.display = "none";
  loadingElement.style.display = "none";
  articleElement.style.display = "flex";
  finishedElement.style.display = "block";
};

// Fetches the WASM file, reporting download progress while
// still allowing the browser to compile it as it streams in.
const fetchWithProgress = async (url) => {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Download failed: ${response.status} ${response.statusText}`);
  }

  // When the response is compressed, Content-Length is the compressed size
  // and cannot be compared to the received (decompressed) bytes.
  const isCompressed = response.headers.has("Content-Encoding");
  const total = isCompressed ? 0 : Number(response.headers.get("Content-Length")) || 0;
  let loaded = 0;

  const progressStream = new TransformStream({
    transform(chunk, controller) {
      loaded += chunk.byteLength;
      showProgress(loaded, total);
      controller.enqueue(chunk);
    },
    flush() {
      statusElement.textContent = "Starting...";
    },
  });

  return new Response(response.body.pipeThrough(progressStream), {
    headers: { "Content-Type": "application/wasm" },
  });
};

try {
  console.log("Loading WebAssembly executable...");
  const go = new Go();
  const result = await WebAssembly.instantiateStreaming(
    fetchWithProgress("web/main.wasm"),
    go.importObject
  );

  hideLoading();

  console.log("Running WebAssembly executable...");
  await go.run(result.instance);

  console.log("Finished WebAssembly executable.");
  showFinished();
} catch (error) {
  showError(error);
}
