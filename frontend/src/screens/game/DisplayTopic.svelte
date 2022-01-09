<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import TopicDisplay from "../../assets/TopicDisplay.svelte";
    import { leader, ws } from "../../store";

    const sendSkip = (_) => {
        $ws.send("SKIP");
    };

    const sendGif = (_) => {
        if (!gifBuffer) {
            toast.push("You need to supply a gif url");
            return;
        }
        $ws.send(`SUBMIT_GIF ${gifBuffer}`);
    };

    let gifBuffer: string;

    let showGifWindow: boolean = false;
</script>

<!-- Display Round Number -->
<TopicDisplay />

{#if gifBuffer}
    <h1>Your submission</h1>
    <p>
        <img width="200px" src={gifBuffer} alt="submitted gif" />
    </p>
{/if}

<input
    type="text"
    placeholder="Enter a gif url"
    on:click={(_) => (showGifWindow = !showGifWindow)}
    bind:value={gifBuffer}
/>
<div id="searchBarWrapper">
    {#if showGifWindow}
        <div id="gifSearchWrapper" />
    {/if}
</div>
<button on:click={sendGif}>Submit!</button>

<!-- Skip Button for Leader -->
{#if $leader}
    <button on:click={sendSkip}>SKIP</button>
{/if}

<style lang="scss">
    button:hover {
        cursor: pointer;
    }

    #gifSearchWrapper {
        width: 50vw;
        height: 10vh;
        position: absolute;
        top: 0;
        left: 0;
        background-color: #131313;
    }
    #searchBarWrapper {
        position: relative;
    }
</style>
