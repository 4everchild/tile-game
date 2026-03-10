//import {addGame} from "./script.js"
import { makeDiv, addDiv } from "./utils.js"
import { colorValue, stateString } from "./utils.js";

export function render(obj){
    const fragment = document.createDocumentFragment();
    addGame(fragment,obj)
    return fragment
}


function addGame(root,game){
    {
    const tmp = makeDiv("drawing-container")
    addFactoryDisplays(tmp,game["factorydisplays"])
    addCenter(tmp,game["center"])
    root.append(tmp)
    }
    {
    const tmp = makeDiv("piles-container")
    addSack(tmp,game["sack"])
    addDiv(tmp,"state", stateString[ game["state"] ] )
    addDiscarded(tmp,game["discarded"])
    addDiv(tmp,"seed", game["seed"] )
    root.append(tmp)
    }

    {
    const tmp = makeDiv("players-container")
    addPlayers(tmp,game["players"])
    root.append(tmp)
    }
}

function addFactoryDisplays(root,factorydisplays){
    const tmp = makeDiv("factory-display-container")
    
    for (const fd of factorydisplays){
        addFactoryDisplay(tmp,fd)
    }
    root.append(tmp)
}

function addFactoryDisplay(root,factorydisplay){
    const tmp = makeDiv("factory-display")

    for (const tile of factorydisplay["tiles"]){
        addTile(tmp,tile)
    }
    
    root.append(tmp)
}

function addCenter(root,center){
    const tmp = makeDiv("center-container")

    if (center["BLUE"] > 0){
        const ttmp = makeDiv("center")
        for(let i =0;i<center["BLUE"];i++){
            addTile(ttmp,1)//index of "BLUE"
        }
        tmp.append(ttmp)
    }
    if (center["YELLOW"] > 0){
        const ttmp = makeDiv("center")
        for(let i =0;i<center["YELLOW"];i++){
            addTile(ttmp,2)//index of "YELLOW"
        }
        tmp.append(ttmp)
    }
    if (center["RED"] > 0){
        const ttmp = makeDiv("center")
        for(let i =0;i<center["RED"];i++){
            addTile(ttmp,3)//index of "RED"
        }
        tmp.append(ttmp)
    }
    if (center["BLACK"] > 0){
        const ttmp = makeDiv("center")
        for(let i =0;i<center["BLACK"];i++){
            addTile(ttmp,4)//index of "BLACK"
        }
        tmp.append(ttmp)
    }
    if (center["GREEN"] > 0){
        const ttmp = makeDiv("center")
        for(let i =0;i<center["GREEN"];i++){
            addTile(ttmp,5)//index of "GREEN"
        }
        tmp.append(ttmp)
    }
    if (center["FIRST"] > 0){
        const ttmp = makeDiv("center")
        for(let i =0;i<center["FIRST"];i++){
            addTile(ttmp,6)//index of "FIRST"
        }
        tmp.append(ttmp)
    }
    
    root.append(tmp)
}

function addTile(root,colornumber){
    let colors = colorValue[colornumber].split(" ")
    const tmp = makeDiv("tile")
    for (const color of colors ){
        tmp.classList.add(color)
    }
    root.append(tmp)
}

function addTileColor(root,color){
    const tmp = makeDiv("tile")
    tmp.classList.add(...color.split(" "))
    root.append(tmp)
}

function addSack(root,sack){
    const tmp = makeDiv("sack-container")

    {
    const ttmp = makeDiv("sack")
    addDiv(ttmp,"amount",sack["BLUE"])

    addTile(ttmp,1)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("sack")
    addDiv(ttmp,"amount",sack["YELLOW"])

    addTile(ttmp,2)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("sack")
    addDiv(ttmp,"amount",sack["RED"])

    addTile(ttmp,3)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("sack")
    addDiv(ttmp,"amount",sack["BLACK"])

    addTile(ttmp,4)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("sack")
    addDiv(ttmp,"amount",sack["GREEN"])

    addTile(ttmp,5)//1 is blue tile

    tmp.append(ttmp)
    }

    root.append(tmp)
}

function addDiscarded(root,discarded){
    const tmp = makeDiv("discarded-container")
    
    {
    const ttmp = makeDiv("discarded")
    addDiv(ttmp,"amount",discarded["BLUE"])

    addTile(ttmp,1)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("discarded")
    addDiv(ttmp,"amount",discarded["YELLOW"])
    
    addTile(ttmp,2)//1 is blue tile

    tmp.append(ttmp)
    }
    
    {
    const ttmp = makeDiv("discarded")
    addDiv(ttmp,"amount",discarded["RED"])
    
    addTile(ttmp,3)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("discarded")
    addDiv(ttmp,"amount",discarded["BLACK"])
    
    addTile(ttmp,4)//1 is blue tile

    tmp.append(ttmp)
    }

    {
    const ttmp = makeDiv("discarded")
    addDiv(ttmp,"amount",discarded["GREEN"])
    
    addTile(ttmp,5)//1 is blue tile

    tmp.append(ttmp)
    }

    root.append(tmp)
}

function addPlayers(root,players){
    let i =0;
    for(const player of players){

        const tmp = makeDiv("player player"+i)
        const ttmp = makeDiv("flex-container-column")

        addName(ttmp,"PLAYER" + (i+1))
        addPlayer(ttmp,player)
        

        tmp.append(ttmp)
        root.append(tmp)

        i++
    }
}




function addPlayer(root,player){
    {
    const tmp = makeDiv("flex-container-row")

    addPatternlines(tmp,player["patternline"])
    addWall(tmp,player["wall"])

    root.append(tmp)
    }
    {
    const tmp = makeDiv("flex-container-row")
    
    addFloor(tmp,player["floorline"])
    addDiv(tmp,"filler")
    addDiv(tmp,"filler")
    addDiv(tmp,"points",player["points"])
    addDiv(tmp,"filler")

    root.append(tmp)
    }
}

function addName(root,name){
    const tmp = makeDiv("flex-container-row")
    
    addDiv(tmp,"filler")
    addDiv(tmp,"player-name","PLAYER: " + name) 
    addDiv(tmp,"filler")

    root.append(tmp)
}

function addPatternlines(root,patternlines){    
    const tmp = makeDiv("patternlines-container")
    for (let i =0;i<5;i++){
        addPatternline(tmp,patternlines[i],i)
    }
    root.append(tmp)
}

function addPatternline(root,patternline,j){
    const tmp = makeDiv("patternline-container")
    for (let i = 0; i<=j; i++){
        if (i<patternline["size"]){
            addTile(tmp,patternline["color"])
        }else{
            addTileColor(tmp,"EMPTY")
        }
    }
    root.append(tmp)
}

function addWall(root,wall){
    const tmp = makeDiv("wall")
    for (const row of wall){
        for (const tile of row){
            addTile(tmp,tile)
        }
    }
    root.append(tmp)
}

function addFloor(root,floor){
    const tmp = makeDiv("floor")
    for (const tile of floor){
        addTile(tmp,tile)
    }
    root.append(tmp)
}
