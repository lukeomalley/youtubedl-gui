<script>
  import Spinner from "svelte-spinner";
  let url = "";
  let downloading = false;

  async function download() {
    downloading = true;

    try {
      await backend.Downloader.DownloadFromYoutube(url);
      downloading = false;
      url = "";
    } catch (err) {
      console.error(err);
      downloading = false;
    }
  }
</script>

<main>
  <div class="input-wrapper">
    <label for="url">Paste Youtube Url Here: </label>
    <input id="url" type="text" bind:value={url} disabled={downloading} />
  </div>

  <div style="margin-bottom: 1rem;" />

  <button disabled={!url.length || downloading} on:click={download}>
    Download from Youtube
  </button>

  {#if downloading}
    <div style="margin-bottom: 1rem;" />
    <Spinner size="50" speed="500" color="#ffffff" thickness="2" gap="40" />
  {/if}
</main>

<style>
  .input-wrapper {
    display: flex;
    flex-direction: column;
  }

  label {
    margin-bottom: 1rem;
  }

  input {
    background: #ffffff50;
    height: 3rem;
    border: none;
    color: #ffffff;
    font-size: 2rem;
  }

  input:focus {
    outline: none;
  }

  button {
    background: #ffffff50;
    padding: 1rem;
    color: white;
    border: none;

    transition: all 0.3s;
  }

  button:focus {
    outline: none;
  }

  button:hover {
    background: #ffffff80;
    cursor: pointer;
  }

  button:disabled {
    background: #ffffff20;
    color: lightgray;
    cursor: default;
  }
</style>
