startX=400
startY=150
toRight=300
toDown=600
toLeft=900
toUp=1200
step=50

curl -X POST http://localhost:17000 -d "figure $startX $startY"
curl -X POST http://localhost:17000 -d "update"

while true; do
    for ((i = 0; i < toRight; i += step)); do
        curl -X POST http://localhost:17000 -d "move $step $step"
        curl -X POST http://localhost:17000 -d "update"
    done

    for ((i = toRight; i < toDown; i += step)); do
        curl -X POST http://localhost:17000 -d "move $((-step)) $step"
        curl -X POST http://localhost:17000 -d "update"
    done

    for ((i = toDown; i < toLeft; i += step)); do
        curl -X POST http://localhost:17000 -d "move $((-step)) $((-step))"
        curl -X POST http://localhost:17000 -d "update"
    done

    for ((i = toLeft; i < toUp; i += step)); do
        curl -X POST http://localhost:17000 -d "move $step $((-step))"
        curl -X POST http://localhost:17000 -d "update"
    done
done