<script lang="ts">
    import TopicDisplay from "../assets/TopicDisplay.svelte";
    import { gifSubmitted, ws } from "../store";
    import { Giphy, Providers } from "./search";
    import type { Provider, SearchResult } from "./search";
import Image from "../assets/Image.svelte";

    let provider: Provider = Giphy; // Make Giphy the default provider

    let searchQuery: string;
    let searchResults: Array<SearchResult> = [];

    let submission: string = "";

    const handleEnter = async (e: KeyboardEvent) => {
        if (e.key !== "Enter") {
            return;
        }
        fetchFirstGifs();
    };

    const fetchFirstGifs = async () => {
        searchResults = await provider.search(searchQuery);
    }

    const fetchGifs = async () => {
        let newResults = await provider.search(searchQuery);
        searchResults = [...searchResults, ...newResults];
    }

    const submitGif = (e: MouseEvent, r: SearchResult) => {
        submission = r.original_url;
        $ws.send(`SUBMIT_GIF ${r.original_url}`);
    };

    const chooseNew = (_: MouseEvent) => {
        $gifSubmitted = false;
    };

    const changeProvider = (_: MouseEvent) => {
        provider = Providers[(Providers.indexOf(provider) + 1) % Providers.length];
        // clear search results
        searchResults = [];
    };

</script>

<TopicDisplay />
{#if !$gifSubmitted}
    <div id="searchWrapper">
        <div id="searchBarWrapper">
            <div id="providerChoice">
                <span id="shownProvider">{provider.name}</span>
                <span id="otherProvider" on:click={changeProvider}>
                    {Providers[
                        (Providers.indexOf(provider) + 1) % Providers.length
                    ].name}
                </span>
            </div>
            <input
                type="text"
                class="gyf-bar"
                placeholder="Search via {provider.name} 🔍"
                on:keypress={handleEnter}
                bind:value={searchQuery}
            />

            <!-- Display Giphy Badge -->
            {#if provider === Giphy}
                <img
                    id="poweredByGiphy"
                    src="/assets/Poweredby_100px-Black_VertLogo.png"
                    alt="Powered by Giphy"
                />
            {/if}
        </div>
        <div id="resultWrapper">
            {#if searchResults.length > 0}
                {#each searchResults as result}
                    <div class="imgContainer" on:click={(e) => submitGif(e, result)} >
                        <Image
                            width="100%"
                            height="100%"
                            src={result.preview_url}
                            alt="gif"
                        />
                    </div>
                {/each}
                <button on:click={fetchGifs}>Load more</button>
            {/if}
        </div>
    </div>
{:else}
    <div id="submissionWrapper">
        <h1>Your Submission</h1>
        <Image width="auto" height="auto" src={submission} alt="Your submission" />
        <button on:click={chooseNew}>Choose another</button>
    </div>
{/if}

<style lang="scss">

    #searchWrapper {
        background-color: #131313;
    }
    #resultWrapper {
        width: 100%;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: 1rem;
        justify-content: center;
        max-height: 50vh;
        overflow-y: scroll;
    }

    button {
        background-color: #24ff00;
        border: none;
        margin-top: 1rem;

        &:hover {
            cursor: pointer;
        }
    }


    .imgContainer {
        width: 8rem;
        height: 8rem;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        border-radius: 1rem;
        background-color: #181818;
        overflow: hidden;

        :global(.imageComponent):hover {
            cursor: pointer;
            opacity: 0.5;
        }
    }

    #shownProvider {
        color: greenyellow;

        &:before {
            content: "▼";
        }
    }

    #otherProvider {
        display: none;
    }

    #providerChoice {
        margin: 1rem;
        width: 5rem;
    }

    #providerChoice:hover span {
        display: block;
        color: white;

        &:before {
            content: "";
        }

        &:hover {
            color: greenyellow;

            &:before {
                content: "🡆";
            }
        }
    }

    #searchBarWrapper {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        max-height: 3rem;

        input {
            margin: 0;
        }
    }

    #poweredByGiphy {
        margin-left: 1em;
        justify-self: flex-end;

        &:hover {
            cursor: default;
            opacity: 1;
        }
    }


    #submissionWrapper {
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

        :global(.imageComponent) {
            height: 10rem;
        }
    }
</style>
