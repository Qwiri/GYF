<script lang="ts">
    import Avatar from "../../assets/Avatar.svelte";
import Image from "../../assets/Image.svelte";
    import TopicDisplay from "../../assets/TopicDisplay.svelte";
    import { leader, votingResults, ws } from "../../store";


    $: maxVotes = $votingResults[0]?.voters.length
</script>

<TopicDisplay />

<h1>Voting results</h1>
<div class="resultWrapper">
    {#each $votingResults as result}
        <div class="votingResult">
            <Image width="auto" height="auto" src={result.url} alt="Image of {result.creator}" />
            <div class="overlayWrapper {result.voters.length === maxVotes ? 'first' : result.voters.length > 0 ? 'second' : ''}">
                {#if result.voters.length === maxVotes}
                    <svg
                        viewBox="0 0 44 46"
                        fill="none"
                        class="coolicon"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            style="fill:none;stroke:#f29e51;stroke-width:6.4252;stroke-miterlimit:4;stroke-dasharray:none"
                            d="M 3.825032,42.708813 22.292464,24.690138 40.643678,42.943806"
                            class="coolicon-1"
                        />
                        <path
                            style="fill:none;stroke:#f29e51;stroke-width:6.4252;stroke-miterlimit:4;stroke-dasharray:none"
                            d="M 4.0025544,23.703704 22.35249,5.4551441 40.759897,23.703704"
                            class="coolicon-2"
                        />
                    </svg>
                {:else if result.voters.length > 0}
                    <svg
                        viewBox="0 0 44 46"
                        class="coolicon"
                        fill="none"
                        version="1.1"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            class="single-coolicon"
                            style="fill:none;stroke:#2d9cdb;stroke-width:6.4252;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                            d="M 3.7523315,33.475853 22.219764,15.457178 40.570978,33.710846"
                        />
                    </svg>
                {/if}
                <span class="{result.voters.length === maxVotes ? 'firstText' : result.voters.length > 0 ? 'secondText' : ''}">{result.creator.toUpperCase()}</span>
                <div class="votedBy overlay">
                    {#each result.voters as voter}
                        <div class="avatar">
                            <Avatar user={voter} width="32px" />
                        </div>
                    {/each}
                </div>
            </div>
        </div>
    {/each}
</div>

<style lang="scss">
    @keyframes cooliconSlideIn {
        0% {
            transform: translateY(100px);
        }
        50% {
            transform: translateY(0);
        }
        100% {
            transform: translateY(-100px);
        }
    }
    @keyframes background-effects {
        from {
        }
        0% {
            backdrop-filter: none;
            background-color: none;
        }
        50% {
            backdrop-filter: blur(3px) grayscale(.8);
            background-color: var(--custom-bg);
        }
        100% {
            backdrop-filter: none;
            background-color: none;
        }
    }


    .single-coolicon {
        --custom-delay: 4s;
        animation: cooliconSlideIn calc(var(--custom-delay))
            cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    }

    .coolicon-1 {
        --custom-delay: 4s;
        animation: cooliconSlideIn calc(var(--custom-delay))
            cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    }
    .coolicon-2 {
        --custom-delay: 3s;
        animation: cooliconSlideIn calc(var(--custom-delay))
            cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    }

    span {
        position: absolute;
        background-color: #131313;
        text-align: left;
        display: inline-flex;
        align-items: center;
        bottom: 0;
        left: 0;
        padding-left: .5rem;
        height: 1rem;
        width: 100%;
        font-weight: bold;
        font-size: .8rem;
        padding: .8rem;
    }

    .resultWrapper {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: 2rem;
        justify-content: center;

    }

    .coolicon {
        width: 20%;
        position: absolute;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        margin: auto;
        filter: drop-shadow(2px 4px 6px black);
    }
    .votingResult {
        position: relative;
        width: clamp(200px, 20vw, 400px);
        height: clamp(200px, 20vw, 400px);
        overflow: hidden;
        background-color: #131313;
        display: flex;
        flex-direction: column;
        justify-content: center;
        border-radius: .5rem;

        :global(.imageComponent) {
            width: 100%;
        }
    }

    .first {
        --custom-bg: #ffda004f;
    }
    .second {
        --custom-bg: #2d9cdb82;
    }

    .firstText {
        color: #FFCB7E
    }
    .secondText {
        color: #2d9cdb;
    }

    .overlayWrapper {
        // backdrop-filter: blur(3px) grayscale(.8);
        position: absolute;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        --custom-delay: 4s;
        animation: background-effects calc(var(--custom-delay)) cubic-bezier(0.250, 0.460, 0.450, 0.940) both;

    }

    .overlay {
        padding: 0.2rem;
        position: absolute;
    }

    .votedBy {
        display: flex;
        bottom: 0;
        right: 0;
    }
</style>
