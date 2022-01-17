<script lang="ts">
    import Avatar from "../../assets/Avatar.svelte";
    import { resetGameValues } from "../../utils";

    import { state, username, stats } from "../../store";
    import { GameState } from "../../types";
    import GameEndUser from "../../assets/GameEndUser.svelte";

    let firstPlace = [];
    let secondPlace = [];
    let honorableMentions = [];

    // create a sorted list of the points in game
    let points = [...new Set(Object.values($stats))].sort((a,b) => b-a);

    Object.entries($stats).forEach(([username, userPoints]) => {
        console.log(username, userPoints);
        if (userPoints === points[0]) {
            firstPlace.push(username);
        } else if (userPoints === points[1]) {
            secondPlace.push(username);
        } else {
            honorableMentions.push(username);
        }
    });
    console.log({ points, firstPlace, secondPlace, honorableMentions });

    function gotoLobby() {
        state.set(GameState.Lobby);
        resetGameValues();
    }
</script>

<div id="wholeScreen">
    <div id="avatarDiv">
        <Avatar user={$username} width="auto" />
    </div>
    <div id="resultDiv">
        <h1>THE WINNERS ARE SET</h1>
        <h2>THANKS FOR PLAYING! 👑</h2>
        <div id="ranking">
            <div id="firstPlace">
                <div id="firstPlaceBadge" class="badge">
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
                </div>
                <div id="firstPlaceAvatars">
                    {#each firstPlace as name}
                        <GameEndUser {name} height="78%" />
                    {/each}
                </div>
                <div class="pointsField">
                    <span class="pointText">{points[0]}</span>
                    <span>votes</span>
                </div>
            </div>
            <div id="secondPlace">
                <div id="secondPlaceBadge" class="badge">
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
                </div>
                <div id="secondPlaceAvatars">
                    {#each secondPlace as name}
                        <GameEndUser {name} height="78%" />
                    {/each}
                </div>
                {#if points[1] !== undefined}
                    <div class="pointsField">
                        <span class="pointText">{points[1]}</span>
                        <span>votes</span>
                    </div>
                {/if}
            </div>
            <div id="honorableMentions">
                <div id="honorableMentionsBadge" class="badge" />
                <div id="honorableMentionsAvatars">
                    {#each honorableMentions as name}
                        <GameEndUser {name} height="78%" />
                    {/each}
                </div>
            </div>
        </div>
        <button on:click={gotoLobby}>Back to Lobby</button>
    </div>
</div>

<style lang="scss">
    h1,
    h2 {
        margin: 0;
    }

    @keyframes cooliconSlideIn {
        0% {
            transform: translateY(100px);
        }
        100% {
            transform: translateY(0);
        }
    }
    .single-coolicon {
        --custom-delay: 2s;
        animation: cooliconSlideIn calc(var(--custom-delay))
            cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    }

    .coolicon-1 {
        --custom-delay: 2s;
        animation: cooliconSlideIn calc(var(--custom-delay))
            cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    }
    .coolicon-2 {
        --custom-delay: 1s;
        animation: cooliconSlideIn calc(var(--custom-delay))
            cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    }

    .badge {
        width: 6rem;
        height: 6rem;
        display: flex;
        justify-content: center;
        align-items: center;
        border-radius: 0.5rem;
        margin: 0.5rem;
        flex-shrink: 0;

        svg {
            height: 50%;
        }
    }

    .pointText {
        color: #24ff00;
        font-size: 1.3rem;
        font-weight: bold;
    }

    #firstPlaceBadge {
        background-color: #ffcb7e;
    }
    #secondPlaceBadge {
        background-color: #4fbfff;
    }
    #honorableMentionsBadge {
        background-color: #545454;
    }

    #wholeScreen {
        min-height: 100vh;
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        justify-content: space-around;
        width: 100vw;

        @media (max-width: 40rem) {
            padding: 1rem;
        }
    }
    #avatarDiv {
        height: 100%;
        width: 25vw;
        display: flex;
        justify-content: center;

        :global(img) {
            height: 100vh;
        }

        @media (max-width: 40rem) {
            display: none;
        }
    }

    #resultDiv {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    #ranking {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        min-width: 30vw;

        #firstPlace,
        #secondPlace,
        #honorableMentions {
            display: flex;
            background-color: #131313;
            border-radius: 0.5rem;
        }

        #firstPlaceAvatars,
        #secondPlaceAvatars,
        #honorableMentionsAvatars {
            min-height: 6rem;
            display: flex;
            flex-wrap: wrap;
            align-items: center;
            background-color: #131313;
            border-radius: 0.5rem;
            row-gap: 1rem;
            padding: 0.5rem 0;
            max-width: 20rem;
        }

        .pointsField {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            width: 6rem;
            height: 6rem;
            background-color: black;
            margin-left: auto;
            border-radius: 0.5rem;
            margin: 0.5rem;
            margin-left: auto;
            flex-shrink: 0;
        }
    }
</style>
