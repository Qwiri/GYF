<script lang="ts">
    import { round } from "../store";
</script>

<div class="topic">
    <!-- Display Round Number -->
    <p>
        {$round.currentRound}/{$round.totalRounds}
    </p>
    <!-- Display Topic -->

    <h2>{$round.topic}</h2>
</div>

<style lang="scss">
    p {
        color: greenyellow;
    }
</style>
