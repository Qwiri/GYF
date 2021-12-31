<script lang="ts">
    import { toast } from "@zerodevx/svelte-toast";
    import { useFocus } from "svelte-navigator";

    import Avatar from "../assets/Avatar.svelte";
    import { username, ws } from "../store";

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

<input
    use:registerFocus
    bind:value={$username}
    on:keypress={onKey}
    name="Username"
    placeholder="Enter Name..."
/>
<Avatar user="{$username}" />
<input type="button" value="JOIN GAME" on:click={connect} />