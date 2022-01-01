<script lang="ts">
    import Avatar from "../../assets/Avatar.svelte";
    import TopicEditor from "../../assets/setup/TopicEditor.svelte";
    import { leader, players, username } from "../../store";
</script>

<!-- Show connected players -->
<div id="playerBar">
    {#each Object.values($players) as player}
        <div class="playerCard">
            <Avatar user={player.name} width="100px" />
            <p class="playerName">
                {#if player.leader}
                    👑
                {/if}
                <span class:self={player.name === $username}>{player.name}</span
                >
            </p>
        </div>
    {/each}
</div>

<!-- Leader specific actions -->
{#if $leader}
    <TopicEditor />
{/if}

<style lang="scss">
    #playerBar {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
        gap: 10px;
        grid-auto-rows: minmax(100px, auto);

        .playerCard {
            margin: 10px;
            padding: 10px;
            border-radius: 10px;
            background-color: #101010;
            color: white;
            font-size: 1.2em;
        }

        .playerName {
            margin-top: 5px;

            .self {
                color: greenyellow;
                font-weight: bold;
            }
        }
    }
</style>
