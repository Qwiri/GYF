<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";

    import Avatar from "../../assets/Avatar.svelte";
    import TopicEditor from "../../assets/setup/TopicEditor.svelte";
    import { leader, players, preferences, username, ws } from "../../store";
    import { copyToClipboard } from "../../utils";

    $: checkedPermTopicList = ($preferences.Permissions & 0b1) === 0b1;

    /**
     * Please refactor this. Thanks! :)
     */
    function getShare(): Array<string> {
        const urlPieces = [
            location.protocol,
            "//",
            location.host,
            location.pathname,
        ];
        let url = urlPieces.join("");

        // remove game ID from pathname
        const i = url.indexOf("/game/") + 6;
        const gameID = url.slice(i);
        url = url.slice(0, i);

        return [url + gameID, url, gameID];
    }

    function copyShareURL(_: MouseEvent) {
        copyToClipboard(getShare()[0]);
        toast.push("Copied invite URL to clipboard!");
    }
</script>

<div id="lobbyWrapper">
    <div id="avatarDiv">
        <Avatar user={$username} width="auto" />

        <!-- share game -->
        <div id="share">
            <div id="shareTxt">
                {getShare()[1]}<span>{getShare()[2]}</span>
            </div>
            <button on:click={copyShareURL}>COPY</button>
        </div>
    </div>
    <div id="lobbyRight">
        <!-- Show connected players -->
        <div class="row">
            <h1>Hi,</h1>
            <h1 class="greenFontColor">{$username}</h1>
            <h1>!</h1>
        </div>
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
        {#if $leader || $preferences.Permissions !== 0 }
            <TopicEditor />
        {/if}
    </div>
</div>

<style lang="scss">
    #playerBar {
        display: grid;
        place-content: center;
        grid-template-columns: repeat(auto-fill, 10rem);
        margin-bottom: 1rem;
        gap: 1rem;

        .playerCard {
            padding: 1rem;
            border-radius: 0.8rem;
            background-color: #101010;
            color: white;
            font-size: 1.2em;
            width: 8rem;
            height: 8rem;
        }

        .playerName {
            margin-top: 0rem;
            overflow: hidden;
            display: flex;
            justify-content: center;
            align-content: center;

            .self {
                color: #24ff00;
                font-weight: bold;
            }
        }
    }

    #share {
        position: absolute;
        bottom: 2rem;
        left: 2rem;
        display: flex;
        align-items: center;
        justify-content: center;

        filter: drop-shadow(0px 2px 4px black);
        border-radius: 0.5rem;
        background-color: #131313;
        padding: 0.5rem;

        #shareTxt {
            width: min-content;
            border: none;
            color: white;
            font-size: 1.2em;
            padding: 0.5rem;
            span {
                color: #ffcb7e;
            }
        }

        // make button nice
        button {
            margin: 0;
            border: none;
            border-radius: 7px;
            color: #131313;
            font-size: 1.2em;
            background-color: #ffcb7e;
            font-weight: bold;

            &:hover {
                cursor: pointer;
            }
        }
    }
    #avatarDiv {
        align-self: flex-end;
        position: relative;
        width: 40vw;
        display: flex;
        justify-content: center;
        z-index: -1;

        :global(img) {
            height: 50vw;
            pointer-events: none;
        }

        @media (max-width: 40rem) {
            display: none;
        }
    }
    #lobbyRight {
        display: flex;
        flex-direction: column;
        justify-content: center;

        @media (min-width: 40em) {
            width: 35vw;
        }
    }
    #lobbyWrapper {
        min-height: 100vh;
        display: flex;
        justify-content: space-evenly;
    }
    .row {
        display: flex;
    }
    h1 {
        font-size: 3.5rem;
    }
    .greenFontColor {
        color: #24ff00;
        margin-left: 0.5ch;
    }
</style>
