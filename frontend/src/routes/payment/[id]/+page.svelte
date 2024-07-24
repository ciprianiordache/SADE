<script>
    import { onMount } from "svelte";
    import { loadStripe } from "@stripe/stripe-js";
    import { page } from "\$app/stores";
    import { apiHost, gatewayKey } from "$lib/utils.js";

    let stripe;
    let error = '';
    let cardNumber, cardExpiry, cardCvc, cardZip
    let mediaId;
    let price = 0;

    $: mediaId = $page.params.id;

    onMount(async () => {
       stripe = await loadStripe(gatewayKey);
       const elements = stripe.elements();

       cardNumber = elements.create('cardNumber');
       cardExpiry = elements.create('cardExpiry');
       cardCvc = elements.create('cardCvc');
       cardZip = elements.create('postalCode', { placeholder: 'ZIP' });

       cardNumber.mount('#card-number');
       cardExpiry.mount('#card-expiry');
       cardCvc.mount('#card-cvc')
       cardZip.mount('#card-postal-code');

       const response = await fetch(apiHost+`/unlock/${mediaId}`, {
           method: 'GET',
           credentials: 'include',
           headers: {
               'Content-Type': 'application/json',
           },
       });
       const media = await response.json();
       if (response.ok) {
           price = media.price;
       } else {
           error = 'Failed to fetch media data!';
       }
    });

    const handleSubmit = async (event) => {
        event.preventDefault();
        const { token, error: stripeError } = await stripe.createToken(cardNumber);

        if (stripeError) {
            error = stripeError.message;
            return;
        }
        const response = await fetch(apiHost + '/process', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
                token: token.id,
                media_id: parseInt(mediaId),
                description: 'Payment for media access.',
                amount: price * 100,
                currency: 'eur',
            }),
        });
        const result = await response.json();
        if (response.ok) {
            alert('Payment successful!');
            window.location.href = `/original?media_id=${mediaId}`;
        } else {
            error = result.error;
        }
    }
</script>

<style>

    .container {
        font-family: monospace;
        display: flex;
        align-items: center;
        justify-content: center;
        height: 90vh;
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

    .StripeElement {
        font-family: monospace;
        box-sizing: border-box;
        height: 40px;
        padding: 10px 12px;
        border: solid 1.5px #D3D3D3;
        background-color: white;
    }

    button {
        font-family: monospace;
        padding: 10px;
        border: none;
        background: coral;
        cursor: pointer;
        height: 35px;
        margin-top: 10px;
        margin-bottom: -10px;
        transition: background-color 0.3s;
    }
    button:hover {
        background: orange;
    }

    .title {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    img {
        width: 100px;
        height: auto;
        margin-bottom: 1rem;
    }

    .StripeElement--focus {
        box-shadow: 0 1px 3px 0 #cfd7df;
    }

    .StripeElement--invalid {
        border-color: #fa755a;
    }

    .StripeElement--webkit-autofill {
        background-color: #fefde5 !important;
    }
</style>
<div class="container">
    <form on:submit={handleSubmit}>
        <div class="title">
            <img src="/icon300.png" alt="Logo" class="logo"/>
        </div>
        <label for="card-number">Card Number</label>
        <div id="card-number" class="StripeElement"></div>

        <label for="card-expiry">Expiration Date</label>
        <div id="card-expiry" class="StripeElement"></div>

        <label for="card-cvc">CVC</label>
        <div id="card-cvc" class="StripeElement"></div>

        <label for="card-postal-code">ZIP Code</label>
        <div id="card-postal-code" class="StripeElement"></div>

        <button type="submit">Pay {price} EUR</button>
    </form>
</div>
{#if error}
    <p style="color: red;">{error}</p>
{/if}