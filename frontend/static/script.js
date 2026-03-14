import {render} from "./render.js"
import { makeDiv, addDiv } from "./utils.js"


const data = document.getElementById("game-data").textContent.trim(); 

let obj;
obj = JSON.parse(data);


let gameroot = document.querySelector(".game-container")


//hover and selected functionality to generate input move
let selected

refresh(gameroot,obj)

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

        tile.addEventListener ("click",() =>{
            if (tile.classList.contains("SELECTED")){
                resetSelected(selected,total)
                return
            }
                
            resetSelected(selected,total)
            selected = container.querySelectorAll(".tile.BLINK")
            if (selected) for (const t of selected){t.classList.add("SELECTED")}
            
            
            // if selecting from center we must select 1st as well
            if (container.classList[0]=="center"){
                for (const t of total) {
                    if (t.classList[1]=="FIRST"){t.classList.add("SELECTED")}
                }
            }

            for (const t of total){ 
                if (!t.classList.contains("SELECTED")){t.classList.add("OPAQUE1")} 
            }
            
        })
    }
}

function resetSelected(selected, total){
    for (const t of total){t.classList.remove("OPAQUE1")}
    if (selected) for (const t of total){
        t.classList.remove("SELECTED")
        selected=null
    }
}

function addHandlersToDraw(){
    let drawTiles = document.querySelector(".drawing-container").querySelectorAll(".tile")
    let fds = document.querySelectorAll(".factory-display");
    let centers = document.querySelectorAll(".center")
    // now each tile in fds and center is hoverable and can select a valid move
    fds.forEach(fd  => addTilesEventListeners(fd, drawTiles))
    centers.forEach(center => addTilesEventListeners(center, drawTiles))

    let fdc = document.querySelector("factory-display-container")
    fdc.addDrawEventListeners(fdc,drawTiles)
}


// TODO for the selection of the patternline on which the current selected tiles could be put
// only available and legal moves should be put there
/*
let player;

const players = document.querySelectorAll(".player")
{
const p = document.querySelector(".state").classList[1]

if (p == "P1"){player=players[0]}
if (p == "P2"){player=players[1]}
if (p == "P3"){player=players[2]}
if (p == "P4"){player=players[3]}
}

let pcs = player.querySelectorAll(".patternline-container")
let wall = player.querySelector(".wall").querySelectorAll(".tile")

//console.log(wall[5])

for (const pc of pcs){
    //console.log(pc)
    pc.addEventListener("mouseenter",() => {

    })
    pc.addEventListener("mouseout",() => {

    })
    
}

*/







function refresh(gameroot,obj){
    selected = null
    //fragment = document.createDocumentFragment();
    //addGame(fragment,obj)
    const fragment = render(obj)
    gameroot.replaceChildren(fragment)

    addHandlersToDraw()
}

//render(gameroot,obj)

