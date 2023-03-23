cd ..
cd internal/parser
python parser.py
cd ..
cd .. 
cp .\internal\parser\*.txt .\cmd
cd cmd
go run main.go