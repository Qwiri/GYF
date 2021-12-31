<script lang="ts">
    import Avatar from "../assets/Avatar.svelte";
    import { players, leader } from "../store";
    import type { Player } from "../types";

    let localPlayers: Array<Player>;
    players.subscribe((n) => {
        localPlayers = Object.values(n);
        console.log("localPlayers", localPlayers);
    });

    console.log("localPlayers (init):", localPlayers);
</script>

<div id="playerBar">
    {#each localPlayers as player}
        <Avatar user={player.name} width="100px" />
        <h2>
            {#if player.leader}
                👑
            {/if}
            {player.name}
        </h2>
    {/each}
</div>

