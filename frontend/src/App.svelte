<script lang="ts">
	import Homepage from "./Homepage.svelte";
	import Game from "./Game.svelte";
	import { Router, Route } from "svelte-navigator";
	import { SvelteToast } from "@zerodevx/svelte-toast";
	import { username } from "./store";

	export let name;
	export let url = "";
</script>

<main>
	<SvelteToast />

	<div class="header">
		<h1>
			{name}!
			<img src="https://i.gifer.com/2iFd.gif" height="64px" alt="cat" />
		</h1>
		{#if $username}
			<h2>Welcome back <span>{$username}</span>!</h2>
		{/if}
	</div>

	<!-- <Homepage /> -->
	<Router {url}>
		<Route path="/">
			<Homepage />
		</Route>
		<Route path="game/:id" let:params>
			<Game id={params.id} />
		</Route>
	</Router>
</main>

<style lang="scss">
	main {
		text-align: center;
		padding: 0;
		max-width: 240px;
		margin: 0 auto;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

	.header {
		h1 {
			color: greenyellow;
			text-transform: uppercase;
			font-size: 4em;
			font-weight: 1000;
		}

		h2 {
			color: grey;

			span {
				color: greenyellow;
			}
		}
	}
</style>
