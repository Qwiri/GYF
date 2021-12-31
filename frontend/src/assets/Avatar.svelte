<script lang="ts">
    import { username } from "../store";

    export let width: string = "100px";
    export let user: string = "";

    // default user to current user
    if (user.length === 0) {
        user = $username;
    }
</script>

<style>
    
</style>

<img
    alt="avatar of {user}"
    width="{width}"
    src="https://avatars.dicebear.com/api/miniavs/{user}.svg"
/>