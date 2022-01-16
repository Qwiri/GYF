<script lang="ts">
import { version } from "../../package.json"
</script>

<div id="footerWrapper" >
    <div id="footerGrid" >
        <p>Frontend version: {version}</p>
        <p>Backend version:</p>
        <p>Made with ❤️ by Tom Doil and Daniel Statzner</p>
        <p>This project is available on <a href="https://github.com/Qwiri/GYF">Github</a></p>
    </div>
</div>

<style lang="scss">

#footerWrapper {
    height: 20vh;
    background-color: black;
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
        
}
#footerGrid {
    width: 80%;
    height: min-content;
    display: grid;
    place-content: center;
    grid-template-rows: repeat(2, 1fr);
    grid-auto-flow: column;
    justify-content: space-between;
    align-content: center;
    row-gap: .5rem;
    justify-items: start;


    p {
        margin: 0;
    }
}


</style>