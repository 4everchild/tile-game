let selected
{// blinking effect on hover for factory displays
let drawTiles = document.querySelector(".drawing-container").querySelectorAll(".tile")

let fds = document.querySelectorAll(".factory-display");
let centers = document.querySelectorAll(".center")

fds.forEach(fd  => {
    let tiles = fd.querySelectorAll(".tile")
    tiles.forEach( tile => {
        tile.addEventListener("mouseenter" , ()=>{
            const groupClass = tile.classList[1]
            group = fd.querySelectorAll("."+groupClass)
            for (const t of group){t.classList.add("BLINK")}
        })
        tile.addEventListener("mouseout" , ()=>{
            fd.querySelectorAll(".BLINK")
            .forEach(t => t.classList.remove("BLINK"));
        })
        tile.addEventListener ("click",() =>{
            for (t of drawTiles){t.classList.remove("OPAQUE1")}

            if (selected) for (t of selected){t.classList.remove("SELECTED")}
            selected = fd.querySelectorAll(".tile.BLINK")
            if (selected) for (t of selected){t.classList.add("SELECTED")}

            for (t of drawTiles){ 
                if (!t.classList.contains("SELECTED")){t.classList.add("OPAQUE1")} 
            }
        }) 
    })
});

centers.forEach(center =>{
    let tiles = center.querySelectorAll(".tile")
    tiles.forEach( tile => {
        tile.addEventListener("mouseenter" , ()=>{
            const groupClass = tile.classList[1]
            group = center.querySelectorAll("."+groupClass)
            for (const t of group){t.classList.add("BLINK")}
        })
        tile.addEventListener("mouseout" , ()=>{
            center.querySelectorAll(".BLINK")
            .forEach(t => t.classList.remove("BLINK"));
        })
        tile.addEventListener ("click",() =>{
            for (t of drawTiles){t.classList.remove("OPAQUE1")}

            if (selected) for (t of selected){t.classList.remove("SELECTED")}
            selected =center.querySelectorAll(".tile.BLINK")
            if (selected) for (t of selected){t.classList.add("SELECTED")}

            for (t of drawTiles){ 
                if (!t.classList.contains("SELECTED")){t.classList.add("OPAQUE1")} 
            }
        }) 
    })
})
    


}