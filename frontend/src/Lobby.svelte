<script lang="ts">
    import { onMount } from 'svelte';
	import {navigate, useFocus} from "svelte-navigator";

    // to set the focus when this route get's opened
    const registerFocus = useFocus();

    export let id;
    let username;
    $: selfLeader = isLeader(username);
    let chatMsg;
    let inputTopic;
    let topics = [];

    let messages = [];

    let players = {};

    // TODO: purge. This is cursed
    $:messages = [players[0], ...messages].slice(1);

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
                break;

            case "CHAT":
                handleChatMessage(args);
                break;

            case "TOPIC_LIST":
                handleTopicList(args);
                break;
        }
        
    }

    const handleTopicList = (args) => {
        topics = [...args];
    }

    const isLeader = (name) => {
        return players[name]?.leader??false;
    }

    const handleChatMessage = (args) => {
        messages = [...messages, {"author": args[0], "msg": args[1]}]
    }

    const handlePlayerList = (args) => {

        //add all player
        players = {};
        args.forEach(player => {
            players[player.name] = player
        });

        if (!connected) {
            connected = true;
        }
    }

    const handlePlayerLeft = (args) => {
        ws.send("LIST 1")
    }

    const handlePlayerJoined = (args) => {
        ws.send("LIST 1");
        console.log(`Player ${args[0]} joined!`);
    }

    const sendMessage = (e) => {
        if (e.keyCode === 13) {
            ws.send(`CHAT ${chatMsg}`);
            chatMsg = "";
        }
    }
    
    const sendTopic = (e) => {
        if (e.keyCode === 13) {
            ws.send(`TOPIC_ADD ${inputTopic}`);
            inputTopic = "";

            // update topic list
            ws.send(`TOPIC_LIST 1`);
        }

    }

    const removeTopic = (e) => {
        const topic = e.srcElement.innerText.slice(0,-1);
        ws.send(`TOPIC_REMOVE ${topic}`);
        ws.send(`TOPIC_LIST 1`);
    }


</script>

<style>
    
</style>

<h1>Lobby!</h1>
<h2>your id is {id}</h2>
{#if !connected}
    <input use:registerFocus name="Username" placeholder="Username" bind:value="{username}" on:keypress="{e => {if (e.keyCode === 13) {connectWithUsername()}}}"/>
    <img alt="user avatar" width="100px" src="https://avatars.dicebear.com/api/miniavs/{username}.svg" />
    <input type="button" value="JOIN GAME" on:click="{connectWithUsername}"/>
{:else}
    <div id="playerBar">
        {#each Object.values(players) as player}
            <img width="100px" src="https://avatars.dicebear.com/api/miniavs/{player.name}.svg" alt="avatar of '{player.name}'"/>
            <h2>
                {#if player.leader}
                    👑
                {/if}
                {player.name}
            </h2>
        {/each}
    </div>
    <!-- topics -->
    {#if isLeader(username)}
        <ul>
            {#each topics as topic}
                <li>
                    <button on:click="{removeTopic}">{topic}❌</button>
                </li>
            {/each}
        </ul>
        <input placeholder="Add topic" type=text on:keypress="{sendTopic}" bind:value="{inputTopic}">
    {/if}
    <!-- chat -->
    <div>
        <div id="chatContent">
            <ul>
            {#each messages as message}
                <li>
                    {#if isLeader(message.author)}
                        👑
                    {/if}
                    <img width="32px" src="https://avatars.dicebear.com/api/miniavs/{message.author}.svg" alt=""/>
                    <span>{message.author}</span>:
                    <span>{message.msg}</span>
                </li>
            {/each}
            </ul>

        </div>
        <input placeholder="Write a chat message" type=text on:keypress="{sendMessage}" bind:value="{chatMsg}">

    </div>
{/if}