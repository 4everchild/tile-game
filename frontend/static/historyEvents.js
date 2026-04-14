let selected = null

export {addHistoryEvents}

function addHistoryEvents(historyroot){
    const history = historyroot.querySelectorAll(".game-history-item")
    console.log(history)
    for (const his of history ){
        addHistoryEvent(his)
    }
}

function addHistoryEvent(his){
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
    })
}
