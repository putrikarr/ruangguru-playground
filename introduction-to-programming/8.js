const CELL_SIZE = 20;
const SNAKE_COLOR = "purple";
const CANVAS_SIZE = 400;
const REDRAW_INTERVAL = 100;
const WIDTH = CANVAS_SIZE / CELL_SIZE;
const HEIGHT = CANVAS_SIZE / CELL_SIZE;

let snakePositionX = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);
let snakePositionY = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let ctx = snakeCanvas.getContext("2d");

        ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
        ctx.fillStyle = SNAKE_COLOR;
        ctx.fillRect(snakePositionX * CELL_SIZE, snakePositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    }, REDRAW_INTERVAL);
}

document.addEventListener("keydown", function (event) {
    if (event.key === "ArrowLeft") {
        snakePositionX = snakePositionX - 1;
        if (snakePositionX < 0) {
            snakePositionX = CANVAS_SIZE / CELL_SIZE - 1;
        }
    } else if (event.key === "ArrowRight") {
        snakePositionX = snakePositionX + 1;
        if (snakePositionX >= WIDTH) {
            snakePositionX = 0;
        }
    } else if (event.key === "ArrowUp") {
        snakePositionY = snakePositionY - 1;
        if (snakePositionY < 0) {
            snakePositionY = CANVAS_SIZE / CELL_SIZE - 1;
        }
    } else if (event.key === "ArrowDown") {
        snakePositionY = snakePositionY + 1;
        if (snakePositionY >= HEIGHT) {
            snakePositionY = 0;
        }
    }
})

