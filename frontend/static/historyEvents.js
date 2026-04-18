let selected = null

export {addHistoryEvents}
import {gameroot, history, refreshGame, addGameHandlers} from "./script.js"

function addHistoryEvents(historyroot){
    const history = historyroot.querySelectorAll(".game-history-item")
    console.log(history)
    for (const his of history ){
        addHistoryEvent(his, history.length)
    }
}

function addHistoryEvent(his,l){
    const hid = Number.parseInt(his.dataset.index)
    console.log(his)
    his.addEventListener("mouseenter" , ()=>{
        console.log(his)
        if(!his.classList.contains("SELECTED")){ his.classList.add("BLINK") }
    })

    his.addEventListener("mouseout" , ()=>{
        console.log(his)
        his.classList.remove("BLINK");
    })

    his.addEventListener ("click",(e) =>{
        e.stopPropagation()
        if(selected){selected.classList.remove("SELECTED")}
        if(selected == his){
            his.classList.remove("SELECTED")
            selected=null
        }else{
            his.classList.remove("BLINK")
            his.classList.add("SELECTED")
            selected = his
        }

        let n = Number.parseInt(his.dataset.index,10)

        refreshGame(gameroot,history.States.at(n))

        if(n == l-1){
            addGameHandlers(gameroot)
        }
    })
}
