<script>
    import Sidebar from "$lib/Sidebar.svelte";
    import { onMount } from 'svelte';
    import Gallery from '$lib/Gallery.svelte';
    import { fetchMediaFiles } from '$lib/utils.js';

    let mediaFiles = [];
    let error = null;
    let loading = true
    let loadingText = "Loading";
    let dots = "";
    let interval;


    const fetchFiles = async () => {
        interval = setInterval(() => {
            dots = dots.length < 3 ? dots + "." : "";
            loadingText = `Loading${dots}`;
        }, 500)
        try {
            mediaFiles = await fetchMediaFiles();
            if (!mediaFiles) {
                error = 'No media files uploaded yet.'
            }
        } catch (err) {
            error = err.message;
        } finally {
            loading = false
        }
        clearInterval(interval);
    }

    onMount(() => {
        fetchFiles()
    });

</script>

<style>
    .content {
        flex: 1;
        font-family: monospace;
        max-height: 90vh;
        overflow-y: auto;
    }
    .main-content {
        padding: 1rem;
    }

    #title {
        position: sticky;
        top: 0;
        background-color: white;
        z-index: 999;
        margin-top: -1rem;
        padding: 0.1rem;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1), 0 2px 4px rgba(0,0,0,0.1);
        border: 1px solid #ddd;
    }
    h2 {
        margin-left: 1rem;
        color: coral;
    }
</style>

<Sidebar />
<div class="content">
    <div id="title"><h2>Uploaded Media</h2></div>
    <div class="main-content">
    {#if loading}
        <p>{loadingText}</p>
        {:else if error}
            <p>{error}</p>
        {:else}
            <Gallery {mediaFiles} />
    {/if}
    </div>
</div>