frame=100 # frame size

curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "bgrect $frame $frame $((800-frame)) $((800-frame))"
curl -X POST http://localhost:17000 -d "update"
