<script>
    import { onMount } from 'svelte';
    import { apiHost } from "$lib/utils.js";

    let error = null;

    function getParam() {
        const params = new URLSearchParams(window.location.search);
        return {
            id: parseInt(params.get("id")),
            first_name: params.get("first_name"),
            last_name: params.get("last_name"),
            email: params.get("email"),
            role: params.get("role"),
            verified: params.get("verified") === 'true'
        };
    }

    async function saveSession(userData) {
        try {
            const response = await fetch(apiHost+'/save', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: 'include',
                body: JSON.stringify(userData)
            });
            if (!response.ok) {
                const result = await response.json();
                throw new Error(result.error || 'Failed to save session data');
            }
        } catch (err) {
            error = err.message;
        }
    }

    onMount(() => {
        const userData = getParam();
        saveSession(userData);
    });
</script>

{#if error}
    <div class="error">{error}</div>
{:else}
    <div class="container">
        <img src="/icon300.png" alt="Logo" id="logo"/>
        <h1>SADE</h1>
        <p>You are now connected!</p>
        <p>Please close this page and return to your app!</p>
    </div>
{/if}

<style>
    .container {
        font-family: monospace;
        text-align: center;
        margin-top: 15%;
    }

    #logo {
        width: 150px;
        height: auto;
    }

    .error {
        color: red;
        text-align: center;
        margin-top: 20%;
    }
</style>
