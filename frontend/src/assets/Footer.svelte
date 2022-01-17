<script lang="ts">
    import { version } from "../../package.json";
    import { backendVersion } from "../store";

    let hover = false;

</script>

<div id="footerWrapper">
    <div id="footerGrid">
        <table id="tabver">
            <tr>
                <td>Frontend</td>
                <td>{version.split(":")[0]}:<a href="https://github.com/Qwiri/GYF/commit/{version.split(":")[1]}" target="_blank" class="aaaaa">{version.split(":")[1]}</a></td>
            </tr>
            {#if $backendVersion}
                <tr>
                    <td>Backend</td>
                    <td>{$backendVersion.split(":")[0]}:<a href="https://github.com/Qwiri/GYF/commit/{$backendVersion.split(":")[1]}" target="_blank" class="aaaaa">{$backendVersion.split(":")[1]}</a></td>
                </tr>
            {/if}
        </table>
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
            <a href="https://github.com/Qwiri/GYF" id="github-link">
                Fork this project on
                <img
                    src="/assets/GitHub-Mark-Light-32px.png"
                    alt="GitHub"
                    id="github-mark"
                />
            </a>
        </p>
    </div>
</div>

<style lang="scss">
    #tabver {
        text-align: left;
    }
    #tabver td:first-child {
        font-weight: bold;
        text-align: right;
    }

    #github-mark {
        height: 1.2rem;
        width: 1.2rem;
        margin-left: 0.5rem;
    }

    .aaaaa {
        color: #22ff00 !important;
    }

    #github-link {
        display: flex;
        align-items: center;
    }

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
        min-height: min-content;
        display: flex;
        justify-content: space-between;
        align-content: center;
        row-gap: 0.5rem;
        flex-wrap: wrap;

        * {
            flex-shrink: 0;
        }
        p {
            margin: 0;
        }
        a {
            text-decoration: none;
            color: rgba(255, 255, 255, 0.808);
        }
    }
    #versionParagraphs {
        text-align: left;
    }
</style>
