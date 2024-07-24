<script>
    import { onMount } from "svelte";
    import {apiHost, uploadMedia} from "$lib/utils.js";
    import {goto} from "$app/navigation";

    let sessionData;
    let file;
    let clientEmail = '';
    let price = '';
    let isLoading = false;
    let loadingText = "Uploading";
    let dots = "";
    let interval;

    onMount(async () => {

        try {
            const response = await fetch(apiHost+'/check', {
                credentials: 'include'
            });
            if (response.ok) {
                sessionData = await response.json()
            } else {
                goto("/");
            }
        } catch (error) {
            console.error('Error fetching session data', error);
        }

    });

    function handleFileUpload(event) {
        file = event.target.files[0];
    }

    async function handleUploadClick() {
        if (!file || !clientEmail || !price) {
            alert('Please fill in all fields and select a file.')
            return;
        }

        isLoading = true;
        interval = setInterval(() => {
            dots = dots.length < 3 ? dots + "." : "";
            loadingText = `Uploading${dots}`;
        }, 500)
        try {
            await uploadMedia(file, clientEmail, price, sessionData.user_id);
            alert('File uploaded!');
            resetForm();
        } catch (error) {
            console.error('An error occurred while uploading file.', error);
            alert('An error occurred while uploading file.');
        } finally {
            isLoading = false;
        }
        clearInterval(interval);
    }

    function resetForm() {
        file = null;
        clientEmail = '';
        price = '';
        document.getElementById('media').value = '';
    }

    function validatePrice(event) {
        let value = event.target.value;
        value = parseFloat(value).toFixed(2);
        if (!isNaN(value)) {
            price = value;
        } else {
            price = '';
        }
    }
</script>

<style>
    .sidebar {
        width: 18%;
        background-color: #f4f4f4;
        padding: 1rem;
        font-family: monospace;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1), 0 2px 4px rgba(0,0,0,0.1);
        border: 1px solid #ddd;
    }

    .sidebar h2 {
        font-family: monospace;
        color: coral;
    }

    .sidebar input, .sidebar button {
        display: block;
        margin: 1rem 0;
        padding: 0.5rem;
        width: 85%;
        font-family: monospace;
    }

    .sidebar button {
        background-color: coral;
        border: none;
        color: black;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .sidebar button:hover {
        background-color: orange;
    }

    .overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
        font-family: monospace;
        color: coral;
        font-size: 1.5rem;
    }
</style>

<div class="sidebar">
    <h2>Upload Media</h2>
    <input id="media" type="file" on:change={handleFileUpload} required>
    <input type="email" bind:value={clientEmail} placeholder="Client Email" required>
    <input type="number" bind:value={price} on:change={validatePrice} placeholder="Price" step="0.01" min="0" required>
    <button on:click={handleUploadClick} disabled={isLoading}>Upload</button>
</div>

{#if isLoading}
    <div class="overlay">{loadingText}</div>
{/if}
