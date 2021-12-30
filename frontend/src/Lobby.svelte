<script lang="ts">
    import { onMount } from 'svelte';
	import {navigate, useFocus} from "svelte-navigator";

    // to set the focus when this route get's opened
    const registerFocus = useFocus();

    export let id;
    let username;

    let players = [];

    let ws;

    let connected = false;
    
    onMount(async () => {

        //connect to the websocket
        ws = new WebSocket(`ws://127.0.0.1:8080/game/socket/${id}`);
        console.log(ws)

        // attack server message handler
        ws.onmessage = handleMessage;
    })

    let connectWithUsername = () => {
        if (ws.readyState === 1) {
            ws.send(`JOIN ${username}`);

        }
    }

    let handleMessage = (msg) => {
        msg = JSON.parse(msg.data);
        const cmd = msg.cmd
        const args = msg.args
        console.log(msg)
        console.log({cmd, args})

        switch (cmd) {
            case "ERROR":
                if (!msg._s && msg.warn == "ERROR game not found") {
                    navigate("/", {replace: true})
                } else if (!msg._s && msg.warn == "ERROR game already started") {
                    // TODO: error handling logic
                    console.log("Game already started");
                }
                return

            case "PLAYER_JOINED":
                handlePlayerJoined(args);
                break;

            case "PLAYER_LEAVE":
                handlePlayerLeft(args);
                break;

            case "LIST":
                handlePlayerList(args);

        }
        
    }

    const handlePlayerList = (args) => {

        //add all player
        players = [...args]
    }

    const handlePlayerLeft = (args) => {
        ws.send("LIST 1")
    }

    const handlePlayerJoined = (args) => {
        if (args[0] === username) {
            connected = true;
        }
        ws.send("LIST 1");
        console.log(`Player ${args[0]} joined!`);
    }


</script>

<style>
    
</style>

<h1>Lobby!</h1>
<h2>your id is {id}</h2>
{#if !connected}
    <input use:registerFocus name="Username" placeholder="Username" bind:value="{username}" />
    <img alt="user avatar" width="100px" src="https://avatars.dicebear.com/api/miniavs/{username}.svg" />
    <input type="button" value="JOIN GAME" on:click="{connectWithUsername}"/>
{:else}
    <div id="playerBar">
        {#each players as player}
            <img width="100px" src="https://avatars.dicebear.com/api/miniavs/{player.name}.svg" alt="avatar of '{player.name}'"/>
            <h2>
                {#if player.leader}
                    👑
                {/if}
                {player.name}
            </h2>
        {/each}
    </div>
{/if}