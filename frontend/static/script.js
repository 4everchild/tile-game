import {render} from "./render.js"
import {setSelected,getSelected, addHandlersToDraw, addPlayerEvents } from "./events.js";
import { renderHistory } from "./renderHistory.js";
import { addHistoryEvents } from "./historyEvents.js";


const data = document.getElementById("game-data").textContent.trim(); 

let gameobj = JSON.parse(data);

export const nPlayers = gameobj.players.length
export const url = window.location.href

export const gameroot = document.querySelector(".game-container")
export const historyroot = document.querySelector(".history-container")

export let history = null

function getActivePlayer(root){
    const state = root.querySelector(".state")
    console.log(state)
    let index = state.textContent.charAt(1)
    return root.querySelector(".player"+index)
}

//console.log(getActivePlayer(gameroot))


//hover and selected functionality to generate input move

await refresh(gameroot, historyroot)




//function addDrawEventListeners

export  async function refreshGame(gameroot,obj){
    setSelected(null)
    const fragment = render(obj)
    gameroot.replaceChildren(fragment)
}

export function addGameHandlers(gameroot){
    addHandlersToDraw(gameroot)
    const player = getActivePlayer(gameroot)
    console.log(player)
    addPlayerEvents(player)
}

export async function refresh(gameroot, historyroot){
    history = await fetchHistory()
    let game = history.States.at(-1)
    await refreshGame(gameroot,game)
    await refreshHistory(historyroot,history)
    addGameHandlers(gameroot)
    addHistoryEvents(historyroot)
}

function refreshHistory(historyroot,history){
    const fragment = renderHistory(history)
    historyroot.replaceChildren(fragment)
}


async function fetchHistory(){
    try {
        const response = await fetch(url+'/history');
        if (!response.ok) throw new Error('Server error');
        const result = await response.json();
        return result
    } catch (err) {
        console.error('Request failed:', err);
    }
} 

//render(gameroot,obj)