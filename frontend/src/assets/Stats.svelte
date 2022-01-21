<script lang="ts">
    import { leader, stats, username, ws } from "../store";
    import Avatar from "./Avatar.svelte";

    // the icon is displayed before the user's rank
    const icons = {
        1: "👑",
    };

    const ranks = {};

    // make sure the same stat values result in the same rank instead of counting 1 up for every player
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

    const kickPlayer = (e: MouseEvent) => {
        const playername: HTMLElement = e.target.dataset.username;
        $ws.send(`KICK ${playername}`);
    }

</script>

<div>
    <table>
        {#each Object.entries($stats) as [user, count]}
            <tr>
                <!-- Icon -->
                <td class="icon">
                    {#if icons[ranks[user]]}
                        {icons[ranks[user]]}
                    {/if}
                </td>

                <!-- Rank -->
                <td class="rank">
                    {#if ranks[user]}
                        {ranks[user]}.
                    {/if}
                </td>

                <td>
                    <!-- Disply Avatar -->
                    <Avatar {user} width="32px" />
                </td>

                <!-- Name -->
                <td class="user">
                    <!-- Display Username -->
                    <span class:self="{user === $username}">{user}</span>:
                </td>

                <!-- Vote Count -->
                <td class="count">
                    {count}
                </td>

                {#if $leader}
                    <td><span class="hover" data-username={user} on:click={kickPlayer}>🥊</span></td>
                {/if}
            </tr>
        {/each}
    </table>
</div>

<style lang="scss">
    div {
        // display: table-cell;
        vertical-align: middle;
        background-color: #131313;
        border-radius: 1rem;
        text-align: left;
        padding: 1rem;
        min-width: 60%;
        max-width: min-content;

        .self {
            color: greenyellow;
        }
        .rank {
            color: lightcoral;
        }

        table {
            border: none;
        }

        td {
            .rank {
                text-align: right;
            }
            .icon {
                text-align: right;
            }
            .user {
                text-align: left;
            }
            .count {
                text-align: left;
                vertical-align: bottom;
            }
        }
        .hover {
            &:hover {
                cursor: pointer;
            }
        }
        
    }
</style>
