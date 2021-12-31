<script lang="ts">
	import Homepage from "./Homepage.svelte";
	import Lobby from "./Lobby.svelte";
	import {Router, Route, Link} from "svelte-navigator";
	import { SvelteToast } from "@zerodevx/svelte-toast";


	export let name;
	export let url = "";
</script>

<main>
	<SvelteToast />
	<h1>{name}!</h1>
	<!-- <Homepage /> -->
	<Router url={url}>
		<Route path="/">
			<Homepage />
		</Route>
		<Route path="game/:id" let:params>
			<Lobby id="{params.id}"/>
		</Route>
	</Router>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>