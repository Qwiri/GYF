<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import TopicDisplay from "../../assets/TopicDisplay.svelte";
    import { leader, players, round, waitingFor, ws } from "../../store";

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
</script>

<!-- Display Round Number -->
<TopicDisplay />

{#if gifBuffer}
    <p>
        <img width="200px" src={gifBuffer} alt="submitted gif" />
    </p>
{/if}

<input type="text" placeholder="Enter a gif url" bind:value={gifBuffer} />
<button on:click={sendGif}>Submit!</button>

<!-- Skip Button for Leader -->
{#if $leader}
    <button on:click={sendSkip}>SKIP</button>
{/if}

<style lang="scss">
    button:hover {
        cursor: pointer;
    }
</style>