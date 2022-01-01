export enum GameState {
    ChooseUsername,
    Lobby,
    SubmitGIF,
    Vote,
    VoteResults
}

export interface Response {
    cmd: string;
    args: Array<any>;
    warn: string;
    _s: boolean;
    _ts: number;
}

export interface Player {
    name: string;
    leader: boolean;
}

export interface Command {
    cmd: string;
    run: void;
}

export interface Round {
    topic: string;
    currentRound: number;
    totalRounds: number;
}

export interface ChatMessage {
    leader: boolean;
    author: string;
    message: string;
}

import { toast } from "@zerodevx/svelte-toast";
import { players } from "./store";

let localPlayers: { [name: string]: Player } = {};
players.subscribe(n => {
    localPlayers = n;
});

export const isLeader = (name: string) => {
    return localPlayers[name]?.leader || false;
};

export const pushWarn = (message: string) => {
    toast.push(message, {
        theme: {
            "--toastBackground": "#F56565",
            "--toastBarBackground": "#C53030",
        },
    });
};
