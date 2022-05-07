<script lang="ts" >

    export let src: string;
    export let alt: string;

    // if no height is given, the loading animation can't be shown
    export let width: string = "100px";
    export let height: string = "100px";

    let loaded = false;

</script>
<div class="imageComponent" id="imageContainer" style="width: {width}; height: {height};">

    <!-- only play animation while image isn't loaded -->
    {#if !loaded}
        <div id="loadingDiv">
            <svg
            viewBox="0 0 81.88971 39.421074"
            version="1.1"
            id="svg"
            xmlns="http://www.w3.org/2000/svg">
            <g
                transform="translate(-50.94151,-81.034871)">
                <path
                style="fill:none;stroke:#24fd00;stroke-width:8;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-opacity:1"
                d="M 71.471849,116.3318 C 90.312046,116.16493 93.35513,85.296553 113.45884,85.036357 133.5728,84.776029 133.96011,116.22354 114.54964,116.45462 95.139169,116.6857 85.289207,84.860857 71.391071,85.075926 51.274686,85.142387 47.684068,116.54248 71.471849,116.3318 Z"
                id="stroke" />
            </g>
            </svg>
        </div>
    {/if}

    <!-- the actual image -->
    <img id="image" loading="lazy" {src} {alt} on:load={e => loaded = true}/>

    
</div>


<style lang="scss">
    #imageContainer {
        position: relative;
    }

    #image {
        vertical-align: top;
        object-fit: contain;
        width: 100%;
        height: 100%;
    }

    #loadingDiv {
        background-color: #131313;
        border-radius: .5rem;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    #svg {
        width: 50%;
    }

    @keyframes animateInfinityIcon {
        from{
            stroke-dashoffset: 0;
        }
        to {
            stroke-dashoffset: 209.14979553222656;
        }
    }

    #stroke {
        stroke-dasharray: 189.14979553222656 20;
        animation-name: animateInfinityIcon;
        animation-timing-function: linear;
        animation-duration: 3s;
        animation-iteration-count: infinite;
        animation-fill-mode: forwards;
    }

</style>