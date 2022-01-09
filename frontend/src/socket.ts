import { toast } from "@zerodevx/svelte-toast";
import { navigate } from "svelte-navigator";
import { chatMessages, leader, players, round, state, stats, submissions, topics, username, waitingFor, votingResults, preferences, gifSubmitted } from "./store";
import { GameState, type Preferences } from "./types";
import type { ChatMessage, Player, Response, VotingResult } from "./types";
import { isLeader, pushInfo, pushSuccess, pushWarn } from "./utils";

let localUsername: string;
username.subscribe(n => localUsername = n);

let localPlayers: { [name: string]: Player };
players.subscribe(n => localPlayers = n);

let localVoteResults: Array<VotingResult>;
votingResults.subscribe(n => localVoteResults = n);

let localState: GameState;
state.subscribe(n => localState = n);

const commands: { [name: string]: (ws: WebSocket, res: Response) => void | any } = {
    JOIN: (_: WebSocket, res: Response) => {
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
        const payloadUser = res.args[0] as string;
        if (payloadUser === localUsername) {
            state.set(GameState.Lobby);
            pushSuccess("Have fun playing GYF! :)");
        } else {
            pushInfo(`👋 Say hi to ${payloadUser}!`);
        }
        return;
    },

    PLAYER_LEAVE: (_: WebSocket, res: Response) => {
        const payloadUser = res.args[0] as string;
        pushInfo("🚪", payloadUser, "left");
    },

    LIST: (_, res: Response) => {
        const payloadPlayers = res.args as Array<Player>;

        let temp: { [name: string]: Player } = {};
        payloadPlayers.forEach((player: Player) => {
            temp[player.name] = player;
        });

        players.set(temp);
    },

    CHAT: (_, res: Response) => {
        const message: ChatMessage = {
            leader: isLeader(res.args[0] as string),
            author: res.args[0] as string,
            message: res.args[1] as string,
        };
        chatMessages.update(msgs => {
            msgs.push(message);
            return msgs;
        });
    },

    TOPIC_LIST: (_, res: Response) => {
        const payloadTopics = res.args as Array<string>;
        topics.set(payloadTopics);
    },

    CHANGE_ROLE: (_: WebSocket, res: Response) => {
        const payloadUser = res.args[0] as string;
        const payloadRole = res.args[1] as string;

        if (payloadUser !== localUsername) {
            return;
        }

        const _leader = payloadRole === "LEADER";
        leader.set(_leader); // update leader in store

        if (_leader) {
            toast.push("You are now the leader!");
        } else {
            pushWarn("You are no longer the leader!");
            topics.set([]); // clear topics
        }
    },

    NEXT_ROUND: (_: WebSocket, res: Response) => {
        gifSubmitted.set(false);
        round.set({
            topic: res.args[0] as string,
            currentRound: res.args[1] as number,
            totalRounds: res.args[2] as number,
        });
    },

    VOTE_START: (_, res: Response) => {
        // clear previous vote results
        votingResults.set([]);

        // load submissions from backend
        const payloadSubmissions = res.args as Array<string>;
        submissions.set(payloadSubmissions);
    },

    VOTE_RESULTS: (_: WebSocket, res: Response) => {
        // reset submissions
        submissions.set([]);

        if (!res._s) {
            return;
        }

        const payloadResults = res.args as Array<VotingResult>;

        // sort voting results
        payloadResults.sort((a: VotingResult, b: VotingResult) => {
            if (a.voters.length > b.voters.length) {
                return -1;
            }
            if (a.voters.length < b.voters.length) {
                return 1;
            }
            return 0;
        });

        votingResults.set(payloadResults);
    },

    STATS: (_, res: Response) => {
        if (!res._s) {
            return;
        }
        
        const payloadStats = res.args[0] as {[name: string]: number};

        // sort stats by value
        const sortable = [];
        for (const name in payloadStats) {
            sortable.push([name, payloadStats[name]]);
        }
        sortable.sort((a, b) => b[1] - a[1]);
        
        // rebuild object
        const objSorted: { [name: string]: number } = {};
        sortable.forEach((item) => (objSorted[item[0]] = item[1]));
        
        // update stats
        stats.set(objSorted);
    },

    STATE: (_, res: Response) => {
        const _state = res.args[0] as number;
        const _gs = GameState[_state];
        if (_gs) {
            if (_gs === "Lobby" && localState === GameState.GameEnd) {
                console.log("Skipped State Change.");
                return;
            }
            state.set(_state);
            console.log("Changed state to", _state, "::", GameState[_state]);
        }
    },

    WAITING_FOR: (_, res: Response) => {
        const payloadWaitingFor = res.args as Array<string>;
        waitingFor.set(payloadWaitingFor);
    },

    PREFERENCES: (_, res: Response) => {
        const payloadPreferences = res.args[0] as Preferences;
        preferences.set(payloadPreferences);
    },

    SUBMIT_GIF: (_, res: Response) => {
        if (!res._s) {
            handleErrors(res);
            return;
        }
        toast.push(`✅ '${res.args[0]}' submitted a GYF. Hurry up!`)
        if (res.args[0] == localUsername) {
            gifSubmitted.set(true);
        }
    },

    GAME_END: (_, res: Response) => {
        pushInfo(`Game ended! Thanks for playing!`);
        state.set(GameState.GameEnd);
    },

    START: (_, res: Response) => {
        // reset
        resetGameValues();
    },
};

export function resetGameValues() {
    round.set({ topic: '', currentRound: 0, totalRounds: 0 });
    stats.set({});
    votingResults.set([]);
    submissions.set([]);
    waitingFor.set([]);
}

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
