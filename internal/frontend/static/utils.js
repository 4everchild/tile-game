let colorValue =[
"EMPTY",
"BLUE",
"YELLOW",
"RED",
"BLACK",
"GREEN",
"FIRST",
"OPAQUE BLUE",
"OPAQUE YELLOW",
"OPAQUE RED",
"OPAQUE BLACK",
"OPAQUE GREEN"
]

let stateString =["SETUP","END","P1","P2","P3","P4"]


export function addDiv(root,classes,text){
    const tmp = document.createElement("div")
    tmp.classList.add(...classes.split(" "));
    tmp.textContent = text

    root.append(tmp)
}

export function makeDiv(classes,text){
    const tmp = document.createElement("div")
    tmp.classList.add(...classes.split(" "));
    tmp.textContent = text

    return tmp
}


export{colorValue, stateString}