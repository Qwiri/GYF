<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import { leader, round, ws } from "../store";

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
<p style="color:greenyellow">
    {$round.currentRound}/{$round.totalRounds}
</p>

<!-- Display Topic -->
<h2>{$round.topic}</h2>
<input type="text" placeholder="Enter a gif url" bind:value={gifBuffer} />
<button on:click={sendGif}>Submit!</button>
{#if gifBuffer}
    <img width="200px" src={gifBuffer} alt="submitted gif" />
{/if}

<!-- Skip Button for Leader -->
{#if $leader}
    <button on:click={sendSkip}>Skip round</button>
{/if}
