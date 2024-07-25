<script lang="ts">
    export let width: string = "100px";
    export let height: string = "auto";
    export let user: string = "";
</script>

<img
    alt="avatar of {user}"
    {width}
    {height}
    src="https://api.dicebear.com/9.x/miniavs/svg?seed={user}"
/>
