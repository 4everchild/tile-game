import {render} from "./render.js"
import { makeDiv, addDiv } from "./utils.js"
import {setSelected,getSelected, addHandlersToDraw } from "./events.js";



const data = document.getElementById("game-data").textContent.trim(); 

let gameobj = JSON.parse(data);

export const nPlayers = gameobj.players.length
export const gameId = Number.parseInt( window.location.href.split("/").at(-1) )

let gameroot = document.querySelector(".game-container")

function getActivePlayer(root){
    const state = root.querySelector(".state")
    console.log(state)
    let index = state.textContent.charAt(1)
    return root.querySelector(".player"+index)
}

//console.log(getActivePlayer(gameroot))


//hover and selected functionality to generate input move

refresh(gameroot,gameobj)








//function addDrawEventListeners

function refresh(gameroot,obj){
    setSelected(null)
    //fragment = document.createDocumentFragment();
    //addGame(fragment,obj)
    const fragment = render(obj)
    gameroot.replaceChildren(fragment)

    addHandlersToDraw(gameroot)
}

//render(gameroot,obj)
