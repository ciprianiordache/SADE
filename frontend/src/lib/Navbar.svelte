<script>
    import { logOut } from "$lib/utils.js";
</script>

<style>
    .navbar {
        background-color: #555555;
        color: white;
        padding: 0.5rem;
        display: flex;
        justify-content: space-between;
        align-items: flex-end;
        font-family: monospace;
    }

    .nav-links {
        display: flex;
        justify-content: left;
        align-items: flex-end;
        gap: 1rem;
    }

    .nav-links a {
        color: white;
        text-decoration: none;
        padding: 1rem 2rem;
        background-color: #555;
        border-radius: 4px;
        transition: background-color 0.3s;
    }

    .nav-links a:hover{
        background-color: #777;
    }

    #logo {
        height: 40px;
        width: auto;
        margin-bottom: 0.2rem;
    }

    #logout {
        width: 35px;
        height: 35px;
        background: url('/logout.png') no-repeat center center;
        background-size: contain;
        border: none;
        cursor: pointer;
        transition: transform 0.3s;
        margin-bottom: 3px;
    }

    #logout:hover {
        transform: scale(1.2);
    }

</style>

<div class="navbar">
    <div class="nav-links">
        <img src="/icon300.png" alt="Logo" id="logo"/>
        <a href="/dashboard">Dashboard</a>
        <a href="/dashboard/transaction">Transactions</a>
        <a href="/dashboard/media">Uploaded Media</a>
    </div>
    <button id="logout" on:click={logOut}></button>
</div>
