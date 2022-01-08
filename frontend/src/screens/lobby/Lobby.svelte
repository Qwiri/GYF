<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";

    import Avatar from "../../assets/Avatar.svelte";
    import TopicEditor from "../../assets/setup/TopicEditor.svelte";
    import { leader, players, preferences, username, ws } from "../../store";

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
        const copyText = document.createElement("textarea");
        copyText.value = getShare()[0];
        document.body.appendChild(copyText);
        copyText.select();
        document.execCommand("copy");
        document.body.removeChild(copyText);

        toast.push("Copied invite URL to clipboard!");
    }

    let checkedAutoSkip: boolean;
    let checkedShuffleTopics: boolean;

    $: checkedAutoSkip = $preferences.AutoSkip;
    $: checkedShuffleTopics = $preferences.ShuffleTopics;

    function clickChangeAutoSkip(event: MouseEvent) {
        $ws.send(
            "CHANGE_PREF " +
                JSON.stringify({
                    key: "AutoSkip",
                    value: !checkedAutoSkip,
                })
        );
    }

    function clickChangeShuffleTopics(event: MouseEvent) {
        $ws.send(
            "CHANGE_PREF " +
                JSON.stringify({
                    key: "ShuffleTopics",
                    value: !checkedShuffleTopics,
                })
        );
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

    <div class="leaderActions">
        <input
            id="cbChangeAutoSkip"
            type="checkbox"
            on:click={clickChangeAutoSkip}
            bind:checked={checkedAutoSkip}
        />
        <label
            for="cbChangeAutoSkip"
            style="color: {checkedAutoSkip ? '#24FF00' : 'salmon'};"
        >
            Auto Skip
        </label>

        <input
            id="cbChangeShuffleTopics"
            type="checkbox"
            on:click={clickChangeShuffleTopics}
            bind:checked={checkedShuffleTopics}
        />
        <label
            for="cbChangeShuffleTopics"
            style="color: {checkedShuffleTopics ? '#24FF00' : 'salmon'};"
        >
            Shuffle Topics
        </label>
    </div>
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
        display: flex;
        flex-wrap: wrap;
        justify-content: space-evenly;
        margin-bottom: 1rem;
        gap: 1rem;

        .playerCard {
            padding: .8rem;
            border-radius: .8rem;
            background-color: #101010;
            color: white;
            font-size: 1.2em;
            width: 10rem;
            height: 10rem;
        }

        .playerName {
            margin-top: 5px;

            .self {
                color: #24FF00;
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
                color: #24FF00;
            }
        }

        // make button nice
        button {
            border: none;
            border-radius: 7px;
            color: #131313;
            font-size: 1.2em;
            background-color: #24FF00;
            font-weight: bold;

            &:hover {
                cursor: pointer;
            }
        }
    }
</style>
