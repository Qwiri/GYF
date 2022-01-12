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

    let gifPreviewWindow: HTMLImageElement;
    let gifPreviewURL: string;

    let searched = false;
    $: if (searchQuery) searched = false;

    let blur = false;

    const fetchFirstGifs = async () => {
        searchResults = []; // #42
        searchResults = await provider.search(searchQuery, true);
        searched = true;
    };

    const fetchGifs = async () => {
        let newResults = await provider.search(searchQuery);
        searchResults = [...searchResults, ...newResults];
        searched = true;
    };

    const submitGif = (e: MouseEvent, r: SearchResult) => {
        submission = r.original_url;
        $ws.send(`SUBMIT_GIF ${r.original_url}`);
    };

    const chooseNew = (_: MouseEvent) => {
        $gifSubmitted = false;
    };

    const changeProvider = (_: MouseEvent) => {
        provider =
            Providers[(Providers.indexOf(provider) + 1) % Providers.length];
        // clear search results
        searchResults = [];
        fetchFirstGifs();
    };

    let timer;

    const debounce = (e: KeyboardEvent) => {
        if (e.key === "Enter") {
            return;
        }

        clearTimeout(timer);
        timer = setTimeout(async () => {
            await fetchFirstGifs();
        }, 300);
    };

    const previewGif = (e: MouseEvent, gif?: SearchResult) => {
        gifPreviewURL = gif?.original_url ?? "";

        if (gifPreviewWindow) {
            gifPreviewWindow.style.left = e.pageX + "px";
            gifPreviewWindow.style.top = e.pageY + "px";
        }
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
                on:keyup={debounce}
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
        {#if gifPreviewURL}
            <img
                id="gifPreviewWindow"
                bind:this={gifPreviewWindow}
                src={gifPreviewURL}
                alt="GYF preview"
                class:blurImage={blur}
                on:load={(_) => (blur = false)}
            />
        {/if}

        <div id="resultWrapper" on:mouseleave={(e) => previewGif(e, undefined)}>
            {#if searchResults.length > 0}
                {#each searchResults as result}
                    <div
                        class="imgContainer"
                        on:mousemove={(e) => previewGif(e, result)}
                        on:mouseleave={(_) => (blur = true)}
                        on:click={(e) => submitGif(e, result)}
                    >
                        <Image
                            width="100%"
                            height="100%"
                            src={result.preview_url}
                            alt="gif"
                        />
                    </div>
                {/each}
                <button on:click={fetchGifs}>Load more</button>
            {:else if searchQuery && searched}
                <div class="noResults">
                    No results found. Try another search.
                </div>
            {/if}
        </div>
    </div>
{:else}
    <div id="submissionWrapper">
        <h1>Your Submission</h1>
        <Image
            width="auto"
            height="auto"
            src={submission}
            alt="Your submission"
        />
        <button on:click={chooseNew}>Choose another</button>
    </div>
{/if}

<style lang="scss">
    .noResults {
        background-color: #ff3838;
        border-radius: 0.8rem;
        padding: 0.4rem;
        color: #131313;
    }

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

    #gifPreviewWindow {
        position: absolute;
        z-index: 1;
        pointer-events: none;
        width: 30rem;
        height: 30rem;
        object-fit: contain;
        background-color: #181818;
        border-radius: 0.5rem;
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

        &:hover {
            cursor: pointer;
        }
    }

    .blurImage {
        filter: blur(10px);
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
