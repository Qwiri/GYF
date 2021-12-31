<script lang="ts">
    import { onMount } from 'svelte';
	import {navigate, useFocus} from "svelte-navigator";
    import { toast } from '@zerodevx/svelte-toast';

    // to set the focus when this route get's opened
    const registerFocus = useFocus();

    export let id;
    let username;
    let selfLeader;
    let chatMsg;
    let inputTopic;
    let topics = [];

    let messages = [];

    let players = {};

    // TODO: purge. This is cursed
    $:messages = [players[0], ...messages].slice(1);

    let ws;

    let connected = false;
    let displayTopic = "";
    let currRounds, maxRounds;
    
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

        // if errors occurred
        if (evalMessage(msg)) { return;}

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

            case "CHANGE_ROLE":
                handleChangeRole(args);
                break;

            case "NEXT_ROUND":
                handleNextRound(args);
                break;
        }
        
    }

    const handleNextRound = (args) => {
        [displayTopic, currRounds, maxRounds] = args;

    }

    const handleChangeRole = (args) => {
        // if self got upgraded to leader
        if (args[0] === username && args[1] === "LEADER") {
            selfLeader = true;

            // get topic list
            ws.send("TOPIC_LIST")
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
    }

    const handlePlayerLeft = (args) => {
        ws.send("LIST")
    }

    const handlePlayerJoined = (args) => {
        ws.send("LIST");
        if (!connected) {
            connected = true;
        }
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
            ws.send(`TOPIC_LIST`);
        }

    }

    const removeTopic = (e) => {
        const topic = e.srcElement.innerText.slice(0,-1);
        ws.send(`TOPIC_REMOVE ${topic}`);
        ws.send(`TOPIC_LIST`);
    }

    // middleware that evaluates websocket messages for errors
    const evalMessage = (msg) => {

        if (!msg._s) {

            if (msg.warn != "") {
                toast.push(msg.warn, {
                    theme: {
                        '--toastBackground': '#F56565',
                        '--toastBarBackground': '#C53030'
                    }
                })
                console.log("AN ERROR OCCURRED");
                console.log(msg.warn);
                console.log("Full log:", msg)

            } else {
                toast.push(`An error occurred but there was no warning message given\n Full log: '${msg.data}`, {
                    theme: {
                        '--toastBackground': '#F56565',
                        '--toastBarBackground': '#C53030'
                    }
                })
                console.log("AN ERROR OCCURRED");
                console.log("Full log:", msg)

            }
            return true;

        }
        return false;
    }


</script>

<style>
    
</style>

{#if !connected}
    <input use:registerFocus name="Username" placeholder="Username" bind:value="{username}" on:keypress="{e => {if (e.keyCode === 13) {connectWithUsername()}}}"/>
    <img alt="user avatar" width="100px" src="https://avatars.dicebear.com/api/miniavs/{username}.svg" />
    <input type="button" value="JOIN GAME" on:click="{connectWithUsername}"/>
{:else if displayTopic}
    <h2>{displayTopic}</h2>
    <p style="color:greenyellow">{currRounds}/{maxRounds}</p>
    {#if selfLeader}
        <button on:click="{ws.send("SKIP")}">Skip round</button>
    {/if}
        
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
    {#if selfLeader}
        <!-- topics -->
        <ul>
            {#each topics as topic}
                <li>
                    <button on:click="{removeTopic}">{topic}❌</button>
                </li>
            {/each}
        </ul>
        <input placeholder="Add topic" type=text on:keypress="{sendTopic}" bind:value="{inputTopic}">

        <!-- start game button -->
        {#if Object.keys(players).length >= 3}
            <button on:click="{ws.send(`START`)}">Start game!</button>
        {:else}
            <button>Need {3-Object.keys(players).length} more players!</button>
        {/if}
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