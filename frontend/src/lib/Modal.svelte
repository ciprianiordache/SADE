<script>
    import {apiHost} from "$lib/utils.js";

    export let isOpen = false;
    export let selectedItem = null;
    export let onClose = () => {};

    let selectedTab = 'preview';

    const handleClose = () => {
        onClose();
    };

    const handleResend = async(email, id) => {
        try {
            const response = await fetch(apiHost + '/resend', {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams({email: email, media_id: id}),
            });

            if(response.ok) {
                alert('Preview email resent successfully!');
            } else {
                alert('Failed to resend preview email.');
            }
        } catch(error) {
            console.error('Error resending preview:', error)
            alert('An error occurred while resending the email.');
        }
    };
</script>

<style>
    .modal {
        display: var(--modal-display, none);
        position: fixed;
        z-index: 1000;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        overflow: auto;
        background-color: rgba(0, 0, 0, 0.5);
    }

    .modal-content {
        background-color: #fefefe;
        margin: 9% auto;
        padding: 20px;
        border: 1px solid #888;
        width: 80%;
        max-width: 600px;
        text-align: center;
        font-family: monospace;
    }

    .close {
        width: 24px;
        height: 24px;
        background: url('/close.png') no-repeat center center;
        background-size: contain;
        border: none;
        cursor: pointer;
        transition: transform 0.3s;
        margin-bottom: 2px;
    }

    .close:hover {
        transform: scale(1.2);
    }

    .modal-media {
        max-width: 100%;
        height: auto;
    }

    .resent {
        font-family: monospace;
        padding: 10px;
        background-color: coral;
        border: none;
        height: 50px;
        width: 400px;
        font-size: 15px;
        color: black;
        cursor: pointer;
        transition: background-color 0.3s;
        margin-top: 0.7rem;
    }

    .resent:hover {
        background-color: orange;
    }

    .navbar {
        background-color: #555555;
        color: white;
        padding: 0.5rem;
        display: flex;
        justify-content: space-between;
        align-items: flex-end;
        font-family: monospace;
    }

    .navbar-button {
        display: flex;
        justify-content: left;
        align-items: flex-end;
        gap: 1rem;
    }

    .navbar-button button {
        color: white;
        text-decoration: none;
        padding: 0.5rem 1.5rem;
        background-color: #555;
        border-radius: 4px;
        border: none;
        transition: background-color 0.3s;
    }

    .navbar-button button:hover {
        background-color: #777;
    }

    #logo {
        width: 50%;
        height: auto;
    }

</style>

<div class="modal" style="--modal-display: {isOpen ? 'block' : 'none'};" on:click={handleClose}>
    <div class="modal-content" on:click|stopPropagation>
        {#if selectedItem}
            <div class="navbar">
                <div class="navbar-button">
                    <button on:click={() => selectedTab = 'preview'}>Preview</button>
                    <button on:click={() => selectedTab = 'original'}>Original</button>
                    <button on:click={() => selectedTab = 'status'}>Status</button>
                </div>
                <button class="close" on:click={handleClose}></button>
            </div>
            {#if selectedTab === 'preview'}
                <h2>Preview</h2>
                {#if selectedItem.type === 'image'}
                    <img src={apiHost +'/'+ selectedItem.preview_path} alt="Preview Image" class="modal-media">
                {:else if selectedItem.type === 'video'}
                    <video src={apiHost +'/'+  selectedItem.preview_path} controls class="modal-media"></video>
                {:else if selectedItem.type === 'audio'}
                    <img src="/icon300.png" alt="Logo" id="logo"/>
                    <audio src={apiHost +'/'+  selectedItem.preview_path} controls></audio>
                {/if}
            {:else if selectedTab === 'original'}
                <h2>Original</h2>
                {#if selectedItem.type === 'image'}
                    <img src={apiHost +'/'+  selectedItem.original_path} alt="Original Image" class="modal-media">
                {:else if selectedItem.type === 'video'}
                    <video src={apiHost +'/'+  selectedItem.original_path} controls class="modal-media"></video>
                {:else if selectedItem.type === 'audio'}
                    <img src="/icon300.png" alt="Logo" id="logo"/>
                    <audio src={apiHost +'/'+  selectedItem.original_path} controls></audio>
                {/if}
            {:else if selectedTab === 'status'}
                <h2>Status</h2>
                <p>{selectedItem.locked ? 'Locked' : 'Bought'}</p>
            {/if}
            <button class="resent" on:click={handleResend(selectedItem.client_email, selectedItem.id)}>Resend Preview Email</button>
        {/if}
    </div>
</div>
