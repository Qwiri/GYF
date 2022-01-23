import { toast } from "@zerodevx/svelte-toast";
import { navigate } from "svelte-navigator";
import { chatMessages, leader, players, round, state, stats, submissions, topics, username, waitingFor, votingResults, preferences, gifSubmitted, backendVersion } from "./store";
import { GameState, type Preferences } from "./types";
import type { ChatMessage, Player, Response, VotingResult } from "./types";
import { isLeader, pushInfo, pushSuccess, pushWarn, resetGameValues } from "./utils";

let localUsername = "";
username.subscribe(n => localUsername = n);

let localState: GameState = GameState.ChooseUsername;
state.subscribe(n => localState = n);

function handleErrors(resp: Response): boolean {
    if (resp.cmd === "ERROR") {
        console.log("errored result:", resp.args);
        console.log("(this may be ignored)");
        return;
    }

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

const commands: { [name: string]: (res: Response) => void | string } = {
    JOIN: (res: Response) => {
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
    },

    PLAYER_LEAVE: (res: Response) => {
        const payloadUser = res.args[0] as string;
        const payloadReason = res.args[1] as string;

        // send kick message
        if (payloadReason === "KICKED") {
            pushInfo("🥊", payloadUser, "was kicked from the game");
        } else {
            pushInfo("🚪", payloadUser, "left");
        }

        // hide ui if we left
        if (payloadUser === localUsername) {
            navigate("/", { replace: true });
            state.set(GameState.ChooseUsername);
            username.set("");
            players.set({});
            waitingFor.set([]);
            submissions.set([]);
            votingResults.set([]);
            stats.set({});
            chatMessages.set([]);
            topics.set([]);
        }
    },

    LIST: (res: Response) => {
        const payloadPlayers = res.args as Array<Player>;

        const temp: { [name: string]: Player } = {};
        payloadPlayers.forEach((player: Player) => {
            temp[player.name] = player;
        });

        players.set(temp);
    },

    CHAT: (res: Response) => {
        if (!res._s) {
            pushWarn(res.warn);
            return;
        }
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

    TOPIC_LIST: (res: Response) => {
        const payloadTopics = res.args as Array<string>;
        topics.set(payloadTopics);
    },

    CHANGE_ROLE: (res: Response) => {
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

    NEXT_ROUND: (res: Response) => {
        gifSubmitted.set(false);
        round.set({
            topic: res.args[0] as string,
            currentRound: res.args[1] as number,
            totalRounds: res.args[2] as number,
        });
    },

    VOTE_START: (res: Response) => {
        // clear previous vote results
        votingResults.set([]);

        // load submissions from backend
        const payloadSubmissions = res.args as Array<string>;
        submissions.set(payloadSubmissions);
    },

    VOTE_RESULTS: (res: Response) => {
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

    STATS: (res: Response) => {
        if (!res._s) {
            return;
        }

        const payloadStats = res.args[0] as { [name: string]: number };

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

    STATE: (res: Response) => {
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

    WAITING_FOR: (res: Response) => {
        const payloadWaitingFor = res.args as Array<string>;
        waitingFor.set(payloadWaitingFor);
    },

    PREFERENCES: (res: Response) => {
        const payloadPreferences = res.args[0] as Preferences;
        preferences.set(payloadPreferences);
    },

    SUBMIT_GIF: (res: Response) => {
        if (!res._s) {
            handleErrors(res);
            return;
        }
        toast.push(`✅ '${res.args[0]}' submitted a GYF. Hurry up!`)
        if (res.args[0] == localUsername) {
            gifSubmitted.set(true);
        }
    },

    GAME_END: () => {
        pushInfo(`Game ended! Thanks for playing!`);
        state.set(GameState.GameEnd);
    },

    START: () => {
        // reset
        resetGameValues();
    },
    VERSION: (res: Response) => {
        const version = res.args[0] as string;
        backendVersion.set(version);
    }
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
    ws.onmessage = (msg: MessageEvent<string>) => {
        console.log("[ws] 📦 ->", msg.data);

        // try to parse message
        const response: Response = JSON.parse(msg.data);

        if (commands[response.cmd]) {
            const result = commands[response.cmd](response);
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

