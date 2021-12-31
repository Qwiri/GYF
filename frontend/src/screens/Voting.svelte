<script>

    import { ws, submissions } from "../store.ts";

    console.log($submissions)

    const submitVote = (e) => {
        $ws.send(`VOTE ${e.srcElement.dataset.url}`)
    }
    
</script>

<div id="submissionsWrapper">
    {#each $submissions as submission}
        <div class="image">
            <img src="{submission}" alt="">
            <div data-url="{submission}" on:click="{submitVote}" class="overlay"></div>
        </div>

    {/each}
</div>

<style lang="scss">
    #submissionsWrapper {
        display: grid;

        .overlay {
            bottom: 0;
            left: 0;
            opacity: 0.5;
            position: absolute;
            right:0;
            top:0;
            z-index: 1;

            &:hover {
                background-color: #181818b6
            }
        }

        .image{
            position: relative;
            width: min-content;
        }
    }
    
</style>

