<script lang="ts">
import { toast } from "@zerodevx/svelte-toast";

    import { chatMessages, ws } from "../store";
    import Avatar from "./Avatar.svelte";

    let buffer: string = "";

    const onKeyDown = (event: KeyboardEvent) => {
        if (event.key !== "Enter") {
            return;
        }
        if (buffer.trim().length === 0) {
            toast.push("Message cannot be empty.")
            return;
        }
        // send chat message and reset buffer
        $ws.send(`CHAT ${buffer}`);
        buffer = "";
    };
</script>

<div>
    <div id="chatContainer">
        <ul>
            {#each $chatMessages as message}
                <li>
                    <Avatar user={message.author} width="24px" />
                    <span class="author">{message.author}</span>:
                    <span class="message">{message.message}</span>
                </li>
            {/each}
        </ul>
        <input
            placeholder="Write a chat message"
            type="text"
            on:keypress={onKeyDown}
            bind:value={buffer}
        />
    </div>
</div>


<style>

    #chatContainer {
        width: 80%;
        float: right;
        margin-right: .5rem;
        border-radius: .5rem
    }
    ul {
        padding: 0;
    }
    li {
        list-style: none;
        width: auto;
        text-align: left;
        overflow-wrap: break-word;
        padding: .5rem;
        padding-top: .2rem;
        padding-bottom: .2rem;
        margin-bottom: .2rem;
    }
    input {
        padding-right: .5rem;
        padding-left: .5rem;
        width: 100%;
    }
    .author {
        color: greenyellow;
    }
    .message {
        color: lightgray;
    }
    Avatar {
        width: 1rem;
        height: 1rem;
    }
</style>
