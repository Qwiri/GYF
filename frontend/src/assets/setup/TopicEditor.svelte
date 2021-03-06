<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";

    import { ws, topics, players, preferences, leader } from "../../store";
    import { copyToClipboard } from "../../utils";

    // other components
    import Swal from "sweetalert2";

    let topicBuffer = "";
    let showManualTopic = false;

    $: checkedAutoSkip = $preferences.AutoSkip;
    $: checkedShuffleTopics = $preferences.ShuffleTopics;

    // permissions
    $: checkedPermTopicList = ($preferences.Permissions & 0b1) === 0b1;
    $: checkedPermTopicCreate = ($preferences.Permissions & 0b10) === 0b10;
    $: checkedPermTopicDelete = ($preferences.Permissions & 0b100) === 0b100;

    const togglePermission = (perm: number) => {
        if (($preferences.Permissions & perm) === perm) {
            $ws.send(`PERM ${$preferences.Permissions & ~perm}`);
        } else {
            $ws.send(`PERM ${$preferences.Permissions | perm}`);
        }
    };

    const sendTopic = (event: KeyboardEvent) => {
        if (event.key === "Enter") {
            // send & clear topic buffer
            $ws.send(`TOPIC_ADD ${topicBuffer}`);
            topicBuffer = "";
            // update topic list from backend
            $ws.send(`TOPIC_LIST`);
        }
    };

    const removeTopic = (event: MouseEvent) => {
        const topic = event.srcElement.dataset.topic;
        $ws.send(`TOPIC_REMOVE ${topic}`);
        $ws.send(`TOPIC_LIST`);
    };

    const startGame = (event: MouseEvent) => {
        $ws.send("START");
    };

    function loadFromFile() {
        const file = this.files[0];
        const reader = new FileReader();

        reader.onload = (e: any) => {
            const data = e.target.result;
            const lines = data.split("\n");

            toast.push(
                `Sending <strong>${lines.length} lines</strong> to backend`
            );
            $ws.send(`TOPIC_ADD_ALL ${JSON.stringify(lines)}`);
        };

        reader.readAsText(file);
    }

    const clearTopics = () => {
        // Swal confirmation dialog to clear all topics
        // on confirmation, execute backend command
        Swal.fire({
            title: "Are you sure?",
            text: "You won't be able to revert this!",
            icon: "warning",
            showCancelButton: true,
            confirmButtonColor: "limegreen",
            cancelButtonColor: "#d33",
            confirmButtonText: "Yes, clear all topics!",
            background: "#131313",
            color: "white"
        }).then((result) => {
            if (result.value) {
                $ws.send("TOPIC_CLEAR");
            }
        });
    };

    const downloadTopics = () => {
        if ($topics.length === 0) {
            toast.push("No topics to download");
            return;
        }

        const data = new Blob([$topics.join("\n")], {
            type: "text/plain;charset=utf-8",
        });
        const url = URL.createObjectURL(data);
        const a = document.createElement("a");
        a.href = url;
        a.download = "topics.txt";
        a.click();
    };

    const saveTopics = () => {
        if ($topics.length === 0) {
            toast.push("No topics to share");
            return;
        }

        const data = JSON.stringify($topics);

        const urlPieces = [location.protocol, "//", location.host].join("");
        const url = `${urlPieces}/?t=${btoa(data)}`;

        Swal.fire({
            title: "Bookmark this URL",
            html: `and <strong>use that URL to create a new game</strong> with the topics you have created:
            <br />
            <input class="gyf-bar" type="text" value="${url}" style="width: 100%" readonly />`,
            showCancelButton: true,
            confirmButtonText: "Copy to clipboard",
            cancelButtonText: "Close",
            cancelButtonColor: "#d33",
            confirmButtonColor: "#28a745",
            reverseButtons: true,
        }).then((result) => {
            if (result.value) {
                copyToClipboard(url);
                toast.push("Copied topic create URL to clipboard");
            }
        });
    };
    function clickChangeAutoSkip(_: MouseEvent) {
        $ws.send(
            "CHANGE_PREF " +
                JSON.stringify({
                    key: "AutoSkip",
                    value: !checkedAutoSkip,
                })
        );
    }
    function clickChangeShuffleTopics(_: MouseEvent) {
        $ws.send(
            "CHANGE_PREF " +
                JSON.stringify({
                    key: "ShuffleTopics",
                    value: !checkedShuffleTopics,
                })
        );
    }
    const saveMenu = (_: MouseEvent) => {
        // sweetalert with two options: "download" and "save template"
        Swal.fire({
            title: "Save menu",
            text: "What do you want to do?",
            icon: "question",
            showCancelButton: true,
            confirmButtonColor: "#48aae2",
            cancelButtonColor: "#48aae2",
            confirmButtonText: "Download",
            cancelButtonText: "Save Topic URL",
        }).then((result) => {
            if (result.value) {
                downloadTopics();
            } else if (result.dismiss.toString() === "cancel") {
                // is there a better way?
                saveTopics();
            }
        });
    };

    function ok(b: boolean): boolean {
        return $leader || b;
    }

    const promptNewPassword = (_: MouseEvent) => {
        Swal.fire({
            title: "Change password",
            text: "Enter a new password or leave empty to remove it",
            input: "password",
            inputAttributes: {
                autocapitalize: "off",
                autocorrect: "off",
                autocomplete: "off",
            },
            showCancelButton: true,
            confirmButtonColor: "#48aae2",
            cancelButtonColor: "#48aae2",
            confirmButtonText: "Change",
            cancelButtonText: "Remove",
        }).then((result) => {
            if (result.isConfirmed) {
                $ws.send(`CHANGE_PASS ${result.value ?? ''}`);
            } else if (result.dismiss.toString() === "cancel") {
                $ws.send("CHANGE_PASS");
            }
        });
    };
</script>

<hr />
<div class="row">
    <h2 class="grayHashtags">#</h2>
    <h2 class="greenText">Topics</h2>
    <h2>({$topics.length})</h2>

    {#if $leader}
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

        <!-- Extra Permission Checkboxes -->
        <input
            id="cbChangePermTopicList"
            type="checkbox"
            on:click={(_) => togglePermission(0b1)}
            bind:checked={checkedPermTopicList}
        />
        <label
            for="cbChangePermTopicList"
            style="color: {checkedPermTopicList ? '#24FF00' : 'salmon'};"
        >
            Perm: List Topics
        </label>

        <input
            id="cbChangePermTopicAdd"
            type="checkbox"
            on:click={(_) => togglePermission(0b10)}
            bind:checked={checkedPermTopicCreate}
        />
        <label
            for="cbChangePermTopicAdd"
            style="color: {checkedPermTopicCreate ? '#24FF00' : 'salmon'};"
        >
            Perm: Add Topics
        </label>

        <input
            id="cbChangePermTopicRemove"
            type="checkbox"
            on:click={(_) => togglePermission(0b100)}
            disabled={!checkedPermTopicList}
            bind:checked={checkedPermTopicDelete}
        />
        <label
            for="cbChangePermTopicRemove"
            style="color: {checkedPermTopicDelete ? '#24FF00' : 'salmon'};"
        >
            Perm: Remove Topics
        </label>
    {/if}
</div>

<!-- Display Topics List -->
{#if ok(checkedPermTopicList)}
    <ul>
        {#if $topics && $topics.length > 0}
            {#each $topics as topic}
                <li>
                    <button>{topic}</button>

                    <!-- Display Delete Button for Leader or if Enhanced Permissions -->
                    {#if $leader || checkedPermTopicDelete}
                        <button
                            class="removeTopicButton"
                            data-topic={topic}
                            on:click={removeTopic}>❌</button
                        >
                    {/if}
                </li>
            {/each}
        {:else}
            <li style="color: salmon;">No topics yet ☹️</li>
        {/if}
    </ul>
{/if}

<div id="actionButtonsWrapper">
    <!-- Topic Create -->
    {#if ok(checkedPermTopicCreate)}
        {#if !showManualTopic}
            <div
                id="manualTopicButton"
                class="actionButton"
                on:click={(_) => (showManualTopic = true)}
            >
                <img src="/assets/addTopic.svg" alt="" />
            </div>
        {:else}
            <input
                placeholder="Add topic (Enter)"
                class="gyf-bar"
                type="text"
                on:blur={(e) => {
                    showManualTopic = false;
                }}
                on:keypress={sendTopic}
                bind:value={topicBuffer}
            />
        {/if}

        <label id="loadFromFileLabel" class="actionButton loadFileButton">
            <input type="file" on:change={loadFromFile} />
            <img src="/assets/import_checkmark.svg" alt="" />
            <span>Import</span>
        </label>
    {/if}

    <!-- Topic List -->
    {#if ok(checkedPermTopicList)}
        <!-- Save Topics Button -->
        <div class="actionButton saveTopicsButton" on:click={saveMenu}>
            <img src="/assets/saveTopics.svg" alt="" />
            <span>Save</span>
        </div>
    {/if}

    <!-- Clear Topics Button -->
    {#if ok(checkedPermTopicDelete)}
        <div id="clearTopicsButton" class="actionButton" on:click={clearTopics}>
            <img src="/assets/nukeTopics.svg" alt="" />
            <span>Nuke Topics</span>
        </div>
    {/if}

    {#if $leader}
        <div
            id="lockLobbyButton"
            class="actionButton"
            on:click={promptNewPassword}
        >
            <img src="/assets/lock.svg" alt="" />
            <span>Lock Lobby</span>
        </div>
    {/if}
</div>
<hr />

<!-- TODO: Move this to Lobby.svelte @Tom -->
{#if $leader}
    <div id="startGameDiv">
        <div id="startGameRow">
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

            <!-- start game button -->
            {#if !$topics || $topics.length <= 0}
                <button id="startGameButton" class="button-secondary"
                    >Need more topics!</button
                >
            {:else if Object.keys($players).length >= 3}
                <button
                    id="startGameButton"
                    class="clickable"
                    on:click={startGame}>Start game!</button
                >
            {:else}
                <button id="startGameButton" class="button-secondary"
                    >Need {3 - Object.keys($players).length} more players!</button
                >
            {/if}
        </div>
    </div>
{/if}

<style lang="scss">
    hr {
        width: 100%;
    }
    h2 {
        margin: 0.5rem 0.5ch;
    }
    .grayHashtags {
        color: #373737;
    }
    .greenText {
        color: #24ff00;
    }
    .row {
        display: flex;
        align-items: center;

        label {
            height: 1rem;
            display: inline-flex;
            align-items: center;
            margin-left: 0.3rem;
        }
        input {
            margin: 0;
        }
    }
    li {
        list-style: none;
        background-color: #1f1f1f;
        width: max-content;
        margin-bottom: 0.5rem;
        padding: 0.5rem;
        border-radius: 0.5rem;
    }

    ul {
        @media (max-width: 40em) {
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        button {
            background-color: transparent;
            color: white;
            font-weight: normal;
            margin: 0;
            padding: 0;

            &:hover {
                cursor: default;
            }
        }
        .removeTopicButton {
            &:hover {
                cursor: pointer;
            }
        }
    }

    #actionButtonsWrapper {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;
        margin-bottom: 1rem;

        @media (max-width: 40em) {
            justify-content: center;
        }
    }

    input[type="file"] {
        position: absolute;
        opacity: 0;
    }
    .clickable:hover {
        cursor: pointer;
    }

    #loadFromFileLabel {
        position: relative;
        display: flex;
        justify-content: center;
        align-items: center;
        overflow: hidden;

        input:hover {
            cursor: pointer;
        }
    }
    #clearTopicsButton {
        --background-color: #e2778b;
        --background-color-transparent: #e2778be0;
    }
    .saveTopicsButton {
        --background-color: #48aae2;
        --background-color-transparent: #48aae2e0;
    }
    .loadFileButton {
        --background-color: #ffcb7e;
        --background-color-transparent: #ffcb7ee0;
    }
    #manualTopicButton {
        --background-color: #27ae60;
        --background-color-transparent: #27ae5fe0;
    }
    #lockLobbyButton {
        --background-color: #ffcb7e;
        --background-color-transparent: #ffcb7ee0;
    }

    .actionButton {
        flex-shrink: 0;

        background-color: transparent;
        border-color: var(--background-color);
        border-style: solid;
        border-width: 2px;
        color: white;
        padding: 0.5rem 1rem;
        border-radius: 0.5rem;

        text-decoration: underline;
        text-decoration-color: var(--background-color);

        font-size: 0.8rem;

        display: flex;
        justify-content: center;
        align-items: center;
        gap: 0.5rem;

        &:hover {
            cursor: pointer;

            animation-name: animateBackground;
            animation-timing-function: linear;
            animation-fill-mode: forwards;
            animation-direction: normal;
            animation-duration: 0.1s;
        }

        span {
            height: 1rem;
            display: inline-flex;
            align-items: center;
        }
    }
    @keyframes animateBackground {
        from {
            background-color: transparent;
            color: white;
        }
        to {
            background-color: var(--background-color-transparent);
            color: black;
        }
    }
    #startGameDiv {
        display: flex;
        justify-content: end;
        align-items: end;
        min-height: 10rem;
    }
    #startGameRow {
        display: flex;
        justify-content: end;
        align-items: center;

        label {
            height: 1rem;
            display: inline-flex;
            align-items: center;
            margin-left: 0.3rem;
        }
        input {
            margin: 0;
        }
    }
    #startGameButton {
        @media (min-width: 40em) {
            font-size: 1.5rem;
        }

        margin-left: 1rem;
    }

    .button-secondary {
        background-color: #ffffff;
        font-weight: normal;
        cursor: default;
    }
</style>
