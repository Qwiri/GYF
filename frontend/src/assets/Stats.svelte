<script lang="ts">
    import { stats, username } from "../store";
    import Avatar from "./Avatar.svelte";

    let ranks = {};
    $: if ($stats) {
        let prev = -1;
        let rank = 0;
        for (const name in $stats) {
            const value = $stats[name];
            if (prev < 0 || value < prev) {
                prev = value;
                rank++;
            }
            ranks[name] = rank;
        }
    }
</script>

<div class="floating">
    <div>
        <ul>
            {#each Object.entries($stats) as [user, count]}
                <li>
                    <span class="rank">{ranks[user] ?? "?"}.</span>

                    <!-- Disply Avatar -->
                    <Avatar {user} width="32px" />

                    <!-- Display Username -->
                    {#if user === $username}
                        <span class="self">{user}</span>:
                    {:else}
                        <span class="username">{user}</span>:
                    {/if}

                    <span class="count">{count}</span>
                </li>
            {/each}
        </ul>
    </div>
</div>

<style lang="scss">
    .floating {
        display: table;
        float: left;

        div {
            display: table-cell;
            vertical-align: middle;
            width: 12%;
            background-color: #131313;
            border-radius: 10px;
            text-align: left;

            ul {
                list-style: none;
                padding: 0;
                margin: 0;
            }

            .username {
                color: greenyellow;
            }
            .self {
                color: coral;
            }
            .rank {
                color: lightcoral;
            }
        }
    }
</style>