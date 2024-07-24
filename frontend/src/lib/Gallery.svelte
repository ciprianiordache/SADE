<script>
    import {apiHost} from "$lib/utils.js";
    import Modal from "$lib/Modal.svelte";

    export let mediaFiles = [];

    let selectedMedia = null;
    let isModalOpen = false;

    const openModal = (media) => {
        selectedMedia = media;
        selectedMedia.preview_path = media.preview_path;
        selectedMedia.original_path = media.original_path;
        isModalOpen = true;
    };

    const closeModal = () => {
        isModalOpen = false;
        selectedMedia = null;
    };
</script>

<style>
    .gallery {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
        padding: 1rem;
    }
    .gallery-item {
        width: 30%;
        background: #f4f4f4;
        padding: 0.5rem;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
        text-align: center;
        font-family: monospace;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        transition: background-color 0.3s;
        border: 1px solid #ddd;
        cursor: pointer;
    }
    .gallery-item:hover {
        background: #c5c5c5;
    }
    .gallery-item img, .gallery-item video {
        max-width: 100%;
        height: auto;
    }

    #logo {
        width: 50%;
        height: auto;
    }

    .audio-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
    }

    .audio-container audio {
        width: 100%;
    }
</style>

<div class="gallery">
    {#each mediaFiles as media}
        <div class="gallery-item" on:click={() => openModal(media)}>
            <h3>{media.client_email}</h3>
            {#if media.type === 'image'}
                <img src={apiHost+`/${media.preview_path}`} alt="Uploaded Image">
            {:else if media.type === 'video'}
                <video src={apiHost+`/${media.preview_path}`} controls></video>
            {:else if media.type === 'audio'}
                <div class="audio-container">
                    <img src="/icon300.png" alt="Logo" id="logo"/>
                    <audio src={apiHost+`/${media.preview_path}`} controls></audio>
                </div>
            {/if}
            <p>Price: {media.price}</p>
        </div>
    {/each}
</div>

<Modal isOpen={isModalOpen} selectedItem={selectedMedia} onClose={closeModal} />

