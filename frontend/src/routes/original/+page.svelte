<script>
    import { onMount } from "svelte";
    import { apiHost, downloadMedia } from "$lib/utils.js";

    let originalData;
    let loading = true;
    let error = null;
    let loadingText = "Loading";
    let dots = "";
    let interval;

    onMount(async () => {
        interval = setInterval(() => {
            dots = dots.length < 3 ? dots + "." : "";
            loadingText = `Loading${dots}`;
        }, 500);

        const params = new URLSearchParams(window.location.search);
        const mediaID = params.get('media_id');

        if (!mediaID) {
            error = "No media id provided!";
            loading = false;
            clearInterval(interval);
            return;
        }
        try {
            const response = await fetch(`${apiHost}/original?media_id=${encodeURIComponent(mediaID)}`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            if (!response.ok) {
                throw new Error('Network response was not ok!');
            } else {
                originalData = await response.json();
                if (!originalData) {
                    throw new Error('Error API response');
                }
            }
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
            clearInterval(interval);
        }
    });

    const handleDownload = async () => {
        await downloadMedia(originalData.id);
    }
</script>

<main>
    <h1>Original Media</h1>
    {#if loading}
        <p>{loadingText}</p>
    {:else if error}
        <p class="error">{error}</p>
    {:else}
        <div class="original">
            <div class="original-item">
                {#if originalData.type === 'image'}
                    <img id="logo" src={`${apiHost}/media/${originalData.original_path}`} alt="Original Image" class="original-image" />
                {:else if originalData.type === 'video'}
                    <video src={`${apiHost}/media/${originalData.original_path}`} controls class="original-video"></video>
                {:else if originalData.type === 'audio'}
                    <img src="/icon300.png" alt="Logo">
                    <audio src={`${apiHost}/media/${originalData.original_path}`} controls class="original-audio"></audio>
                {/if}
                <button on:click={handleDownload}>Download the original</button>
            </div>
        </div>
    {/if}
</main>

<style>
    main {
        font-family: monospace;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 90vh;
    }
    .original {
        display: flex;
        flex-wrap: wrap;
        gap: 1em;
    }
    .original-item {
        display: flex;
        flex-direction: column;
        border-color: whitesmoke;
        border-style: solid;
        border-width: 1px;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
        padding: 1em;
    }
    .original-image {
        max-width: 900px;
        max-height: 900px;
        object-fit: cover;
        margin-bottom: 10px;
    }
    .original-video, .original-audio {
        max-width: 900px;
    }
    button {
        margin-top: 10px;
        font-family: monospace;
        padding: 10px;
        border: none;
        background: coral;
        cursor: pointer;
        height: 45px;
        font-size: 23px;
        transition: background-color 0.3s;
    }

    button:hover {
        background: orange;
    }
    .error {
        color: red;
    }
</style>
