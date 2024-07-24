<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { apiHost } from "$lib/utils.js";

    let message = 'Check Your Email';
    let subMessage = 'A link has been sent to your email address. Please check your inbox and click on the link to connect on the app. You will be redirected to the dashboard if the link is validated.';

    async function check() {
        try {
            const response = await fetch(`${apiHost}/check`, {
                credentials: 'include',
            });
            if(response.ok) {
                const result = await response.json();
                if (result.message === "You are now connected!") {
                    if(result.role === 'admin') {
                        goto('/dashboard');
                    } else {
                        message = 'Unauthorized';
                        subMessage = 'You do not have permission to view this page.';
                    }
                } else {
                    console.log('No connection message received');
                }
            } else {
                console.log('Failed to check. Status:', response.status);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    }

    onMount(() => {
        check();
        const interval = setInterval(check, 2000);
        return () => clearInterval(interval);
    });
</script>

<div class="container">
    <img src="/icon300.png" alt="Logo" id="logo"/>
    <h1>{message}</h1>
    <p>{subMessage}</p>
</div>

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
</style>
