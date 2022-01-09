import { writable } from "svelte/store";
import type { Writable } from "svelte/store";
import type { ChatMessage, GameState, Player, Preferences, Round, VotingResult } from "./types";

export const username: Writable<string> = writable('');

export const players: Writable<{ [name: string]: Player }> = writable({});
export const waitingFor: Writable<Array<string>> = writable([]);
export const gifSubmitted: Writable<boolean> = writable(false);
export const submissions: Writable<Array<string>> = writable([]);
export const votingResults: Writable<Array<VotingResult>> = writable([]);
export const stats: Writable<{ [name: string]: number }> = writable({});

export const chatMessages: Writable<Array<ChatMessage>> = writable([]);

export const leader: Writable<boolean> = writable(false);
export const topics: Writable<Array<string>> = writable([]);
export const preferences: Writable<Preferences> = writable({
    AutoSkip: true,
    MaxPlayers: -1,
    MinPlayers: -1,
    MaxTopics: -1,
    MinTopics: -1,
    ShuffleTopics: true,
});

export const round: Writable<Round> = writable({
    topic: '',
    currentRound: 0,
    totalRounds: 0
});

export const ws: Writable<WebSocket> = writable();

export const state: Writable<GameState> = writable(0);