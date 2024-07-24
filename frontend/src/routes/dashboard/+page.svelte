<script>
    import Sidebar from '$lib/Sidebar.svelte';
    import {onMount} from "svelte";
    import {apiHost} from "$lib/utils.js";


    let firstName = '';
    let lastName = '';

    onMount(async () => {
        try {
            const response = await fetch(`${apiHost}/check`, {
                credentials: 'include',
            });
            if (response.ok) {
                const result = await response.json();
                firstName = result.first_name || '';
                lastName = result.last_name || '';
            } else {
                console.log('Failed to fetch user details. Status:', response.status);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    });

</script>

<Sidebar />

<style>
    .main-content {
        flex: 1;
        margin-left: 30%;
        margin-top: 20%;
        padding: 1rem;
        font-family: monospace;
    }
</style>

<div class="main-content">
    <h2>SADE Home</h2>
    <p>Welcome {firstName} {lastName}!</p>
</div>
