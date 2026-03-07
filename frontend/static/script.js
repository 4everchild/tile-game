//hover and selected functionality to generate input move
let selected

let drawTiles = document.querySelector(".drawing-container").querySelectorAll(".tile")

let fds = document.querySelectorAll(".factory-display");
let centers = document.querySelectorAll(".center")

function addTilesEventListeners(container, total){
    let tiles = container.querySelectorAll(".tile")
    tiles.forEach( tile => {
        tile.addEventListener("mouseenter" , ()=>{
            const groupClass = tile.classList[1]
            group = container.querySelectorAll("."+groupClass)
            for (const t of group){t.classList.add("BLINK")}
        })
        tile.addEventListener("mouseout" , ()=>{
            container.querySelectorAll(".BLINK")
            .forEach(t => t.classList.remove("BLINK"));
        })
        tile.addEventListener ("click",() =>{
            for (t of total){t.classList.remove("OPAQUE1")}

            if (selected) for (t of total){t.classList.remove("SELECTED")}
            selected = container.querySelectorAll(".tile.BLINK")
            if (selected) for (t of selected){t.classList.add("SELECTED")}
            
            
            // if selecting from center we must select 1st as well
            console.log(container.classList)
            if (container.classList[0]=="center"){
                for (t of total) {
                    if (t.classList[1]=="FIRST"){t.classList.add("SELECTED")}
                }
            }

            for (t of total){ 
                if (!t.classList.contains("SELECTED")){t.classList.add("OPAQUE1")} 
            }
            
        }) 
    })
}

// now each tile in fds and center is hoverable and can select a valid move
fds.forEach(fd  => addTilesEventListeners(fd, drawTiles))
centers.forEach(center => addTilesEventListeners(center, drawTiles))



// now for the selection of the patternline on which the current selected tiles could be put
// only available and legal moves should be put there
let player;

let players = document.querySelectorAll(".player")
p = document.querySelector(".state").classList[1]

if (p == "P1"){player=players[0]}
if (p == "P2"){player=players[1]}
if (p == "P3"){player=players[2]}
if (p == "P4"){player=players[3]}

let pcs = player.querySelectorAll(".patternline-container")
console.log(pcs)



function render(data){
    console.log(data)
}