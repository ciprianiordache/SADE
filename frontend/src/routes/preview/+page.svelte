<script>
    import { onMount} from "svelte";
    import { apiHost } from "$lib/utils.js";


    let previewData = null;
    let loading = true;
    let loadingText = "Loading";
    let dots = "";
    let error = null;
    let interval;

    onMount(async () => {
        interval = setInterval(() => {
            dots = dots.length < 3 ? dots + "." : "";
            loadingText = `Loading${dots}`;
        }, 500);

        const params = new URLSearchParams(window.location.search);
        const previewPath = params.get("preview_path");
        if (!previewPath) {
            error = "No preview path specified.";
            loading = false;
            clearInterval(interval);
            return;
        }
        try {
            const response = await fetch(`${apiHost}/preview?preview_path=${encodeURIComponent(previewPath)}`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            if (data) {
                previewData = data;
            } else {
                throw new Error('Error API response');
            }
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
            clearInterval(interval);
        }
    });

    const handlePayment = (mediaId) => {
        window.location.href = `/payment/${mediaId}`;
    };
</script>

<main>
    <h1>Preview Media</h1>
    {#if loading}
        <p>{loadingText}</p>
    {:else if error}
        <p class="error">{error}</p>
    {:else}
        <div class="preview">
            <div class="preview-item">
                <h2>Price: {previewData.price}</h2>
                {#if previewData.type === 'image'}
                    <img id="logo" src={`${apiHost}/media/${previewData.preview_path}`} alt="Preview Image" class="preview-image" />
                {:else if previewData.type === 'video'}
                    <video src={`${apiHost}/media/${previewData.preview_path}`} controls class="preview-video"></video>
                {:else if previewData.type === 'audio'}
                    <img src="/icon300.png" alt="Logo">
                    <audio src={`${apiHost}/media/${previewData.preview_path}`} controls class="preview-audio"></audio>
                {/if}
                <button on:click={() => handlePayment(previewData.id)}>Pay to unlock the original file.</button>
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
    .preview {
        display: flex;
        flex-wrap: wrap;
        gap: 1em;
    }
    .preview-item {
        display: flex;
        flex-direction: column;
        border-color: whitesmoke;
        border-style: solid;
        border-width: 1px;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
        padding: 1em;
    }
    .preview-image {
        max-width: 900px;
        max-height: 900px;
        object-fit: cover;
        margin-bottom: 10px;
    }
    .preview-video, .preview-audio {
        max-width: 900px;
    }
    button {
        font-family: monospace;
        margin-top: 10px;
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
