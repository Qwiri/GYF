<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";

    import { ws, topics, players, preferences } from "../../store";
    import { copyToClipboard } from "../../utils";

    // other components
    import Swal from "sweetalert2";

    let topicBuffer: string;

    let showManualTopic = false;

    let checkedAutoSkip: boolean;
    let checkedShuffleTopics: boolean;

    $: checkedAutoSkip = $preferences.AutoSkip;
    $: checkedShuffleTopics = $preferences.ShuffleTopics;

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
            confirmButtonColor: "#3085d6",
            cancelButtonColor: "#d33",
            confirmButtonText: "Yes, clear all topics!",
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
        const url = `http://localhost:5000/?t=${btoa(data)}`;

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
            cancelButtonText: "Save Template URL",
        }).then((result) => {
            if (result.value) {
                downloadTopics();
            } else if (result.dismiss.toString() === "cancel") { // is there a better way?
                saveTopics();
            }
        });
    }
</script>

<hr />
<div class="row">
    <h2 class="grayHashtags">#</h2>
    <h2 class="greenText">Topics</h2>
    <h2>({$topics.length})</h2>
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
<!-- display topics -->
<ul>
    {#each $topics as topic}
        <li>
            <button>{topic}</button>
            <button class="removeTopicButton" data-topic={topic} on:click={removeTopic}>❌</button>
        </li>
    {/each}
</ul>

<div id="actionButtonsWrapper">
    <!-- Manual Topic Textbox -->
    {#if !showManualTopic}
        <div
            id="manualTopicButton"
            class="actionButton"
            on:click={(e) => (showManualTopic = true)}
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

    <!-- Save Topics Button -->
    <div class="actionButton saveTopicsButton" on:click={saveMenu}>
        <img src="/assets/saveTopics.svg" alt="" />
        <span>Save</span>
    </div>

    <!-- Clear Topics Button -->
    <div id="clearTopicsButton" class="actionButton" on:click={clearTopics}>
        <img src="/assets/nukeTopics.svg" alt="" />
        <span>Nuke Topics</span>
    </div>
</div>
<hr />

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
        {#if Object.keys($players).length >= 3}
            <button id="startGameButton" class="clickable" on:click={startGame}>Start game!</button>
        {:else}
            <button id="startGameButton">Need {3 - Object.keys($players).length} more players!</button>
        {/if}

    </div>

</div>

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
            margin-left: .3rem;
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
        border-radius: .5rem;
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
    }
    .saveTopicsButton,
    #downloadTopicsButton {
        --background-color: #48aae2;
    }
    .loadFileButton {
        --background-color: #ffcb7e;
    }
    #manualTopicButton {
        --background-color: #27ae60;
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
        }

        span {
            height: 1rem;
            display: inline-flex;
            align-items: center;
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
            margin-left: .3rem;
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
</style>
