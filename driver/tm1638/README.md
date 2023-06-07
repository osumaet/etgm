# TM1638 chip driver

Draft version

tm1638.go contains implementation of very basic IC's functions like opening and closing connection, brightness control, keyboard buffer reading, display memory clearing, byte sending.
It can't contain any functions for displaying symbols or LEDs driving because different schemas of connection between IC and indicators possible. Indication function should be implemented soon.