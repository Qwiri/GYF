<script>
import Avatar from "../../assets/Avatar.svelte";
import Stats from "../../assets/Stats.svelte";

import { state, stats, waitingFor } from "../../store";
import { GameState } from "../../types";

// components
import DisplayTopic from "../game/DisplayTopic.svelte";
import Voting from "../game/Voting.svelte";
import VotingResults from "../game/VotingResults.svelte";
</script>

<!-- Player Leaderboard -->
{#if $stats && Object.keys($stats).length > 0}
    <Stats />
{/if}

<!-- Game -->
{#if $state == GameState.SubmitGIF}
    <DisplayTopic />
{:else if $state == GameState.Vote}
    <Voting />
{:else if $state == GameState.VoteResults}
    <VotingResults />
{/if}

<!-- Waiting For Block -->
{#if $waitingFor && $waitingFor.length > 0}
    <hr />
    <h3>
        Waiting for 
        <span class="waiting">{$waitingFor.length}</span> 
        more people
    </h3>
    {#each $waitingFor as player}
        <Avatar user={player} width="32px" />
    {/each}
    <hr />
{/if}

<style lang="scss">
.waiting {
    color: greenyellow;
}
</style>