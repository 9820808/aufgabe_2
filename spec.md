# Spezifikation: Composite-Demonstrator „Verschachtelbares Fenster“

## 1. Ziel

Der Demonstrator zeigt das Composite-Pattern: Ein Baum aus Fensterteilen
wird über dieselbe Schnittstelle (`Item`) angesprochen wie ein einzelnes
Teil. Der Aufrufer ruft `Price()` bzw. `Display()` nur an der Wurzel auf
und erhält das Ergebnis für den gesamten Baum.

## 2. Begriffe und Rollen

| Pattern-Rolle | Typ      | Kennung (Display) | Beschreibung                        |
|---------------|----------|-------------------|-------------------------------------|
| Component     | `Item`   | –                 | Gemeinsame Schnittstelle            |
| Composite     | `Window` | `Window`          | Fenster, enthält Rahmen + 4 Felder  |
| Leaf          | `Glass`  | `Glass Panel`     | Glasscheibe                         |
| Leaf          | `Frame`  | `Frame`           | Fensterrahmen                       |

## 3. Schnittstelle `Item`

```go
type Item interface {
    Price() int
    Display() string
}
```

- **S-1** `Glass`, `Frame` und `Window` implementieren `Item` über
  Pointer-Receiver (`*Glass`, `*Frame`, `*Window`).

## 4. Struktur eines Fensters

- **S-2** Ein `Window` besteht aus genau **einem** `Frame` und genau
  **vier** Feldern.
- **S-3** Jedes Feld ist entweder ein `Glass` oder ein `Window`. Die vier
  Felder dürfen gemischt sein (z.B. Feld 0–2 `Glass`, Feld 3 `Window`).
- **S-4** Die Felder sind von 0 bis 3 nummeriert, gezählt von links nach
  rechts, oben nach unten:

  ```
  +-----+-----+
  |  0  |  1  |
  +-----+-----+
  |  2  |  3  |
  +-----+-----+
  ```

- **S-5** Die Wurzel eines Baums ist immer ein `Window`.
- **S-6** Erzeugt wird ein Fenster über
  `NewWindow(frame *Frame, panes [4]Item) (*Window, error)`.
  Die Anzahl 4 ist durch den Array-Typ garantiert. Ein Fehler wird
  zurückgegeben, wenn
  - `frame` `nil` ist, oder
  - ein Feld `nil` ist, oder
  - ein Feld weder `*Glass` noch `*Window` ist (z.B. ein `*Frame`).
- **S-7** Nach der Erzeugung ist ein Fenster unveränderlich. Dadurch ist
  der Baum ohne Locking thread-safe lesbar. Es gibt keinen Store.

## 5. Preise

- **S-8** `Glass.Price()` liefert **30**.
- **S-9** `Frame.Price()` liefert **50**.
- **S-10** `Window.Price()` hat keinen Eigenpreis. Es liefert die Summe aus
  `Frame.Price()` und den `Price()`-Werten der vier Felder (rekursiv).
- Folgerung: Ein Fenster mit vier Glasfeldern kostet 50 + 4·30 = **170**.

## 6. Anzeige

- **S-11** Jedes Element gibt eine Zeile mit seiner Kennung aus (siehe
  Abschnitt 2). Jede Zeile endet mit `\n`.
- **S-12** `Window.Display()` liefert zuerst die eigene Zeile, danach in
  dieser Reihenfolge die Ausgaben von: `Frame`, Feld 0, Feld 1, Feld 2,
  Feld 3 (Tiefensuche, Pre-Order: jedes Kind steht direkt unter seinem
  Elternteil).
- **S-13** Die Ausgaben der Kinder werden pro Zeile um **zwei Leerzeichen**
  eingerückt. Dadurch entspricht die Einrückung der Tiefe im Baum.
- **S-14** Leaves (`Glass`, `Frame`) geben nur ihre eigene Zeile ohne
  Einrückung aus; die Einrückung übernimmt das umschließende `Window`.

Beispiel (Fenster mit Feld 3 als Unterfenster):

```
Window
  Frame
  Glass Panel
  Glass Panel
  Glass Panel
  Window
    Frame
    Glass Panel
    Glass Panel
    Glass Panel
    Glass Panel
```

## 7. Programm

- **S-15** Nicht interaktive Kommandozeilen-Anwendung ohne Parameter,
  Start mit `go run ./src`.
- **S-16** Das Programm baut drei Beispielbäume und gibt für jeden eine
  Überschrift, `Display()` und `Price()` aus.

### Beispielbäume

| Nr. | Name             | Aufbau                                                                                   | Preis |
|-----|------------------|------------------------------------------------------------------------------------------|-------|
| 1   | Einfaches Fenster | 4× Glass                                                                                | 170   |
| 2   | Geteiltes Feld   | Feld 0–2 Glass, Feld 3 Window (4× Glass) — Beispiel aus der Aufgabe                     | 310   |
| 3   | Sprossenfenster  | Feld 0 Window (4× Glass), Feld 1 Glass, Feld 2 Beispiel 2 (verschachtelt), Feld 3 Glass | 590   |

Rechnung Beispiel 3: 50 + 170 + 30 + 310 + 30 = 590.

## 8. Verzeichnisstruktur

```
go.mod
spec.md            diese Spezifikation
src/main.go        Einstieg, baut und gibt die Beispielbäume aus
src/window/        Item, Glass, Frame, Window
test/              automatisierte Tests (go test ./...)
```

## 9. Tests

- **T-1** Preise der Leaves (S-8, S-9).
- **T-2** Preise der drei Beispielbäume (S-10, Abschnitt 7).
- **T-3** Exakte `Display()`-Ausgabe für Leaves und Beispiel 2 (S-11–S-14).
- **T-4** `NewWindow` liefert Fehler bei `nil`-Rahmen, `nil`-Feld und
  `Frame` als Feld (S-6).
- **T-5** Ein `Window` wird als `Item` verwendet (Wurzel und Feld werden
  gleich behandelt, S-1/S-3).
