import { toast } from "@zerodevx/svelte-toast";
import { navigate } from "svelte-navigator";
import { chatMessages, leader, players, round, state, stats, submissions, topics, username, waitingFor } from "./store";
import { ChatMessage, GameState, isLeader, Player, pushWarn, Response } from "./types";

let localUsername: string;
username.subscribe(n => localUsername = n);

let localPlayers: { [name: string]: Player };
players.subscribe(n => localPlayers = n);

const commands: { [name: string]: (ws: WebSocket, res: Response) => void | any } = {
    JOIN: (ws: WebSocket, res: Response) => {
        // errored responses only occurr as responses to the JOIN command
        if (!res._s) {
            switch (res.warn) {
                case "game already started":
                case "game not found":
                    navigate("/?warn=" + res.warn, { replace: true });
                    return;
                default:
                    pushWarn(res.warn);
            }
        }

        // self joined?
        if (res.args[0] === localUsername) {
            state.set(GameState.Lobby);
            toast.push("Have fun playing GYF! :)", { theme: { "--toastBackground": "#64F4A0", "--toastBarBackground": "#3BDD7F" } });
        } else {
            toast.push(`Say hi to ${res.args[0]} 👋!`);
        }

        ws.send("LIST");
        return;
    },

    PLAYER_LEAVE: (ws: WebSocket, res: Response) => {
        ws.send("LIST");
        console.log("Player", res.args[0], "left");
    },

    LIST: (_, res: Response) => {
        let temp: { [name: string]: Player } = {};
        res.args.forEach((player: Player) => {
            temp[player.name] = player;
        });
        players.set(temp);
    },

    CHAT: (_, res: Response) => {
        const message: ChatMessage = {
            leader: isLeader(res.args[0]),
            author: res.args[0],
            message: res.args[1],
        };
        chatMessages.update(msgs => {
            msgs.push(message);
            return msgs;
        });
    },

    TOPIC_LIST: (_, res: Response) => {
        topics.set([...res.args])
    },

    CHANGE_ROLE: (ws: WebSocket, res: Response) => {
        if (res.args[0] !== localUsername) {
            return;
        }
        const _leader = res.args[1] === "LEADER";
        leader.set(_leader); // update leader in store

        if (_leader) {
            toast.push("You are now the leader!");
            ws.send("TOPIC_LIST"); // request topic list
        } else {
            toast.push("You are no longer the leader!");
        }
    },

    NEXT_ROUND: (ws: WebSocket, res: Response) => {
        state.set(GameState.SubmitGIF);
        ws.send("STATS"); // request player stats for each round

        round.set({
            topic: res.args[0],
            currentRound: res.args[1],
            totalRounds: res.args[2],
        });

        waitingFor.set(Object.keys(localPlayers));
    },

    SUBMIT_GIF: (_, res: Response) => {
        if (!res._s) {
            pushWarn(res.warn);
            return;
        }
        waitingFor.set(res.args.slice(1));
        toast.push(`${res.args[0]} submitted a gif`);
    },

    VOTE_START: (_, res: Response) => {
        state.set(GameState.Vote);
        submissions.set(res.args);
        // waiting for all players
        waitingFor.set(Object.keys(localPlayers));
    },

    VOTE_RESULTS: (ws: WebSocket, res: Response) => {
        state.set(GameState.VoteResults);
        ws.send("STATS");

        // TODO: Display results
    },

    VOTE: (_, res: Response) => {
        const voter: string = res.args.shift();
        toast.push(`${voter} voted!`);
        waitingFor.set(res.args);
    },

    STATS: (_, res: Response) => {
        if (!res._s) {
            return;
        }
        const sortable = [];
        for (const name in res.args[0]) {
            sortable.push([name, res.args[0][name]]);
        }
        sortable.sort((a, b) => b[1] - a[1]);
        const objSorted: { [name: string]: number } = {};
        sortable.forEach((item) => (objSorted[item[0]] = item[1]));
        stats.set(objSorted);
    },

    STATE: (_, res: Response) => {
        const _state = res.args[0];
        if (GameState[_state]) {
            state.set(_state);
            console.log("changed state to", _state);
        }
    },
};


export function hijack(ws: WebSocket) {
    ws.onerror = (e) => {
        console.error("WebSocket error", e);
    };

    ws.onopen = () => {
        console.log("connection opened");

        // Auto Join Game
        const query: URLSearchParams = new URLSearchParams(
            document.location.search
        );
        if (query.has("name")) {
            const user = query.get("name");
            username.set(user);

            ws.send(`JOIN ${user}`);
        }
    };

    ws.onclose = () => {
        console.log("connection closed");
    }

    // brutally hijack server message handler
    ws.onmessage = (msg: MessageEvent<any>) => {
        console.log("[ws] 📦 ->", msg.data);

        // try to parse message
        const response: Response = JSON.parse(msg.data);

        if (commands[response.cmd]) {
            const result = commands[response.cmd](ws, response);
            if (result) {
                console.error("error executing", response.cmd, result);
            }
        } else {
            console.log("Unknown command", response);
            handleErrors(response);
        }
    };

    // intercept sending messages
    // FOR EDUCATIONAL PURPOSES ONLY
    const s = ws.send;
    ws.send = (data: string) => {
        console.log("[ws] 👋 <-", data);
        return s.call(ws, data);
    }
}

function handleErrors(resp: Response): boolean {
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
}
