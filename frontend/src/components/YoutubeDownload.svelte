<script>
  import { onMount } from "svelte";
  import Spinner from "svelte-spinner";

  let homeDirs = [];
  let downloading = false;
  let format = 0;
  let dir = "yt";
  let url = "";

  onMount(async () => {
    homeDirs = await backend.Downloader.ListHomeDirectories();
  });

  async function download() {
    downloading = true;

    try {
      await backend.Downloader.DownloadFromYoutube(url, format, dir);
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
    <input
      id="url"
      type="text"
      bind:value={url}
      disabled={downloading}
      placeholder="Paste youtube url here..."
    />
    <select bind:value={format}>
      <option value={0}>Audio Only</option>
      <option value={1}>Video and Audio</option>
    </select>

    <select bind:value={dir}>
      {#each homeDirs as dir}
        <option value={dir}>{dir}</option>
      {/each}
    </select>
  </div>

  <div style="margin-bottom: 1rem;" />

  <div class="button-wrapper">
    <button disabled={!url.length || downloading} on:click={download}>
      Download from Youtube
    </button>
  </div>

  {#if downloading}
    <div style="margin-bottom: 1rem;" />
    <Spinner size="50" speed="500" color="#ffffff" thickness="2" gap="40" />
  {/if}
</main>

<style>
  .input-wrapper {
    display: flex;
    gap: 1rem;
  }

  input {
    background: #ffffff30;
    padding: 0 1rem;
    width: 500px;
    height: 3rem;
    border: none;
    color: #ffffff;
    font-size: 1rem;
  }

  select {
    padding: 0 1rem;
    background: #ffffff30;
    border: none;
    font-size: 1rem;
    transition: all 0.3s;
  }

  select:focus {
    outline: none;
    background: #ffffff50;
  }

  input:focus {
    outline: none;
    background: #ffffff50;
  }

  button {
    background: #ffffff30;
    padding: 1rem;
    color: white;
    border: none;
    transition: all 0.3s;
  }

  button:focus {
    outline: none;
    background: #ffffff50;
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
