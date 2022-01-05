<script lang="ts" >
import TopicDisplay from "../assets/TopicDisplay.svelte";
import { gifSubmitted, ws } from "../store";


    let searchQuery;
    let provider = "Giphy";
    let searchResults = [];

    let submission = "";


    const handleEnter = async (e: KeyboardEvent) => {
        if (e.key !== "Enter") {
            return;
        }
        searchResults = [];

        if (provider === "Giphy") {
            let res = await fetch(`https://api.giphy.com/v1/gifs/search?api_key=epo51yrPMWiwryp1w5xbOEK9gJUpGbIX&q=${searchQuery}`).catch(e => console.log(e));
            let body = await res.json().catch(e => console.log(e))
            searchResults = body.data;

        } else if (provider === "Tenor") {
            let res = await fetch(`https://g.tenor.com/v1/search?q=${searchQuery}&key=LIDSRZULELA&limiT=8`).catch(e => console.log(e));
            let body = await res.json().catch(e => console.log(e))
            searchResults = body.results;

        }

    }

    const handleGifClick = (e: MouseEvent, r) => {
        submission = returnGifFullUrl(r);
        $ws.send(`SUBMIT_GIF ${returnGifFullUrl(r)}`)
    }

    const returnGifPreviewUrl = (r) => {
        if (provider === "Giphy") {
            return r.images.preview_gif.url;
        } else if (provider === "Tenor") {
            return r.media[0].nanogif.url;
        }

    }
    const returnGifFullUrl = (r) => {
        if (provider === "Giphy") {
            return r.images.original.url;
        } else if (provider === "Tenor") {
            return r.media[0].gif.url;
        }

    }

    const changeProvider = (e: MouseEvent) => {
        searchResults = [];
        if (provider === "Giphy") {
            provider = "Tenor";

        } else if (provider === "Tenor") {
            provider = "Giphy";
        }


    }

</script>
<TopicDisplay />
{#if !$gifSubmitted}
    <div id="searchWrapper">
        <div id="searchBarWrapper">
            <div id="providerChoice">
                <span id="shownProvider">{provider}</span>
                <span id="otherProvider" on:click="{changeProvider}">{provider === "Giphy" ? "Tenor" : "Giphy"}</span>
            </div>
            <input type="text" placeholder="Search via {provider}" on:keypress={handleEnter} bind:value={searchQuery}>
            {#if provider === "Giphy"}
                <img id="poweredByGiphy" src="/assets/Poweredby_100px-Black_VertLogo.png" alt="Powered by Giphy" />
            {/if}
        </div>
        <div id="resultWrapper">
            {#if searchResults !== []}
                {#each searchResults as r}
                <div class="imgContainer">
                    <img src="{returnGifPreviewUrl(r)}" on:click={e => handleGifClick(e, r)} alt="gif">
                </div>
                {/each}
            {/if}
        </div>
    </div>
{:else}
    <div id="submissionWrapper">
        <h1>Your Submission</h1>
        <img id="submission" src={submission} alt="Your submission" />
        <button on:click={e => {$gifSubmitted = false}}>Choose another</button>
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
        background-color: #24FF00;
        margin-top: 1rem;

        &:hover {
            cursor: pointer;
        }
    }

    img:hover {
        cursor: pointer;
        opacity: .5;
    }

    .imgContainer {
        width: 8rem;
        height: 8rem;
        display: flex;
        flex-direction: column;
        justify-content: center;
        border-radius: 1rem;
        background-color: #181818;
        overflow: hidden;
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
        color:white;

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
        justify-self: flex-end;

        &:hover {
            cursor: default;
            opacity: 1;
        }
    }

    #submission {
        height: 10rem;
    }

    #submissionWrapper {
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
</style>