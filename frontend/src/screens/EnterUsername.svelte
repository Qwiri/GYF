<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import { useFocus } from "svelte-navigator";

    import Avatar from "../assets/Avatar.svelte";
    import { username, ws } from "../store";

    export let id: string;
    const registerFocus = useFocus();

    const connect = () => {
        // check if a name was entered
        if ($username.trim().length === 0) {
            toast.push("Please enter a username");
            return;
        }
        if ($ws.readyState === 1) {
            $ws.send(`JOIN ${$username}`);
        }
    };

    const onKey = (event: KeyboardEvent) => {
        // check if the key is enter
        if (event.key !== "Enter") {
            return;
        }
        connect();
    };
</script>

<div class="header">
    <h1>
        GYF
        <img src="https://i.gifer.com/2iFd.gif" height="64px" alt="cat" />
    </h1>
    <h2>ANY DESCRIPTION HERE</h2>
    <input type="button" value="LEARN HOW TO PLAY" />
</div>

<div class="invite">
    <p>You were invited to game <span>{id}</span></p>
    <div class="avatar">
        <Avatar user={$username} width="100px" />
    </div>
    <input
        type="text"
        use:registerFocus
        bind:value={$username}
        on:keypress={onKey}
        name="Username"
        placeholder="Enter Name..."
    />
    <input type="button" value="JOIN GAME" on:click={connect} />
</div>

<style lang="scss">
    .invite {
        margin-top: 6em;

        input[type="button"] {
            background-color: #fff500;
            border-radius: 5px;
            font-weight: bold;
        }

        input[type="text"] {
            background-color: #131313;
            border: none;
            color: white;
        }

        input {
            font-size: 1.3rem;
        }

        p {
            color: grey;
            span {
                color: #ffcb7e;
            }
        }
    }
    .header {
        input[type="button"] {
            background-color: #131313;
            color: grey;
            border: none;
            padding: 1em;
            border-radius: 7px;
        }

        h1 {
            font-size: 6em;
            color: greenyellow;
            text-transform: uppercase;
            font-weight: 1000;

            margin: 0;
            margin-top: 5rem;
        }
        h2 {
            margin-top: 0;
            font-weight: normal;
            color: grey;
        }
    }
</style>
