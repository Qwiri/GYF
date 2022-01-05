<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";

    import Avatar from "../../assets/Avatar.svelte";
    import TopicEditor from "../../assets/setup/TopicEditor.svelte";
    import { leader, players, username } from "../../store";

    /**
     * Please refactor this. Thanks! :)
     */
    function getShare(): Array<string> {
        const urlPieces = [
            location.protocol,
            "//",
            location.host,
            location.pathname
        ];
        let url = urlPieces.join("");

        // remove game ID from pathname
        const i = url.indexOf("/game/") + 6;
        const gameID = url.slice(i);
        url = url.slice(0, i);

        return [
            url + gameID, 
            url, 
            gameID
        ];
    }

    function copyShareURL(_: MouseEvent) {
        const copyText = document.createElement("textarea");
        copyText.value = getShare()[0];
        document.body.appendChild(copyText);
        copyText.select();
        document.execCommand("copy");
        document.body.removeChild(copyText);

        toast.push("Copied invite URL to clipboard!");
    }
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
                <span class:self={player.name === $username}>
                    {player.name}
                </span>
            </p>
        </div>
    {/each}
</div>

<!-- Leader specific actions -->
{#if $leader}
    <TopicEditor />
{/if}

<!-- share game -->
<div id="share">
    <div id="shareTxt">
        {getShare()[1]}<span>{getShare()[2]}</span>
    </div>
    <button on:click={copyShareURL}>COPY</button>
</div>

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

    #share {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        #shareTxt {
            width: min-content;
            border: none;
            border-radius: 7px;
            color: white;
            font-size: 1.2em;
            background-color: #131313;
            padding: 0.5rem;
            span {
                color: greenyellow;
            }
        }

        // make button nice
        button {
            border: none;
            border-radius: 7px;
            color: #131313;
            font-size: 1.2em;
            background-color: greenyellow;
            font-weight: bold;

            &:hover {
                cursor: pointer;
            }
        }
    }
</style>
