<script lang="ts">
    import TopicDisplay from "../../assets/TopicDisplay.svelte";
    import { ws, submissions, gifSubmitted } from "../../store";

    let gifVoted;

    const submitVote = (e: MouseEvent, i: number) => {
        $ws.send(`VOTE ${e.srcElement.dataset.url}`);

        if (gifVoted === undefined) {
            gifVoted = i;
        }
    };
</script>

<TopicDisplay />

<div id="submissionsWrapper">
    {#each $submissions as submission, i}
        <div class="image">
            {#if i === gifVoted}
                <svg id="votedSvg" viewBox="0 0 44 40" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M0.333008 12.2C0.33247 9.05906 1.59484 6.0498 3.8361 3.84929C6.07737 1.64878 9.10927 0.441829 12.2497 0.499995C15.9705 0.480235 19.521 2.05821 21.9997 4.83333C24.4784 2.05821 28.0288 0.480235 31.7497 0.499995C34.8901 0.441829 37.922 1.64878 40.1632 3.84929C42.4045 6.0498 43.6669 9.05906 43.6663 12.2C43.6663 23.8047 29.8452 32.5667 21.9997 39.5C14.1715 32.5082 0.333008 23.8133 0.333008 12.2Z" fill="#219653"/>
                </svg>
            {/if}

            <img class="img" src={submission} alt="" />
            <div class="background" />
            <div id="{i === gifVoted ? 'votedGif' : ''}" data-url={submission} on:click={e => submitVote(e, i)} class="overlay" />
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
            right: 0;
            top: 0;
            z-index: 2;
            border-radius: 0.5rem;

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
        .image {
            position: relative;
            display: flex;
            flex-direction: column;
            justify-content: center;
            width: clamp(200px, 20vw, 400px);
            height: clamp(200px, 20vw, 400px);
            overflow: hidden;
        }

        #votedSvg {
            position: absolute;
            top: 0;
            bottom: 0;
            left: 0;
            right: 0;
            margin: auto;
            width: 4rem;
            z-index: 2;

            animation: plop-heart .5s cubic-bezier(0.075, 0.82, 0.165, 1) 0s 1 alternate forwards;
        }

        #votedGif {
            background-color: #2196537d;
            animation: plop-heart-bg .5s cubic-bezier(0.075, 0.82, 0.165, 1) 0s 1 alternate forwards;
        }

    }

    @keyframes plop-heart {
        from {
            width: 0;
            opacity: 0;

        }
        to {
            width: 4rem;
            opacity: 1;

        }
    }
    @keyframes plop-heart-bg {
        from {
            opacity: 0;

        }
        to {
            opacity: 1;

        }
    }
</style>
