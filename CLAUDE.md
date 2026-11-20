# CLAUDE.md

## Projekt

Entwickle einen Demonstrator für das Composite-Pattern anhand eines in sich
verschachtelbaren Fensters um. Ziel des Demonstrators ist es, die
Funktionsweise des Composite-Patterns, im Speziellen den Zugriff auf den
Baum, als wäre es ein einzelnes Objekt, zu zeigen.

Mit folgender Struktur:

```
               +------------+
         +---->| Item       |
   1..n  |     +------------+
         |     | price()    |
         |     | display()  |
         |     +------------+
         |            ^
         |            |
         |            +-----------------+----------------+
         |            |                 |                |
         |            |                 |                |
         |     +------------+    +----------- +    +------------+
         +-----| Window     |    | Glass      |    | Frame      |
               +------------+    +------------+    +------------+
               | price()    |    | price()    |    | price()    |
               | display()  |    | display()  |    | display()  |
               +------------+    +------------+    +------------+
```

Ist ein Baum aufzubauen, wobei die root ein `Window` ist und jedes Window
genau ein `Frame` als auch entweder vier mal `Glass` oder `Window` haben
kann:

```
      +---------------+---------------+
      |               |               |
      |               |               |
      |               |               |
      |               |               |
      |               |               |
      |               |               |
      |               |               |
      +---------------+-------+-------+
      |               |       |       |
      |               |       |       |
      |               |       |       |
      |               +-------+-------+
      |               |       |       |
      |               |       |       |
      |               |       |       |
      +---------------+-------+-------+
```
(Hier als Beispiel eine Unterteilung, wo das dritte (gezählt wird von 0..3
links nach rechts, oben nach unten) `Item` ein weiteres Window ist, 0, 1 und
2 sind `Glass`.

Es soll zwei metoden von `Item` erstellt werden:

* price() -> sammelt aus allen Kindern des Baumes die Ausgabe von price()
  zusammen und gibt die Summe zurück als integer.
* display() -> sammelt aus allen Kindern des Baumes (Breite vor Tiefe) die
  Ausgaben von display und fügt sie zu einem String zusammen. Jedes Element
  gibt eine Zeile mit der eigenen Kennung aus (zB: "Glass Panel" für
  `Glass`). Die Topologie des Baumes soll durch Einrückung sichtbar gemacht
  werden.

Das Programm ist eine nicht interaktive Kommandozeilen-Anwendung ohne
Parameter. Erstelle dazu drei von dir gewählte aussagekräftige Beispielbäume
und lasse price() und display() davon ausgeben.

## Arbeitsweise

Es wird Spec-Driven Development für die Entwicklung angewendet. Das
bedeutet, dass spec.md die maßgebliche Quelle ist und Code und Tests daraus
abgeleitet werden, nicht umgekehrt.

## Directory-Struktor

```
spec.md		Spezifikation
src		Quellcode
test		automatisierte Tests
```

## Konventionen

- Die Implementierungssprache ist go
- Wende go "objektorientiert" an, verwende structs und methoden
- Zuweisung der Funktionen zu structs über Pointer
- Interfaces nur bei echtem Bedarf
- Keine neuen Dependencies ohne Nachfrage
- Keine Secrets im Code (auch bei Prototypen nicht)
- In-Memory-Store muss thread-safe sein

## Regeln

- Keine generierten Dateien von Hand editieren
- Commit-Messages: kurz und imperativ
- Bei Unklarheit fragen statt raten
- Prototyp-Fokus: lieber einfach und klar als elegant und komplex
