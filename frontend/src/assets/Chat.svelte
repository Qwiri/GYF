<script lang="ts">
import { toast } from "@zerodevx/svelte-toast";

    import { chatMessages, ws } from "../store";
    import { isLeader } from "../types";
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
    <ul>
        {#each $chatMessages as message}
            <li>
                {#if isLeader(message.author)}
                    <span class="role">👑</span>
                {/if}
                <Avatar user={message.author} width="32px" />
                <span class="author">{message.author}</span>:
                <span class="message">{message.message}</span>
            </li>
        {/each}
    </ul>
</div>

<input
    placeholder="Write a chat message"
    type="text"
    on:keypress={onKeyDown}
    bind:value={buffer}
/>

<style>
    li {
        list-style: none;
    }
    .author {
        color: greenyellow;
    }
    .message {
        color: lightgray;
    }
</style>
