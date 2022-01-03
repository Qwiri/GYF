<script lang="ts">
	import Homepage from "./Homepage.svelte";
	import Game from "./Game.svelte";
	import { Router, Route } from "svelte-navigator";
	import { SvelteToast } from "@zerodevx/svelte-toast";
	import { state } from "./store";
	
	let title: string = "GYF";
	$: title = `GYF - ${$state}`;
	$: {
		document.title = title;
	}

	export let url = "";
</script>

<title>{title}</title>

<main>
	<SvelteToast />

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
</style>
