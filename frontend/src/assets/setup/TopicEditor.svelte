<script lang="ts">
import { toast } from "@zerodevx/svelte-toast";

    import { ws, topics, players } from "../../store";

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

            toast.push(`Sending <strong>${lines.length} lines</strong> to backend`);
            $ws.send(`TOPIC_ADD_ALL ${JSON.stringify(lines)}`);
        };

        reader.readAsText(file);
    };

    const clearTopics = () => {
        $ws.send("TOPIC_CLEAR");
    };
</script>

<style>
    li {
        list-style: none;
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
        background-color: white;
    }

</style>

<!-- display topics -->
<ul>
    {#each $topics as topic}
        <li>
            <button on:click={removeTopic}>{topic}❌</button>
        </li>
    {/each}
</ul>

<input
    placeholder="Add topic"
    type="text"
    on:keypress={sendTopic}
    bind:value={topicBuffer}
/>
<input type="file" on:change="{loadFromFile}" />
<input id="clearTopicsButton" type="button" value="Clear Topics" on:click="{clearTopics}" />

<!-- start game button -->
{#if Object.keys($players).length >= 3}
    <button class="clickable" on:click={startGame}>Start game!</button>
{:else}
    <button>Need {3 - Object.keys($players).length} more players!</button>
{/if}
