<script lang="ts">
    import Avatar from "../assets/Avatar.svelte";
    import { players } from "../store";
</script>

<div id="playerBar">
    {#each Object.values($players) as player}
        <Avatar user={player.name} width="100px" />
        <h2>
            {#if player.leader}
                👑
            {/if}
            {player.name}
        </h2>
    {/each}
</div>

