export {getSelected, setSelected, addHandlersToDraw, addPlayerEvents}
import { nPlayers, url, refresh, gameroot } from "./script.js"

import { colorValue } from "./utils.js"
let selected

function getSelected(){
    return selected
}

function setSelected(val){
    selected = val
}


function addSelected(tile){
    selected.push(tile)
}

function resetSelected(total){
    for (const t of total){
        t.classList.remove("OPAQUE1")
        t.classList.remove("SELECTED")
    }
}

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

            if (tile.classList.contains("SELECTED")){
                setSelected(null)
                resetSelected(total)
                return
            }
            
            setSelected(null)
            resetSelected(total)

            setSelected( Array.from(container.querySelectorAll(".tile.BLINK")) )
            for (const t of selected){t.classList.add("SELECTED")}
            
            
            // if selecting from center we must select 1st as well
            if (container.classList[0]=="center"){
                for (const t of total) {
                    if (t.classList[1]=="FIRST"){
                        t.classList.add("SELECTED")
                        addSelected(t)
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


function addHandlersToDraw(root){
    let drawTiles = root.querySelector(".drawing-container").querySelectorAll(".tile")
    let fds = root.querySelectorAll(".factory-display");
    let centers = root.querySelectorAll(".center")
    // now each tile in fds and center is hoverable and can select a valid move
    fds.forEach(fd  => addTilesEventListeners(fd, drawTiles))
    centers.forEach(center => addTilesEventListeners(center, drawTiles))

    const dc = root.querySelector(".drawing-container")
    addDrawingContainerEvents(dc, drawTiles)
    

    // TODO fix player selection
    /*
    const pls = root.querySelectorAll(".patternline-container")
    const wallTiles = root.querySelector(".wall").querySelectorAll(".tile")
    
    for (let i=0;i<5;i++){
        addPatternlinesEvents(pls[i],Array.from(wallTiles).slice(i*5,i*5+5),i)
    }

    const floor = root.querySelector(".floor")
    addFloorEvents(floor)
    */
}

function addPlayerEvents(player){
    // TODO fix player selection
    const pls = player.querySelectorAll(".patternline-container")
    const wallTiles = player.querySelector(".wall").querySelectorAll(".tile")
    
    for (let i=0;i<5;i++){
        addPatternlinesEvents(pls[i],Array.from(wallTiles).slice(i*5,i*5+5),i)
    }

    const floor = player.querySelector(".floor")
    addFloorEvents(floor)
}

function addFloorEvents(floor){
    floor.addEventListener("click", async() => {
    //console.log(selected)
        if(selected != null){
            const move = {group: getGroupSelected(),color: getColorNumberSelected(),row: 5}// TODO perform request here
            console.log(move)
            try {
                const response = await fetch(url+'/move', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(move)
                });

                if (!response.ok) throw new Error('Server error');

                const result = await response.json();
                refresh(gameroot,result)

            } catch (err) {
                console.error('Request failed:', err);
            }
        }else{
            console.log("### not valid ###\n")
        }
    })
}

function addPatternlinesEvents(pl,wallTiles){
    const plid = Number.parseInt(pl.dataset.index)
    pl.addEventListener("click", async () =>{
        if(isSelectedMoveValidForPatternline(pl,wallTiles)){
            const move = {group: getGroupSelected(),color: getColorNumberSelected(),row: plid }// TODO perform request here
            console.log(move)
            try {
                const response = await fetch(url+'/move', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(move)
                });

                if (!response.ok) throw new Error('Server error');

                const result = await response.json();
                refresh(gameroot,result)
                console.log('Success:', result);

            } catch (err) {
                console.error('Request failed:', err);
            }
        }else{
            console.log("### not valid ###\n")
        }


    })
}

function getGroupSelected(){
    if(!selected){return null}

    for (const tile of selected){
        const parent = tile.parentElement
        if(parent.dataset.index){
            return Number.parseInt(parent.dataset.index)
        }else{
            return 2*nPlayers +1
        }
    }
}

function getColorSelected(){
    console.log(selected)
    if(!selected){return null}
    for (const tile of selected){
        return tile.classList[1]
    }
    return null;
}

function getColorNumberSelected(){
    console.log(selected)
    if(!selected){return null}
    for (const tile of selected){
        return colorValue.indexOf(tile.classList[1])
    }
    return null;
}


//only for patternline event, not for general purpouse
function isSelectedMoveValidForPatternline(patternline,wallTiles){
    console.log(selected)
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

function addDrawingContainerEvents(dc,total){
    dc.addEventListener("click", (e) => {
        setSelected(null)
        resetSelected(total)
    })
}

