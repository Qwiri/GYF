<script lang="ts">
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
</script>

<style>
    li {
        list-style: none;
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

<!-- start game button -->
{#if Object.keys($players).length >= 3}
    <button on:click={startGame}>Start game!</button>
{:else}
    <button>Need {3 - Object.keys($players).length} more players!</button>
{/if}
