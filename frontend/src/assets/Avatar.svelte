<script lang="ts">
    export let width: string = "100px";
    export let user: string = "";
</script>

<img
    alt="avatar of {user}"
    {width}
    src="https://avatars.dicebear.com/api/miniavs/{user}.svg"
/>
