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
            toast.push("Message cannot be empty.");
            return;
        }
        // send chat message and reset buffer
        $ws.send(`CHAT ${buffer}`);
        buffer = "";
    };
</script>

<div>
    <div id="chatContainer">
        <div id="messageContainer">
            <ul>
                {#each $chatMessages as message}
                    <li>
                        <Avatar user={message.author} width="24px" />
                        <span class="author">{message.author}</span>:
                        <span class="message">{message.message}</span>
                    </li>
                {/each}
            </ul>
            <div id="chatGradient" />
        </div>
        <input
            placeholder="Write a chat message"
            class="gyf-bar"
            type="text"
            on:keypress={onKeyDown}
            bind:value={buffer}
        />
    </div>
</div>

<style lang="scss">
    #chatContainer {
        width: 80%;
        float: right;
        margin-right: 0.5rem;
        border-radius: 0.5rem;
        display: flex;
        flex-direction: column;

        /* turn the chat arround, if on mobile */
        @media (max-width: 40rem) {
            flex-direction: column-reverse;
        }
    }

    #messageContainer {
        position: relative;
    }

    #chatGradient {
        background: linear-gradient(0, transparent 65%, #181818) ;
        pointer-events: none;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
    }
    ul {
        padding: 0;
        display: flex;
        flex-direction: column;
        max-height: 35vh;
        overflow-y: scroll;
        scroll-snap-align: end;
        @media (max-width: 40rem) {
            flex-direction: column-reverse;
        }
    }
    li {
        list-style: none;
        width: auto;
        text-align: left;
        overflow-wrap: break-word;
        padding: 0.5rem;
        padding-top: 0.2rem;
        padding-bottom: 0.2rem;
        margin-bottom: 0.2rem;
    }
    input {
        padding-right: 0.5rem;
        padding-left: 0.5rem;
        width: 100%;
    }
    .author {
        color: greenyellow;
    }
    .message {
        color: lightgray;
    }
</style>
