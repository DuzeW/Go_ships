package dumpster

controllerHTTP.GameStatus()

controllerHTTP.GameStatus()
char := opBoard.Listen(context.TODO())
result := controllerHTTP.Fire(char)
if result == "miss" {
miss = append(miss, char)
}
if result == "hit" || result == "sunk" {
hit = append(hit, char)
}
if result == "sunk" {
for i := 0; i < len(hit); i++ {
x, y := coordToInts(hit[i])
for j := -1; j < 2; j++ {
for k := -1; k < 2; k++ {
if x+j <= 9 && x+j >= 0 && y+k <= 9 && y+k >= 0 {
miss = append(miss, intsToCoord(x+j, y+k))
}
}
}
}
}
for i := 0; i < len(miss); i++ {
x, y := coordToInts(miss[i])
opStates[x][y] = gui.Miss
}
for i := 0; i < len(hit); i++ {
x, y := coordToInts(hit[i])
opStates[x][y] = gui.Hit
}
txt.SetText(fmt.Sprintf("Twój ruch Coordinate: %s %s Pozostały czas %d Skuteczność strzałów: dobra", char, result, controllerHTTP.Timer))
opBoard.SetStates(opStates)
ui.Log("Coordinate: %s", char)
for i := 0; i < len(controllerHTTP.OppShots); i++ {
x, y := coordToInts(controllerHTTP.OppShots[i])
states[x][y] = gui.Miss
}
for i := 0; i < len(controllerHTTP.OppShots); i++ {
x, y := coordToInts(controllerHTTP.OppShots[i])
for j := 0; j < len(myCoords); j++ {
if controllerHTTP.OppShots[i] == myCoords[j] {
states[x][y] = gui.Hit
}
}
}
board.SetStates(states)
controllerHTTP.GameStatus()
if controllerHTTP.Status != "game_in_progress" {
break
}