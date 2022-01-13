import { toast } from "@zerodevx/svelte-toast";
import { players, preferences, round, stats, submissions, votingResults, waitingFor } from "./store";
import type { Player } from "./types";

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

export const pushSuccess = (...message: string[]) => {
    toast.push(message.join(" "), {
        theme: {
            "--toastBackground": "#64F4A0",
            "--toastBarBackground": "#3BDD7F",
        },
    });
}

export const pushInfo = (...message: string[]) => {
    toast.push(message.join(" "), {
        theme: {
            // make blue
            "--toastBackground": "#008BEF",
            "--toastBarBackground": "#0071C1",
        },
    });
}

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

export function resetGameValues() {
    round.set({ topic: '', currentRound: 0, totalRounds: 0 });
    stats.set({});
    votingResults.set([]);
    submissions.set([]);
    waitingFor.set([]);
}
