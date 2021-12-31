<script lang="ts">
    import { onMount } from "svelte";
    import { navigate } from "svelte-navigator";

    import { toast } from "@zerodevx/svelte-toast";

    // components
    import DisplayTopic from "./screens/DisplayTopic.svelte";
    import EnterUsername from "./screens/EnterUsername.svelte";
    import Chat from "./assets/Chat.svelte";
    import TopicEditor from "./assets/setup/TopicEditor.svelte";

    import {
        chatMessages,
        connected,
        leader,
        players,
        round,
        topics,
        username,
        ws,
        waitingFor,
submissions,
    } from "./store";

    let votingBuffer: boolean;

    import { Response, Player, isLeader, pushWarn } from "./types";
    import Lobby from "./screens/Lobby.svelte";
    import Voting from "./screens/Voting.svelte";

    export let id;

    onMount(async () => {
        // connect to the websocket
        ws.set(new WebSocket(`ws://127.0.0.1:8080/game/socket/${id}`));
        console.log(ws);

        $ws.onopen = () => {
            // TODO: refactor this
            const query: URLSearchParams = new URLSearchParams(
                document.location.search
            );
            if (query.has("name")) {
                $username = query.get("name");
                $ws.send(`JOIN ${$username}`); // auto join
            }
        };

        // brutally attack server message handler
        $ws.onmessage = (msg) => {
            const response: Response = JSON.parse(msg.data);

            const commands: { [name: string]: (res: Response) => void } = {
                JOIN: (res: Response) => {
                    if (res._s) {
                        return;
                    }
                    switch (res.warn) {
                        case "game already started":
                        case "game not found":
                            navigate("/?warn=" + res.warn, { replace: true });
                            return;
                        default:
                            pushWarn(res.warn);
                    }
                },

                PLAYER_JOINED: (res: Response) => {
                    $ws.send("LIST");
                    console.log("Player", res.args[0], "joined");
                    // update connected
                    if (res.args[0] === $username) {
                        $connected = true;
                    } else {
                        toast.push(`Player '${res.args[0]}' joined`);
                    }
                },

                PLAYER_LEAVE: (res: Response) => {
                    $ws.send("LIST");
                    console.log("Player", res.args[0], "left");
                },

                LIST: (res: Response) => {
                    let temp: { [name: string]: Player } = {};
                    res.args.forEach((player: Player) => {
                        temp[player.name] = player;
                    });
                    $players = temp;
                },

                CHAT: (res: Response) =>
                    ($chatMessages = [
                        ...$chatMessages,
                        {
                            leader: isLeader(res.args[0]),
                            author: res.args[0],
                            message: res.args[1],
                        },
                    ]),

                TOPIC_LIST: (res: Response) => topics.set([...res.args]),

                CHANGE_ROLE: (res: Response) => handleChangeRole(res),

                NEXT_ROUND: (res: Response) => handleNextRound(res),

                SUBMIT_GIF: (res: Response) => handleSubmitGif(res),

                VOTE_START: (res: Response) => handleVote(res),
            };

            if (commands[response.cmd]) {
                commands[response.cmd](response);
            } else {
                console.log("Unknown command", response);
                handleErrors(response);
            }
        };
    });

    const handleVote = (res: Response) => {
        votingBuffer=true;
        $submissions = res.args;
        console.log("switch to voting");
    }

    const handleNextRound = (res: Response) => {
        votingBuffer = false;
        round.set({
            topic: res.args[0],
            currentRound: res.args[1],
            totalRounds: res.args[2],
        })
        $waitingFor = Object.keys($players);
    }

    const handleSubmitGif = (res: Response) => {
        if (!res._s) {
            pushWarn(res.warn);
            return;
        }
        toast.push(`${res.args[0]} submitted a gif`);
        $waitingFor = res.args.slice(1);
    }

    const handleChangeRole = (res: Response) => {
        if (res.args[0] === $username) {
            leader.set(res.args[1] === "LEADER");
            if ($leader) {
                toast.push("You are now the leader!");
                // request topic list
                $ws.send("TOPIC_LIST");
            } else {
                toast.push("You are no longer the leader!");
            }
        }
    };

    // middleware that evaluates websocket messages for errors
    const handleErrors = (resp: Response) => {
        if (!resp._s) {
            if (resp.warn != "") {
                console.log("AN ERROR OCCURRED");
                console.log(resp.warn);
                console.log("Full log:", resp);

                pushWarn(resp.warn);
            } else {
                console.log("AN ERROR OCCURRED");
                console.log("Full log:", resp);

                pushWarn(
                    "An error occurred but there was no warning message given"
                );
            }
        }
        return !resp._s;
    };
</script>

<!-- Display Choose Username if not connected to a lobby -->
{#if !$connected}
    <EnterUsername />
{:else}
    <Lobby />
    {#if votingBuffer}
        <Voting />
        
    {:else if $round.topic}
        <DisplayTopic />
    {:else if $leader}
        <TopicEditor />
    {/if}
    <Chat />
{/if}
