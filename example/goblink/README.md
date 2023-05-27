
# GoBlink (Blink + Goroutine)

## About

That Blink does not use MCU lock to make pause between LED on and off.
Ofcouse it's not most optimal way to do that. But it's funny.

## Build & Flash

Build with [TinyGo](https://tinygo.org/)   
Use `tinygo flash` with `-scheduler=tasks` with proper port and target.