import { makeDiv, addDiv } from "./utils.js"


export function renderHistory(obj){
    const fragment = document.createDocumentFragment();
    addHistory(fragment,obj)
    return fragment
}

function addHistory(root, history){
    console.log(history)
    let i=0
    for (const game in history.States){
        const tmp = makeDiv("game-history-item", "TURN "+i )
        tmp.dataset.index = i
        root.append(tmp)
        i++
    }
}