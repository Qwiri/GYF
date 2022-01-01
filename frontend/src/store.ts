import { writable, Writable } from "svelte/store";
import type { ChatMessage, GameState, Player, Round } from "./types";

export const username: Writable<string> = writable('');

export const players: Writable<{[name: string]: Player}> = writable({});
export const waitingFor: Writable<Array<string>> = writable([]);
export const submissions: Writable<Array<string>> = writable([]);
export const stats: Writable<{[name: string]: number}> = writable({});

export const chatMessages: Writable<Array<ChatMessage>> = writable([]);

export const leader: Writable<Boolean> = writable(false);
export const topics: Writable<Array<string>> = writable([]);

export const round: Writable<Round> = writable({
    topic: '',
    currentRound: 0,
    totalRounds: 0
});

export const ws: Writable<WebSocket> = writable();

export const state: Writable<GameState> = writable(0);