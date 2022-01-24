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
    args: Array<unknown>;
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
    Permissions: number;
}

export interface GifFetchError {
    statusCode: number,
    statusText: string,
    redirected: boolean,
    json?: unknown
}