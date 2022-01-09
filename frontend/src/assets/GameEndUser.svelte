<script lang="ts">
    import { username } from "../store";

    import Avatar from "./Avatar.svelte";

    export let name;
    export let height = "auto";
</script>

<div class="avatar">
    <Avatar user={name} width="auto" {height} />
    <span class:self={name === $username}>{name}</span>
</div>

<style lang="scss">
    .self {
        color: greenyellow;
    }
    .avatar {
        height: 5rem;
        width: 5rem;
        overflow: hidden;
        display: flex;
        flex-direction: column;
    }
</style>
