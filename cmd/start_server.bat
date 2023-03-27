
pip install requests
pip install bs4
pip install lxml
cd ..
cd internal/parser
python parser.py
cd ..
cd .. 
copy .\internal\parser\*.txt .\cmd
cd cmd
go run server.go

