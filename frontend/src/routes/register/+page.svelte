<script>
    import { goto } from '$app/navigation';
    import { apiHost } from "$lib/utils.js";

    let email = '';
    let firstName = '';
    let lastName = '';
    let password = '';

    async function handleRegister(event) {
        event.preventDefault()
        try {
            const response = await fetch(apiHost+'/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams({ email, first_name: firstName, last_name: lastName, password:  password})
            });
            const result = await response.json();
            if(response.ok) {
                goto('/message');
            } else {
                alert(`${result.error}`);
            }
        } catch (error) {
            console.error('Error: ', error);
            alert('An unexpected error occurred!!!')
        }
    }
</script>


<div class="register_container">
    <form on:submit={handleRegister}>
        <div class="title">
            <img src="/icon300.png" alt="Logo" class="logo"/>
        </div>
        <h1>Welcome</h1>
        <h2>Register to SADE</h2>
        <label for="email">Email</label>
        <input id="email" type="email" bind:value={email} required>
        <label for="password">Password</label>
        <input id="password" type="password" bind:value={password} required>
        <label for="first_name">First Name</label>
        <input id="first_name" type="text" bind:value={firstName} required>
        <label for="last_name">Last Name</label>
        <input id="last_name" type="text" bind:value={lastName} required>
        <button type="submit">Register</button>
    </form>

</div>

<style>
    .register_container {
        font-family: monospace;
        display: flex;
        align-items: center;
        justify-content: center;
        height: 90vh;
    }
    .title {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    h2 {
        font-size: large;
    }

    form {
        display: flex;
        flex-direction: column;
        width: 350px;
        background-color: whitesmoke;
        padding: 30px;
        border-color: whitesmoke;
        border-style: solid;
        border-width: 1px;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    }
    img {
        width: 100px;
        height: auto;
        margin-bottom: 1rem;
    }
    input {
        font-family: monospace;
        margin-top: 10px;
        margin-bottom: 15px;
        padding: 8px;
        font-size: 20px;
        border: solid 1.5px #D3D3D3;
        background-color: white;
        height: 40px;
    }

    h1 {
        margin-bottom: -5px;
    }


    button {
        font-family: monospace;
        padding: 10px;
        border: none;
        background: coral;
        cursor: pointer;
        height: 45px;
        margin-bottom: -10px;
        transition: background-color 0.3s;
    }
    button:hover {
        background: orange;
    }
</style>
