export enum GameState {
    ChooseUsername,
    Lobby = 1 << 0,
    SubmitGIF = 1 << 1,
    Vote = 1 << 2,
    VoteResults = 1 << 3,
    GameEnd = 1 << 16,
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

export interface VotingResult {
    creator: string;
    url: string;
    voters: Array<string>;
}

export interface Preferences {
    AutoSkip: boolean;
    MinPlayers: number;
    MaxPlayers: number;
    MinTopics: number;
    MaxTopics: number;
    ShuffleTopics: boolean;
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

export function base64DecodeUnicode(str) {
    // Convert Base64 encoded bytes to percent-encoding, and then get the original string.
    const percentEncodedStr = atob(str)
        .split("")
        .map(function (c) {
            return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("");
    return decodeURIComponent(percentEncodedStr);
}

export function copyToClipboard(str) {
    const copyText = document.createElement("textarea");
    copyText.value = str;
    document.body.appendChild(copyText);
    copyText.select();
    document.execCommand("copy");
    document.body.removeChild(copyText);
}