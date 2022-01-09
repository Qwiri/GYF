<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";

    import { ws, topics, players } from "../../store";
    import { copyToClipboard } from "../../types";

    // other components
    import Swal from "sweetalert2";

    let topicBuffer: string;

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
        const topic = event.srcElement.innerText.slice(0, -1);
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
        $ws.send("TOPIC_CLEAR");
    };

    const downloadTopics = (_: MouseEvent) => {
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

    const saveTopics = (_: MouseEvent) => {
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
</script>

<!-- display topics -->
<ul>
    {#each $topics as topic}
        <li>
            <button on:click={removeTopic}>{topic}❌</button>
        </li>
    {/each}
</ul>

<div class="topicWrapper">
    <!-- Manual Topic Textbox -->
    <input
        placeholder="Add topic (Enter)"
        class="gyf-bar"
        type="text"
        on:keypress={sendTopic}
        bind:value={topicBuffer}
    />

    <!-- Load Topics from File Button -->
    <input type="file" on:change={loadFromFile} />

    <!-- Topic Actions -->
    <ul>
        <li>
            <!-- Clear Topics Button -->
            <input
                id="clearTopicsButton"
                type="button"
                value="💣 Nuke Topics"
                on:click={clearTopics}
            />
        </li>
        <li>
            <!-- Download Topics Button -->
            <input
                type="button"
                value="Download Topics"
                on:click={downloadTopics}
            />
        </li>
        <li>
            <!-- Save Topics Button -->
            <input type="button" value="Save Topics" on:click={saveTopics} />
        </li>
    </ul>
</div>

<!-- start game button -->
{#if Object.keys($players).length >= 3}
    <button class="clickable" on:click={startGame}>Start game!</button>
{:else}
    <button>Need {3 - Object.keys($players).length} more players!</button>
{/if}

<style>
    li {
        list-style: none;
    }

    .topicWrapper {
        background-color: #131313;
        border-radius: 10px;
    }

    input[type="button"]:hover {
        cursor: pointer;
    }
    input[type="file"] {
        width: 100%;
        max-width: 20rem;
    }
    .clickable:hover {
        cursor: pointer;
    }

    #clearTopicsButton {
        background-color: salmon;
    }
</style>
