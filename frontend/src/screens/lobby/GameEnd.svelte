<script lang="ts">
import { resetGameValues } from "../../socket";

import { state } from "../../store";
import { GameState } from "../../types";

function gotoLobby() {
    state.set(GameState.Lobby);
    resetGameValues();
}
</script>

<h1>Game Ended!</h1>
<button on:click="{gotoLobby}">Back to Lobby</button>