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
            console.log(selected)
            
            

            if (tile.classList.contains("SELECTED")){
                selected=null
                resetSelected(selected,total)
                return
            }
            
            selected=null
            resetSelected(selected,total)
            console.log(selected)

            

            selected = Array.from(container.querySelectorAll(".tile.BLINK"))
            for (const t of selected){t.classList.add("SELECTED")}
            
            
            // if selecting from center we must select 1st as well
            if (container.classList[0]=="center"){
                for (const t of total) {
                    if (t.classList[1]=="FIRST"){
                        t.classList.add("SELECTED")
                        selected.push(t)
                    }
                }
            }

            for (const t of total){ 
                if (!t.classList.contains("SELECTED")){t.classList.add("OPAQUE1")} 
            }
            console.log(getColorSelected())

        })
    }
}

function resetSelected(selected, total){
    for (const t of total){
        t.classList.remove("OPAQUE1")
        t.classList.remove("SELECTED")
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
    
    const pls = root.querySelectorAll(".patternline-container")
    const wallTiles = root.querySelector(".wall").querySelectorAll(".tile")
    
    for (let i=0;i<5;i++){
        addPatternlinesEvents(pls[i],Array.from(wallTiles).slice(i*5,i*5+5),i)
    }

    const floor = root.querySelector(".floor")
    addFloorEvents(floor)
}

function addFloorEvents(floor){
    floor.addEventListener("click", () => {
    console.log(selected)
        if(selected != null){
            console.log("### is valid! ###\n")
        }else{
            console.log("### not valid ###\n")
        }
    })
}

function addPatternlinesEvents(pl,wallTiles){
    console.log(pl)
    console.log(wallTiles)
    pl.addEventListener("click", () =>{
        //console.log(isSelectedMoveValid())
        console.log(pl)
        console.log(wallTiles)
        if(isSelectedMoveValidForPatternline(pl,wallTiles)){
            //perform request here
            console.log("### is valid! ###\n")
        }else{
            console.log("### not valid ###\n")
        }


    })
}


//only for patternline event, not for general purpouse
function isSelectedMoveValidForPatternline(patternline,wallTiles){
    let color = getColorSelected()
    if (!color){
        console.log("### 1")
        return false
    } 
    // means !selected
    if (color == "FIRST"){
        console.log("### 2")
        return false
    }

    for (const tile of wallTiles){
        console.log(tile)
        if (color == tile.classList[1] && !tile.classList.contains("OPAQUE")){
            console.log("### 3")
            return false
        }
    }

    // selected already cannot be empty
    if (getColorPatternline(patternline) != color && getColorPatternline(patternline) !="EMPTY"){
        return false 
    }

    if (countTilesInPatterline(patternline,"EMPTY") == 0){
        return false
    }

    return true 
}

function getColorSelected(){
    if(!selected){return null}
    for (const tile of selected){
        return tile.classList[1]
    }
    return null;
}

// TODO look here
function getColorPatternline(pl){
    const tiles = pl.querySelectorAll(".tile")
    for (const tile of tiles){
        if (tile.classList[1]=="EMPTY"){continue}
        return tile.classList[1]
    }
    return "EMPTY"
}

function isFirstSelected(){
    for (const tile of selected){
        if (tile.classList[1]=="FIRST"){return true}
    }
    return false
}

function countTilesInPatterline(pl,color){
    let i =0
    const tiles = pl.querySelectorAll(".tile")
    for (const tile of tiles){
        if (tile.classList[1] == color){i++}
    }
    return i
}


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

