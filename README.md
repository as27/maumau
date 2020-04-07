# maumau
Simples MauMau Spiel für mehrere Spieler.

Das Spiel wird durch den Server verwaltet. Der Server stellt Websockets zur Verfügung und übermittelt darüber laufend den aktuellen Spielstand.

## Definitionen

### Stapel

Der Stapel das sind die gemischten Karten. Vom Stapel ziehen die Spieler immer die oberste Karte.

### Haufen

Auf den Haufen werden die unterschiedlichen Karten abgelegt. Dabei dürfen nur bestimmte Karten auf den Stapel gelegt werden.

## Events

Für die Umsetzung werden verschiedene Events definiert. 

- [x] newGame: Neues Spiel
- [ ] Erste Karte Aufdecken
- [x] newPlayer(name): Spieler hinzufügen
- [x] pushCardToStack: Karte auf den Stapel legen (wird nach dem Mischen der Karten verwendet)
- [ ] setActivePlayer: Spieler am Zug
- [ ] playerReady: Spieler ist fertig mit dem Zug
- [x] popCardFromStack: Spieler nimmt oberste Karte vom Stapel
- [x] pushCardToHeap: Spieler legt Karte auf den Haufen
- [ ] gameOver: Spieler hat gewonnen

## GUI

Umsetzung erfolg über HTML

- [x] HTML Template Struktur
- [x] vue.js Rendering via JSON Input
- [x] Anbindung an Wesocket
- [ ] Abfangen von nicht gültigen Zügen
- [ ] Anmeldebildschirm
- [ ] Undo und Redo
- [ ] Neues Spiel

Definition der einzelnen Phasen für den Client

- [ ] Status: Warte Anmeldung Partner
- [ ] Status: Warte auf Zug
- [ ] Status: Spiele Karte
- [ ] Status: Nimm eine Karte
- [ ] Status: Wähle Farbe

## Server

- [x] Game State für Spieler
- [x] Game Start
- [ ] Validierung der Züge 
- [x] Undo und Redo Handler
- [x] Neues Spiel Handler