<script lang="ts">
    import { leader, round, ws } from "../store";

    const sendSkip = (_) => {
        $ws.send("SKIP");
    };
</script>

<!-- Display Round Number -->
<p style="color:greenyellow">
    {$round.currentRound}/{$round.totalRounds}
</p>

<!-- Display Topic -->
<h2>{$round.topic}</h2>

<!-- Skip Button for Leader -->
{#if $leader}
    <button on:click={sendSkip}>Skip round</button>
{/if}
