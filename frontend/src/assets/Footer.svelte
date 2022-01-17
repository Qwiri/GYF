<script lang="ts">
    import { version } from "../../package.json";

    let hover = false;
</script>

<div id="footerWrapper">
    <div id="footerGrid">
        <p>
            <b>Frontend:</b>
            {version}
            <b>| Backend:</b>
        </p>
        <p>
            Made with ❤️ by <a
                target="_blank"
                href="https://github.com/Qwiri"
                on:mouseenter={(e) => (hover = true)}
                on:mouseleave={(e) => (hover = false)}
                >{hover ? "🥝" : "Qwiri"}</a
            >.
        </p>
        <p>
            Fork this project on <a href="https://github.com/Qwiri/GYF"
                >GitHub</a
            >
        </p>
    </div>
</div>

<style lang="scss">
    #footerWrapper {
        height: 15vh;
        background-color: black;
        color: rgba(255, 255, 255, 0.808);
        display: flex;
        align-items: center;
        justify-content: center;
    }
    #footerGrid {
        width: 80%;
        height: min-content;
        display: grid;
        place-content: center;
        grid-template-columns: repeat(3, 1fr);
        grid-auto-flow: column;
        justify-content: space-between;
        align-content: center;
        row-gap: 0.5rem;
        justify-items: center;

        p {
            margin: 0;
        }
        a {
            text-decoration: none;
            color: rgba(255, 255, 255, 0.808);
        }
    }
</style>
