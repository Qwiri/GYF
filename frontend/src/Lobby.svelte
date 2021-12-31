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
    } from "./store";

    import type { Response, Player } from "./types";
    import Lobby from "./screens/Lobby.svelte";

    export let id;

    onMount(async () => {
        // connect to the websocket
        ws.set(new WebSocket(`ws://127.0.0.1:8080/game/socket/${id}`));
        console.log(ws);

        // brutally attack server message handler
        $ws.onmessage = (msg) => {
            const response: Response = JSON.parse(msg.data);

            // if errors occurred
            if (handleErrors(response)) {
                return;
            }

            const commands: { [name: string]: (res: Response) => void } = {
                ERROR: (res: Response) => {
                    if (res.warn === "game not found") {
                        navigate("/", { replace: true });
                        return;
                    }
                    // TODO: error handling logic
                },

                PLAYER_JOINED: (res: Response) => {
                    $ws.send("LIST");
                    console.log("Player", res.args[0], "joined");
                    // update connected
                    if (res.args[0] === $username) {
                        $connected = true;
                    }
                },

                PLAYER_LEAVE: (res: Response) => {
                    $ws.send("LIST");
                    console.log("Player", res.args[0], "left");
                },

                LIST: (res: Response) => {
                    let temp: {[name: string]: Player} = {};
                    res.args.forEach((player: Player) => {
                        temp[player.name] = player;
                    });
                    $players = temp;
                },

                CHAT: (res: Response) =>
                    ($chatMessages = [
                        ...$chatMessages,
                        {
                            author: res.args[0],
                            message: res.args[1],
                        },
                    ]),

                TOPIC_LIST: (res: Response) => topics.set([...res.args]),

                CHANGE_ROLE: (res: Response) => handleChangeRole(res),

                NEXT_ROUND: (res: Response) =>
                    round.set({
                        topic: res.args[0],
                        currentRound: res.args[1],
                        totalRounds: res.args[2],
                    }),
            };

            if (commands[response.cmd]) {
                commands[response.cmd](response);
            }
        };
    });

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
                toast.push(resp.warn, {
                    theme: {
                        "--toastBackground": "#F56565",
                        "--toastBarBackground": "#C53030",
                    },
                });
                console.log("AN ERROR OCCURRED");
                console.log(resp.warn);
                console.log("Full log:", resp);
            } else {
                toast.push(
                    "An error occurred but there was no warning message given",
                    {
                        theme: {
                            "--toastBackground": "#F56565",
                            "--toastBarBackground": "#C53030",
                        },
                    }
                );
                console.log("AN ERROR OCCURRED");
                console.log("Full log:", resp);
            }
            return true;
        }
        return !resp._s;
    };
</script>

<!-- Display Choose Username if not connected to a lobby -->
{#if !$connected}
    <EnterUsername />
{:else}
    <Lobby />
    {#if $round.topic}
        <DisplayTopic />
    {:else}
        {#if $leader}
            <TopicEditor />
        {/if}
    {/if}
    <Chat />
{/if}
