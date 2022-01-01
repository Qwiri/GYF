<script>
    import { ws, submissions } from "../../store.ts";

    const submitVote = (e) => {
        $ws.send(`VOTE ${e.srcElement.dataset.url}`)
    }
</script>

<div id="submissionsWrapper">
    {#each $submissions as submission}
        <div class="image">
            <img class="img" src="{submission}" alt="">
            <div class="background"></div>
            <div data-url="{submission}" on:click="{submitVote}" class="overlay"></div>
        </div>

    {/each}
</div>

<style lang="scss">
    #submissionsWrapper {
        display: flex;
        flex-wrap: wrap;
        gap: 2rem;
        justify-items: center;
        justify-content: center;

        .overlay {
            bottom: 0;
            left: 0;
            position: absolute;
            right:0;
            top:0;
            z-index: 2;

            &:hover {
                box-shadow: inset 0px 0px 10px 1000px rgb(173 255 47 / 50%);
                backdrop-filter: grayscale(1);
            }
        }

        .img {
            width: clamp(200px, 20vw, 400px);
            z-index: 1;
        }
        .background {
            position: absolute;
            top: 0;
            bottom: 0;
            left: 0;
            right: 0;
            background-color: #131313;
            border-radius: 0.5rem;
            z-index: -1;
        }
        .image{
            position: relative;
            display: flex;
            flex-direction: column;
            justify-content: center;
            width: clamp(200px, 20vw, 400px);
            height: clamp(200px, 20vw, 400px);
        }
    }
</style>
