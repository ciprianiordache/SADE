<script>
    import Sidebar from "$lib/Sidebar.svelte";
    import Table from "$lib/Table.svelte";
    import { onMount } from "svelte";
    import { apiHost, formatDate } from "$lib/utils.js";

    let data = [];
    let columns = ["id", "media_id", "client_email", "amount", "currency", "description", "created_at"];
    let columnHeaders = ["ID", "Media ID", "Client Email", "Amount", "Currency", "Description", "Created At"];
    let loading = true;
    let error = null;

    onMount(async () => {
        try {
            const response = await fetch(`${apiHost}/transaction`, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            if (!response.ok) {
                throw new Error('Failed to fetch transactions');
            }
            const transactions = await response.json();
            if(transactions === null) {
                error = 'Empty data';
            }
            data = transactions.data.map(t => ({
                id: t.id,
                media_id: t.media_id,
                client_email: t.client_email,
                amount: (t.amount / 100).toFixed(2),
                currency: t.currency.toUpperCase(),
                description: t.description,
                created_at: formatDate(t.created_at),
            }));
        } catch (err) {
            console.error("Error fetching transactions:", err);
            error = "No transaction made yet.";
        } finally {
            loading = false;
        }
    });
</script>

<style>
    .content {
        flex: 1;
        font-family: monospace;
        max-height: 90vh;
        overflow-y: auto;
    }

    #title {
        position: sticky;
        top: 0;
        background-color: white;
        z-index: 999;
        padding: 0.1rem;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1), 0 2px 4px rgba(0,0,0,0.1);
        border: 1px solid #ddd;
    }
    h2 {
        margin-left: 1rem;
        color: coral;
    }

    .error {
        margin-left: 1rem;
    }
</style>

<Sidebar />

<div class="content">
    <div class="main-content">
        <div id="title"><h2>Transactions</h2></div>
        {#if loading}
            <p>Loading...</p>
        {:else if error}
            <p class="error">{error}</p>
        {:else}
            <Table {data} {columns} {columnHeaders} />
        {/if}
    </div>
</div>