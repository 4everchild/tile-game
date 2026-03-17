import {render} from "./render.js"
import { makeDiv, addDiv } from "./utils.js"





const data = document.getElementById("game-data").textContent.trim(); 

let gameobj = JSON.parse(data);


let gameroot = document.querySelector(".game-container")

function getActivePlayer(root){
    const state = root.querySelector(".state")
    console.log(state)
    let index = state.textContent.charAt(1)
    return root.querySelector(".player"+index)
}

//console.log(getActivePlayer(gameroot))


//hover and selected functionality to generate input move
let selected

refresh(gameroot,gameobj)

function addTilesEventListeners(container, total){
    let tiles = container.querySelectorAll(".tile")
    for (const tile of tiles){
        if (tile.classList.contains("EMPTY")){continue}

        tile.addEventListener("mouseenter" , ()=>{
            const groupClass = tile.classList[1]
            let group = container.querySelectorAll("."+groupClass)
            for (const t of group){t.classList.add("BLINK")}
        })

        tile.addEventListener("mouseout" , ()=>{
            container.querySelectorAll(".BLINK")
            .forEach(t => t.classList.remove("BLINK"));
        })

        tile.addEventListener ("click",(e) =>{
            e.stopPropagation()
            resetSelected(selected,total)

            if (tile.classList.contains("SELECTED")){return}

            selected = container.querySelectorAll(".tile.BLINK")
            for (const t of selected){t.classList.add("SELECTED")}
            
            
            // if selecting from center we must select 1st as well
            if (container.classList[0]=="center"){
                for (const t of total) {
                    if (t.classList[1]=="FIRST"){t.classList.add("SELECTED")}
                }
            }

            for (const t of total){ 
                if (!t.classList.contains("SELECTED")){t.classList.add("OPAQUE1")} 
            }
            //console.log(getColorSelected())

        })
    }
}

function resetSelected(selected, total){
    for (const t of total){t.classList.remove("OPAQUE1")
        t.classList.remove("SELECTED")
        selected=null
    }
}

function addHandlersToDraw(root){
    let drawTiles = root.querySelector(".drawing-container").querySelectorAll(".tile")
    let fds = root.querySelectorAll(".factory-display");
    let centers = root.querySelectorAll(".center")
    // now each tile in fds and center is hoverable and can select a valid move
    fds.forEach(fd  => addTilesEventListeners(fd, drawTiles))
    centers.forEach(center => addTilesEventListeners(center, drawTiles))

    const dc = root.querySelector(".drawing-container")
    addDrawingContainerEvents(dc,selected, drawTiles)

    //const pls = root.querySelectorAll(".patternline-container")
    //const wallTiles = root.querySelector(".wall").querySelectorAll(".tile")
    //for (let i=0;i<5;i++){
    //    addPatternlinesEvents(pls[i],Array.from(wallTiles).slice(i*5,i*5+5),i)
    //}
}
/*
function addPatternlinesEvents(pl,wallTiles,i){
    console.log(pl)
    console.log(wallTiles)
    console.log(i)
    pl.addEventListener("mouseenter", () =>{
        //tiles = pl.querySelectorAll("tile")

        // if can't place it shoud stop
        // if it can place it should set the tiles opaque, also the floorline
    })
}
*/

/*
function canPlace(color,patternline,wallTiles,index){
    if (index>5){return false}
    if (index==5){return true}
    
    //plTiles = 
}

function getColorSelected(){
    console.log(selected)
}

*/
function addDrawingContainerEvents(dc,selected,total){
    dc.addEventListener("click", (e) => {
        resetSelected(selected,total)
    })
}



//function addDrawEventListeners

function refresh(gameroot,obj){
    selected = null
    //fragment = document.createDocumentFragment();
    //addGame(fragment,obj)
    const fragment = render(obj)
    gameroot.replaceChildren(fragment)

    addHandlersToDraw(gameroot)
}

//render(gameroot,obj)

