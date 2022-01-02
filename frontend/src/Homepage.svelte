<script lang="ts">
    import { navigate, useFocus } from "svelte-navigator";
    import { pushWarn } from "./types";

    const query: URLSearchParams = new URLSearchParams(window.location.search);
    if (query.has("warn")) {
        const warn = query.get("warn");

        if (["game already started", "game not found"].includes(warn)) {
            pushWarn(warn);
        }

        window.history.pushState({}, "", `${window.location.pathname}`); // remove warn from URL
    }

    // to set the focus when this route get's opened
    const registerFocus = useFocus();

    const connectToGame = (id: string) => {
        navigate(`/game/${id}`, { replace: false });
    };

    const createGame = async () => {
        const url: string = "http://127.0.0.1:8080/game/create";

        // create a new game room
        const res: Response = await fetch(url, { method: "GET" });

        // check if backend responded as expected
        if (!res.ok) {
            throw new Error(`Request failed with status ${res.status}`);
        }
        const body = await res.json();

        // connect to the game room
        connectToGame(body.ID);
    };
</script>

<input use:registerFocus type="button" value="NEW GAME" on:click={createGame} />
