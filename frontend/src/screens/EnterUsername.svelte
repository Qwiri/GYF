<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import { useFocus } from "svelte-navigator";

    import Avatar from "../assets/Avatar.svelte";
    import Header from "../assets/Header.svelte";
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
        name="Username"
        placeholder="Enter Name..."
    />
    <input type="button" value="JOIN GAME" on:click={connect} />
</div>

<style lang="scss">
    .invite {
        input[type="button"] {
            background-color: #fff500;
            border-radius: 5px;
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
</style>
