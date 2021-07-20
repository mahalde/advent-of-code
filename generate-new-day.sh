DEFAULTDAY=$(date +'%-d')
DEFAULTYEAR=$(date +'%-Y')
DAY=${1:-$DEFAULTDAY}
YEAR=${2:-$DEFAULTYEAR}

echo "Creating new file structure for day $DAY of year $YEAR"

if [ $DAY -lt 10 ]; then
  mkdir $YEAR/day-0$DAY
  cp template $YEAR/day-0$DAY/day0$DAY.go
  touch $YEAR/day-0$DAY/README.md
  echo "# Advent of Code $YEAR: Day 0$DAY" >$YEAR/day-0$DAY/README.md
else
  mkdir $YEAR/day-$DAY
  cp template $YEAR/day-$DAY/day$DAY.go
  touch $YEAR/day-$DAY/README.md
  echo "# Advent of Code $YEAR: Day $DAY" >$YEAR/day-$DAY/README.md

fi

echo "Files successfully created :)"
echo "Puzzle input still needs to be fetched"
git add $YEAR/
