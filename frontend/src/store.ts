import { writable, Writable } from "svelte/store";
import type { ChatMessage, Player, Round } from "./types";

export const username: Writable<string> = writable('');

export const players: Writable<{[name: string]: Player}> = writable({});

export const chatMessages: Writable<Array<ChatMessage>> = writable([]);

export const leader: Writable<Boolean> = writable(false);
export const topics: Writable<Array<string>> = writable([]);
export const connected: Writable<Boolean> = writable(false);

export const round: Writable<Round> = writable({
    topic: '',
    currentRound: 0,
    totalRounds: 0
});

export const ws: Writable<WebSocket> = writable();
