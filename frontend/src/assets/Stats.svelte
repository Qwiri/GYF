<script lang="ts">
    import { stats, username } from "../store";
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

                <!-- Avatar + Name -->
                <td class="user">
                    <!-- Disply Avatar -->
                    <Avatar {user} width="32px" />
                    <!-- Display Username -->
                    <span class:self="{user === $username}">{user}</span>:
                </td>

                <!-- Vote Count -->
                <td class="count">
                    {count}
                </td>
            </tr>
        {/each}
    </table>
</div>

<style lang="scss">
    div {
        display: table-cell;
        vertical-align: middle;
        width: 12%;
        background-color: #131313;
        border-radius: 10px;
        text-align: left;
        padding: 10px;

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
        
    }
</style>
