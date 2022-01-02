<script lang="ts">
    import { onMount } from "svelte";

    // components
    import EnterUsername from "./screens/EnterUsername.svelte";
    import Chat from "./assets/Chat.svelte";
    import GroupGame from "./screens/group/GroupGame.svelte";
    import Lobby from "./screens/lobby/Lobby.svelte";

    import { ws, state } from "./store";
    import { hijack } from "./socket";
    import { GameState } from "./types";

    export let id;

    onMount(async () => {
        console.log("connecting to socket ...");
        const socket = new WebSocket(`ws://127.0.0.1:8080/game/socket/${id}`);
        console.log(socket);

        // add listeners
        hijack(socket);

        ws.set(socket);
    });
</script>

<!-- Display Choose Username if not connected to a lobby -->
{#if $state == GameState.ChooseUsername}
    <EnterUsername {id} />
{:else}
    {#if $state == GameState.Lobby}
        <div id="wholeScreen">
            <div id="lobby">
                <Lobby />
            </div>
            <div id="chat">
                <Chat />
            </div>
        </div>
    {:else}
        <GroupGame />
    {/if}
{/if}

<style>
    #wholeScreen {
        display: flex;
        flex-direction: row;
        height: 100vh;
        align-items: center;
        justify-content: flex-end;
    }

    #lobby {
        width: 50vw;
    }
    #chat {
        width: 25vw;
        align-self: flex-end;
    }
</style>