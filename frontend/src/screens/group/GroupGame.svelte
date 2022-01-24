<script>
    import Avatar from "../../assets/Avatar.svelte";
    import Chat from "../../assets/Chat.svelte";
    import Stats from "../../assets/Stats.svelte";

    import { leader, state, stats, waitingFor, ws } from "../../store";
    import { GameState } from "../../types";

    // components
    import Voting from "../game/Voting.svelte";
    import VotingResults from "../game/VotingResults.svelte";
    import SearchGif from "../SearchGif.svelte";
    import GameEnd from "../lobby/GameEnd.svelte";
    import { toast } from "@zerodevx/svelte-toast";
    import Swal from "sweetalert2";

    function skip() {
        toast.push("Skipping...");
        if ($state === GameState.VotingResults) {
            $ws.send("NEXT_ROUND");
        } else {
            $ws.send("SKIP");
        }
    }

    function endGame() {
        Swal.fire({
            title: "Are you sure?",
            text: "Do you really want to end the game?",
            icon: "warning",
            showCancelButton: true,
            confirmButtonColor: "limegreen",
            cancelButtonColor: "#d33",
            confirmButtonText: "Yes!",
            background: "#131313",
            color: "white"
        }).then((result) => {
            if (result.value) {
                $ws.send("END_GAME");
            }
        });
    }
</script>

<!-- Player Leaderboard -->
<div id="wholeScreen">
    {#if $state == GameState.GameEnd}
        <GameEnd />
    {:else}
            
        <div id="statsWindow" class="screenSub">
            {#if $stats && Object.keys($stats).length > 0}
                <Stats />
            {/if}
        </div>
        <div id="screenMain">
            <!-- Game -->
            {#if $state == GameState.SubmitGIF}
                <SearchGif />
            {:else if $state == GameState.Vote}
                <Voting />
            {:else if $state == GameState.VoteResults}
                <VotingResults />
            {/if}
            <!-- Waiting For Block -->
            {#if $waitingFor && $waitingFor.length > 0}
                <hr />
                <p>
                    Waiting for
                    <span class="waiting">{$waitingFor.length}</span>
                    more {$waitingFor.length == 1 ? "player" : "players"} ⏱
                </p>
                {#each $waitingFor as player}
                    <Avatar user={player} width="32px" />
                {/each}
                <hr />
            {/if}

            <!-- Skip Button -->
            {#if $leader}
                {#if $waitingFor && $waitingFor.length > 0}
                    <button class="btn-force" on:click={skip}
                        >Force Continue</button
                    >
                {:else}
                    <button class="btn-continue" on:click={skip}
                        >Continue</button
                    >
                {/if}
                <!-- End game button -->
                <button class="btn-force" on:click={endGame}>End Game</button>
            {/if}
        </div>
        <div id="chatContainer" class="screenSub">
            <Chat />
        </div>
    {/if}
</div>

<style lang="scss">
    hr {
        border-color: #010101;
    }

    #wholeScreen {
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        min-height: 100vh;
        align-items: center;
        justify-content: center;
    }

@media (min-width: 40rem) {
    .screenSub {

        width: 25vw;
    }
    #screenMain {

        width: 50vw;
    }
}

    #screenMain {
        flex-grow: 1;
        flex-shrink: 1;
        flex-basis: 20rem;
    }

    #chatContainer {
        align-self: flex-end;
    }
    .waiting {
        color: greenyellow;
    }

    button {
        font-weight: bold;
        border: none;
        border-radius: 7px;

        font-size: 1.3rem;
        margin-top: 1em;

        &:hover {
            cursor: pointer;
        }
    }

    .btn-force {
        background-color: salmon;
    }
    .btn-continue {
        background-color: lightgreen;
    }
</style>
