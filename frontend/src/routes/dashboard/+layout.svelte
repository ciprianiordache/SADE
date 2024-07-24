<script context="module">
    export async function load() {
        const sessionValid = await checkSession();
        if (!sessionValid) {
            return {
                status: 302,
                redirect: '/'
            };
        }
    }
</script>

<script>
    import Navbar from '$lib/Navbar.svelte';
    import { onMount } from "svelte";
    import { checkSession } from "$lib/utils.js";

    let loading = true;
    let sessionValid = false;
    let loadingText = "Loading";
    let dots = "";
    let interval;

    onMount(async () => {
        interval = setInterval(() => {
            dots = dots.length < 3 ? dots + "." : "";
            loadingText = `Loading${dots}`;
        }, 500);
        sessionValid = await checkSession();
        loading = false;

        clearInterval(interval);
    })

</script>

<style>
    .container {
        display: flex;
        height: calc(100vh - 80px);
    }

    .loading {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 96vh;
        font-family: monospace;
        font-size: 1.5rem;
    }
</style>

{#if loading}
    <div class="loading">{loadingText}</div>
{:else}
    <Navbar />
    <div class="container">
        <slot></slot>
    </div>
{/if}