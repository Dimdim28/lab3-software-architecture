windowLength=800
halfFigure=150
startPosition=$halfFigure
finishPosition=$((windowLength - halfFigure))
maxDistance=$((finishPosition - startPosition))
step=50
second=0.01 # one second

curl -X POST http://localhost:17000 -d "figure $startPosition $startPosition"
curl -X POST http://localhost:17000 -d "update"

position=$startPosition
distance=0
sleep $second

while true; do
    while (( position > startPosition )); do
        if ((distance - step < 0)); then
            curl -X POST http://localhost:17000 -d "move $((-distance)) $((-distance))"
            position=$startPosition
            distance=0
        else
            curl -X POST http://localhost:17000 -d "move $((-step)) $((-step))"
            position=$((position - step))
            distance=$((distance - step))
        fi
            curl -X POST http://localhost:17000 -d "update"
            sleep $second
    done

    while (( position < finishPosition )); do
        if ((distance + step > maxDistance)); then
            curl -X POST http://localhost:17000 -d "move $((maxDistance-distance)) $((maxDistance-distance))"
            position=$finishPosition
            distance=$maxDistance
        else
            curl -X POST http://localhost:17000 -d "move $step $step"
            position=$((position + step))
            distance=$((distance + step))
        fi
            curl -X POST http://localhost:17000 -d "update"
            sleep $second
    done
done