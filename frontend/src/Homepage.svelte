<script lang="ts">
	import {navigate, useFocus} from "svelte-navigator";

    // to set the focus when this route get's opened
    const registerFocus = useFocus();

    let connectToGame = (id) => {
        navigate(`/game/${id}`, {replace: false, });
    }

    let createGame = async () => {
        let url = "http://127.0.0.1:8080/game/create"

        // create a new game room
        const res = await fetch(url, { method: "GET" })

        // check if backend responded as expected
        if (!res.ok) {
            throw new Error(`Request failed with status ${res.status}`)
        }
        const body = await res.json()

        // connect to the game room
        connectToGame(body.ID);
    }
</script>

<style>
    
</style>

<!-- <input use:registerFocus name="Username" placeholder="Username" bind:value="{username}" />
<img alt="user avatar" width="100px" src="https://avatars.dicebear.com/api/miniavs/{username}.svg" /> -->
<input use:registerFocus type="button" value="NEW GAME" on:click="{createGame}"/>