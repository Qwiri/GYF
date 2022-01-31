<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import { useFocus } from "svelte-navigator";

    import Avatar from "../assets/Avatar.svelte";
    import Header from "../assets/Header.svelte";
    import { username, ws } from "../store";

    export let id: string;
    const registerFocus = useFocus();

    let passwordBuffer = "";

    const connect = () => {
        // check if a name was entered
        if ($username.trim().length === 0) {
            toast.push("Please enter a username");
            return;
        }
        // check if name contains spaces
        if ($username.indexOf(" ") !== -1) {
            toast.push("Username cannot contain spaces");
            return;
        }
        if ($ws.readyState === 1) {
            $ws.send(`JOIN ${$username} ${passwordBuffer}`);
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

<Header />

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
        placeholder="Enter Name..."
    />
    <br />
    <input
        type="password"
        bind:value={passwordBuffer}
        placeholder="Lobby Password (optional)"
        autocomplete="off"
    />
    <br />
    <input type="button" value="JOIN GAME" on:click={connect} />
</div>

<style lang="scss">
    .invite {
        input[type="button"] {
            background-color: #24ff00;
            border-radius: 5px;
            border: none;
            font-weight: bold;
            &:hover {
                cursor: pointer;
            }
        }

        input[type="text"] {
            background-color: #131313;
            border: none;
            color: white;
        }

        input[type="password"] {
            background-color: #131313;
            border: none;
            color: salmon;
            font-size: 0.8rem;

            text-security: disc; 
            -webkit-text-security:disc;
        }

        input {
            font-size: 1.3rem;
        }

        p {
            color: grey;
            span {
                color: #24ff00;
            }
        }
    }
</style>
