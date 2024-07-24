<script>
    export let data = [];
    export let columns = [];
    export let columnHeaders = [];
</script>

<table>
    <thead>
    <tr class="title">
        {#each columnHeaders as header}
            <th>{header}</th>
        {/each}
    </tr>
    </thead>
    <tbody>
    {#each data as row}
        <tr>
            {#each columns as column}
                <td>{row[column]}</td>
            {/each}
        </tr>
    {/each}
    </tbody>
</table>

<style>
    table {
        width: 100%;
        border-collapse: collapse;
        text-align: center;
    }

    td {
        padding: 0.5rem;
        border: 1px solid #ddd;
    }


    th {
        background-color: #f4f4f4;
        padding: 0.8rem;
    }

    .title {
        box-shadow: 0 2px 4px rgba(0,0,0,0.1), 0 2px 4px rgba(0,0,0,0.1);
        border: 1px solid #ddd;
    }

</style>
